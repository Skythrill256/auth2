package models

import "database/sql"

func RunMigrations(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        email VARCHAR(255) NOT NULL UNIQUE,
        password VARCHAR(255),
        is_verified BOOLEAN DEFAULT FALSE,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        google_id VARCHAR(255) UNIQUE,
        github_id BIGINT UNIQUE,
        facebook_id BIGINT UNIQUE,
        microsoft_id VARCHAR(255) UNIQUE
    )`

	_, err := db.Exec(query)
	return err
}
