package models

import "database/sql"

func RunMigrations(db *sql.DB) error {
	// Create users table
	usersQuery := `CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        email VARCHAR(255) NOT NULL UNIQUE,
        password VARCHAR(255),
        is_verified BOOLEAN DEFAULT FALSE,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        google_id VARCHAR(255) UNIQUE,
        github_id BIGINT UNIQUE,
        facebook_id BIGINT UNIQUE,
        microsoft_id VARCHAR(255) UNIQUE,
        linkedin_id BIGINT UNIQUE
    )`

	_, err := db.Exec(usersQuery)
	if err != nil {
		return err
	}

	// Create user_profile table
	profileQuery := `CREATE TABLE IF NOT EXISTS user_profile (
        id SERIAL PRIMARY KEY,
        user_id INTEGER NOT NULL UNIQUE,
        name VARCHAR(255),
        avatar VARCHAR(1024),
        bio TEXT,
        phone_number VARCHAR(20),
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
    )`

	_, err = db.Exec(profileQuery)
	if err != nil {
		return err
	}

	// Create user_extra_info table
	extraInfoQuery := `CREATE TABLE IF NOT EXISTS user_extra_info (
        id SERIAL PRIMARY KEY,
        user_id INTEGER NOT NULL,
        key VARCHAR(255) NOT NULL,
        value TEXT,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
        UNIQUE(user_id, key)
    )`

	_, err = db.Exec(extraInfoQuery)
	return err
}
