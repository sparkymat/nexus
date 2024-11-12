package presenter

import (
	"github.com/google/uuid"
	"github.com/sparkymat/nexus/internal/dbx"
)

type Object struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	IsTemplate bool   `json:"isTemplate"`

	Properties []Property `json:"properties"`
}

func ObjectFromModel(m dbx.Object, p []dbx.Property, pom map[uuid.UUID]dbx.Object) Object {
	o := Object{
		ID:         m.ID.String(),
		Name:       m.Name,
		IsTemplate: m.IsTemplate,
	}

	props := []Property{}

	for _, prop := range p {
		objVal := pom[prop.ID]
		pprop := PropertyFromModel(prop, objVal)

		props = append(props, pprop)
	}

	o.Properties = props

	return o
}