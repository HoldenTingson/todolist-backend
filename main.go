package main

import (
	"database/sql"
	"todolist/task"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:FoscpuOervYyUzBuitTDuoKhSPmYBBcR@tcp(shuttle.proxy.rlwy.net:53166)/railway")
	if err != nil {
		panic(err)
	}

	repository := task.NewRepository(db)
	service := task.NewService(repository)
	handler := task.NewHandler(service)

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"https://thunderous-kringle-8a9eb1.netlify.app"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "X-Requested-With", "Content-Length"}
	config.AllowCredentials = true
	router.Use(cors.New(config))

	router.GET("/task", handler.GetTask)
	router.GET("/task/:id", handler.GetTaskById)
	router.POST("/task", handler.CreateTask)
	router.PATCH("/task/:id", handler.UpdateTask)
	router.DELETE("/task/:id", handler.DeleteTask)

	router.Run(":8080")
}
