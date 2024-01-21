package repo

import "github.com/jackc/pgx/v5"

type coreRepoImpl struct {
	pgx.Conn
}
