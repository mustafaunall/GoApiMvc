package main

import (
	"LearnGo/config"
	"LearnGo/helper"
	"LearnGo/model"
	"LearnGo/router"
	"github.com/rs/zerolog/log"
	"net/http"
)

// @title Test Api
// @version 1.0
// @description Api Description
// @host localhost:8080
// @BasePath /api/v1
func main() {
	log.Info().Msg("Started Server!")
	// Database
	db := config.DatabaseConnection()

	err := db.Table("products").AutoMigrate(&model.Product{})
	if err != nil {
		return
	}

	// Router
	route := router.NewRouter(db)

	server := &http.Server{
		Addr:    ":8080",
		Handler: route,
	}

	serverErr := server.ListenAndServe()
	helper.ErrorPanic(serverErr)
}
