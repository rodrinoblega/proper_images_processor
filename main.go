package main

import (
	"flag"
	"github.com/rodrinoblega/proper_images_processor/src/adapters"
	"github.com/rodrinoblega/proper_images_processor/src/configuration"
	"github.com/rodrinoblega/proper_images_processor/src/instrumentation"
	"strconv"
	"time"
)

func main() {
	amountOfImagesRequested := flag.Int("amount", 1, "an int")
	amountOfThreads := flag.Int("threads", 1, "an int")
	flag.Parse()
	instrumentation.LogMessage("Amount of images to find: " + strconv.Itoa(*amountOfImagesRequested))
	instrumentation.LogMessage("Amount of threads to use: " + strconv.Itoa(*amountOfThreads))

	start := time.Now()

	service := adapters.NewImagesProcessorService(configuration.CreateImagesProcessorService())

	error := service.Execute(*amountOfImagesRequested, *amountOfThreads)
	if error != nil {
		instrumentation.LogMessage("The application has the following error: " + error.Error())
	}

	end := time.Since(start)
	instrumentation.LogMessage("The application took " + end.String() + " to finish")

}
