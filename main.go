package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/krhone/go-quest/controllers"
	"github.com/krhone/go-quest/models"
)

const (
	defaultPort      = "8008"
	idleTimeout      = 30 * time.Second
	writeTimeout     = 180 * time.Second
	readHeaderimeout = 10 * time.Second
	readTimeout      = 10 * time.Second
)

// @title Sample Quest API
// @version 1.0
// @description This is a module Quest
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8008
// @BasePath /api
func main() {
	err := godotenv.Load()
	check(err)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	handler := controllers.New()

	server := &http.Server{
		Addr:              "0.0.0.0:" + port,
		Handler:           handler,
		IdleTimeout:       idleTimeout,
		WriteTimeout:      writeTimeout,
		ReadHeaderTimeout: readHeaderimeout,
		ReadTimeout:       readTimeout,
	}

	models.ConnectDatabase()

	err = server.ListenAndServe()

	check(err)
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
