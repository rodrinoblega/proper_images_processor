package configuration

import (
	"github.com/rodrinoblega/proper_images_processor/src/frameworks"
	"github.com/rodrinoblega/proper_images_processor/src/usecases"
	"net/http"
)

const URL = "https://icanhas.cheezburger.com/"
const DIRECTORY = "images/"

func CreateImagesProcessorService() *usecases.ImageProcessorUseCase {
	return usecases.NewImageProcessorUseCase(
		createImageFinderUseCase(),
		createImageDownloaderUseCase(),
	)
}

func createImageFinderUseCase() *usecases.ImageFinderUseCase {
	return usecases.NewImageFinderUseCase(frameworks.NewImageFinderClientImpl(URL, &http.Client{}))
}

func createImageDownloaderUseCase() *usecases.ImageDownloaderAndStorerUseCase {
	return usecases.NewImageDownloaderAndStorerUseCase(
		frameworks.NewImageDownloaderClientImpl(&http.Client{}),
		&frameworks.ImageStorerClientImpl{Directory: DIRECTORY},
	)
}
