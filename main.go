package main

import (
    "github.com/boltdb/bolt"
    "github.com/docopt/docopt-go"

    "github.com/awonak/todo/service"
)

func main() {
    usage := `Todo - The App!

Usage:
  todo [-d <db>] [--database=<db>]
  todo -h | --help
  todo --version

Options:
  -h --help        Show this screen.
  --database=<db>  Path to database file [default: ./todo.db].`

    // Parse arguments
    arguments, _ := docopt.Parse(usage, nil, true, "Todo 1.0", false)
    dbpath := arguments["--database"].(string)

    // Load database
    db, err := bolt.Open(dbpath, 0600, nil)
    if err != nil {
        panic(err)
    }
    defer db.Close()

    // Fire up the service
    svc := service.TodoService{}
    svc.Run(db)
}
