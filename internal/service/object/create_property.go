package object

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/sparkymat/nexus/internal/dbx"
)

type CreatePropertyOptions struct {
	ObjectID      uuid.UUID
	Name          string
	PropertyType  dbx.PropertyType
	StringValue   *string
	IntegerValue  *int64
	FloatValue    *float64
	BooleanValue  *bool
	DateValue     *time.Time
	ObjectValueID *uuid.UUID
}

func (s *Service) CreateProperty(ctx context.Context, options CreatePropertyOptions) (dbx.Property, *dbx.Object, error) {
	if _, err := s.dbx.FetchObjectByID(ctx, options.ObjectID); err != nil {
		return dbx.Property{}, nil, fmt.Errorf("object not found: %w", err)
	}

	var objectValuePtr *dbx.Object

	if options.ObjectValueID != nil {
		objectValue, err := s.dbx.FetchObjectByID(ctx, *options.ObjectValueID)
		if err != nil {
			return dbx.Property{}, nil, fmt.Errorf("object value not found: %w", err)
		}

		objectValuePtr = &objectValue
	}

	createParams := dbx.CreatePropertyParams{
		ObjectID:     options.ObjectID,
		PropertyType: options.PropertyType,
	}

	switch options.PropertyType {
	case dbx.PropertyTypeString:
		if options.StringValue != nil {
			createParams.StringValue = pgtype.Text{String: *options.StringValue, Valid: true}
		}
	case dbx.PropertyTypeInteger:
		if options.IntegerValue != nil {
			createParams.IntegerValue = pgtype.Int8{Int64: *options.IntegerValue, Valid: true}
		}
	case dbx.PropertyTypeBoolean:
		if options.BooleanValue != nil {
			createParams.BooleanValue = pgtype.Bool{Bool: *options.BooleanValue, Valid: true}
		}
	case dbx.PropertyTypeDate:
		if options.DateValue != nil {
			createParams.DateValue = pgtype.Date{Time: *options.DateValue, Valid: true}
		}
	case dbx.PropertyTypeFloat:
		if options.FloatValue != nil {
			createParams.FloatValue = pgtype.Float8{Float64: *options.FloatValue, Valid: true}
		}
	case dbx.PropertyTypeImage:
		if options.StringValue != nil {
			createParams.ImagePath = pgtype.Text{String: *options.StringValue, Valid: true}
		}
	case dbx.PropertyTypeObject:
		if options.ObjectValueID != nil {
			createParams.ObjectValueID = uuid.NullUUID{UUID: *options.ObjectValueID, Valid: true}
		}
	}

	prop, err := s.dbx.CreateProperty(ctx, createParams)
	if err != nil {
		return dbx.Property{}, nil, fmt.Errorf("failed to create property: %w", err)
	}

	return prop, objectValuePtr, nil
}
