package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort                string
	DBHost                 string
	DBPort                 string
	DBUser                 string
	DBPassword             string
	DBName                 string
	JWTSecret              string
	EmailHost              string
	EmailPort              string
	EmailSender            string
	EmailUsername          string
	EmailPass              string
	GoogleClientID         string
	GoogleClientSecret     string
	GoogleRedirectURL      string
	GithubClientID         string
	GithubClientSecret     string
	GithubRedirectURL      string
	GitLabClientID         string
	GitLabClientSecret     string
	GitLabRedirectURL      string
	FacebookClientID       string
	FacebookClientSecret   string
	FacebookRedirectURL    string
	MicrosoftClientID      string
	MicrosoftClientSecret  string
	MicrosoftRedirectURL   string
	LinkedInClientID       string
	LinkedInClientSecret   string
	LinkedInRedirectURL    string
	AmazonClientID         string
	AmazonClientSecret     string
	AmazonRedirectURL      string
	BitbucketClientID      string
	BitbucketClientSecret  string
	BitbucketRedirectURL   string
	FoursquareClientID     string
	FoursquareClientSecret string
	FoursquareRedirectURL  string
	HerokuClientID         string
	HerokuClientSecret     string
	HerokuRedirectURL      string
	InstagramClientID      string
	InstagramClientSecret  string
	InstagramRedirectURL   string
	JiraClientID           string
	JiraClientSecret       string
	JiraRedirectURL        string
	SlackClientID          string
	SlackClientSecret      string
	SlackRedirectURL       string
	SpotifyClientID        string
	SpotifyClientSecret    string
	SpotifyRedirectURL     string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		AppPort:                os.Getenv("APP_PORT"),
		DBHost:                 os.Getenv("DB_HOST"),
		DBPort:                 os.Getenv("DB_PORT"),
		DBUser:                 os.Getenv("DB_USER"),
		DBPassword:             os.Getenv("DB_PASSWORD"),
		DBName:                 os.Getenv("DB_NAME"),
		JWTSecret:              os.Getenv("JWT_SECRET"),
		EmailHost:              os.Getenv("EMAIL_HOST"),
		EmailPort:              os.Getenv("EMAIL_PORT"),
		EmailSender:            os.Getenv("EMAIL_SENDER"),
		EmailUsername:          os.Getenv("EMAIL_USERNAME"),
		EmailPass:              os.Getenv("EMAIL_PASSWORD"),
		GoogleClientID:         os.Getenv("GOOGLE_CLIENT_ID"),
		GoogleClientSecret:     os.Getenv("GOOGLE_CLIENT_SECRET"),
		GoogleRedirectURL:      os.Getenv("GOOGLE_REDIRECT_URL"),
		GithubClientID:         os.Getenv("GITHUB_CLIENT_ID"),
		GithubClientSecret:     os.Getenv("GITHUB_CLIENT_SECRET"),
		GithubRedirectURL:      os.Getenv("GITHUB_REDIRECT_URL"),
		GitLabClientID:         os.Getenv("GITLAB_CLIENT_ID"),
		GitLabClientSecret:     os.Getenv("GITLAB_CLIENT_SECRET"),
		GitLabRedirectURL:      os.Getenv("GITLAB_REDIRECT_URL"),
		FacebookClientID:       os.Getenv("FACEBOOK_CLIENT_ID"),
		FacebookClientSecret:   os.Getenv("FACEBOOK_CLIENT_SECRET"),
		FacebookRedirectURL:    os.Getenv("FACEBOOK_REDIRECT_URL"),
		MicrosoftClientID:      os.Getenv("MICROSOFT_CLIENT_ID"),
		MicrosoftClientSecret:  os.Getenv("MICROSOFT_CLIENT_SECRET"),
		MicrosoftRedirectURL:   os.Getenv("MICROSOFT_REDIRECT_URL"),
		LinkedInClientID:       os.Getenv("LINKEDIN_CLIENT_ID"),
		LinkedInClientSecret:   os.Getenv("LINKEDIN_CLIENT_SECRET"),
		LinkedInRedirectURL:    os.Getenv("LINKEDIN_REDIRECT_URL"),
		AmazonClientID:         os.Getenv("AMAZON_CLIENT_ID"),
		AmazonClientSecret:     os.Getenv("AMAZON_CLIENT_SECRET"),
		AmazonRedirectURL:      os.Getenv("AMAZON_REDIRECT_URL"),
		BitbucketClientID:      os.Getenv("BITBUCKET_CLIENT_ID"),
		BitbucketClientSecret:  os.Getenv("BITBUCKET_CLIENT_SECRET"),
		BitbucketRedirectURL:   os.Getenv("BITBUCKET_REDIRECT_URL"),
		FoursquareClientID:     os.Getenv("FOURSQUARE_CLIENT_ID"),
		FoursquareClientSecret: os.Getenv("FOURSQUARE_CLIENT_SECRET"),
		FoursquareRedirectURL:  os.Getenv("FOURSQUARE_REDIRECT_URL"),
		HerokuClientID:         os.Getenv("HEROKU_CLIENT_ID"),
		HerokuClientSecret:     os.Getenv("HEROKU_CLIENT_SECRET"),
		HerokuRedirectURL:      os.Getenv("HEROKU_REDIRECT_URL"),
		InstagramClientID:      os.Getenv("INSTAGRAM_CLIENT_ID"),
		InstagramClientSecret:  os.Getenv("INSTAGRAM_CLIENT_SECRET"),
		InstagramRedirectURL:   os.Getenv("INSTAGRAM_REDIRECT_URL"),
		JiraClientID:           os.Getenv("JIRA_CLIENT_ID"),
		JiraClientSecret:       os.Getenv("JIRA_CLIENT_SECRET"),
		JiraRedirectURL:        os.Getenv("JIRA_REDIRECT_URL"),
		SlackClientID:          os.Getenv("SLACK_CLIENT_ID"),
		SlackClientSecret:      os.Getenv("SLACK_CLIENT_SECRET"),
		SlackRedirectURL:       os.Getenv("SLACK_REDIRECT_URL"),
		SpotifyClientID:        os.Getenv("SPOTIFY_CLIENT_ID"),
		SpotifyClientSecret:    os.Getenv("SPOTIFY_CLIENT_SECRET"),
		SpotifyRedirectURL:     os.Getenv("SPOTIFY_REDIRECT_URL"),
	}
}
