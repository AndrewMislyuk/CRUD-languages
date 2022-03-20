package main

import (
	"log"

	"github.com/AndrewMislyuk/CRUD-languages/pkg/database"
	_ "github.com/lib/pq"
)

func main() {
	db, err := database.NewPostgresConnection(database.ConnectionInfo{
		Host:     "localhost",
		Port:     5432,
		Username: "root",
		DBName:   "crud-languages",
		SSLMode:  "disable",
		Password: "root",
	})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
