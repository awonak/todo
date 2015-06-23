package main

import (
    "github.com/awonak/todo/service"
)

func main() {

    // not much here; it'll grow as we externalize config and add options
    svc := service.TodoService{}
    svc.Run()
}
