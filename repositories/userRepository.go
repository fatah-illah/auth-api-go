package repositories

import (
	"auth-api/models"
	dbcontext "auth-api/repositories/dbContext"
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	dbHandler *sql.DB
	// transaction field can be removed if not used, or implement it appropriately
	jwtSecret []byte
}

func NewUserRepository(dbHandler *sql.DB, jwtSecret []byte) *UserRepository {
	return &UserRepository{
		dbHandler: dbHandler,
		jwtSecret: jwtSecret,
	}
}

func (ur UserRepository) CreateUser(ctx *gin.Context, user *models.User) (*models.User, *models.ResponseError) {
	store := dbcontext.New(ur.dbHandler)

	newUserID, err := store.CreateUser(ctx, dbcontext.CreateUserParams{
		Username:     user.Username,
		PasswordHash: user.PasswordHash,
		Email:        user.Email,
	})

	if err != nil {
		// You might want to add specific error handling based on the type of error
		// (e.g., duplicate username or invalid data)
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	user.ID = newUserID
	return user, nil
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (ur *UserRepository) generateToken() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() // Token akan kedaluwarsa setelah 1 jam

	tokenString, err := token.SignedString(ur.jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (ur UserRepository) AuthenticateUser(ctx *gin.Context, username, password string) (string, error) {
	store := dbcontext.New(ur.dbHandler)

	// Ambil user berdasarkan username dari database
	user, err := store.GetUserByUsername(ctx, username)
	if err != nil {
		return "", err
	}

	// Verifikasi password dengan yang ada di database
	// Anda perlu menambahkan logika untuk memeriksa hash dari password
	// Misalnya menggunakan bcrypt atau algoritma hashing lainnya
	isValidPassword := checkPasswordHash(password, user.PasswordHash)
	if !isValidPassword {
		return "", errors.New("incorrect password")
	}

	// Jika password valid, buat token akses untuk pengguna
	token, err := ur.generateToken()
	if err != nil {
		return "", err
	}

	tokenParams := dbcontext.CreateTokenForUserParams{
		UserID:      sql.NullInt32{Int32: user.ID, Valid: true},
		AccessToken: token,
		ExpiresAt:   sql.NullTime{Time: time.Now().Add(time.Hour * 24), Valid: true},
	}

	_, err = store.CreateTokenForUser(ctx, tokenParams)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (ur *UserRepository) GetProtectedResource(ctx *gin.Context, resourceID int32) (models.ProtectedResource, error) {
	store := dbcontext.New(ur.dbHandler)
	return store.GetProtectedResource(ctx, resourceID)
}

func (ur *UserRepository) VerifyToken(ctx *gin.Context, token string) (bool, error) {
	store := dbcontext.New(ur.dbHandler) // Membuat store berdasarkan dbHandler

	userID, err := store.VerifyAccessToken(ctx, token)
	if err != nil {
		// Jika terjadi error saat query, misalnya karena token tidak ditemukan atau sudah kedaluwarsa.
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	// Jika user ID valid dan tidak null, berarti token valid.
	return userID.Valid, nil
}
