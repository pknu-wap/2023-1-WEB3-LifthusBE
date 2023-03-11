package auth

import (
	"lifthus-auth/service/session"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// NewSessionHandler godoc
// @Router       /session/new [post]
// @Summary      gets new connection and assign a lifthus session token.
// @Description  at the same time user connects to lifthus newly, the client requests new session token.
// @Description  and the server returns session id with session token in cookie.
// @ Description then the client send the session id to Hus auth server.
// @Tags         auth
// @Success      201 "returns session id with session token in cookie"
// @Failure      500 "failed to create new session"
func (ac authApiController) NewSessionHandler(c echo.Context) error {
	// get lifthus_st from cookie
	lifthus_st, err := c.Cookie("lifthus_st")
	if err != nil {
		if err != http.ErrNoCookie {
			log.Println("[F] getting lifthus_st from cookie failed: ", err)
			return c.NoContent(http.StatusInternalServerError)
		}
	}
	// if there is session cookie already, just maintain the session.
	// it makes the session to be maintained through the whole tabs of the browser unlike using session storage.
	if lifthus_st != nil {
		return c.NoContent(http.StatusOK)
	}
	/*
		although unlikely, this may cause malfunction. if the session is deleted from database before cookie,
		the client should clear the cookie self.
	*/

	// and if there is no cookie, create new session and return session id

	// create new session and get the session id
	sid, err := session.CreateSession(c.Request().Context(), ac.Client)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	// create new jwt session token with session id
	st := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sid": sid,
		"uid": nil, // it will be omitted actually.
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	hsk := []byte(os.Getenv("HUS_SECRET_KEY"))
	stSigned, err := st.SignedString(hsk)
	if err != nil {
		log.Println("[F] signing session token failed: ", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	// set cookie with session id
	cookie := &http.Cookie{
		Name:     "lifthus_st",
		Value:    stSigned,
		Path:     "/",
		Secure:   false,
		HttpOnly: true,
		Domain:   os.Getenv("LIFTHUS_DOMAIN"),
		SameSite: http.SameSiteDefaultMode,
	}
	c.SetCookie(cookie)

	return c.String(http.StatusCreated, sid)
}
