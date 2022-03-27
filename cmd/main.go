package main

import (
	"log"
	"net/http"
	"time"

	"github.com/AndrewMislyuk/CRUD-languages/internal/repository/psql"
	"github.com/AndrewMislyuk/CRUD-languages/internal/service"
	"github.com/AndrewMislyuk/CRUD-languages/internal/transport/rest"
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

	languagesRepo := psql.NewLanguages(db)
	languagesService := service.NewLanguages(languagesRepo)
	handler := rest.NewHandler(languagesService)

	handler.InitRouter()

	log.Println("SERVER STARTED AT", time.Now().Format(time.RFC3339))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
