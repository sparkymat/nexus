package presenter

import "github.com/sparkymat/nexus/internal/dbx"

type Property struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	PropertyType string   `json:"propertyType"`
	StringValue  *string  `json:"stringValue"`
	IntValue     *int32   `json:"intValue"`
	FloatValue   *float64 `json:"floatValue"`
	BoolValue    *bool    `json:"boolValue"`
	DateValue    *string  `json:"dateValue"`
	ObjectValue  *Object  `json:"object"`
}

func PropertyFromModel(m dbx.Property, obj dbx.Object) Property {
	p := Property{
		ID:           m.ID.String(),
		Name:         m.Name,
		PropertyType: string(m.PropertyType),
	}

	switch m.PropertyType {
	case dbx.PropertyTypeString:
		if m.StringValue.Valid {
			p.StringValue = &m.StringValue.String
		}
	case dbx.PropertyTypeInteger:
		if m.IntegerValue.Valid {
			p.IntValue = &m.IntegerValue.Int32
		}
	case dbx.PropertyTypeFloat:
		if m.FloatValue.Valid {
			p.FloatValue = &m.FloatValue.Float64
		}
	case dbx.PropertyTypeBoolean:
		if m.BooleanValue.Valid {
			p.BoolValue = &m.BooleanValue.Bool
		}
	case dbx.PropertyTypeImage:
		if m.ImagePath.Valid {
			p.StringValue = &m.ImagePath.String
		}
	case dbx.PropertyTypeDate:
		if m.DateValue.Valid {
			dateValue := m.DateValue.Time.Format("2006-01-02")
			p.DateValue = &dateValue
		}
	case dbx.PropertyTypeObject:
		if m.ObjectValueID.Valid {
			object := ObjectFromModel(obj, nil, nil)
			p.ObjectValue = &object
		}
	}

	return p
}
