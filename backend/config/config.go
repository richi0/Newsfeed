package config

import (
	"log"

	"gorm.io/gorm"
)

type Application struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	DB       *gorm.DB
}

func NewApplication(errorLog *log.Logger, infoLog *log.Logger, db *gorm.DB) *Application {
	return &Application{errorLog, infoLog, db}
}
