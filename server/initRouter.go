package server

import (
	"auth-api/controllers"

	"github.com/gin-gonic/gin"
)

func InitRouter(controllerMgr *controllers.ControllersManager, jwtSecret string) *gin.Engine {
	router := gin.Default()

	// Endpoint untuk pendaftaran dan masuk yang tidak memerlukan autentikasi JWT
	router.POST("/user/register", controllerMgr.UserController.CreateUser)
	router.POST("/user/login", controllerMgr.UserController.Login)

	// Grup rute yang memerlukan otentikasi JWT
	authRoute := router.Group("/user/auth") // Ubah basis URL untuk menghindari konflik dengan endpoint di atas
	authRoute.Use(JWTAuthMiddleware(jwtSecret))
	{
		// Contoh endpoint yang memerlukan autentikasi JWT
		// authRoute.GET("/profile", controllerMgr.UserController.GetProfile)

		// Menambahkan endpoint GetProtectedResource
		authRoute.GET("/resource/:resourceID", controllerMgr.UserController.GetProtectedResource)
	}

	return router
}
