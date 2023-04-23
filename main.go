package main

import (
	"log"

	"github.com/luis501angel/go-course/db"
	"github.com/luis501angel/go-course/handlers"
)

func main() {
	if !db.CheckConnection() {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Handlers()
}
