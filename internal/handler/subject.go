package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Subject struct {
}

type SubjectCreateRequest struct {
	Name string `json:"name"`
}

func (s *Subject) Create(c echo.Context) error {
	var subjectCreateRequest SubjectCreateRequest

	if err := c.Bind(&subjectCreateRequest); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

func (s *Subject) Destroy(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}

func (s *Subject) Index(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}

func (s *Subject) Show(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}

func (s *Subject) Update(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}
