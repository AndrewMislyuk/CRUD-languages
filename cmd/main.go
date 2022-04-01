package main

import (
	"net/http"

	"github.com/AndrewMislyuk/CRUD-languages/internal/repository/psql"
	"github.com/AndrewMislyuk/CRUD-languages/internal/service"
	"github.com/AndrewMislyuk/CRUD-languages/internal/transport/rest"
	"github.com/AndrewMislyuk/CRUD-languages/pkg/database"
	"github.com/sirupsen/logrus"
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
		logrus.Fatal(err)
	}
	defer db.Close()

	languagesRepo := psql.NewLanguages(db)
	languagesService := service.NewLanguages(languagesRepo)
	handler := rest.NewHandler(languagesService)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: handler.InitRouter(),
	}

	logrus.Infoln("Server has been running...")

	if err := srv.ListenAndServe(); err != nil {
		logrus.Fatal(err)
	}
}
