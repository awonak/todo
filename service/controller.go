package service

import (
    "github.com/gin-gonic/gin"
)

type TodoService struct{}

func (tr *TodoService) Run() {

    // initialize the resource and inject our db connection
    todoResource := &TodoResource{db: Db{}}

    router := gin.Default()

    // Static files
    // router.Static("/assets", "./assets")
    // router.StaticFile("/favicon.ico", "./resources/favicon.ico")

    // Simple group: api
    api := router.Group("/api")
    {
        api.GET("/tasks", todoResource.List)
        api.POST("/tasks", todoResource.Create)
        api.GET("/tasks/:id", todoResource.Read)
        // api.POST("/tasks/:id", todoResource.Update)
        api.DELETE("/tasks/:id", todoResource.Delete)
    }

    router.Run(":8080")
}
