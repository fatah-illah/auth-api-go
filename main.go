package main

import (
	"auth-api/config"
	"auth-api/server"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	log.Println("Starting User Authentication REST API!")

	log.Println("Initializing Configuration!")
	config := config.InitConfig(getConfigFileName())

	// jwtSecret := config.GetString("security.jwt_secret")
	// repoManager := repositories.NewRepositoriesManager(dbHandler, []byte(jwtSecret))

	log.Println("Initializing DataBase!")
	dbHandler := server.InitDatabase(config)
	log.Println(dbHandler)

	// // Test insert to users, using goroutine
	// ctx := context.Background()
	// queries := repositories.New(dbHandler)

	// newUser, err := queries.CreateUser(ctx,
	// 	repositories.CreateUserParams{
	// 		Username:     "ucum",
	// 		PasswordHash: "hashed_password_5",
	// 		Email:        "ucum@gmail.com",
	// 	},
	// )

	// if err != nil {
	// 	log.Fatal("Error: ", err)
	// }

	// log.Println(newUser)

	log.Println("Initializing HTTP Server!")
	httpServer := server.InitHttpServer(config, dbHandler)

	httpServer.Start()
}

func getConfigFileName() string {
	env := os.Getenv("ENV")

	if env != "" {
		return "config" + env
	}
	// return string-nya nama file toml-nya
	return "config"
}
