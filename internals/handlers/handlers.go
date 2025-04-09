package handlers

import (
	"bytes"
	"encoding/json"
	"html/template"
	"io"
	"net/http"
	"strconv"

	"github.com/Skythrill256/auth-service/internals/config"
	"github.com/Skythrill256/auth-service/internals/db"
	"github.com/Skythrill256/auth-service/internals/models"
	"github.com/Skythrill256/auth-service/internals/services"
	"github.com/Skythrill256/auth-service/internals/utils"
	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	Repository *db.Repository
	Config     *config.Config
}

func NewHandler(repository *db.Repository, config *config.Config) *Handler {
	return &Handler{
		Repository: repository,
		Config:     config,
	}
}

func (h *Handler) SignUpUser(w http.ResponseWriter, r *http.Request) {
	var user utils.UserDTO
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	err = services.SignUpUser(user, h.Repository, h.Config)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User Registered, Please verify your email"})
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var userDTO utils.UserDTO
	err := json.NewDecoder(r.Body).Decode(&userDTO)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	// Get the client's IP address
	ipAddress := r.Header.Get("X-Real-IP")
	if ipAddress == "" {
		ipAddress = r.Header.Get("X-Forwarded-For")
	}
	if ipAddress == "" {
		ipAddress = r.RemoteAddr
	}

	token, err := services.LoginUser(userDTO, h.Repository, h.Config, ipAddress)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func (h *Handler) VerifyEmail(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	if token == "" {
		http.Error(w, "Token is required", http.StatusBadRequest)
		return
	}

	err := services.VerifyEmail(token, h.Repository, h.Config)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Email Verified Successfully"})
}

func (h *Handler) GoogleOAuthConsentRedirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, services.GoogleOAuthConsentURL(h.Config), http.StatusTemporaryRedirect)
}

func (h *Handler) GithubOAuthConsentRedirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, services.GithubOAuthConsentURL(h.Config), http.StatusTemporaryRedirect)
}

func (h *Handler) FacebookOAuthConsentRedirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, services.FacebookOAuthConsentURL(h.Config), http.StatusTemporaryRedirect)
}

func (h *Handler) MicrosoftOAuthConsentRedirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, services.MicrosoftOAuthConsentURL(h.Config), http.StatusTemporaryRedirect)
}

func (h *Handler) LinkedinOAuthConsentRedirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, services.LinkedinOAuthConsentURL(h.Config), http.StatusTemporaryRedirect)
}

func (h *Handler) AmazonOAuthConsentURL(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, services.AmazonOAuthConsentURL(h.Config), http.StatusTemporaryRedirect)
}

func (h *Handler) AmazonLogin(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Code is required", http.StatusBadRequest)
		return
	}
	ipAddress := r.Header.Get("X-Real-IP")
	if ipAddress == "" {
		ipAddress = r.Header.Get("X-Forwarded-For")
	}
	if ipAddress == "" {
		ipAddress = r.RemoteAddr
	}
	token, err := services.AmazonLogin(h.Config, h.Repository, code, ipAddress)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func (h *Handler) GetProfile(w http.ResponseWriter, r *http.Request) {
	email, ok := utils.GetUserEmailFromContext(r.Context())
	if !ok || email == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	user, err := h.Repository.GetUserByEmail(email)
	userProfile, err := h.Repository.GetUserProfile(user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userProfile)
}

func (h *Handler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	email, ok := utils.GetUserEmailFromContext(r.Context())
	if !ok || email == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	user, err := h.Repository.GetUserByEmail(email)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	var req utils.UserProfileDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = services.UpdateUserProfile(user.ID, req.Name, req.Avatar, req.Bio, req.PhoneNumber, h.Repository)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Profile updated successfully"})
}

func (h *Handler) GoogleLogin(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Code is required", http.StatusBadRequest)
		return
	}
	ipAddress := r.Header.Get("X-Real-IP")
	if ipAddress == "" {
		ipAddress = r.Header.Get("X-Forwarded-For")
	}
	if ipAddress == "" {
		ipAddress = r.RemoteAddr
	}
	token, err := services.GoogleLogin(h.Config, h.Repository, code, ipAddress)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func (h *Handler) GithubLogin(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Code is required", http.StatusBadRequest)
		return
	}
	ipAddress := r.Header.Get("X-Real-IP")
	if ipAddress == "" {
		ipAddress = r.Header.Get("X-Forwarded-For")
	}
	if ipAddress == "" {
		ipAddress = r.RemoteAddr
	}
	token, err := services.GithubLogin(h.Config, h.Repository, code, ipAddress)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func (h *Handler) FacebookLogin(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Code is required", http.StatusBadRequest)
		return
	}
	ipAddress := r.Header.Get("X-Real-IP")
	if ipAddress == "" {
		ipAddress = r.Header.Get("X-Forwarded-For")
	}
	if ipAddress == "" {
		ipAddress = r.RemoteAddr
	}
	token, err := services.FacebookLogin(h.Config, h.Repository, code, ipAddress)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func (h *Handler) MicrosoftLogin(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Code is required", http.StatusBadRequest)
		return
	}
	ipAddress := r.Header.Get("X-Real-IP")
	if ipAddress == "" {
		ipAddress = r.Header.Get("X-Forwarded-For")
	}
	if ipAddress == "" {
		ipAddress = r.RemoteAddr
	}
	token, err := services.MicrosoftLogin(h.Config, h.Repository, code, ipAddress)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func (h *Handler) LinkedinLogin(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Code is required", http.StatusBadRequest)
		return
	}
	ipAddress := r.Header.Get("X-Real-IP")
	if ipAddress == "" {
		ipAddress = r.Header.Get("X-Forwarded-For")
	}
	if ipAddress == "" {
		ipAddress = r.RemoteAddr
	}
	token, err := services.LinkedinLogin(h.Config, h.Repository, code, ipAddress)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func (h *Handler) GetUserById(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Id is required", http.StatusBadRequest)
	}
	userId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	user, err := services.GetUserByID(userId, h.Repository)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (h *Handler) ForgotPassword(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	if email == "" {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}
	err := services.ForgotPassword(email, h.Repository, h.Config)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Reset Password link sent to your email"})
}

