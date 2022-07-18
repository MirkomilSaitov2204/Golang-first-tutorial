package config

import "html/template"

//AppConfig
type AppConfig struct {
	TemplateCache map[string]*template.Template
}
