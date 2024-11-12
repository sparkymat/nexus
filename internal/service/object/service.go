package object

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sparkymat/nexus/internal/dbx"
)

func New(db *pgxpool.Pool, dbx *dbx.Queries) *Service {
	return &Service{
		db:  db,
		dbx: dbx,
	}
}

type Service struct {
	db  *pgxpool.Pool
	dbx *dbx.Queries
}
