package database

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	// Necessary
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/oprimogus/cardapiogo/config/logger"
	"os"
)

var (
	instance *PostgresDatabase
	log      = logger.GetLogger("Postgres")
)

// PostgresDatabase struct
type PostgresDatabase struct {
	db *sql.DB
}

// GetInstance of PostgresDatabase
func GetInstance() *PostgresDatabase {
	if instance == nil {
		instance = createInstance()
	}
	return instance
}

// func NewInstance() *PostgresDatabase {
// 	return createInstance()
// }

func createInstance() *PostgresDatabase {
	var err error
	database := &PostgresDatabase{}
	strConnection := database.createStringConn()
	database.db, err = database.getConnection(strConnection)
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

// Many return many of rows
func (d PostgresDatabase) Many(ctx context.Context, query string, params ...interface{}) (*sql.Rows, error) {
	return d.db.QueryContext(ctx, query, params...)
}

// One return one row
func (d PostgresDatabase) One(ctx context.Context, query string, params ...interface{}) *sql.Row {
	return d.db.QueryRowContext(ctx, query, params...)
}

// Exec SQL function
func (d PostgresDatabase) Exec(ctx context.Context, query string, params ...interface{}) (sql.Result, error) {
	return d.db.ExecContext(ctx, query, params...)
}

func (d PostgresDatabase) createStringConn() string {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUsername, dbPassword, dbName)
}

func (d PostgresDatabase) getConnection(connStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("database: could not open connection: %w", err)
	}
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("database: could not establish connection: %w", err)
	}
	return db, nil
}

func (d PostgresDatabase) migrate() error {
	sourceURL := os.Getenv("MIGRATION_SOURCE_URL")
	dbName := os.Getenv("DB_NAME")
	log.Info("starting migration execution")
	driver, err := postgres.WithInstance(d.db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("database: could not create migration connection: %w", err)
	}
	log.Infof("Executing migrations on path: %s", sourceURL)	
	m, err := migrate.NewWithDatabaseInstance(
		"file://"+sourceURL,
		dbName, driver,
	)
	if m != nil {
		err = m.Up()
		if err != nil && err.Error() != "no change" {
			return fmt.Errorf("database: error when executing database migration: %w", err)
		}
	}
	log.Info("finalizing migrations!")
	return nil
}

// GetDB return a sql.DB pointer
func (d PostgresDatabase) GetDB() *sql.DB {
	return d.db
}

// Close connection with database
func (d PostgresDatabase) Close() {
	err := d.db.Close()
	if err != nil {
		log.Error(err)
	}
}
