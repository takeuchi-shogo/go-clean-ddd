package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/takeuchi-shogo/go-clean-ddd/pkg/infra/config"
	"github.com/takeuchi-shogo/go-clean-ddd/pkg/infra/database"
)

func main() {
	var migrationDir string
	var seedDir string

	flag.StringVar(&migrationDir, "migration-dir", "migrations", "Directory with migration files")
	flag.StringVar(&seedDir, "seed-dir", "seeds", "Directory with seed files")
	flag.Parse()

	c := config.New("LOCAL")
	// Connect to the database using GORM
	db, err := database.NewDB(*c)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	// Get the underlying *sql.DB instance
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Could not get *sql.DB instance: %v", err)
	}
	defer sqlDB.Close()

	driver, err := mysql.WithInstance(sqlDB, &mysql.Config{})
	if err != nil {
		log.Fatalf("Could not create database driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", migrationDir),
		"mysql", driver)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("An error occurred while syncing the database: %v", err)
	}

	log.Println("Migration completed successfully")

	if err := seedDatabase(sqlDB, seedDir); err != nil {
		log.Fatalf("Seeding failed: %v", err)
	}

	log.Println("Seeding completed successfully")
}

func seedDatabase(db *sql.DB, seedDir string) error {
	files, err := os.ReadDir(seedDir)
	if err != nil {
		return fmt.Errorf("could not read seed directory: %v", err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		content, err := os.ReadFile(fmt.Sprintf("%s/%s", seedDir, file.Name()))
		if err != nil {
			return fmt.Errorf("could not read seed file %s: %v", file.Name(), err)
		}

		if _, err := db.Exec(string(content)); err != nil {
			return fmt.Errorf("could not execute seed file %s: %v", file.Name(), err)
		}

		log.Printf("Executed seed file: %s", file.Name())
	}

	return nil
}
