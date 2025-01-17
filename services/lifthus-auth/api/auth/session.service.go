package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"lifthus-auth/common/db"
	"lifthus-auth/common/dto"
	"lifthus-auth/common/helper"
	"lifthus-auth/common/lifthus"
	"lifthus-auth/ent"
	"strconv"

	"lifthus-auth/common/service/session"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// SessionHandler godoc
// @Tags         auth
// @Router       /session-v2 [get]
// @Summary		 validates session. publishes new one if it isn't. refreshes expired session.
//
// @Success      200 "Ok, session refreshed, session info JSON returned"
// @Success      201 "Created, new session issued, redirect to cloudhus and do connect"
// @Failure      500 "Internal Server Error"
func (ac authApiController) SessionHandler(c echo.Context) error {
	/*
		1. get session token from Authorization header
	*/
	lst, err := helper.GetHeaderLST(c) // returns error if the format is invalid or it's empty.

	/*
		2. if there was no problem while getting the token from header, validate it.
	*/
	var ls *ent.Session
	if err == nil {
		ls, err = session.ValidateSessionQueryUser(c.Request().Context(), lst)
	}

	var nlst string // new session token
	/*
		3-1. try refreshing the session if valid or expired but valid
	*/
	if err == nil || session.IsExpiredValid(err) {
		ls, nlst, err = session.RefreshSessionHard(c.Request().Context(), ls)
	}

	// if refresh succeeded, return the refreshed session token
	if err == nil {
		c = helper.SetAuthHeader(c, nlst)
		// returning sessoin user info
		var uinf *dto.SessionUserInfo
		// if session is signed by any user, return the user info
		if ls.Edges.User != nil {
			lsu := ls.Edges.User
			uinf = &dto.SessionUserInfo{
				UID:        strconv.FormatUint(lsu.ID, 10),
				Registered: lsu.Registered,
				Username:   lsu.Username,
				Usercode:   lsu.Usercode,
			}
		}
		// the client will get OK sign and that is all. no more thing to do.
		return c.JSON(http.StatusOK, uinf)
	}
	log.Printf("issuing new session because of %v", err)
	/*
		3-2. issue new session.
		first, after validation above, the session may turn out to be invalid.
		second, the refresh may have failed.
		in both cases, err won't be nil and the flow comes here.
		then issue a new session.
	*/
	ns, nlst, err := session.CreateSession(c.Request().Context())
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to issue new session")
	}

	c = helper.SetAuthHeader(c, nlst)

	// the client will get Created sign.
	// in this case, the client must redirect to Cloudhus themselves to connect the sessions.
	return c.String(http.StatusCreated, ns.ID.String())
}

// GetSIDHandler godoc
// @Tags         auth
// @Router       /sid [get]
// @Summary		 returns client's SID. should be encrypted later.
//
// @Success      200 "Ok, session ID"
// @Failure      401 "Unauthorized, the token is expired"
// @Failure      500 "Internal Server Error"
func (ac authApiController) GetSIDHandler(c echo.Context) error {
	lst, err := helper.GetHeaderLST(c)
	if err != nil {
		return c.String(http.StatusUnauthorized, "no valid token")
	}
	sid, _, exp, err := session.ValidateSession(c.Request().Context(), lst)
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to validate session")
	} else if exp {
		c.Response().Header().Set("WWW-Authenticate", `Bearer realm="auth.lifthus.com", error="expired_token"`)
		return c.String(http.StatusUnauthorized, "expired_token")
	}
	return c.String(http.StatusOK, sid.String())
}

