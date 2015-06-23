package service

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
    "time"
)

type Db map[int32]*Todo

type Todo struct {
    Id          int32  `json:"id"`
    Created     int32  `json:"created"`
    Title       string `json:"title"`
    Description string `json:"description"`
    Done        bool   `json:"done"`
}

type TodoResource struct {
    db      Db    `json:"db"`
    counter int32 `json:"counter"`
}

func (tr *TodoResource) List(c *gin.Context) {
    output := make([]*Todo, 0)
    for _, todo := range tr.db {
        output = append(output, todo)
    }
    c.JSON(http.StatusOK, output)
}

func (tr *TodoResource) Create(c *gin.Context) {
    var todo Todo

    c.Bind(&todo)
    tr.counter += 1 // very not thread safe
    todo.Id = tr.counter
    todo.Created = int32(time.Now().Unix())
    tr.db[tr.counter] = &todo
    c.JSON(http.StatusCreated, todo)
}

func (tr *TodoResource) Read(c *gin.Context) {
    idStr := c.Params.ByName("id")
    idInt, _ := strconv.Atoi(idStr)
    id := int32(idInt)

    if todo, ok := tr.db[id]; !ok {
        c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
    } else {
        c.JSON(http.StatusOK, todo)
    }
}

func (tr *TodoResource) Delete(c *gin.Context) {
    idStr := c.Params.ByName("id")
    idInt, _ := strconv.Atoi(idStr)
    id := int32(idInt)

    if todo, ok := tr.db[id]; !ok {
        c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
    } else {
        delete(tr.db, todo.Id)
        c.Data(http.StatusNoContent, "application/json", make([]byte, 0))
    }
}
