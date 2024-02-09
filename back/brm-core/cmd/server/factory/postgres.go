package factory

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/spf13/viper"
	"os"
	"time"
)

func ConnectToPostgres(ctx context.Context) (*pgx.Conn, error) {
	coreRepoUrl := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		viper.GetString("postgres_core.username"),
		os.Getenv("POSTGRES_CORE_PASSWORD"),
		viper.GetString("postgres_core.host"),
		viper.GetInt("postgres_core.port"),
		viper.GetString("postgres_core.dbname"),
		viper.GetString("postgres_core.sslmode"))

	// 30 attempts to connect to postgres starting in docker container
	for i := 0; i < 30; i++ {
		conn, err := pgx.Connect(ctx, coreRepoUrl)
		if err != nil {
			time.Sleep(time.Second)
		} else {
			return conn, nil
		}
	}

	return nil, errors.New("unable to connect to postgres core repo")
}
