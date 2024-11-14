package api

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sparkymat/nexus/internal"
	"github.com/sparkymat/nexus/internal/dbx"
	"github.com/sparkymat/nexus/internal/handler/api/presenter"
	"github.com/sparkymat/nexus/internal/service/object"
)

type ObjectsCreateRequest struct {
	Name       string  `json:"name"`
	IsTemplate bool    `json:"isTemplate"`
	TemplateID *string `json:"templateId"`
}

func ObjectsCreate(s internal.Services) echo.HandlerFunc {
	return wrapWithAuth(func(c echo.Context, _ dbx.User) error {
		var request ObjectsCreateRequest
		if err := c.Bind(&request); err != nil {
			return renderError(c, http.StatusUnprocessableEntity, "failed to parse request", err)
		}

		var templateID *uuid.UUID

		if request.TemplateID != nil {
			templateIDValue, err := uuid.Parse(*request.TemplateID)
			if err != nil {
				return renderError(c, http.StatusUnprocessableEntity, "failed to parse template id", err)
			}
			templateID = &templateIDValue
		}

		object, properties, err := s.Object.CreateObject(
			c.Request().Context(),
			object.CreateObjectOptions{
				Name:       request.Name,
				IsTemplate: request.IsTemplate,
				TemplateID: templateID,
			},
		)
		if err != nil {
			return renderError(c, http.StatusInternalServerError, "failed to create object", err)
		}

		presentedObject := presenter.ObjectFromModel(object, properties, nil)

		return c.JSON(http.StatusCreated, presentedObject)
	})
}
