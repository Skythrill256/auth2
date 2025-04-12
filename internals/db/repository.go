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
	query := `INSERT INTO users (email, password, is_verified, google_id, github_id, gitlab_id, facebook_id, microsoft_id, linkedin_id, amazon_id, bitbucket_id, foursquare_id, heroku_id, instagram_id) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) RETURNING id`
	err = tx.QueryRow(query, user.Email, user.Password, user.IsVerified,
		user.GoogleID, user.GithubID, user.GitLabID, user.FacebookID, user.MicrosoftID, user.LinkedinID, user.AmazonID, user.BitbucketID, user.FoursquareID, user.HerokuID, user.InstagramID).Scan(&user.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (repo *Repository) GetUserByID(id int) (*models.User, error) {
	var user models.User
	query := `SELECT id, email, password, is_verified, created_at, updated_at, google_id, github_id, facebook_id, microsoft_id, linkedin_id, amazon_id FROM users WHERE id=$1`
	err := repo.DB.QueryRow(query, id).Scan(&user.ID, &user.Email, &user.Password, &user.IsVerified, &user.CreatedAt, &user.UpdatedAt, &user.GoogleID, &user.GithubID, &user.FacebookID, &user.MicrosoftID, &user.LinkedinID, &user.AmazonID)
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
	query := `SELECT id, email, password, is_verified, created_at, updated_at, google_id, github_id, facebook_id, microsoft_id, linkedin_id, amazon_id FROM users WHERE email=$1`
	err := repo.DB.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Password, &user.IsVerified, &user.CreatedAt, &user.UpdatedAt, &user.GoogleID, &user.GithubID, &user.FacebookID, &user.MicrosoftID, &user.LinkedinID, &user.AmazonID)
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

func (repo *Repository) CreateLoginRecord(userID int, ipAddress string) error {
	query := `INSERT INTO login_records (user_id, ip_address) VALUES ($1, $2)`
	_, err := repo.DB.Exec(query, userID, ipAddress)
	return err
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

func (repo *Repository) GetUserByAmazonID(amazonID string) (*models.User, error) {
	var user models.User
	query := `SELECT id, email, password, is_verified, created_at, updated_at, amazon_id
	FROM users WHERE amazon_id = $1`

	err := repo.DB.QueryRow(query, amazonID).Scan(
		&user.ID, &user.Email, &user.Password, &user.IsVerified, &user.CreatedAt, &user.UpdatedAt, &user.AmazonID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (repo *Repository) GetUserByBitbucketID(bitbucketID string) (*models.User, error) {
	var user models.User
	query := `SELECT id, email, password, is_verified, created_at, updated_at, bitbucket_id FROM users WHERE bitbucket_id=$1`
	err := repo.DB.QueryRow(query, bitbucketID).Scan(&user.ID, &user.Email, &user.Password, &user.IsVerified, &user.CreatedAt, &user.UpdatedAt, &user.BitbucketID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (repo *Repository) GetUserByFoursquareID(foursquareID string) (*models.User, error) {
	var user models.User
	query := `SELECT id, email, password, is_verified, created_at, updated_at, foursquare_id FROM users WHERE foursquare_id=$1`
	err := repo.DB.QueryRow(query, foursquareID).Scan(&user.ID, &user.Email, &user.Password, &user.IsVerified, &user.CreatedAt, &user.UpdatedAt, &user.FoursquareID)
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

func (repo *Repository) GetUserByGitLabID(gitlabID int64) (*models.User, error) {
	var user models.User
	query := `SELECT id, email, password, is_verified, created_at, updated_at, gitlab_id
	FROM users WHERE gitlab_id = $1`

	err := repo.DB.QueryRow(query, gitlabID).Scan(
		&user.ID, &user.Email, &user.Password, &user.IsVerified, &user.CreatedAt, &user.UpdatedAt, &user.GitLabID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (repo *Repository) GetUserByHerokuID(herokuID string) (*models.User, error) {
	var user models.User
	query := `SELECT id, email, password, is_verified, created_at, updated_at, heroku_id
	FROM users WHERE heroku_id = $1`

	err := repo.DB.QueryRow(query, herokuID).Scan(
		&user.ID, &user.Email, &user.Password, &user.IsVerified, &user.CreatedAt, &user.UpdatedAt, &user.HerokuID)
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

func (repo *Repository) GetUserLoginHistory(userID int) ([]models.LoginRecord, error) {
	query := `SELECT id, user_id, ip_address, login_time, created_at FROM login_records 
	         WHERE user_id = $1 ORDER BY login_time DESC`

	rows, err := repo.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []models.LoginRecord
	for rows.Next() {
		var record models.LoginRecord
		err := rows.Scan(&record.ID, &record.UserID, &record.IPAddress, &record.LoginTime, &record.CreatedAt)
		if err != nil {
			return nil, err
		}
		records = append(records, record)
	}

	return records, nil
}

func (repo *Repository) GetUserByInstagramID(instagramID string) (*models.User, error) {
	var user models.User
	query := `SELECT id, email, password, is_verified, created_at, updated_at, instagram_id
	FROM users WHERE instagram_id = $1`

	err := repo.DB.QueryRow(query, instagramID).Scan(
		&user.ID, &user.Email, &user.Password, &user.IsVerified, &user.CreatedAt, &user.UpdatedAt, &user.InstagramID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
