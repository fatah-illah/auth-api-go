package controllers

import (
	"auth-api/models"
	"auth-api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *services.UserService
}

// declare constructor
func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

// method
func (uc UserController) CreateUser(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, responseErr := uc.userService.CreateUser(ctx, &user)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User successfully created", "user": response})
}

func (uc UserController) Login(ctx *gin.Context) {
	// Membuat struct khusus untuk menerima request login
	var loginRequest struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, responseErr := uc.userService.Login(ctx, loginRequest.Username, loginRequest.Password)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}

func (uc UserController) GetProtectedResource(ctx *gin.Context) {
	// Mengambil resourceID dari parameter URL.
	resourceIDStr := ctx.Param("resourceID")
	resourceID, err := strconv.ParseInt(resourceIDStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid resource ID"})
		return
	}

	// Memeriksa token dari header untuk otentikasi.
	accessToken := ctx.GetHeader("Authorization")
	if accessToken == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Access token required"})
		return
	}

	// Memanggil fungsi GetProtectedResource dengan parameter ctx, accessToken, dan resourceID.
	resource, responseErr := uc.userService.GetProtectedResource(ctx, accessToken, int32(resourceID))
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"resource": resource})
}
