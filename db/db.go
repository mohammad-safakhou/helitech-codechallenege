package db

import (
	"codechallenge/config"
	"codechallenge/utils"
	"context"
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"time"
)

type DB struct {
	migrationFS embed.FS
	Connection  *sql.DB
	dbMigration *migrate.Migrate
}

func (d *DB) Migrate() error {
	source, err := iofs.New(d.migrationFS, config.AppConfig.Databases.Postgres.MigrationPath)
	if err != nil {
		return err
	}

	connectionString := utils.PostgresURI(
		config.AppConfig.Databases.Postgres.Host,
		config.AppConfig.Databases.Postgres.Port,
		config.AppConfig.Databases.Postgres.User,
		config.AppConfig.Databases.Postgres.Pass,
		config.AppConfig.Databases.Postgres.DatabaseName,
		config.AppConfig.Databases.Postgres.SslMode,
	)
	m, err := migrate.NewWithSourceInstance("iofs", source, connectionString)
	if err != nil {
		return err
	}
	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}
	return nil
}

func New(migrationFS embed.FS) *DB {
	codeDB, err := utils.PostgresConnection(
		config.AppConfig.Databases.Postgres.Host,
		config.AppConfig.Databases.Postgres.Port,
		config.AppConfig.Databases.Postgres.User,
		config.AppConfig.Databases.Postgres.Pass,
		config.AppConfig.Databases.Postgres.DatabaseName,
		config.AppConfig.Databases.Postgres.SslMode,
		config.AppConfig.Databases.Postgres.MaxOpenConns,
		config.AppConfig.Databases.Postgres.MaxIdleConns,
		config.AppConfig.Databases.Postgres.Timeout)
	if err != nil {
		panic(fmt.Sprintf("could not get DB connection: %s", err.Error()))
	}

	ctx, cancelCtx := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelCtx()
	var temp int
	row := codeDB.QueryRowContext(
		ctx,
		"SELECT 1",
	)
	err = row.Scan(&temp)
	if err != nil {
		panic(fmt.Sprintf("could not get DB: %s", err.Error()))
	}

	return &DB{
		migrationFS: migrationFS,
		Connection:  codeDB,
	}
}
