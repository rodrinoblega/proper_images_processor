package adapters

import (
	"errors"
	"github.com/rodrinoblega/proper_images_processor/src/usecases"
)

type ImagesProcessorService struct {
	ImagesProcessor *usecases.ImageProcessorUseCase
}

func NewImagesProcessorService(imageProcessorUseCase *usecases.ImageProcessorUseCase) *ImagesProcessorService {
	return &ImagesProcessorService{
		ImagesProcessor: imageProcessorUseCase,
	}
}

func (imagesProcessorService *ImagesProcessorService) Execute(amountOfImagesRequested, amountOfThreads int) error {
	err := validateInputParams(amountOfImagesRequested, amountOfThreads)

	if err == nil {
		err = imagesProcessorService.ImagesProcessor.Execute(amountOfImagesRequested, amountOfThreads)
	}

	return err

}

func validateInputParams(amountOfImages int, amountOfThreads int) error {
	if amountOfImages > 100 {
		return errors.New("amount of images should not be more than 100")
	}

	if amountOfThreads < 1 || amountOfThreads > 5 {
		return errors.New("amount of threads should be between 1 and 5")
	}

	return nil
}
