package rec

import (
	"lifthus-auth/common/guard"
	"net/http"
	"routine/ent"

	"github.com/labstack/echo/v4"
)

type RecApiControllerParams struct {
	DbClient   *ent.Client
	HttpClient *http.Client
}

// NewRecApiController returns Echo instance comprising of rec api routes to main.
func NewRecApiController(recApi *echo.Echo, params RecApiControllerParams) *echo.Echo {
	recApiController := newRecApiController(params)

	/* REC */
	// create program rec
	recApi.POST("/routine/rec/program/weekly", recApiController.createWeeklyProgramRec, guard.UserGuard)
	// query program rec

	// query routineact recs by given range of date or specific date
	recApi.GET("/routine/rec/routineact", recApiController.queryRoutineActRecs, guard.UserGuard)

	// update routineact rec
	recApi.PUT("/routine/rec/routineact", recApiController.updateRoutineActRec, guard.UserGuard)

	return recApi
}

// newAuthApiController returns a new authApiController that implements every auth api features.
func newRecApiController(params RecApiControllerParams) recApis {
	return &recApiController{dbClient: params.DbClient, httpClient: params.HttpClient}
}

// authApiController defines what auth api has to have and implements authApis interface at service file.
type recApiController struct {
	dbClient   *ent.Client
	httpClient *http.Client
}

// authApis interface defines what auth api has to handle
type recApis interface {
	createWeeklyProgramRec(c echo.Context) error

	queryRoutineActRecs(c echo.Context) error

	updateRoutineActRec(c echo.Context) error
}
