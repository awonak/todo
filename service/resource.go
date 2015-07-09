package service

import (
    "encoding/json"
    "net/http"
    "strconv"
    "time"

    "github.com/boltdb/bolt"
    "github.com/gin-gonic/gin"
)

var bucket = []byte("todos")

type BoltModel interface {
    Key() []byte
    Value() []byte
}

type Todo struct {
    Id          uint64 `json:"id"`
    Created     int32  `json:"created"`
    Title       string `json:"title"`
    Description string `json:"description"`
    Done        bool   `json:"done"`
}

func (t *Todo) Key() []byte {
    return []byte(string(t.Id))
}

func (t *Todo) Value() []byte {
    jsonTodo, _ := json.Marshal(t)
    return []byte(jsonTodo)
}

type TodoResource struct {
    db *bolt.DB
}

func NewTodoResource(db *bolt.DB) *TodoResource {
    // Make sure our bucket exists
    db.Update(func(tx *bolt.Tx) error {
        tx.CreateBucketIfNotExists(bucket)
        return nil
    })
    return &TodoResource{db: db}
}

func (tr *TodoResource) List(c *gin.Context) {
    // List all todo records
    var todos []*Todo
    err := tr.db.View(func(tx *bolt.Tx) error {

        b := tx.Bucket(bucket)
        cur := b.Cursor()

        for k, v := cur.First(); k != nil; k, v = cur.Next() {
            var todo *Todo
            json.Unmarshal(v, &todo)
            todos = append(todos, todo)
        }

        return nil
    })

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, todos)
}

func (tr *TodoResource) Create(c *gin.Context) {
    // Parse the form data
    var todo Todo
    c.Bind(&todo)

    todo.Created = int32(time.Now().Unix())

    // Write to the database
    err := tr.db.Update(func(tx *bolt.Tx) error {

        b := tx.Bucket(bucket)

        if todo.Id == 0 {
            id, _ := b.NextSequence()
            todo.Id = id
        }

        return b.Put(todo.Key(), todo.Value())
    })

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Respond with the results
    c.JSON(http.StatusCreated, todo)
}

func (tr *TodoResource) Read(c *gin.Context) {
    // Parse the form data for todo key
    id, _ := strconv.Atoi(c.Param("id"))
    todo := Todo{Id: uint64(id)}

    // retrieve the data
    err := tr.db.View(func(tx *bolt.Tx) error {

        b := tx.Bucket(bucket)

        jsonTodo := b.Get(todo.Key())
        json.Unmarshal(jsonTodo, &todo)

        return nil
    })

    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
        return
    }

    c.JSON(http.StatusOK, todo)
}

func (tr *TodoResource) Update(c *gin.Context) {
    // Parse the form data
    var todo Todo
    c.Bind(&todo)

    // Write to the database
    err := tr.db.Update(func(tx *bolt.Tx) error {
        b := tx.Bucket(bucket)
        return b.Put(todo.Key(), todo.Value())
    })

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Respond with the results
    c.JSON(http.StatusCreated, todo)
}

func (tr *TodoResource) Delete(c *gin.Context) {
    // Parse the form data for todo key
    id, _ := strconv.Atoi(c.Param("id"))
    todo := Todo{Id: uint64(id)}

    // Delete the record
    err := tr.db.Update(func(tx *bolt.Tx) error {
        b := tx.Bucket(bucket)
        b.Delete(todo.Key())
        return nil
    })

    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.Data(http.StatusNoContent, "application/json", []byte(""))
}
