package api

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sparkymat/nexus/internal"
	"github.com/sparkymat/nexus/internal/dbx"
	"github.com/sparkymat/nexus/internal/handler/api/presenter"
)

func ObjectsShow(s internal.Services) echo.HandlerFunc {
	return wrapWithAuthForMember(func(c echo.Context, _ dbx.User, id uuid.UUID) error {
		object, properties, pom, err := s.Object.FetchObject(c.Request().Context(), id)
		if err != nil {
			return renderError(c, http.StatusInternalServerError, "Failed to fetch object", err)
		}

		presentedObject := presenter.ObjectFromModel(object, properties, pom)

		return c.JSON(http.StatusOK, presentedObject)
	})
}
