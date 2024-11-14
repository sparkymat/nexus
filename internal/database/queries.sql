-- name: CreateUser :one
INSERT INTO users (
  name, email, encrypted_password
) VALUES (
  @name::text, @email::text, @encrypted_password::text
) RETURNING *;

-- name: FetchUserByEmail :one
SELECT * FROM users
  WHERE email = @email::text LIMIT 1;

-- name: CreateObject :one
INSERT INTO objects (
  name, is_template, template_id
) VALUES (
@name::text, @is_template::boolean, sqlc.narg('template_id')
) RETURNING *;

-- name: CreateProperty :one
INSERT INTO properties (
  object_id, name, property_type, string_value, integer_value, float_value, date_value, boolean_value, object_value_id, image_path, template_id
) VALUES (
  @object_id::uuid, @name::text, @property_type::property_type, sqlc.narg('string_value'), sqlc.narg('integer_value'), sqlc.narg('float_value'), sqlc.narg('date_value'), sqlc.narg('boolean_value'), sqlc.narg('object_value_id'), sqlc.narg('image_path'), sqlc.narg('template_id')
) RETURNING *;

-- name: FetchPropertiesByObjectID :many
SELECT * FROM properties
  WHERE object_id = @object_id::uuid;

-- name: FetchObjectByID :one
SELECT * FROM objects
  WHERE id = @id::uuid
  LIMIT 1;
