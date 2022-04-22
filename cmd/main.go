package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/AndrewMislyuk/CRUD-languages/internal/config"
	"github.com/AndrewMislyuk/CRUD-languages/internal/repository/psql"
	"github.com/AndrewMislyuk/CRUD-languages/internal/service"
	"github.com/AndrewMislyuk/CRUD-languages/internal/transport/rest"
	"github.com/AndrewMislyuk/CRUD-languages/pkg/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

const (
	CONFIG_DIR  = "configs"
	CONFIG_FILE = "main"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)
}

// @title CRUD API Languages
// @version 1.0
// @description API about programming languages

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal(err)
	}

	cfg, err := config.New(CONFIG_DIR, CONFIG_FILE)
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Printf("%+v", cfg)

	db, err := database.NewPostgresConnection(database.ConnectionInfo{
		Host:     cfg.DB.Host,
		Port:     cfg.DB.Port,
		Username: cfg.DB.Username,
		DBName:   cfg.DB.Name,
		SSLMode:  cfg.DB.SSLMode,
		Password: cfg.DB.Password,
	})
	if err != nil {
		logrus.Fatal(err)
	}
	defer db.Close()

	languagesRepo := psql.NewRepository(db)
	languagesService := service.NewService(languagesRepo)
	handler := rest.NewHandler(languagesService)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
		Handler: handler.InitRouter(),
	}

	logrus.Infoln("Server has been running...")

	if err := srv.ListenAndServe(); err != nil {
		logrus.Fatal(err)
	}
}
