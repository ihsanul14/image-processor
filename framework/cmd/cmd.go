package cmd

import (
	"image-processor/application"
	"image-processor/entity/processor"
	"image-processor/framework/router"
	"image-processor/usecase"
)

func Run() {
	entity := processor.NewMemory()
	uc := usecase.NewUsecase(entity)
	app := application.NewApplication(uc)
	router := router.NewRouter(app)
	router.Run()
}
