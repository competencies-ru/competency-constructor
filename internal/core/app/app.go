package app

import "github.com/competencies-ru/competency-constructor/internal/core/app/service"

type (
	Application struct {
		Services Services
	}

	Services struct {
		UgsnService service.UgsnHandler
	}
)
