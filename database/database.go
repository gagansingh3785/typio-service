package database

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/gagansingh3785/typio-service/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	zlog "github.com/rs/zerolog/log"
)

type Database struct {
	DB *sqlx.DB
}

func NewDatabase(cfg *config.Config) (*Database, error) {
	db, err := sqlx.Connect("postgres", cfg.GetDBURL())
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.DB.MaxOpenConnections)
	db.SetMaxIdleConns(cfg.DB.MaxIdleConnections)
	db.SetConnMaxIdleTime(time.Duration(cfg.DB.ConnMaxIdleTimeInSeconds) * time.Second)
	db.SetConnMaxLifetime(time.Duration(cfg.DB.ConnMaxLifetimeInSeconds) * time.Second)

	return &Database{DB: db}, nil
}

func (d *Database) CreateMigrationFiles(name string, migrationsPath string) error {
	zlog.Info().Msgf("Creating migration files for %s", name)
	eventTimestamp := time.Now().Unix()

	// Create up migration file with timestamp
	upMigrationFilePath := filepath.Join(
		migrationsPath,
		fmt.Sprintf("%d_%s.up.sql", eventTimestamp, name),
	)
	upMigrationFile, err := os.Create(upMigrationFilePath)
	if err != nil {
		return err
	}
	defer upMigrationFile.Close()

	// Create down migration file with timestamp
	downMigrationFilePath := filepath.Join(
		migrationsPath,
		fmt.Sprintf("%d_%s.down.sql", eventTimestamp, name),
	)
	downMigrationFile, err := os.Create(downMigrationFilePath)
	if err != nil {
		return err
	}
	defer downMigrationFile.Close()

	zlog.Info().Msg("Migration files created successfully")

	return nil
}

func (d *Database) RunMigrations(migrationPath string) error {
	zlog.Info().Msg("Running migrations...")
	dbDriver, err := postgres.WithInstance(d.DB.DB, &postgres.Config{})
	if err != nil {
		return err
	}
	absMigrationPath, err := filepath.Abs(migrationPath)
	if err != nil {
		return err
	}
	// Convert to forward slashes for Windows compatibility
	sourceURL := "file://" + filepath.ToSlash(absMigrationPath)
	m, err := migrate.NewWithDatabaseInstance(sourceURL, "postgres", dbDriver)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil {
		return err
	}
	zlog.Info().Msg("Migrations applied successfully")
	return nil
}

func (d *Database) RollbackLastMigration(migrationPath string) error {
	zlog.Info().Msg("Rolling back the last migration...")
	dbDriver, err := postgres.WithInstance(d.DB.DB, &postgres.Config{})
	if err != nil {
		return err
	}

	absMigrationPath, err := filepath.Abs(migrationPath)
	if err != nil {
		return err
	}

	sourceURL := "file://" + filepath.ToSlash(absMigrationPath)
	m, err := migrate.NewWithDatabaseInstance(sourceURL, "postgres", dbDriver)
	if err != nil {
		return err
	}

	if err := m.Steps(-1); err != nil {
		return err
	}

	zlog.Info().Msg("Migration rolled back successfully")

	return nil
}
