package repo

import "github.com/jackc/pgx/v5"

type leadRepoImpl struct {
	pgx.Conn
}
