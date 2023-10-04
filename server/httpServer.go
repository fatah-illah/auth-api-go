package server

import (
	"auth-api/controllers"
	"auth-api/repositories"
	"auth-api/services"
	"database/sql"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type HttpServer struct {
	config *viper.Viper
	router *gin.Engine
	// userController *controllers.UserController
	ControllersManager controllers.ControllersManager
}

func InitHttpServer(config *viper.Viper, dbHandler *sql.DB) HttpServer {
	jwtSecret := config.GetString("security.jwt_secret")                                     // Get JWT secret from the config
	repositoriesManager := repositories.NewRepositoriesManager(dbHandler, []byte(jwtSecret)) // Pass the JWT secret here

	servicesManager := services.NewServicesManager(repositoriesManager)

	controllersManager := controllers.NewControllersManager(servicesManager)

	router := InitRouter(controllersManager, jwtSecret)

	return HttpServer{
		config:             config,
		router:             router,
		ControllersManager: *controllersManager,
	}
}

// Running gin HttpServer
func (hs HttpServer) Start() {
	err := hs.router.Run(hs.config.GetString("http.server_address"))

	if err != nil {
		log.Fatalf("Error while starting HTTP Server: %v", err)
	}
}

func JWTAuthMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header not provided"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Next()
	}
}
