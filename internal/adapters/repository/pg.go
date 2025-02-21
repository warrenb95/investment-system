package repository

import (
	"embed"
	"fmt"
	"net/http"

	"github.com/go-pg/migrations/v8"
	"github.com/go-pg/pg/v10"
	_ "github.com/lib/pq" // Required for PostgreSQL
	"github.com/sirupsen/logrus"
)

//go:embed migrations/*.sql
var pgMigrations embed.FS

// PostgresRepository struct for database interactions
type PostgresRepository struct {
	db     *pg.DB
	logger *logrus.Logger
}

// NewPostgresRepository initializes a new repository
func NewPostgresRepository(db *pg.DB, logger *logrus.Logger) (*PostgresRepository, error) {
	p := &PostgresRepository{db: db, logger: logger}
	return p, p.runMigrations()
}

func (r *PostgresRepository) runMigrations() error {
	// run migrations
	collection := migrations.NewCollection()
	err := collection.DiscoverSQLMigrationsFromFilesystem(http.FS(pgMigrations), "migrations")
	if err != nil {
		r.logger.WithError(err).Error("migration discovery")
		return err
	}

	// start the migrations
	_, _, err = collection.Run(r.db, "init")
	if err != nil {
		r.logger.WithError(err).Error("Starting PostgreSQL migrations")
		return fmt.Errorf("starting PostgreSQL migrations: %w", err)
	}

	oldVersion, newVersion, err := collection.Run(r.db, "up")
	if err != nil {
		r.logger.WithError(err).Error("Running PostgreSQL migrations")
		return fmt.Errorf("running PostgreSQL migrations: %w", err)
	}

	if newVersion != oldVersion {
		r.logger.WithFields(logrus.Fields{
			"old_version": oldVersion, "new_version": newVersion,
		}).Info("new database migration")
	} else {
		r.logger.WithField("old_version", oldVersion).Info("migration version")
	}

	return err
}
