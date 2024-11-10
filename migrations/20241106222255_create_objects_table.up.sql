CREATE TABLE objects (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  name text NOT NULL,
  alternate_names_csv text NOT NULL DEFAULT '',
  is_template boolean NOT NULL DEFAULT FALSE,
  template_id uuid REFERENCES objects(id) ON DELETE SET NULL,
  created_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX objects_name_idx ON objects (name);
CREATE INDEX objects_alternate_names_idx ON objects (alternate_names_csv);
