package usecases

import (
	"github.com/rodrinoblega/proper_images_processor/src/domain"
	"github.com/rodrinoblega/proper_images_processor/src/instrumentation"
	"strconv"
)

type ImageDownloaderClient interface {
	Download(url string) ([]byte, error)
}

type ImageStorerClient interface {
	Store(index int, bytes []byte)
}

type ImageDownloaderAndStorerUseCase struct {
	ImageDownloaderClient ImageDownloaderClient
	ImageStorerClient     ImageStorerClient
}

func NewImageDownloaderAndStorerUseCase(imageDownloaderClient ImageDownloaderClient, imageStorerClient ImageStorerClient) *ImageDownloaderAndStorerUseCase {
	return &ImageDownloaderAndStorerUseCase{
		ImageDownloaderClient: imageDownloaderClient,
		ImageStorerClient:     imageStorerClient,
	}
}

func (imageDownloaderUseCase *ImageDownloaderAndStorerUseCase) execute(image domain.Image, index int, wg *SemaphoredWaitGroup) {
	imageBytes, err := imageDownloaderUseCase.ImageDownloaderClient.Download(image.Url)
	if err != nil {
		instrumentation.LogMessage("There was an error downloading an image. URL: " + image.Url + ". Error: " + err.Error())
	}

	imageDownloaderUseCase.ImageStorerClient.Store(index, imageBytes)

	if err == nil {
		instrumentation.LogMessage("Success saving image " + strconv.Itoa(index+1))
	}

	wg.Done()
}
