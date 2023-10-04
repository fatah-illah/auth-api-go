package services

import (
	"auth-api/models"
	"auth-api/repositories"
	"database/sql"
	"errors"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func isValidPassword(password string) bool {
	if len(password) < 8 {
		return false
	}
	containsUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	containsLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	containsNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
	containsSpecial := regexp.MustCompile(`[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>?~]`).MatchString(password)

	return containsUpper && containsLower && containsNumber && containsSpecial
}

func (us UserService) CreateUser(ctx *gin.Context, user *models.User) (*models.User, *models.ResponseError) {
	if strings.TrimSpace(user.Username) == "" {
		return nil, &models.ResponseError{
			Message: "Username cannot be empty",
			Status:  400, // Bad Request
		}
	}

	if !isValidEmail(user.Email) {
		return nil, &models.ResponseError{
			Message: "Invalid email format",
			Status:  400, // Bad Request
		}
	}

	if !isValidPassword(user.PasswordHash) {
		return nil, &models.ResponseError{
			Message: "Password must be at least 8 characters long and include at least 1 uppercase, 1 lowercase, 1 number, and 1 special character",
			Status:  400, // Bad Request
		}
	}

	// Hash the password before passing to the repository
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return nil, &models.ResponseError{
			Message: "Error hashing password",
			Status:  500, // Internal Server Error
		}
	}
	user.PasswordHash = string(hashedPassword)

	return us.userRepository.CreateUser(ctx, user)
}

func (us UserService) Login(ctx *gin.Context, username, password string) (string, *models.ResponseError) {
	// Verifikasi kredensial pengguna dengan repository
	token, err := us.userRepository.AuthenticateUser(ctx, username, password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// Jika error karena user tidak ditemukan
			return "", &models.ResponseError{
				Message: "Invalid username or password",
				Status:  401, // Unauthorized
			}
		} else {
			// Jika error lain dari database atau proses otentikasi
			return "", &models.ResponseError{
				Message: "Internal server error",
				Status:  500, // Internal Server Error
			}
		}
	}
	return token, nil
}

func (us UserService) GetProtectedResource(ctx *gin.Context, token string, resourceID int32) (models.ProtectedResource, *models.ResponseError) {
	// Memeriksa validitas token.
	isValid, tokenErr := us.userRepository.VerifyToken(ctx, token)
	if !isValid || tokenErr != nil {
		return models.ProtectedResource{}, &models.ResponseError{
			Message: "Invalid or expired token",
			Status:  401, // Unauthorized
		}
	}

	resource, err := us.userRepository.GetProtectedResource(ctx, resourceID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.ProtectedResource{}, &models.ResponseError{
				Message: "Resource not found",
				Status:  404, // Not Found
			}
		}
		return models.ProtectedResource{}, &models.ResponseError{
			Message: "Internal server error: " + err.Error(),
			Status:  500, // Internal Server Error
		}
	}
	return resource, nil
}
