package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/melardev/GoGinGonicApiGormCrud/controllers"
	"github.com/melardev/GoGinGonicApiGormCrud/infrastructure"
	"github.com/melardev/GoGinGonicApiGormCrud/models"
	"github.com/melardev/GoGinGonicApiGormCrud/seeds"
	"os"
)

func migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Todo{})
}

func main() {
	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
		os.Exit(0)
	}

	database := infrastructure.OpenDbConnection()
	defer database.Close()
	migrate(database)
	seeds.Seed(database)

	goGonicEngine := gin.Default()
	goGonicEngine.Use(cors.Default())
	goGonicEngine.GET("/api/todos", controllers.GetAllTodos).
		// GET("/api/todos/completed", controllers.GetAllPendingTodos).
		// GET("/api/todos/pending", controllers.GetAllCompletedTodos).
		GET("/api/todos/:id", controllers.GetTodoById)

	// This is how you should do it, the above was just to get started :)
	apiGroup := goGonicEngine.Group("/api")
	apiGroup.POST("/todos", controllers.CreateTodo)
	apiGroup.PUT("/todos/:id", controllers.UpdateTodo)
	apiGroup.PATCH("/todos/:id", controllers.CreateTodo)

	apiGroup.DELETE("/todos", controllers.DeleteAllTodos)
	apiGroup.DELETE("/todos/:id", controllers.DeleteTodo)

	goGonicEngine.Run(":8080")
}
