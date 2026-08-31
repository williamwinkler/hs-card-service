package migrations

import (
	"database/sql"
	"embed"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/opentelemetry/tracing"
)

const defaultDatabaseURL = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"

//go:embed sql/*.sql
var migrationFiles embed.FS

type Database struct {
	Db *gorm.DB
}

func SetupDatabase() (*Database, error) {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = defaultDatabaseURL
	}

	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return &Database{}, fmt.Errorf("failed connecting to postgres: %w", err)
	}
	if err := db.Use(tracing.NewPlugin(
		tracing.WithoutQueryVariables(),
		tracing.WithoutServerAddress(),
		tracing.WithQueryFormatter(databaseOperation),
	)); err != nil {
		return &Database{}, fmt.Errorf("enable database telemetry: %w", err)
	}

	if err := pingDatabase(db); err != nil {
		return &Database{}, err
	}

	if err := applyMigrations(db); err != nil {
		return &Database{}, err
	}

	return &Database{Db: db}, nil
}

func databaseOperation(query string) string {
	fields := strings.Fields(query)
	if len(fields) == 0 {
		return ""
	}
	return fields[0]
}

func pingDatabase(db *gorm.DB) error {
	var result int
	if err := db.Raw("SELECT 1").Scan(&result).Error; err != nil {
		return fmt.Errorf("postgres ping failed: %w", err)
	}
	return nil
}

func applyMigrations(db *gorm.DB) error {
	if err := db.Exec(`
		CREATE TABLE IF NOT EXISTS schema_migrations (
			version TEXT PRIMARY KEY,
			applied_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
		)
	`).Error; err != nil {
		return fmt.Errorf("failed creating schema_migrations: %w", err)
	}

	entries, err := migrationFiles.ReadDir("sql")
	if err != nil {
		return fmt.Errorf("failed reading migration directory: %w", err)
	}

	var names []string
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".sql") {
			continue
		}
		names = append(names, entry.Name())
	}
	sort.Strings(names)

	for _, name := range names {
		var count int64
		if err := db.Raw("SELECT COUNT(*) FROM schema_migrations WHERE version = ?", name).Scan(&count).Error; err != nil {
			return fmt.Errorf("failed checking migration %s: %w", name, err)
		}
		if count > 0 {
			continue
		}

		content, err := migrationFiles.ReadFile("sql/" + name)
		if err != nil {
			return fmt.Errorf("failed reading migration %s: %w", name, err)
		}

		if err := executeSQLStatements(db, string(content)); err != nil {
			return fmt.Errorf("failed running migration %s: %w", name, err)
		}

		if err := db.Exec("INSERT INTO schema_migrations (version, applied_at) VALUES (?, ?)", name, time.Now().UTC()).Error; err != nil {
			return fmt.Errorf("failed storing migration %s: %w", name, err)
		}
		log.Printf("Applied migration %s", name)
	}

	return nil
}

func executeSQLStatements(db *gorm.DB, migrationSQL string) error {
	statements := strings.Split(migrationSQL, ";")
	for _, statement := range statements {
		trimmed := strings.TrimSpace(statement)
		if trimmed == "" {
			continue
		}
		if err := db.Exec(trimmed).Error; err != nil {
			if err == sql.ErrNoRows {
				continue
			}
			return err
		}
	}
	return nil
}
