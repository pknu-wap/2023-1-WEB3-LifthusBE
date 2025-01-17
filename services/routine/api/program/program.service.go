package program

import (
	"net/http"
	"routine/common/db"
	"routine/common/dto"
	"routine/ent"
	"strconv"

	"github.com/labstack/echo/v4"
)

// queryProgramBySlug godoc
// @Router       /program/{slug} [get]
// @Param slug path string true "program slug"
// @Summary      gets Program slug from path and returns corresponding Program
// @Tags         program
// @Success      200 "returns program"
// @Failure      400 "invalid request"
// @Failure      404 "program not found"
// @Failure      500 "failed to query program"
func (pc programApiController) queryProgramBySlug(c echo.Context) error {
	slug := c.Param("slug")
	program, err := db.QueryProgramBySlug(pc.dbClient, c.Request().Context(), slug)
	if ent.IsNotFound(err) {
		return c.String(http.StatusNotFound, err.Error())
	} else if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, program)
}

// queryProgramsByName godoc
// @Router       /program [get]
// @Param title query string true "program title"
// @Param skip query int false "skip"
// @Summary      gets Program name from query-string and returns corresponding Programs
// @Tags         program
// @Success      200 "returns programs"
// @Failure      400 "invalid request"
// @Failure      500 "failed to query programs"
func (pc programApiController) queryProgramsByName(c echo.Context) error {
	programTitle := c.QueryParam("title")
	skipStr := c.QueryParam("skip")
	// convert skip to int if it exists
	skip, err := strconv.Atoi(skipStr)
	if skipStr == "" {
		skip = 0
	} else {
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
	}
	programs, err := db.QueryProgramsByName(pc.dbClient, c.Request().Context(), programTitle, skip)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, programs)
}

// createWeeklyProgram godoc
// @Router       /program/weekly [post]
// @Param createWeeklyProgramDto body dto.CreateWeeklyProgramDto true "createWeeklyProgram DTO"
// @Summary      gets CreateWeeklyProgramDto from body and returns created program's ID
// @Tags         program
// @Success      201 "returns created program's ID"
// @Failure      400 "invalid body"
// @Failure      401 "unauthorized"
// @Failure 	 403 "forbidden"
// @Failure      500 "failed to create program"
func (pc programApiController) createWeeklyProgram(c echo.Context) error {
	createProgramDto := new(dto.CreateWeeklyProgramDto)
	if err := c.Bind(createProgramDto); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	aid := createProgramDto.Author
	uid := c.Get("uid").(uint64)
	if aid != uid {
		return c.String(http.StatusForbidden, "you are not allowed to create program for others")
	}

	pid, err := db.CreateWeeklyProgram(pc.dbClient, c.Request().Context(), createProgramDto)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, pid)
}
