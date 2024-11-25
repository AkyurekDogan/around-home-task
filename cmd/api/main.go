/*
The main package for the service
*/
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/AkyurekDogan/around-home-task/internal/app/domain"
	"github.com/AkyurekDogan/around-home-task/internal/app/handler"
	"github.com/AkyurekDogan/around-home-task/internal/app/infrastructure/drivers"
	"github.com/AkyurekDogan/around-home-task/internal/app/infrastructure/middlewares"
	"github.com/AkyurekDogan/around-home-task/internal/app/infrastructure/repository"
	"github.com/AkyurekDogan/around-home-task/internal/app/service"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
)

const (
	// ENV environment file path
	ENV = ".env"
	//ENV_CNF_PATH config path
	ENV_CNF_PATH = "CONFIG_PATH"
)

// main entry point
func main() {
	// load environment variables
	err := godotenv.Load(ENV)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	// use environment variable to get the config path
	appEnvConfigPath := os.Getenv(ENV_CNF_PATH)
	if appEnvConfigPath == "" {
		log.Fatalf("%s environment variable must be set", ENV_CNF_PATH)
	}
	// unmarshall the config file and get all settings in the configuration file.
	yamlFile, err := os.ReadFile(appEnvConfigPath)
	if err != nil {
		log.Fatalf("Error reading configuration YAML file: %v", err)
	}
	var config domain.Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("Error unmarshalling YAML file: %v", err)
	}

	// initialize db connector

	// read driver
	dbDriverRead := drivers.NewPostgres(
		config.Database.Username,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Database,
	)

	// initialize repository
	dbRepoPartnerSkill := repository.NewPartnerSkill(dbDriverRead)
	dbRepoPartnerRating := repository.NewPartnerRating(dbDriverRead)
	dbRepoPartner := repository.NewPartner(dbDriverRead)
	dbRepoMatch := repository.NewMatch(dbDriverRead)

	// initialize services
	srvPartnerRating := service.NewPartnerRating(dbRepoPartnerRating)
	srvPartnerSkill := service.NewPartnerSkill(dbRepoPartnerSkill)
	srvPartner := service.NewPriceService(dbRepoPartner, srvPartnerSkill, srvPartnerRating)
	matchService := service.NewMatch(dbRepoMatch)

	// handlers
	handlerMatch := handler.NewMatch(matchService)
	handlerPartner := handler.NewPartner(srvPartner)
	// Create a new router
	r := chi.NewRouter()
	r.Use(middlewares.AddHeaderMiddleware())
	// Define the endpoints
	r.Get("/match", handlerMatch.Get)
	r.Get("/partner", handlerPartner.Get)
	// Start the HTTP server
	err = http.ListenAndServe(config.Server.Host, r)
	if err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