// SignInPropagationHandler godoc
// @Tags         auth
// @Router       /hus/signin [patch]
// @Summary		 processes user sign-in propagation from cloudhus.
// @Description  the "signin_propagation" token should be included in the request body.
// @Success      200 "Ok, session signed"
// @Failure      400 "Bad Request"
// @Failure	  500 "Internal Server Error"
func (ac authApiController) SignInPropagationHandler(c echo.Context) error {
	sipBytes, err := io.ReadAll(c.Request().Body)
	if err != nil {
		log.Println("failed to read body")
		return c.String(http.StatusInternalServerError, "failed to read body")
	}
	sip := string(sipBytes)

	// parse the token
	sipClaims, expired, err := helper.ParseJWTWithHMAC(sip)
	if expired || err != nil || sipClaims["pps"].(string) != "signin_propagation" {
		log.Println("wrong purpose or expired or erroneous")
		return c.String(http.StatusBadRequest, "invalid token")
	}

	// get the session ID
	sid := sipClaims["csid"].(string)
	suuid, err := uuid.Parse(sid)
	if err != nil {
		log.Println("csid is not a valid uuid")
		return c.String(http.StatusBadRequest, "invalid token")
	}
	hsid := sipClaims["hsid"].(string)
	hsuuid, err := uuid.Parse(hsid)
	if err != nil {
		log.Println("hsid is not a valid uuid")
		return c.String(http.StatusBadRequest, "invalid token")
	}

	// query the session
	ls, err := db.QuerySessionByID(c.Request().Context(), suuid)
	if err != nil || ls == nil {
		log.Println("failed to query session")
		return c.String(http.StatusInternalServerError, "failed to query session")
	}
	if ls.Hsid != nil && *ls.Hsid != hsuuid {
		log.Println("hsid mismatch")
		return c.String(http.StatusBadRequest, "invalid token")
	}

	// marshal the "user" to json bytes
	cuB, err := json.Marshal(sipClaims["user"].(map[string]interface{}))
	if err != nil {
		log.Println("failed to marshal user dto")
		return c.String(http.StatusBadRequest, "invalid token")
	}
	// unmarshal it to HusConnUser
	var cu *dto.HusConnUser
	err = json.Unmarshal(cuB, &cu)
	if err != nil {
		log.Println("failed to unmarshal user dto")
		return c.String(http.StatusBadRequest, "invalid token")
	}

	// query the user and register if not exists
	lu, err := db.QueryUserByID(c.Request().Context(), cu.Uid)
	if err != nil {
		log.Println("failed to query user")
		return c.String(http.StatusInternalServerError, "failed to query user")
	}
	if lu == nil {
		_, err := db.RegisterUser(c.Request().Context(), *cu)
		if err != nil {
			log.Println("failed to register user")
			return c.String(http.StatusInternalServerError, "failed to register user")
		}
	}
	/* REFRESH THE USER INFO WITH THE LATEST ONE (not implemented yet) */

	_, err = ls.Update().SetUID(cu.Uid).SetSignedAt(time.Now()).Save(c.Request().Context())
	if err != nil {
		log.Println("failed to update session")
		return c.String(http.StatusInternalServerError, "failed to update session")
	}

	return c.String(http.StatusOK, "signed")
}

// SignOutPropagationHandler godoc
// @Tags         auth
// @Router       /hus/signout [patch]
// @Summary		 processes user sign-out propagation from cloudhus.
// @Description  the "signout_propagation" token should be included in the request body.
// @Success      200 "Ok, session signed"
// @Failure      400 "Bad Request"
// @Failure	  500 "Internal Server Error"
func (ac authApiController) SignOutPropagationHandler(c echo.Context) error {
	sopBytes, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to read body")
	}
	sop := string(sopBytes)

	// parse the token
	sopClaims, expired, err := helper.ParseJWTWithHMAC(sop)
	if expired || err != nil || sopClaims["pps"].(string) != "signout_propagation" {
		return c.String(http.StatusBadRequest, "invalid token")
	}

	hsid := sopClaims["hsid"].(string)
	hsuuid, err := uuid.Parse(hsid)
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid token")
	}

	// query the latest related session
	ls, err := db.QuerySessionByHsid(c.Request().Context(), hsuuid)
	if err != nil || ls == nil {
		return c.String(http.StatusInternalServerError, "failed to query session")
	}

	// sign out of the session
	_, err = ls.Update().ClearUID().ClearSignedAt().Save(c.Request().Context())
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to update session")
	}

	return c.String(http.StatusOK, "signed out")
}

