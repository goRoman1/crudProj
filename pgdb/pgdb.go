package pgdb

import (
	"context"
	"crudProj/config"
	"fmt"
	"github.com/jackc/pgx/v4"
	_ "github.com/lib/pq"
	"os"
)

func Dial(cfg *config.Config)*pgx.Conn {
	var err error
	if cfg.DriverName == "" {
		return nil
	}

	pgUrl := fmt.Sprintf("%s://%s:%s@%s:%s/%s", cfg.DriverName, cfg.DbUser,cfg.DbPassword,cfg.DbHost, cfg.DbPort, cfg.DbName)
//	urlExample := "postgres://user:mypassword@localhost:5432/postgres"
	conn, err := pgx.Connect(context.Background(), pgUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Successfully connected!")

	return conn
}
