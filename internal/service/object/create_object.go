package object

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sparkymat/nexus/internal/dbx"
)

type CreateObjectOptions struct {
	Name       string
	IsTemplate bool
	TemplateID *uuid.UUID
}

func (s *Service) CreateObject(ctx context.Context, options CreateObjectOptions) (dbx.Object, []dbx.Property, error) {
	templateProperties := []dbx.Property{}

	if options.TemplateID != nil {
		var err error

		templateProperties, err = s.dbx.FetchPropertiesByObjectID(ctx, *options.TemplateID)
		if err != nil {
			return dbx.Object{}, nil, fmt.Errorf("failed to fetch properties for template object: %w", err)
		}
	}

	tx, err := s.db.Begin(ctx)
	if err != nil {
		return dbx.Object{}, nil, fmt.Errorf("failed to start transaction: %w", err)
	}

	defer tx.Rollback(ctx)

	qtx := s.dbx.WithTx(tx)

	createParams := dbx.CreateObjectParams{
		Name:       options.Name,
		IsTemplate: options.IsTemplate,
	}

	if options.TemplateID != nil {
		createParams.TemplateID = uuid.NullUUID{Valid: true, UUID: *options.TemplateID}
	}

	obj, err := qtx.CreateObject(ctx, createParams)
	if err != nil {
		return dbx.Object{}, nil, fmt.Errorf("failed to create object: %w", err)
	}

	properties := []dbx.Property{}

	for _, templateProp := range templateProperties {
		prop, err := qtx.CreateProperty(ctx, dbx.CreatePropertyParams{
			ObjectID:      obj.ID,
			Name:          templateProp.Name,
			PropertyType:  templateProp.PropertyType,
			StringValue:   templateProp.StringValue,
			IntegerValue:  templateProp.IntegerValue,
			FloatValue:    templateProp.FloatValue,
			BooleanValue:  templateProp.BooleanValue,
			DateValue:     templateProp.DateValue,
			ObjectValueID: templateProp.ObjectValueID,
			ImagePath:     templateProp.ImagePath,
			TemplateID:    uuid.NullUUID{UUID: templateProp.ID, Valid: true},
		})
		if err != nil {
			return dbx.Object{}, nil, fmt.Errorf("failed to create property '%s': %w", templateProp.Name, err)
		}

		properties = append(properties, prop)
	}

	if err := tx.Commit(ctx); err != nil {
		return dbx.Object{}, nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return obj, properties, nil
}
