package main

import (
	"log"
	"patungan/handler"
	"patungan/user"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:root@tcp(127.0.0.1:3306)/patungan?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)

	router.Run(":8888")
	
	// userInput := user.RegisterUserInput{}
	// userInput.Name = "Tes simpan dari service"
	// userInput.Email = "contoh@gmil.com"
	// userInput.Occupation = "Student"
	// userInput.Password = "password"

	// userService.RegisterUser(userInput)

}
// input dari user
// handler, mapping input dari user -> struct input
// service : melakukan mapping dari struct input ke struct User
// repository
// db
