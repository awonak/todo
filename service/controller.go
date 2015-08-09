package service

import (
    "github.com/boltdb/bolt"
    "github.com/gin-gonic/gin"
)

type TodoService struct{}

func (tr *TodoService) Run(db *bolt.DB) {

    // initialize the resource and inject our db connection
    todoResource := NewTodoResource(db)

    router := gin.Default()

    // Static files
    router.Static("/fonts", "./client/dist/fonts")
    router.Static("/images", "./client/dist/images")
    router.Static("/scripts", "./client/dist/scripts")
    router.Static("/styles", "./client/dist/styles")
    router.StaticFile("/favicon.ico", "./client/app/favicon.ico")
    router.StaticFile("/", "./client/dist/index.html")

    // Simple group: api
    api := router.Group("/api")
    {
        api.GET("/tasks", todoResource.List)
        api.POST("/tasks", todoResource.Create)
        api.GET("/tasks/:id", todoResource.Read)
        api.POST("/tasks/:id", todoResource.Update)
        api.DELETE("/tasks/:id", todoResource.Delete)
    }

    router.Run(":8080")
}
