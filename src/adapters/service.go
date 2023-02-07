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

func (imagesProcessorService *ImagesProcessorService) Execute(amountOfImagesRequested int) error {
	err := validateInputParams(amountOfImagesRequested)

	if err == nil {
		err = imagesProcessorService.ImagesProcessor.Execute(amountOfImagesRequested)
	}

	return err

}

func validateInputParams(amountOfImages int) error {
	if amountOfImages > 100 {
		return errors.New("amount of images should not be more than 100")
	}

	return nil
}
