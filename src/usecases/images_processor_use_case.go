package usecases

import (
	"github.com/rodrinoblega/proper_images_processor/src/domain"
	"github.com/rodrinoblega/proper_images_processor/src/instrumentation"
)

type ImageProcessorUseCase struct {
	ImagesFinderUseCase     *ImageFinderUseCase
	ImagesDownloaderUseCase *ImageDownloaderAndStorerUseCase
}

func NewImageProcessorUseCase(imagesFinderUseCase *ImageFinderUseCase, imagesDownloaderUseCase *ImageDownloaderAndStorerUseCase) *ImageProcessorUseCase {
	return &ImageProcessorUseCase{
		ImagesFinderUseCase:     imagesFinderUseCase,
		ImagesDownloaderUseCase: imagesDownloaderUseCase,
	}
}

func (imageProcessor *ImageProcessorUseCase) Execute(amountOfImagesRequested, amountOfThreads int) error {
	var imagesUrlAccumulated []domain.Image

	instrumentation.LogMessage("Finding images...")
	imageUrlFound, err := imageProcessor.ImagesFinderUseCase.execute(imagesUrlAccumulated, amountOfImagesRequested)
	if err != nil {
		return err
	}

	instrumentation.LogMessage("Downloading images...")
	wg := &SemaphoredWaitGroup{Sem: make(chan bool, amountOfThreads)}
	for index, image := range imageUrlFound {
		wg.Add(1)
		go imageProcessor.ImagesDownloaderUseCase.execute(image, index, wg)
	}
	wg.Wait()

	instrumentation.LogMessage("Done!")
	return err
}
