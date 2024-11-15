package object

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/sparkymat/nexus/internal/dbx"
)

func (s *Service) FetchObject(ctx context.Context, id uuid.UUID) (dbx.Object, []dbx.Property, map[uuid.UUID]dbx.Object, error) {
	object, err := s.dbx.FetchObjectByID(ctx, id)
	if err != nil {
		return dbx.Object{}, nil, nil, fmt.Errorf("failed to fetch object: %w", err)
	}

	properties, err := s.dbx.FetchPropertiesByObjectID(ctx, id)
	if err != nil {
		return dbx.Object{}, nil, nil, fmt.Errorf("failed to fetch properties: %w", err)
	}

	propertyObjectIDs := lo.Map(properties, func(p dbx.Property, _ int) uuid.NullUUID { return p.ObjectValueID })
	notNullPropertyObjectIDs := lo.Filter(propertyObjectIDs, func(p uuid.NullUUID, _ int) bool { return p.Valid })
	objectIDs := lo.Map(notNullPropertyObjectIDs, func(c uuid.NullUUID, _ int) uuid.UUID { return c.UUID })

	valueObjects, err := s.dbx.FetchObjectsByID(ctx, objectIDs)
	if err != nil {
		return dbx.Object{}, nil, nil, fmt.Errorf("failed to fetch value objects: %w", err)
	}

	pom := lo.SliceToMap(valueObjects, func(o dbx.Object) (uuid.UUID, dbx.Object) { return o.ID, o })

	return object, properties, pom, nil
}
