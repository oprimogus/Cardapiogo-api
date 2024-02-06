package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	err := PopulateLocalDatabase()
	if err != nil {
		panic(err)
	}
}

func PopulateLocalDatabase() error {
	db, err := getSQLDBConnection(createStringConn())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic("Fail connect in local database")
	}

	// Verifica se os dados de teste já existem
	if !checkTestDataExists(db) {
		insertTestData(db)
	} else {
		fmt.Println("Os dados de teste já existem no banco de dados.")
	}
	return nil
}

func createStringConn() string {
	dbHost := "localhost"
	dbPort := "5435"
	dbUsername := "cardapiogo"
	dbPassword := "cardapiogo"
	dbName := "postgres"
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUsername, dbPassword, dbName)
}

func getSQLDBConnection(connStr string) (*sql.DB, error) {
	sqlDB, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("database: could not open sql connection: %w", err)
	}
	return sqlDB, nil
}

func checkTestDataExists(db *sql.DB) bool {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM users LIMIT 1);`

	err := db.QueryRow(query).Scan(&exists)
	if err != nil {
		log.Fatalf("Erro ao verificar a existência de dados de teste: %v", err)
	}

	return exists
}

func insertTestData(db *sql.DB) {

	// Caminho para o arquivo .sql com os dados de teste
	sqlFilePath := "internal/infra/database/sql/mocks/mockUserData.sql"

	// Lê e executa o SQL do arquivo
	err := executeSQLFile(db, sqlFilePath)
	if err != nil {
		log.Fatalf("Erro ao executar arquivo SQL: %v", err)
	}

	fmt.Println("Arquivo SQL executado e dados de teste inseridos com sucesso.")
}

func executeSQLFile(db *sql.DB, filePath string) error {
	// Lê o conteúdo do arquivo SQL
	query, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("erro ao ler o arquivo SQL: %w", err)
	}

	// Executa o SQL no banco de dados
	_, err = db.Exec(string(query))
	if err != nil {
		return fmt.Errorf("erro ao executar o SQL: %w", err)
	}

	return nil
}
