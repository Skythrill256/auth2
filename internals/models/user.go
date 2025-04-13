package models

import (
	"time"
)

type User struct {
	ID           int       `json:"id"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	IsVerified   bool      `json:"is_verified"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	GoogleID     *string   `json:"google_id,omitempty"`
	GithubID     *int64    `json:"github_id,omitempty"`
	GitLabID     *int64    `json:"gitlab_id,omitempty"`
	FacebookID   *int64    `json:"facebook_id,omitempty"`
	MicrosoftID  *string   `json:"microsoft_id,omitempty"`
	LinkedinID   *int64    `json:"linkedin_id,omitempty"`
	AmazonID     *string   `json:"amazon_id,omitempty"`
	BitbucketID  *string   `json:"bitbucket_id,omitempty"`
	FoursquareID *string   `json:"foursquare_id,omitempty"`
	HerokuID     *string   `json:"heroku_id,omitempty"`
	InstagramID  *string   `json:"instagram_id,omitempty"`
	JiraID       *string   `json:"jira_id,omitempty"`
	SlackID      *string   `json:"slack_id,omitempty"`
	SpotifyID    *string   `json:"spotify_id,omitempty"`
	YahooID      *string   `json:"yahoo_id,omitempty"`
}

type UserProfile struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	Name        string    `json:"name"`
	Avatar      string    `json:"avatar,omitempty"`
	Bio         string    `json:"bio,omitempty"`
	PhoneNumber string    `json:"phone_number,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UserExtraInfo struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Key       string    `json:"key"`
	Value     string    `json:"value"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
