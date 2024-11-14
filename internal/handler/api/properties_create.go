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

type PropertiesCreateRequest struct {
	Name          string           `json:"name"`
	PropertyType  dbx.PropertyType `json:"propertyType"`
	StringValue   *string          `json:"stringValue"`
	IntegerValue  *int64           `json:"integerValue"`
	FloatValue    *float64         `json:"floatValue"`
	BooleanValue  *bool            `json:"booleanValue"`
	DateValue     *string          `json:"dateValue"`
	ObjectValueID *string          `json:"objectValueId"`
}

func PropertiesCreate(s internal.Services) echo.HandlerFunc {
	return wrapWithAuthForChild(func(c echo.Context, _ dbx.User, parentID uuid.UUID) error {
		fileHeader, err := c.FormFile("image_file")
		if err != nil {
			return renderError(c, http.StatusBadRequest, "missing file", err)
		}

		file, err := fileHeader.Open()
		if err != nil {
			return renderError(c, http.StatusBadRequest, "failed to open file", err)
		}
		defer file.Close()

		var request PropertiesCreateRequest
		if err = c.Bind(&request); err != nil {
			return renderError(c, http.StatusUnprocessableEntity, "failed to parse request", err)
		}

		property, object, err := s.Object.CreateProperty(
			c.Request().Context(),
			object.CreatePropertyOptions{
				Name:     request.Name,
				ObjectID: parentID,
			},
		)
		if err != nil {
			return renderError(c, http.StatusInternalServerError, "failed to create property", err)
		}

		renderedProperty := presenter.PropertyFromModel(property, object)

		return c.JSON(http.StatusCreated, renderedProperty)
	})
}
