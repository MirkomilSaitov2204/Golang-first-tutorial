package config

import (
	"html/template"
	"log"
)

//AppConfig
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
}
