package storage

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Database struct {
	DB *sql.DB
}

func NewDatabase() (*Database, error) {
	dsn := "host=localhost user=postgres dbname=food_blog port=5432 sslmode=disable TimeZone=UTC"
	// return sql.Open("postgres", dsn)

	// Open a database connection
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	// Ping the database to ensure a successful connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Database connection established")

	return &Database{db}, nil
}

func CreateExtenstion(database *Database) error {
	_, err := database.DB.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`)
	if err != nil {
		return fmt.Errorf("failed to create extenstion: %s", err)
	}
	return nil
}

func CreateRecipesTable(database *Database) error {
	_, err := database.DB.Exec(`CREATE TABLE IF NOT EXISTS recipies(
		id uuid NOT NULL,
		title VARCHAR NOT NULL,
		description VARCHAR NOT NULL, 
		ingredients VARCHAR NOT NULL,
		instructions VARCHAR NOT NULL,
		PRIMARY KEY (id)
	)`)
	if err != nil {
		return fmt.Errorf("failed to create recipies table %s", err)
	}
	return nil
}