func (h *Handler) ResetPassword(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		token := r.URL.Query().Get("token")
		if token == "" {
			http.Error(w, "Token is required", http.StatusBadRequest)
			return
		}

		tmpl, err := template.New("reset").Parse(`
			<!DOCTYPE html>
			<html>
			<head>
				<title>Reset Password</title>
			</head>
			<body>
				<h2>Reset Password</h2>
				<form method="POST" action="/reset-password">
					<input type="hidden" name="token" value="{{.Token}}">
					<p><label>New Password:</label></p>
					<p><input type="password" name="password" required></p>
					<p><button type="submit">Reset Password</button></p>
				</form>
			</body>
			</html>
		`)
		if err != nil {
			http.Error(w, "Error loading form", http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, struct{ Token string }{Token: token})
	} else if r.Method == "POST" {
		token := r.FormValue("token")
		newPassword := r.FormValue("password")

		if token == "" || newPassword == "" {
			http.Error(w, "All fields are required", http.StatusBadRequest)
			return
		}

		email, err := utils.ParseJWT(token, h.Config.JWTSecret)
		if err != nil {
			http.Error(w, "Invalid or expired token", http.StatusBadRequest)
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Error hashing password", http.StatusInternalServerError)
			return
		}

		err = h.Repository.UpdateUserPassword(email, string(hashedPassword))
		if err != nil {
			http.Error(w, "Error updating password", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Password successfully reset"})
	}
}

func (h *Handler) CreateUserExtraInfo(w http.ResponseWriter, r *http.Request) {
	email, ok := utils.GetUserEmailFromContext(r.Context())
	if !ok || email == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	user, err := h.Repository.GetUserByEmail(email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Read the request body
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	// Try to decode as array first
	var reqArray []utils.UserExtraInfoDTO
	if err := json.NewDecoder(bytes.NewReader(bodyBytes)).Decode(&reqArray); err != nil {
		// If not an array, try single object
		var req utils.UserExtraInfoDTO
		if err := json.NewDecoder(bytes.NewReader(bodyBytes)).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		reqArray = []utils.UserExtraInfoDTO{req}
	}

	// Create slice to store all info objects
	infos := make([]*models.UserExtraInfo, len(reqArray))
	for i, req := range reqArray {
		infos[i] = &models.UserExtraInfo{
			UserID: user.ID,
			Key:    req.Key,
			Value:  req.Value,
		}
	}

	// Create all records in a transaction
	for _, info := range infos {
		if err := h.Repository.CreateUserExtraInfo(info); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(infos)
}

func (h *Handler) GetUserExtraInfo(w http.ResponseWriter, r *http.Request) {
	email, ok := utils.GetUserEmailFromContext(r.Context())
	if !ok || email == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	user, err := h.Repository.GetUserByEmail(email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	key := r.URL.Query().Get("key")
	if key == "" {
		// If no key is provided, return all extra info
		infoList, err := h.Repository.GetAllUserExtraInfo(user.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(infoList)
		return
	}

	// If key is provided, return specific info
	info, err := h.Repository.GetUserExtraInfo(user.ID, key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if info == nil {
		http.Error(w, "Info not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
}

func (h *Handler) UpdateUserExtraInfo(w http.ResponseWriter, r *http.Request) {
	email, ok := utils.GetUserEmailFromContext(r.Context())
	if !ok || email == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	user, err := h.Repository.GetUserByEmail(email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var req utils.UserExtraInfoDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	info := &models.UserExtraInfo{
		UserID: user.ID,
		Key:    req.Key,
		Value:  req.Value,
	}

	if err := h.Repository.UpdateUserExtraInfo(info); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
}

func (h *Handler) DeleteUserExtraInfo(w http.ResponseWriter, r *http.Request) {
	email, ok := utils.GetUserEmailFromContext(r.Context())
	if !ok || email == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	user, err := h.Repository.GetUserByEmail(email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "Key is required", http.StatusBadRequest)
		return
	}

	if err := h.Repository.DeleteUserExtraInfo(user.ID, key); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) GetLoginHistory(w http.ResponseWriter, r *http.Request) {
	email, ok := utils.GetUserEmailFromContext(r.Context())
	if !ok || email == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	user, err := h.Repository.GetUserByEmail(email)
	if err != nil {
		http.Error(w, "Failed to get user", http.StatusInternalServerError)
		return
	}

	loginRecords, err := h.Repository.GetUserLoginHistory(user.ID)
	if err != nil {
		http.Error(w, "Failed to get login history", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(loginRecords); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) BitbucketLogin(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Code is required", http.StatusBadRequest)
		return
	}

	ipAddress := r.Header.Get("X-Real-IP")
	if ipAddress == "" {
		ipAddress = r.Header.Get("X-Forwarded-For")
	}
	if ipAddress == "" {
		ipAddress = r.RemoteAddr
	}

	token, err := services.BitbucketLogin(h.Config, h.Repository, code, ipAddress)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
