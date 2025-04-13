package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Skythrill256/auth-service/internals/config"
	"github.com/Skythrill256/auth-service/internals/db"
	"github.com/Skythrill256/auth-service/internals/models"
	"github.com/Skythrill256/auth-service/internals/utils"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/amazon"
	"golang.org/x/oauth2/bitbucket"
	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/foursquare"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/gitlab"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/heroku"
	"golang.org/x/oauth2/instagram"
	"golang.org/x/oauth2/linkedin"
	"golang.org/x/oauth2/microsoft"
	"golang.org/x/oauth2/slack"
	"golang.org/x/oauth2/spotify"
	"golang.org/x/oauth2/yahoo"
)

func GetGoogleOAuthConfig(cfg *config.Config) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     cfg.GoogleClientID,
		ClientSecret: cfg.GoogleClientSecret,
		RedirectURL:  cfg.GoogleRedirectURL,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}
}

func GetGithubOAuthConfig(cfg *config.Config) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     cfg.GithubClientID,
		ClientSecret: cfg.GithubClientSecret,
		RedirectURL:  cfg.GithubRedirectURL,
		Scopes:       []string{"user"},
		Endpoint:     github.Endpoint,
	}
}

func GetFacebookOAuthConfig(cfg *config.Config) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     cfg.FacebookClientID,
		ClientSecret: cfg.FacebookClientSecret,
		RedirectURL:  cfg.FacebookRedirectURL,
		Scopes:       []string{"email"},
		Endpoint:     facebook.Endpoint,
	}
}

func GetAmazonOAuthConfig(cfg *config.Config) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     cfg.AmazonClientID,
		ClientSecret: cfg.AmazonClientSecret,
		RedirectURL:  cfg.AmazonRedirectURL,
		Scopes:       []string{"profile"}, // Basic profile information
		Endpoint:     amazon.Endpoint,
	}
}

func GetMicrosoftOAuthConfig(cfg *config.Config) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     cfg.MicrosoftClientID,
		ClientSecret: cfg.MicrosoftClientSecret,
		RedirectURL:  cfg.MicrosoftRedirectURL,
		Scopes:       []string{"openid", "profile", "email", "offline_access", "User.Read"},
		Endpoint:     microsoft.AzureADEndpoint("common"),
	}
}

func GetLinkedinOAuthConfig(cfg *config.Config) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     cfg.LinkedInClientID,
		ClientSecret: cfg.LinkedInClientSecret,
		RedirectURL:  cfg.LinkedInRedirectURL,
		Scopes:       []string{"r_emailaddress", "r_liteprofile"},
		Endpoint:     linkedin.Endpoint,
	}
}

func GetBitbucketOAuthConfig(cfg *config.Config) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     cfg.BitbucketClientID,
		ClientSecret: cfg.BitbucketClientSecret,
		RedirectURL:  cfg.BitbucketRedirectURL,
		Scopes:       []string{"account", "email"},
		Endpoint:     bitbucket.Endpoint,
	}
}

func GetFoursquareOAuthConfig(cfg *config.Config) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     cfg.FoursquareClientID,
		ClientSecret: cfg.FoursquareClientSecret,
		RedirectURL:  cfg.FoursquareRedirectURL,
		Scopes:       []string{},
		Endpoint:     foursquare.Endpoint,
	}
}

func GetGitLabOAuthConfig(cfg *config.Config) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     cfg.GitLabClientID,
		ClientSecret: cfg.GitLabClientSecret,
		RedirectURL:  cfg.GitLabRedirectURL,
		Scopes:       []string{"read_user", "email"},
		Endpoint:     gitlab.Endpoint,
	}
}

func GetHerokuOAuthConfig(cfg *config.Config) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     cfg.HerokuClientID,
		ClientSecret: cfg.HerokuClientSecret,
		RedirectURL:  cfg.HerokuRedirectURL,
		Scopes:       []string{"global"},
		Endpoint:     heroku.Endpoint,
	}
}

func GetInstagramOAuthConfig(cfg *config.Config) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     cfg.InstagramClientID,
		ClientSecret: cfg.InstagramClientSecret,
		RedirectURL:  cfg.InstagramRedirectURL,
		Scopes:       []string{"basic"}, // Basic profile information
		Endpoint:     instagram.Endpoint,
	}
}

