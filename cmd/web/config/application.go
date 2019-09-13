package config

import (
	"log"
)

// Application configuration
type Application struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
}
