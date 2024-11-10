CREATE TYPE property_type AS ENUM ('string', 'integer', 'float', 'boolean','object', 'image');
CREATE TABLE properties (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  name text NOT NULL,
  string_value text,
  integer_value integer,
  float_value double precision,
  boolean_value boolean,
  object_id uuid REFERENCES objects(id) ON DELETE CASCADE,
  image_path text,
  template_id uuid REFERENCES properties(id) ON DELETE SET NULL,
  created_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX properties_name_idx ON properties (name);
CREATE INDEX properties_string_idx ON properties (string_value);
CREATE INDEX properties_integer_idx ON properties (integer_value);
CREATE INDEX properties_float_idx ON properties (float_value);
CREATE INDEX properties_boolean_idx ON properties (boolean_value);
CREATE INDEX properties_object_id_idx ON properties (object_id);
CREATE INDEX properties_image_path_idx ON properties (image_path);
CREATE INDEX properties_template_id_idx ON properties (template_id);
