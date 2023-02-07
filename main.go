package main

import (
	"github.com/rodrinoblega/proper_images_processor/src/adapters"
	"github.com/rodrinoblega/proper_images_processor/src/configuration"
	"github.com/rodrinoblega/proper_images_processor/src/instrumentation"
	"time"
)

func main() {
	start := time.Now()

	service := adapters.NewImagesProcessorService(configuration.CreateImagesProcessorService())

	error := service.Execute(10)
	if error != nil {
		instrumentation.LogMessage("The application has the following error: " + error.Error())
	}

	end := time.Since(start)
	instrumentation.LogMessage("The application took " + end.String() + " to finish")

}