func GetJiraOAuthConfig(cfg *config.Config) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     cfg.JiraClientID,
		ClientSecret: cfg.JiraClientSecret,
		RedirectURL:  cfg.JiraRedirectURL,
		Scopes:       []string{"read:jira-user"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://auth.atlassian.com/authorize",
			TokenURL: "https://auth.atlassian.com/oauth/token",
		},
	}
}

func GetSlackOAuthConfig(cfg *config.Config) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     cfg.SlackClientID,
		ClientSecret: cfg.SlackClientSecret,
		RedirectURL:  cfg.SlackRedirectURL,
		Scopes:       []string{"identity.basic", "identity.email"},
		Endpoint:     slack.Endpoint,
	}
}

func GetSpotifyOAuthConfig(cfg *config.Config) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     cfg.SpotifyClientID,
		ClientSecret: cfg.SpotifyClientSecret,
		RedirectURL:  cfg.SpotifyRedirectURL,
		Scopes:       []string{"user-read-email", "user-read-private"},
		Endpoint:     spotify.Endpoint,
	}
}

func GetYahooOAuthConfig(cfg *config.Config) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     cfg.YahooClientID,
		ClientSecret: cfg.YahooClientSecret,
		RedirectURL:  cfg.YahooRedirectURL,
		Scopes:       []string{"openid", "email", "profile"},
		Endpoint:     yahoo.Endpoint,
	}
}

func GoogleOAuthConsentURL(cfg *config.Config) string {
	oauthConfig := GetGoogleOAuthConfig(cfg)
	return oauthConfig.AuthCodeURL("state")
}

func GithubOAuthConsentURL(cfg *config.Config) string {
	oauthConfig := GetGithubOAuthConfig(cfg)
	return oauthConfig.AuthCodeURL("state")
}

func FacebookOAuthConsentURL(cfg *config.Config) string {
	oauthConfig := GetFacebookOAuthConfig(cfg)
	return oauthConfig.AuthCodeURL("state")
}

func MicrosoftOAuthConsentURL(cfg *config.Config) string {
	oauthConfig := GetMicrosoftOAuthConfig(cfg)
	return oauthConfig.AuthCodeURL("state")
}

func LinkedinOAuthConsentURL(cfg *config.Config) string {
	oauthConfig := GetLinkedinOAuthConfig(cfg)
	return oauthConfig.AuthCodeURL("state")
}

func AmazonOAuthConsentURL(cfg *config.Config) string {
	oauthConfig := GetAmazonOAuthConfig(cfg)
	return oauthConfig.AuthCodeURL("state")
}

func BitbucketOAuthConsentURL(cfg *config.Config) string {
	oauthConfig := GetBitbucketOAuthConfig(cfg)
	return oauthConfig.AuthCodeURL("state")
}

func FoursquareOAuthConsentURL(cfg *config.Config) string {
	oauthConfig := GetFoursquareOAuthConfig(cfg)
	return oauthConfig.AuthCodeURL("state")
}

func GitLabOAuthConsentURL(cfg *config.Config) string {
	oauthConfig := GetGitLabOAuthConfig(cfg)
	return oauthConfig.AuthCodeURL("state")
}

func HerokuOAuthConsentURL(cfg *config.Config) string {
	oauthConfig := GetHerokuOAuthConfig(cfg)
	return oauthConfig.AuthCodeURL("state")
}

func InstagramOAuthConsentURL(cfg *config.Config) string {
	oauthConfig := GetInstagramOAuthConfig(cfg)
	return oauthConfig.AuthCodeURL("state")
}

func JiraOAuthConsentURL(cfg *config.Config) string {
	oauthConfig := GetJiraOAuthConfig(cfg)
	return oauthConfig.AuthCodeURL("state")
}

func SlackOAuthConsentURL(cfg *config.Config) string {
	oauthConfig := GetSlackOAuthConfig(cfg)
	return oauthConfig.AuthCodeURL("state")
}

func SpotifyOAuthConsentURL(cfg *config.Config) string {
	oauthConfig := GetSpotifyOAuthConfig(cfg)
	return oauthConfig.AuthCodeURL("state")
}

