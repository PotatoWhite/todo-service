package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"todo-service/pkg/todo"
)

type Config struct {
	Port   int    // port to listen on
	DBHost string // mysql host
	DBUser string // mysql user
	DBPass string // mysql password
}

// main entry point
func main() {
	// config
	cfg := Config{
		Port:   8080,
		DBHost: "localhost:3306",
		DBUser: "potato",
		DBPass: "1234",
	}

	// bootstrap
	db := mysql.New(mysql.Config{
		DSN: fmt.Sprintf("%s:%s@tcp(%s)/todolist?charset=utf8mb4&parseTime=True&loc=Local", cfg.DBUser, cfg.DBPass, cfg.DBHost),
	})

	gormDB, err := gorm.Open(db, &gorm.Config{})
	if err != nil {
		panic(err)
	}

	gormDB.AutoMigrate(&todo.TodoEntity{})

	// services & handlers
	todoService := todo.NewService(gormDB)
	todoHandler := todo.NewHandler(todoService)

	// routes
	r := gin.Default()
	todoHandler.RegisterRoute(r.Group("/todos"))

	// start server
	if err := r.Run(fmt.Sprintf(":%d", cfg.Port)); err != nil {
		panic(err)
	}
}
