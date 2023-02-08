package configuration

import (
	"github.com/rodrinoblega/proper_images_processor/src/frameworks"
	"github.com/rodrinoblega/proper_images_processor/src/usecases"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestConfiguration_ProdConfiguration(t *testing.T) {
	//when
	conf := CreateImagesProcessorService()

	//then
	assert.Equal(t, createImageProcessorUseCaseExpected(), conf)
}

func createImageProcessorUseCaseExpected() *usecases.ImageProcessorUseCase {
	return &usecases.ImageProcessorUseCase{
		ImagesFinderUseCase: &usecases.ImageFinderUseCase{
			ImageFinderClient: &frameworks.ImageFinderClientImpl{Url: "https://icanhas.cheezburger.com/", Client: &http.Client{}},
		},
		ImagesDownloaderUseCase: &usecases.ImageDownloaderAndStorerUseCase{
			ImageDownloaderClient: &frameworks.ImageDownloaderClientImpl{Client: &http.Client{}},
			ImageStorerClient:     &frameworks.ImageStorerClientImpl{Directory: "images/"},
		},
	}
}