func FoursquareLogin(cfg *config.Config, repository *db.Repository, code string, ipAddress string) (string, error) {
	oauthConfig := GetFoursquareOAuthConfig(cfg)

	oauthToken, err := oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return "", err
	}

	// Foursquare API v2 requires the OAuth token as a parameter
	userInfoURL := fmt.Sprintf("https://api.foursquare.com/v2/users/self?oauth_token=%s&v=20231201", oauthToken.AccessToken)
	resp, err := http.Get(userInfoURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var foursquareResponse struct {
		Response struct {
			User struct {
				ID    string `json:"id"`
				Email string `json:"contact,omitempty"`
			} `json:"user"`
		} `json:"response"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&foursquareResponse); err != nil {
		return "", err
	}

	foursquareID := foursquareResponse.Response.User.ID
	email := foursquareResponse.Response.User.Email

	if foursquareID == "" {
		return "", errors.New("failed to get Foursquare ID from response")
	}

	user, err := repository.GetUserByFoursquareID(foursquareID)
	if err != nil {
		return "", err
	}

	if user == nil {
		newUser := &models.User{
			Email:        email,
			IsVerified:   true,
			FoursquareID: &foursquareID,
		}
		err := repository.CreateUser(newUser)
		if err != nil {
			return "", err
		}
		user = newUser
	}

	// Log the login attempt
	err = repository.CreateLoginRecord(user.ID, ipAddress)
	if err != nil {
		return "", err
	}

	jwtToken, err := utils.GenerateJWT(user.Email, cfg.JWTSecret)
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func GoogleLogin(cfg *config.Config, repository *db.Repository, code string, ipAddress string) (string, error) {
	oauthConfig := GetGoogleOAuthConfig(cfg)

	oauthToken, err := oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return "", err
	}

	client := oauthConfig.Client(context.Background(), oauthToken)

	resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var googleUser map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&googleUser); err != nil {
		return "", err
	}

	email, ok := googleUser["email"].(string)
	if !ok {
		return "", errors.New("failed to get email from Google response")
	}
	googleID, ok := googleUser["sub"].(string)
	if !ok {
		return "", errors.New("failed to get Google ID from response")
	}

	user, err := repository.GetUserByGoogleID(googleID)
	if err != nil {
		return "", err
	}

	if user == nil {
		newUser := &models.User{
			Email:      email,
			IsVerified: true,
			GoogleID:   &googleID,
		}
		err := repository.CreateUser(newUser)
		if err != nil {
			return "", err
		}
		user = newUser
	}

	// Log the login attempt
	err = repository.CreateLoginRecord(user.ID, ipAddress)
	if err != nil {
		return "", err
	}

	jwtToken, err := utils.GenerateJWT(user.Email, cfg.JWTSecret)
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func GithubLogin(cfg *config.Config, repository *db.Repository, code string, ipAddress string) (string, error) {
	oauthConfig := GetGithubOAuthConfig(cfg)

	oauthToken, err := oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return "", err
	}

	client := oauthConfig.Client(context.Background(), oauthToken)

	resp, err := client.Get("https://api.github.com/user/emails")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var githubEmails []map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&githubEmails); err != nil {
		return "", err
	}

	var email string
	for _, githubEmail := range githubEmails {
		if primary, ok := githubEmail["primary"].(bool); ok && primary {
			if emailStr, ok := githubEmail["email"].(string); ok {
				email = emailStr
				break
			}
		}
	}

	if email == "" {
		return "", errors.New("failed to get primary email from Github response")
	}

	resp, err = client.Get("https://api.github.com/user")
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	var githubUser map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&githubUser); err != nil {
		return "", err
	}

	githubID, ok := githubUser["id"].(float64)
	if !ok {
		return "", errors.New("failed to get Github ID from response")
	}

	githubIDInt := int64(githubID)

	user, err := repository.GetUserByGithubID(githubIDInt)
	if err != nil {
		return "", err
	}

	if user == nil {
		newUser := &models.User{
			Email:      email,
			IsVerified: true,
			GithubID:   &githubIDInt,
		}
		err := repository.CreateUser(newUser)
		if err != nil {
			return "", err
		}
		user = newUser
	}

	// Log the login attempt
	err = repository.CreateLoginRecord(user.ID, ipAddress)
	if err != nil {
		return "", err
	}

	jwtToken, err := utils.GenerateJWT(user.Email, cfg.JWTSecret)
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func FacebookLogin(cfg *config.Config, repository *db.Repository, code string, ipAddress string) (string, error) {
	oauthConfig := GetFacebookOAuthConfig(cfg)

	oauthToken, err := oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return "", err
	}

	client := oauthConfig.Client(context.Background(), oauthToken)

	resp, err := client.Get("https://graph.facebook.com/me?fields=email")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var facebookUser map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&facebookUser); err != nil {
		return "", err
	}

	fmt.Println(facebookUser)

	email, ok := facebookUser["email"].(string)
	if !ok {
		return "", errors.New("failed to get email from Facebook response")
	}

	facebookID, ok := facebookUser["id"].(string)
	if !ok {
		return "", errors.New("failed to get Facebook ID from response")
	}

	facebookIDInt, err := strconv.ParseInt(facebookID, 10, 64)
	if err != nil {
		return "", err
	}

	user, err := repository.GetUserByFacebookID(facebookIDInt)
	if err != nil {
		return "", err
	}

	if user == nil {
		newUser := &models.User{
			Email:      email,
			IsVerified: true,
			FacebookID: &facebookIDInt,
		}
		err := repository.CreateUser(newUser)
		if err != nil {
			return "", err
		}
		user = newUser
	}

	// Log the login attempt
	err = repository.CreateLoginRecord(user.ID, ipAddress)
	if err != nil {
		return "", err
	}

	jwtToken, err := utils.GenerateJWT(user.Email, cfg.JWTSecret)
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func MicrosoftLogin(cfg *config.Config, repository *db.Repository, code string, ipAddress string) (string, error) {
	oauthConfig := GetMicrosoftOAuthConfig(cfg)

	oauthToken, err := oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return "", err
	}

	accessToken := oauthToken.AccessToken
	if accessToken == "" {
		return "", errors.New("failed to retrieve access token")
	}

	req, err := http.NewRequest("GET", "https://graph.microsoft.com/v1.0/me", nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var microsoftUser struct {
		ID    string `json:"id"`
		Email string `json:"mail"`
		UPN   string `json:"userPrincipalName"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&microsoftUser); err != nil {
		return "", err
	}

	if microsoftUser.Email == "" {
		microsoftUser.Email = microsoftUser.UPN
	}

	if microsoftUser.Email == "" {
		return "", errors.New("failed to retrieve email from Microsoft API")
	}

	user, err := repository.GetUserByMicrosoftID(microsoftUser.ID)
	if err != nil {
		return "", err
	}

	if user == nil {
		newUser := &models.User{
			Email:       microsoftUser.Email,
			IsVerified:  true,
			MicrosoftID: &microsoftUser.ID,
		}
		if err := repository.CreateUser(newUser); err != nil {
			return "", err
		}
		user = newUser
	}

	// Log the login attempt
	err = repository.CreateLoginRecord(user.ID, ipAddress)
	if err != nil {
		return "", err
	}

	jwtToken, err := utils.GenerateJWT(user.Email, cfg.JWTSecret)
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func AmazonLogin(cfg *config.Config, repository *db.Repository, code string, ipAddress string) (string, error) {
	oauthConfig := GetAmazonOAuthConfig(cfg)

	oauthToken, err := oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return "", err
	}

	client := oauthConfig.Client(context.Background(), oauthToken)

	// Get user profile from Amazon
	resp, err := client.Get("https://api.amazon.com/user/profile")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var amazonUser struct {
		UserID string `json:"user_id"`
		Email  string `json:"email"`
		Name   string `json:"name"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&amazonUser); err != nil {
		return "", err
	}

	if amazonUser.Email == "" {
		return "", errors.New("failed to get email from Amazon response")
	}

	user, err := repository.GetUserByAmazonID(amazonUser.UserID)
	if err != nil {
		return "", err
	}

	if user == nil {
		newUser := &models.User{
			Email:      amazonUser.Email,
			IsVerified: true,
			AmazonID:   &amazonUser.UserID,
		}
		err := repository.CreateUser(newUser)
		if err != nil {
			return "", err
		}
		user = newUser
	}

	// Log the login attempt
	err = repository.CreateLoginRecord(user.ID, ipAddress)
	if err != nil {
		return "", err
	}

	jwtToken, err := utils.GenerateJWT(user.Email, cfg.JWTSecret)
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func BitbucketLogin(cfg *config.Config, repository *db.Repository, code string, ipAddress string) (string, error) {
	oauthConfig := GetBitbucketOAuthConfig(cfg)

	oauthToken, err := oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return "", err
	}

	client := oauthConfig.Client(context.Background(), oauthToken)

	// Get user profile from Bitbucket
	resp, err := client.Get("https://api.bitbucket.org/2.0/user")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var bitbucketUser struct {
		UUID        string `json:"uuid"`
		Username    string `json:"username"`
		AccountID   string `json:"account_id"`
		DisplayName string `json:"display_name"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&bitbucketUser); err != nil {
		return "", err
	}

	// Get user email from Bitbucket
	resp, err = client.Get("https://api.bitbucket.org/2.0/user/emails")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var emailResponse struct {
		Values []struct {
			Email     string `json:"email"`
			IsPrimary bool   `json:"is_primary"`
		} `json:"values"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&emailResponse); err != nil {
		return "", err
	}

	var email string
	for _, e := range emailResponse.Values {
		if e.IsPrimary {
			email = e.Email
			break
		}
	}

	if email == "" {
		return "", errors.New("failed to get email from Bitbucket response")
	}

	user, err := repository.GetUserByBitbucketID(bitbucketUser.UUID)
	if err != nil {
		return "", err
	}

	if user == nil {
		newUser := &models.User{
			Email:       email,
			IsVerified:  true,
			BitbucketID: &bitbucketUser.UUID,
		}
		err := repository.CreateUser(newUser)
		if err != nil {
			return "", err
		}
		user = newUser
	}

	// Log the login attempt
	err = repository.CreateLoginRecord(user.ID, ipAddress)
	if err != nil {
		return "", err
	}

	jwtToken, err := utils.GenerateJWT(user.Email, cfg.JWTSecret)
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func LinkedinLogin(cfg *config.Config, repository *db.Repository, code string, ipAddress string) (string, error) {
	oauthConfig := GetLinkedinOAuthConfig(cfg)

	oauthToken, err := oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return "", err
	}

	client := oauthConfig.Client(context.Background(), oauthToken)

	resp, err := client.Get("https://api.linkedin.com/v2/emailAddress?q=members&projection=(elements*(handle~))")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var emailResponse struct {
		Elements []struct {
			Handle struct {
				EmailAddress string `json:"emailAddress"`
			} `json:"handle~"`
		} `json:"elements"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&emailResponse); err != nil {
		return "", err
	}

	if len(emailResponse.Elements) == 0 {
		return "", errors.New("failed to get email from LinkedIn response")
	}

	email := emailResponse.Elements[0].Handle.EmailAddress

	resp, err = client.Get("https://api.linkedin.com/v2/me")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var linkedinUser struct {
		ID string `json:"id"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&linkedinUser); err != nil {
		return "", err
	}

	linkedinID, err := strconv.ParseInt(linkedinUser.ID, 10, 64)
	if err != nil {
		return "", err
	}

	user, err := repository.GetUserByLinkedinID(linkedinID)
	if err != nil {
		return "", err
	}

	if user == nil {
		newUser := &models.User{
			Email:      email,
			IsVerified: true,
			LinkedinID: &linkedinID,
		}
		err := repository.CreateUser(newUser)
		if err != nil {
			return "", err
		}
		user = newUser
	}

	// Log the login attempt
	err = repository.CreateLoginRecord(user.ID, ipAddress)
	if err != nil {
		return "", err
	}

	jwtToken, err := utils.GenerateJWT(user.Email, cfg.JWTSecret)
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func GitLabLogin(cfg *config.Config, repository *db.Repository, code string, ipAddress string) (string, error) {
	oauthConfig := GetGitLabOAuthConfig(cfg)

	oauthToken, err := oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return "", err
	}

	client := oauthConfig.Client(context.Background(), oauthToken)

	resp, err := client.Get("https://gitlab.com/api/v4/user")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var gitlabUser struct {
		ID    int64  `json:"id"`
		Email string `json:"email"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&gitlabUser); err != nil {
		return "", err
	}

	if gitlabUser.Email == "" {
		return "", errors.New("failed to get email from GitLab response")
	}

	user, err := repository.GetUserByGitLabID(gitlabUser.ID)
	if err != nil {
		return "", err
	}

	if user == nil {
		newUser := &models.User{
			Email:      gitlabUser.Email,
			IsVerified: true,
			GitLabID:   &gitlabUser.ID,
		}
		err := repository.CreateUser(newUser)
		if err != nil {
			return "", err
		}
		user = newUser
	}

	// Log the login attempt
	err = repository.CreateLoginRecord(user.ID, ipAddress)
	if err != nil {
		return "", err
	}

	jwtToken, err := utils.GenerateJWT(user.Email, cfg.JWTSecret)
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func HerokuLogin(cfg *config.Config, repository *db.Repository, code string, ipAddress string) (string, error) {
	oauthConfig := GetHerokuOAuthConfig(cfg)

	oauthToken, err := oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return "", err
	}

	client := oauthConfig.Client(context.Background(), oauthToken)

	// Get user info from Heroku API
	resp, err := client.Get("https://api.heroku.com/account")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Set required headers for Heroku API
	resp.Request.Header.Set("Accept", "application/vnd.heroku+json; version=3")

	var herokuUser struct {
		ID    string `json:"id"`
		Email string `json:"email"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&herokuUser); err != nil {
		return "", err
	}

	if herokuUser.Email == "" {
		return "", errors.New("failed to get email from Heroku response")
	}

	user, err := repository.GetUserByHerokuID(herokuUser.ID)
	if err != nil {
		return "", err
	}

	if user == nil {
		newUser := &models.User{
			Email:      herokuUser.Email,
			IsVerified: true,
			HerokuID:   &herokuUser.ID,
		}
		err := repository.CreateUser(newUser)
		if err != nil {
			return "", err
		}
		user = newUser
	}

	// Log the login attempt
	err = repository.CreateLoginRecord(user.ID, ipAddress)
	if err != nil {
		return "", err
	}

	jwtToken, err := utils.GenerateJWT(user.Email, cfg.JWTSecret)
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func InstagramLogin(cfg *config.Config, repository *db.Repository, code string, ipAddress string) (string, error) {
	oauthConfig := GetInstagramOAuthConfig(cfg)

	oauthToken, err := oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return "", err
	}

	// Get user info from Instagram API
	resp, err := http.Get(fmt.Sprintf("https://graph.instagram.com/me?fields=id,username&access_token=%s", oauthToken.AccessToken))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var instagramUser struct {
		ID       string `json:"id"`
		Username string `json:"username"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&instagramUser); err != nil {
		return "", err
	}

	if instagramUser.ID == "" {
		return "", errors.New("failed to get Instagram ID from response")
	}

	user, err := repository.GetUserByInstagramID(instagramUser.ID)
	if err != nil {
		return "", err
	}

	if user == nil {
		newUser := &models.User{
			Email:       instagramUser.Username + "@instagram.com",
			IsVerified:  true,
			InstagramID: &instagramUser.ID,
		}
		err := repository.CreateUser(newUser)
		if err != nil {
			return "", err
		}
		user = newUser
	}

	// Log the login attempt
	err = repository.CreateLoginRecord(user.ID, ipAddress)
	if err != nil {
		return "", err
	}

	jwtToken, err := utils.GenerateJWT(user.Email, cfg.JWTSecret)
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func JiraLogin(cfg *config.Config, repository *db.Repository, code string, ipAddress string) (string, error) {
	oauthConfig := GetJiraOAuthConfig(cfg)

	oauthToken, err := oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return "", err
	}

	client := oauthConfig.Client(context.Background(), oauthToken)

	// Get user info from Jira API
	resp, err := client.Get("https://api.atlassian.com/me")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var jiraUser struct {
		AccountID string `json:"account_id"`
		Email     string `json:"email"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&jiraUser); err != nil {
		return "", err
	}

	if jiraUser.Email == "" {
		return "", errors.New("failed to get email from Jira response")
	}

	user, err := repository.GetUserByJiraID(jiraUser.AccountID)
	if err != nil {
		return "", err
	}

	if user == nil {
		newUser := &models.User{
			Email:      jiraUser.Email,
			IsVerified: true,
			JiraID:     &jiraUser.AccountID,
		}
		err := repository.CreateUser(newUser)
		if err != nil {
			return "", err
		}
		user = newUser
	}

	// Log the login attempt
	err = repository.CreateLoginRecord(user.ID, ipAddress)
	if err != nil {
		return "", err
	}

	jwtToken, err := utils.GenerateJWT(user.Email, cfg.JWTSecret)
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func SlackLogin(cfg *config.Config, repository *db.Repository, code string, ipAddress string) (string, error) {
	oauthConfig := GetSlackOAuthConfig(cfg)

	oauthToken, err := oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return "", err
	}

	// Get user info from Slack API
	resp, err := http.Get(fmt.Sprintf("https://slack.com/api/users.identity?token=%s", oauthToken.AccessToken))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var slackResponse struct {
		OK   bool `json:"ok"`
		User struct {
			ID    string `json:"id"`
			Email string `json:"email"`
			Name  string `json:"name"`
		} `json:"user"`
		Team struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"team"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&slackResponse); err != nil {
		return "", err
	}

	if !slackResponse.OK {
		return "", errors.New("failed to get user info from Slack API")
	}

	if slackResponse.User.ID == "" {
		return "", errors.New("failed to get Slack user ID from response")
	}

	user, err := repository.GetUserBySlackID(slackResponse.User.ID)
	if err != nil {
		return "", err
	}

	if user == nil {
		newUser := &models.User{
			Email:      slackResponse.User.Email,
			IsVerified: true,
			SlackID:    &slackResponse.User.ID,
		}
		err := repository.CreateUser(newUser)
		if err != nil {
			return "", err
		}
		user = newUser
	}

	// Log the login attempt
	err = repository.CreateLoginRecord(user.ID, ipAddress)
	if err != nil {
		return "", err
	}

	jwtToken, err := utils.GenerateJWT(user.Email, cfg.JWTSecret)
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func SpotifyLogin(cfg *config.Config, repository *db.Repository, code string, ipAddress string) (string, error) {
	oauthConfig := GetSpotifyOAuthConfig(cfg)

	oauthToken, err := oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return "", err
	}

	// Get user info from Spotify API
	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/me", nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("Authorization", "Bearer "+oauthToken.AccessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var spotifyUser struct {
		ID          string `json:"id"`
		Email       string `json:"email"`
		DisplayName string `json:"display_name"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&spotifyUser); err != nil {
		return "", err
	}

	if spotifyUser.ID == "" {
		return "", errors.New("failed to get Spotify ID from response")
	}

	user, err := repository.GetUserBySpotifyID(spotifyUser.ID)
	if err != nil {
		return "", err
	}

	if user == nil {
		newUser := &models.User{
			Email:      spotifyUser.Email,
			IsVerified: true,
			SpotifyID:  &spotifyUser.ID,
		}
		err := repository.CreateUser(newUser)
		if err != nil {
			return "", err
		}
		user = newUser
	}

	// Log the login attempt
	err = repository.CreateLoginRecord(user.ID, ipAddress)
	if err != nil {
		return "", err
	}

	jwtToken, err := utils.GenerateJWT(user.Email, cfg.JWTSecret)
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func YahooLogin(cfg *config.Config, repository *db.Repository, code string, ipAddress string) (string, error) {
	oauthConfig := GetYahooOAuthConfig(cfg)

	oauthToken, err := oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return "", err
	}

	// Get user info from Yahoo API
	req, err := http.NewRequest("GET", "https://api.login.yahoo.com/openid/v1/userinfo", nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("Authorization", "Bearer "+oauthToken.AccessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var yahooUser struct {
		Sub           string `json:"sub"`
		Email         string `json:"email"`
		Name          string `json:"name"`
		FamilyName    string `json:"family_name"`
		GivenName     string `json:"given_name"`
		EmailVerified bool   `json:"email_verified"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&yahooUser); err != nil {
		return "", err
	}

	if yahooUser.Sub == "" {
		return "", errors.New("failed to get Yahoo ID from response")
	}

	user, err := repository.GetUserByYahooID(yahooUser.Sub)
	if err != nil {
		return "", err
	}

	if user == nil {
		newUser := &models.User{
			Email:      yahooUser.Email,
			IsVerified: yahooUser.EmailVerified,
			YahooID:    &yahooUser.Sub,
		}
		err := repository.CreateUser(newUser)
		if err != nil {
			return "", err
		}
		user = newUser
	}

	// Log the login attempt
	err = repository.CreateLoginRecord(user.ID, ipAddress)
	if err != nil {
		return "", err
	}

	jwtToken, err := utils.GenerateJWT(user.Email, cfg.JWTSecret)
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}