// SignOutHandler godoc
// @Tags         auth
// @Router       /session/signout [patch]
// @Summary		 gets sign-out request from the client and propagates it to Cloudhus.
// @Success      200 "Ok, signed out of the session"
// @Failure      400 "Bad Request"
// @Failure	  401 "Unauthorized, the token is expired or the session is not signed"
// @Failure	  500 "Internal Server Error"
func (ac authApiController) SignOutHandler(c echo.Context) error {
	// get lifthus_st from the cookie
	lst, err := helper.GetHeaderLST(c)
	if err != nil {
		log.Println("failed to get cookie")
		return c.String(http.StatusBadRequest, "no valid token")
	}

	ls, err := session.ValidateSessionQueryUser(c.Request().Context(), lst)
	if session.IsExpiredValid(err) {
		log.Println("session expired")
		c.Response().Header().Set("WWW-Authenticate", `Bearer realm="auth.lifthus.com", error="expired_token"`)
		return c.String(http.StatusUnauthorized, "expired_token")
	} else if err != nil {
		log.Println("failed to validate session")
		return c.String(http.StatusInternalServerError, "failed to validate session")
	} else if ls.Edges.User == nil {
		log.Println("the session is not signed")
		c.Response().Header().Set("WWW-Authenticate", `Bearer realm="auth.lifthus.com", error="not_signed"`)
		return c.String(http.StatusUnauthorized, "the session is not signed")
	}

	propagCh := make(chan error) // for waiting propagation goroutine
	txCh := make(chan error)     // for waiting transaction goroutine

	go func() {
		sot, err := helper.SignedHusTotalSignOutToken(ls.Hsid.String())
		if err != nil {
			propagCh <- fmt.Errorf("failed to generate token")
			return
		}
		// request to the Cloudhus endpoint
		req, err := http.NewRequest(http.MethodPatch, "https://auth.cloudhus.com/auth/hus/signout"+"", strings.NewReader(sot))
		if err != nil {
			propagCh <- fmt.Errorf("failed to create request")
			return
		}
		resp, err := lifthus.Http.Do(req)
		if err != nil {
			propagCh <- fmt.Errorf("propagation failed")
			return
		}
		defer resp.Body.Close()
		// propagation failed or succeeded
		if resp.StatusCode == http.StatusOK {
			propagCh <- nil
			return
		}
		propagCh <- fmt.Errorf("propagation failed")
	}()

	tx, err := db.Client.Tx(c.Request().Context())
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to start transaction")
	}

	go func() {
		ls, err = tx.Session.UpdateOne(ls).ClearUID().ClearSignedAt().SetTid(uuid.New()).Save(c.Request().Context())
		if err != nil {
			db.Rollback(tx, err)
			txCh <- err
			return
		}
		txCh <- nil
	}()

	propagErr := <-propagCh
	txErr := <-txCh

	if propagErr != nil || txErr != nil {
		_ = db.Rollback(tx, nil)
		log.Println("failed to sign out propag or tx failed")
		return c.String(http.StatusInternalServerError, "failed to sign out")
	}

	err = tx.Commit()
	if err != nil {
		_ = db.Rollback(tx, err)
		log.Println("failed to commit tx")
		return c.String(http.StatusInternalServerError, "failed to sign out")
	}

	lst, err = session.SessionToToken(c.Request().Context(), ls)
	if err != nil {
		log.Println("failed to tokenize the session")
		return c.String(http.StatusInternalServerError, "failed to sign out")
	}

	c = helper.SetAuthHeader(c, lst)

	// from context get uid
	uid, ok := c.Get("uid").(uint64)
	if ok {
		log.Printf("user %d signed out", uid)
	}
	return c.String(http.StatusOK, "signed_out")
}
