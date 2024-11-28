package integration_tests

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

var testDB *sql.DB

func TestMain(m *testing.M) {
	log.Println("Starting test setup...")

	if err := setupTestDB(); err != nil {
		log.Fatalf("Failed to setup test database: %v", err)
	}

	code := m.Run()

	if testDB != nil {
		cleanupDatabase(testDB)
		testDB.Close()
	}

	os.Exit(code)
}

func getProjectRoot() (string, error) {
	// Start with the current working directory
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// Walk up the directory tree until we find go.mod
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("could not find project root (no go.mod found)")
		}
		dir = parent
	}
}

func setupTestDB() error {
	projectRoot, err := getProjectRoot()
	if err != nil {
		return fmt.Errorf("could not get project root: %v", err)
	}

	err = godotenv.Load(filepath.Join(projectRoot, ".env"))
	if err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
	}

	// Get database connection string
	testDsn := os.Getenv("TEST_DSN")
	if testDsn == "" {
		return fmt.Errorf("TEST_DSN environment variable is not set")
	}

	log.Printf("Attempting to connect to database...")

	testDB, err = sql.Open("postgres", testDsn)
	if err != nil {
		return fmt.Errorf("cannot connect to database: %v", err)
	}

	if err = testDB.Ping(); err != nil {
		return fmt.Errorf("could not ping database: %v", err)
	}

	log.Printf("Successfully connected to database")

	if err = runMigrations(testDB); err != nil {
		return fmt.Errorf("failed to run migrations: %v", err)
	}

	if err = verifyTables(testDB); err != nil {
		return fmt.Errorf("failed to verify tables: %v", err)
	}

	return nil
}

func runMigrations(db *sql.DB) error {
	migrationsDir := "test_migrations"
	log.Printf("Running migrations from: %s", migrationsDir)

	files, err := os.ReadDir(migrationsDir)
	if err != nil {
		return fmt.Errorf("could not read migrations directory: %v", err)
	}

	var migrationFiles []string
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".sql") {
			migrationFiles = append(migrationFiles, file.Name())
		}
	}
	sort.Strings(migrationFiles)

	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("could not begin transaction: %v", err)
	}
	defer tx.Rollback()

	for _, fileName := range migrationFiles {
		log.Printf("Executing migration: %s", fileName)
		content, err := os.ReadFile(filepath.Join(migrationsDir, fileName))
		if err != nil {
			return fmt.Errorf("could not read migration %s: %v", fileName, err)
		}

		_, err = tx.Exec(string(content))
		if err != nil {
			return fmt.Errorf("could not execute migration %s: %v", fileName, err)
		}

		log.Printf("Successfully executed migration: %s", fileName)
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("could not commit transaction: %v", err)
	}

	return nil
}

func verifyTables(db *sql.DB) error {
	tables := []string{"users", "roles", "meetings", "subjects", "role_relations"}

	for _, table := range tables {
		var exists bool
		query := `
            SELECT EXISTS (
                SELECT FROM information_schema.tables 
                WHERE table_schema = 'public' 
                AND table_name = $1
            );
        `

		err := db.QueryRow(query, table).Scan(&exists)
		if err != nil {
			return fmt.Errorf("error checking table %s: %v", table, err)
		}

		if !exists {
			return fmt.Errorf("table %s was not created", table)
		}

		log.Printf("Verified table exists: %s", table)
	}

	return nil
}

func cleanupDatabase(db *sql.DB) {
	log.Println("Cleaning up test database...")
	tables := []string{"meetings", "role_relations", "roles", "subjects", "users"}
	for _, table := range tables {
		if _, err := db.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s CASCADE", table)); err != nil {
			log.Printf("Error dropping table %s: %v", table, err)
		}
	}
}
