package main

import (
	"go-photoUser/app"
	"go-photoUser/controller"
	"go-photoUser/middlewares"
	"go-photoUser/migrations"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := app.InitDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	migrations.Migrate(db)
	// seeders.Seed(db)

	router := gin.Default()

	// Rute CRUD Photo
	router.GET("/products", middlewares.AuthMiddleware(), controller.ListProducts(db))
	router.GET("/products/:id", middlewares.AuthMiddleware(), controller.GetPhoto(db))
	router.POST("/products", middlewares.AuthMiddleware(), controller.CreatePhoto(db))
	router.PUT("/products/:id", middlewares.AuthMiddleware(), controller.UpdatePhoto(db))
	router.DELETE("/products/:id", middlewares.AuthMiddleware(), controller.DeletePhoto(db))

	// Rute Login dan Register User
	router.POST("/login", controller.Login(db))
	router.POST("/register", controller.Register(db))
	router.PUT("/user/:id", controller.UpdateUsername(db))
	router.DELETE("/user/:id", controller.DeleteUser(db))

	// jalankan server
	router.Run(":5000")

}
