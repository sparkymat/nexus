// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: queries.sql

package dbx

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createObject = `-- name: CreateObject :one
INSERT INTO objects (
  name, is_template, template_id
) VALUES (
$1::text, $2::boolean, $3
) RETURNING id, name, alternate_names_csv, is_template, template_id, created_at, updated_at
`

type CreateObjectParams struct {
	Name       string
	IsTemplate bool
	TemplateID uuid.NullUUID
}

func (q *Queries) CreateObject(ctx context.Context, arg CreateObjectParams) (Object, error) {
	row := q.db.QueryRow(ctx, createObject, arg.Name, arg.IsTemplate, arg.TemplateID)
	var i Object
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.AlternateNamesCsv,
		&i.IsTemplate,
		&i.TemplateID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createProperty = `-- name: CreateProperty :one
INSERT INTO properties (
  object_id, name, property_type, string_value, integer_value, float_value, date_value, boolean_value, object_value_id, image_path, template_id
) VALUES (
  $1::uuid, $2::text, $3::property_type, $4, $5, $6, $7, $8, $9, $10, $11
) RETURNING id, name, object_id, template_id, property_type, string_value, integer_value, float_value, object_value_id, date_value, boolean_value, image_path, created_at, updated_at
`

type CreatePropertyParams struct {
	ObjectID      uuid.UUID
	Name          string
	PropertyType  PropertyType
	StringValue   pgtype.Text
	IntegerValue  pgtype.Int8
	FloatValue    pgtype.Float8
	DateValue     pgtype.Date
	BooleanValue  pgtype.Bool
	ObjectValueID uuid.NullUUID
	ImagePath     pgtype.Text
	TemplateID    uuid.NullUUID
}

func (q *Queries) CreateProperty(ctx context.Context, arg CreatePropertyParams) (Property, error) {
	row := q.db.QueryRow(ctx, createProperty,
		arg.ObjectID,
		arg.Name,
		arg.PropertyType,
		arg.StringValue,
		arg.IntegerValue,
		arg.FloatValue,
		arg.DateValue,
		arg.BooleanValue,
		arg.ObjectValueID,
		arg.ImagePath,
		arg.TemplateID,
	)
	var i Property
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.ObjectID,
		&i.TemplateID,
		&i.PropertyType,
		&i.StringValue,
		&i.IntegerValue,
		&i.FloatValue,
		&i.ObjectValueID,
		&i.DateValue,
		&i.BooleanValue,
		&i.ImagePath,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  name, email, encrypted_password
) VALUES (
  $1::text, $2::text, $3::text
) RETURNING id, name, email, encrypted_password, created_at, updated_at
`

type CreateUserParams struct {
	Name              string
	Email             string
	EncryptedPassword string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser, arg.Name, arg.Email, arg.EncryptedPassword)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.EncryptedPassword,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const fetchObjectByID = `-- name: FetchObjectByID :one
SELECT id, name, alternate_names_csv, is_template, template_id, created_at, updated_at FROM objects
  WHERE id = $1::uuid
  LIMIT 1
`

func (q *Queries) FetchObjectByID(ctx context.Context, id uuid.UUID) (Object, error) {
	row := q.db.QueryRow(ctx, fetchObjectByID, id)
	var i Object
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.AlternateNamesCsv,
		&i.IsTemplate,
		&i.TemplateID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const fetchPropertiesByObjectID = `-- name: FetchPropertiesByObjectID :many
SELECT id, name, object_id, template_id, property_type, string_value, integer_value, float_value, object_value_id, date_value, boolean_value, image_path, created_at, updated_at FROM properties
  WHERE object_id = $1::uuid
`

func (q *Queries) FetchPropertiesByObjectID(ctx context.Context, objectID uuid.UUID) ([]Property, error) {
	rows, err := q.db.Query(ctx, fetchPropertiesByObjectID, objectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Property
	for rows.Next() {
		var i Property
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.ObjectID,
			&i.TemplateID,
			&i.PropertyType,
			&i.StringValue,
			&i.IntegerValue,
			&i.FloatValue,
			&i.ObjectValueID,
			&i.DateValue,
			&i.BooleanValue,
			&i.ImagePath,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const fetchUserByEmail = `-- name: FetchUserByEmail :one
SELECT id, name, email, encrypted_password, created_at, updated_at FROM users
  WHERE email = $1::text LIMIT 1
`

func (q *Queries) FetchUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, fetchUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.EncryptedPassword,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
