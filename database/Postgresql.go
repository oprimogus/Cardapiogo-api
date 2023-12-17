package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"github.com/oprimogus/cardapiogo/config/logger"
)

var (
	instance *PostgresDatabase
	log      = logger.GetLogger("Postgres")
)

// PostgresDatabase struct
type PostgresDatabase struct {
	pool *pgxpool.Pool
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
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUsername, dbPassword, dbName)
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
	sourceURL := os.Getenv("MIGRATION_SOURCE_URL")
	dbName := os.Getenv("DB_NAME")
	log.Info("starting migration execution")

	driver, err := postgres.WithInstance(d.sqlDB, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("database: could not create migration driver: %w", err)
	}

	log.Infof("Executing migrations on path: %s", sourceURL)
	m, err := migrate.NewWithDatabaseInstance(
		"file://"+sourceURL,
		dbName, driver,
	)
	if m != nil {
		err = m.Up()
		if err != nil && err != migrate.ErrNoChange {
			return fmt.Errorf("database: error when executing database migration: %w", err)
		}
	}
	log.Info("finalizing migrations!")
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
