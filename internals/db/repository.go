package db

import (
	"database/sql"
	"errors"

	"github.com/Skythrill256/auth-service/internals/models"
)

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{DB: db}
}

func (repo *Repository) CreateUser(user *models.User) error {
	tx, err := repo.DB.Begin()
	if err != nil {
		return err
	}

	// Insert into users table
	query := `INSERT INTO users (email, password, is_verified, google_id, github_id, facebook_id, microsoft_id, linkedin_id) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
	err = tx.QueryRow(query, user.Email, user.Password, user.IsVerified,
		user.GoogleID, user.GithubID, user.FacebookID, user.MicrosoftID, user.LinkedinID).Scan(&user.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (repo *Repository) GetUserByID(id int) (*models.User, error) {
	var user models.User
	query := `SELECT id, email, password, is_verified, created_at, updated_at, google_id, github_id, facebook_id, microsoft_id, linkedin_id FROM users WHERE id=$1`
	err := repo.DB.QueryRow(query, id).Scan(&user.ID, &user.Email, &user.Password, &user.IsVerified, &user.CreatedAt, &user.UpdatedAt, &user.GoogleID, &user.GithubID, &user.FacebookID, &user.MicrosoftID, &user.LinkedinID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (repo *Repository) GetUserProfile(userID int) (*models.UserProfile, error) {
	var profile models.UserProfile
	query := `SELECT id, user_id, name, avatar, bio, phone_number, created_at, updated_at
	FROM user_profile WHERE user_id=$1`

	err := repo.DB.QueryRow(query, userID).Scan(
		&profile.ID, &profile.UserID, &profile.Name, &profile.Avatar,
		&profile.Bio, &profile.PhoneNumber, &profile.CreatedAt, &profile.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &profile, nil
}

func (repo *Repository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	query := `SELECT id, email, password, is_verified, created_at, updated_at, google_id, github_id, facebook_id, microsoft_id, linkedin_id FROM users WHERE email=$1`
	err := repo.DB.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Password, &user.IsVerified, &user.CreatedAt, &user.UpdatedAt, &user.GoogleID, &user.GithubID, &user.FacebookID, &user.MicrosoftID, &user.LinkedinID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (repo *Repository) VerifyUserEmail(email string) error {
	query := `UPDATE users SET is_verified = true, updated_at = CURRENT_TIMESTAMP WHERE email = $1`
	_, err := repo.DB.Exec(query, email)
	if err != nil {
		return err
	}
	return nil
}

func (repo *Repository) GetUserByGoogleID(googleID string) (*models.User, error) {
	var user models.User
	query := `SELECT id, email, password, is_verified, created_at, updated_at, google_id
	FROM users WHERE google_id = $1`

	err := repo.DB.QueryRow(query, googleID).Scan(
		&user.ID, &user.Email, &user.Password, &user.IsVerified, &user.CreatedAt, &user.UpdatedAt, &user.GoogleID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (repo *Repository) GetUserByGithubID(githubID int64) (*models.User, error) {
	var user models.User
	query := `SELECT id, email, password, is_verified, created_at, updated_at, github_id
	FROM users WHERE github_id = $1`

	err := repo.DB.QueryRow(query, githubID).Scan(
		&user.ID, &user.Email, &user.Password, &user.IsVerified, &user.CreatedAt, &user.UpdatedAt, &user.GithubID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (repo *Repository) GetUserByMicrosoftID(microsoftID string) (*models.User, error) {
	var user models.User
	query := `SELECT id, email, password, is_verified, created_at, updated_at, microsoft_id
	FROM users WHERE microsoft_id = $1`

	err := repo.DB.QueryRow(query, microsoftID).Scan(
		&user.ID, &user.Email, &user.Password, &user.IsVerified, &user.CreatedAt, &user.UpdatedAt, &user.MicrosoftID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (repo *Repository) UpdateProfile(userID int, name, avatar, bio, phoneNumber string) error {
	// Check if profile exists
	var exists bool
	err := repo.DB.QueryRow(`SELECT EXISTS(SELECT 1 FROM user_profile WHERE user_id = $1)`, userID).Scan(&exists)
	if err != nil {
		return err
	}

	if exists {
		// Update existing profile
		query := `UPDATE user_profile SET name = $1, avatar = $2, bio = $3, phone_number = $4, updated_at = CURRENT_TIMESTAMP 
		WHERE user_id = $5`
		_, err = repo.DB.Exec(query, name, avatar, bio, phoneNumber, userID)
	} else {
		// Create new profile
		query := `INSERT INTO user_profile (user_id, name, avatar, bio, phone_number) VALUES ($1, $2, $3, $4, $5)`
		_, err = repo.DB.Exec(query, userID, name, avatar, bio, phoneNumber)
	}

	return err
}

func (repo *Repository) GetUserByFacebookID(facebookID int64) (*models.User, error) {
	var user models.User
	query := `SELECT id, email, password, is_verified, created_at, updated_at, facebook_id FROM users WHERE facebook_id = $1`
	err := repo.DB.QueryRow(query, facebookID).Scan(&user.ID, &user.Email, &user.Password, &user.IsVerified, &user.CreatedAt, &user.UpdatedAt, &user.FacebookID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (repo *Repository) GetUserByLinkedinID(linkedinID int64) (*models.User, error) {
	var user models.User
	query := `SELECT id, email, password, is_verified, created_at, updated_at, linkedin_id FROM users WHERE linkedin_id = $1`
	err := repo.DB.QueryRow(query, linkedinID).Scan(&user.ID, &user.Email, &user.Password, &user.IsVerified, &user.CreatedAt, &user.UpdatedAt, &user.LinkedinID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (repo *Repository) ForgotPassword(email string) error {
	query := `UPDATE users SET password = $1 WHERE email = $2`
	_, err := repo.DB.Exec(query, email, "password")
	if err != nil {
		return err
	}
	return nil
}
func (repo *Repository) UpdateUserPassword(email string, newPassword string) error {
	query := `UPDATE users SET password = $1, updated_at = CURRENT_TIMESTAMP WHERE email = $2`

	_, err := repo.DB.Exec(query, newPassword, email)
	if err != nil {
		return err
	}
	return nil
}

func (repo *Repository) CreateUserExtraInfo(info *models.UserExtraInfo) error {
	query := `INSERT INTO user_extra_info (user_id, key, value) VALUES ($1, $2, $3) RETURNING id, created_at, updated_at`
	err := repo.DB.QueryRow(query, info.UserID, info.Key, info.Value).Scan(&info.ID, &info.CreatedAt, &info.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (repo *Repository) GetUserExtraInfo(userID int, key string) (*models.UserExtraInfo, error) {
	var info models.UserExtraInfo
	query := `SELECT id, user_id, key, value, created_at, updated_at FROM user_extra_info WHERE user_id = $1 AND key = $2`
	err := repo.DB.QueryRow(query, userID, key).Scan(&info.ID, &info.UserID, &info.Key, &info.Value, &info.CreatedAt, &info.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &info, nil
}

func (repo *Repository) GetAllUserExtraInfo(userID int) ([]models.UserExtraInfo, error) {
	query := `SELECT id, user_id, key, value, created_at, updated_at FROM user_extra_info WHERE user_id = $1`
	rows, err := repo.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var infoList []models.UserExtraInfo
	for rows.Next() {
		var info models.UserExtraInfo
		err := rows.Scan(&info.ID, &info.UserID, &info.Key, &info.Value, &info.CreatedAt, &info.UpdatedAt)
		if err != nil {
			return nil, err
		}
		infoList = append(infoList, info)
	}
	return infoList, rows.Err()
}

func (repo *Repository) UpdateUserExtraInfo(info *models.UserExtraInfo) error {
	query := `UPDATE user_extra_info SET value = $1, updated_at = CURRENT_TIMESTAMP WHERE user_id = $2 AND key = $3 RETURNING id`
	err := repo.DB.QueryRow(query, info.Value, info.UserID, info.Key).Scan(&info.ID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *Repository) DeleteUserExtraInfo(userID int, key string) error {
	query := `DELETE FROM user_extra_info WHERE user_id = $1 AND key = $2`
	result, err := repo.DB.Exec(query, userID, key)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}
