package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"newsfeed/config"
	"newsfeed/handlers"
	"newsfeed/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	readEnv("../.env")
	dsn := fmt.Sprintf("host=localhost user=postgres password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Zurich",
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB_NAME"),
		os.Getenv("POSTGRES_PORT"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app := config.NewApplication(errorLog, infoLog, db)
	db.AutoMigrate(&models.Feed{}, &models.News{})

	data := models.NewModels(app)
	mux := handlers.RegisterRouts(app, data)

	log.Println("Start server on :4000")
	err = http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
