package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"

	"github.com/oprimogus/cardapiogo/internal/config"
	logger "github.com/oprimogus/cardapiogo/pkg/log"
)

var (
	instance *PostgresDatabase
	log      = logger.NewLogger("Postgres")
)

// PostgresDatabase struct
type PostgresDatabase struct {
	pool  *pgxpool.Pool
	sqlDB *sql.DB
}

// GetInstance of PostgresDatabase
func GetInstance() *PostgresDatabase {
	if instance == nil {
		instance = createInstance()
	}
	return instance
}

func createInstance() *PostgresDatabase {
	database := &PostgresDatabase{}
	strConnection := database.createStringConn()

	// Open connection with pgx
	var err error
	database.pool, err = database.getPgxConnection(strConnection)
	if err != nil {
		log.Error(err)
		panic(err)
	}

	// Open connection with database/sql for migration
	database.sqlDB, err = database.getSQLDBConnection(strConnection)
	if err != nil {
		log.Error(err)
		panic(err)
	}

	err = database.migrate()
	if err != nil {
		log.Error(err)
		panic(err)
	}
	return database
}

func (d PostgresDatabase) createStringConn() string {
	config := config.GetInstance()
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable search_path=public",
		config.Database.Host(),
		config.Database.Port(),
		config.Database.User(),
		config.Database.Password(),
		config.Database.Name(),
	)
}

func (d PostgresDatabase) getPgxConnection(connStr string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		return nil, fmt.Errorf("database: could not open pgx connection: %w", err)
	}
	return pool, nil
}

func (d PostgresDatabase) getSQLDBConnection(connStr string) (*sql.DB, error) {
	sqlDB, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("database: could not open sql connection: %w", err)
	}
	return sqlDB, nil
}

func (d PostgresDatabase) migrate() error {
	err := setWorkingDirToProjectRoot()
	if err != nil {
		return err
	}
	sourceURL := "file://internal/database/migrations"
	config := config.GetInstance()
	log.Info("starting migration execution")
	driver, err := postgres.WithInstance(d.sqlDB, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("database: could not create migration driver: %w", err)
	}
	migrator, err := migrate.NewWithDatabaseInstance(sourceURL, config.Database.Name(), driver)
	if err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("database: Could not create migrator: %w", err)
	}
	err = migrator.Up()
	if err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("database: Could not apply migrations: %w", err)
	}
	if err == migrate.ErrNoChange {
		log.Info("No migrations to run.")
	} else {
		log.Info("Migrations applied successfully.")
	}
	return nil
}

// GetDB return a pgxpool.Pool pointer
func (d PostgresDatabase) GetDB() *pgxpool.Pool {
	return d.pool
}

// Close connection with database
func (d PostgresDatabase) Close() {
	d.pool.Close()
	if d.sqlDB != nil {
		d.sqlDB.Close()
	}
}

func setWorkingDirToProjectRoot() error {
	currentDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("could not get current directory: %w", err)
	}

	// Caminha para o diretório pai até encontrar o arquivo go.mod (indicando a raiz do projeto)
	for {
		if _, err := os.Stat(filepath.Join(currentDir, "go.mod")); err == nil {
			break
		}

		parentDir := filepath.Dir(currentDir)
		if parentDir == currentDir {
			return fmt.Errorf("could not find project root (go.mod not found)")
		}

		currentDir = parentDir
	}

	// Define o diretório de trabalho como a raiz do projeto
	if err := os.Chdir(currentDir); err != nil {
		return fmt.Errorf("could not change to project root: %w", err)
	}

	return nil
}
