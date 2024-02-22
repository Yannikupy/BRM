package repo

import (
	"github.com/jackc/pgx/v5"
)

type adRepoImpl struct {
	pgx.Conn
}
