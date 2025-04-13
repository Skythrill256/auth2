package main

import (
	"log"
	"net/http"

	"github.com/Skythrill256/auth-service/internals/config"
	"github.com/Skythrill256/auth-service/internals/db"
	"github.com/Skythrill256/auth-service/internals/handlers"
	"github.com/Skythrill256/auth-service/internals/models"
	"github.com/Skythrill256/auth-service/internals/utils"
	"github.com/gorilla/mux"
)

func main() {
	cfg := config.LoadConfig()

	conn, err := db.Connect(cfg)
	if err != nil {
		log.Fatal("Error connecting to database")
	}
	defer conn.Close()
	err = models.RunMigrations(conn)
	if err != nil {
		log.Fatal("Error running migrations")
	}

	repository := db.NewRepository(conn)
	handler := handlers.NewHandler(repository, cfg)

	router := mux.NewRouter()
	router.HandleFunc("/signup", handler.SignUpUser).Methods("POST")
	router.HandleFunc("/login", handler.Login).Methods("POST")
	router.HandleFunc("/verify-email", handler.VerifyEmail).Methods("GET")
	router.HandleFunc("/forget-password", handler.ForgotPassword).Methods("GET")
	router.HandleFunc("/reset-password", handler.ResetPassword).Methods("GET", "POST")

	router.HandleFunc("/auth/google", handler.GoogleOAuthConsentRedirect).Methods("GET")
	router.HandleFunc("/auth/google/callback", handler.GoogleLogin).Methods("GET")

	router.HandleFunc("/auth/github", handler.GithubOAuthConsentRedirect).Methods("GET")
	router.HandleFunc("/auth/github/callback", handler.GithubLogin).Methods("GET")

	router.HandleFunc("/auth/facebook", handler.FacebookOAuthConsentRedirect).Methods("GET")
	router.HandleFunc("/auth/facebook/callback", handler.FacebookLogin).Methods("GET")

	router.HandleFunc("/auth/microsoft", handler.MicrosoftOAuthConsentRedirect).Methods("GET")
	router.HandleFunc("/auth/microsoft/callback", handler.MicrosoftLogin).Methods("GET")

	router.HandleFunc("/auth/linkedin", handler.LinkedinOAuthConsentRedirect).Methods("GET")
	router.HandleFunc("/auth/linkedin/callback", handler.LinkedinLogin).Methods("GET")

	router.HandleFunc("/auth/amazon", handler.AmazonOAuthConsentURL).Methods("GET")
	router.HandleFunc("/auth/amazon/callback", handler.AmazonLogin).Methods("GET")

	router.HandleFunc("/auth/bitbucket", handler.BitbucketOAuthConsentRedirect).Methods("GET")
	router.HandleFunc("/auth/bitbucket/callback", handler.BitbucketLogin).Methods("GET")

	router.HandleFunc("/auth/foursquare", handler.FoursquareOAuthConsentRedirect).Methods("GET")
	router.HandleFunc("/auth/foursquare/callback", handler.FoursquareLogin).Methods("GET")

	router.HandleFunc("/auth/gitlab", handler.GitLabOAuthConsentRedirect).Methods("GET")
	router.HandleFunc("/auth/gitlab/callback", handler.GitLabLogin).Methods("GET")

	router.HandleFunc("/auth/heroku", handler.HerokuOAuthConsentRedirect).Methods("GET")
	router.HandleFunc("/auth/heroku/callback", handler.HerokuLogin).Methods("GET")

	router.HandleFunc("/auth/instagram", handler.InstagramOAuthConsentRedirect).Methods("GET")
	router.HandleFunc("/auth/instagram/callback", handler.InstagramLogin).Methods("GET")

	router.HandleFunc("/auth/jira", handler.JiraOAuthConsentRedirect).Methods("GET")
	router.HandleFunc("/auth/jira/callback", handler.JiraLogin).Methods("GET")

	router.HandleFunc("/auth/slack", handler.SlackOAuthConsentRedirect).Methods("GET")
	router.HandleFunc("/auth/slack/callback", handler.SlackLogin).Methods("GET")

	router.HandleFunc("/auth/spotify", handler.SpotifyOAuthConsentRedirect).Methods("GET")
	router.HandleFunc("/auth/spotify/callback", handler.SpotifyLogin).Methods("GET")

	router.HandleFunc("/auth/yahoo", handler.YahooOAuthConsentRedirect).Methods("GET")
	router.HandleFunc("/auth/yahoo/callback", handler.YahooLogin).Methods("GET")

	router.HandleFunc("/get-user", handler.GetUserById).Methods("GET")

	// Profile management routes with authentication middleware
	profileRouter := router.PathPrefix("/profile").Subrouter()
	profileRouter.Use(utils.AuthMiddleware(cfg.JWTSecret))
	profileRouter.HandleFunc("", handler.GetProfile).Methods("GET")
	profileRouter.HandleFunc("", handler.UpdateProfile).Methods("PUT")

	// User Extra Info routes with authentication middleware
	extraInfoRouter := router.PathPrefix("/extra-info").Subrouter()
	extraInfoRouter.Use(utils.AuthMiddleware(cfg.JWTSecret))
	extraInfoRouter.HandleFunc("", handler.CreateUserExtraInfo).Methods("POST")
	extraInfoRouter.HandleFunc("", handler.GetUserExtraInfo).Methods("GET")
	extraInfoRouter.HandleFunc("", handler.UpdateUserExtraInfo).Methods("PUT")
	extraInfoRouter.HandleFunc("", handler.DeleteUserExtraInfo).Methods("DELETE")

	// Login History routes with authentication middleware
	loginHistoryRouter := router.PathPrefix("/login-history").Subrouter()
	loginHistoryRouter.Use(utils.AuthMiddleware(cfg.JWTSecret))
	loginHistoryRouter.HandleFunc("", handler.GetLoginHistory).Methods("GET")

	log.Println("Server is running on port", cfg.AppPort)
	log.Fatal(http.ListenAndServe(":"+cfg.AppPort, router))
}
