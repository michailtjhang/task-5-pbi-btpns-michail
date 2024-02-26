package main

import (
	"go-photoUser/configs"
	"go-photoUser/handlers"
	"go-photoUser/middlewares"
	"go-photoUser/migrations"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := configs.InitDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	migrations.Migrate(db)
	// seeders.Seed(db)

	router := gin.Default()

	// Rute CRUD Photo
	router.GET("/products", middlewares.AuthMiddleware(), handlers.ListProducts(db))
	router.GET("/products/:id", handlers.GetPhoto(db))
	router.POST("/products", handlers.CreatePhoto(db))
	router.PUT("/products/:id", handlers.UpdatePhoto(db))
	router.DELETE("/products/:id", handlers.DeletePhoto(db))

	// Rute Login dan Register User
	router.POST("/login", handlers.Login(db))
	router.POST("/register", handlers.Register(db))
	router.PUT("/user/:id", handlers.UpdateUsername(db))
	router.DELETE("/user/:id", handlers.DeleteUser(db))

	// jalankan server
	router.Run(":5000")

}
