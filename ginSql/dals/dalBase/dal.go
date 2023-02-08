package dalbase

import (
	"time"

	"github.com/jackc/pgx"
)

var connectionPool *pgx.ConnPool

func connectToDB() *pgx.ConnPool {
	// This database connection are from a sample database online
	// using this way to be easier for others to test this project
	// since it is a demo project
	conf := pgx.ConnConfig{
		Host:     "psql-mock-database-cloud.postgres.database.azure.com",
		User:     "miqcymvpmzslcwiixonyohhc@psql-mock-database-cloud",
		Password: "xejlvutptlsroaprbwkcftkx",
		Port:     5432,
		Database: "tasks1675287007220blncgcabxwzlhpws",
	}

	confPool := pgx.ConnPoolConfig{
		ConnConfig:     conf,
		MaxConnections: 5,                //Default
		AcquireTimeout: time.Second * 10, // 10 seconds timeout
	}

	pool, err := pgx.NewConnPool(confPool)

	if err != nil {
		return nil
	}

	return pool
}

func GetDB() *pgx.ConnPool {
	if connectionPool == nil {
		connectionPool = connectToDB()
	}

	return connectionPool
}
