package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/rodrinoblega/proper_images_processor/src/adapters"
	"github.com/rodrinoblega/proper_images_processor/src/frameworks"
	"github.com/rodrinoblega/proper_images_processor/src/usecases"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

const TestDirectory = "images_test/"

func TestFunctionalImagesProcessor_ImagesFoundGreaterThanRequestedImages(t *testing.T) {
	//SET UP
	os.RemoveAll(TestDirectory)
	os.Mkdir(TestDirectory, 0755)
	amountOfImages := 30

	//when
	service := adapters.NewImagesProcessorService(NewImagesProcessorUseCaseWithMockedClients("image_finder_mock.html"))
	err := service.Execute(amountOfImages)

	//then
	files, _ := os.ReadDir(TestDirectory)
	var count int
	for _, _ = range files {
		count++
	}

	assert.Nil(t, err)
	assert.Equal(t, 30, count)
}

func TestImageFinderUseCase_ImagesFoundLessThanRequestedImages(t *testing.T) {
	//SET UP
	os.RemoveAll(TestDirectory)
	os.Mkdir(TestDirectory, 0755)
	amountOfImages := 10

	//when
	service := adapters.NewImagesProcessorService(NewImagesProcessorUseCaseWithMockedClients("image_finder_mock_with_no_records.html"))
	err := service.Execute(amountOfImages)

	//then
	files, _ := os.ReadDir(TestDirectory)
	var count int
	for _, _ = range files {
		count++
	}

	assert.Nil(t, err)
	assert.Equal(t, 0, count)
}

func TestFunctionalImagesProcessor_With500Images(t *testing.T) {
	//SET UP
	os.RemoveAll(TestDirectory)
	os.Mkdir(TestDirectory, 0755)
	amountOfImages := 500

	//when
	service := adapters.NewImagesProcessorService(NewImagesProcessorUseCaseWithMockedClients("image_finder_mock.html"))
	err := service.Execute(amountOfImages)

	//then
	assert.NotNil(t, err)
	assert.Equal(t, "amount of images should not be more than 100", err.Error())
}

func NewImagesProcessorUseCaseWithMockedClients(fileName string) *usecases.ImageProcessorUseCase {
	return usecases.NewImageProcessorUseCase(
		createImageFinderUseCaseWithMockedClient(fileName),
		createImageDownloaderUseCaseWithMockedClient(),
	)
}

func createImageFinderUseCaseWithMockedClient(fileName string) *usecases.ImageFinderUseCase {
	return usecases.NewImageFinderUseCase(ImageFinderClientMock{mockDirectory: fileName})
}

func createImageDownloaderUseCaseWithMockedClient() *usecases.ImageDownloaderAndStorerUseCase {
	return usecases.NewImageDownloaderAndStorerUseCase(
		ImageDownloaderClientMock{},
		&frameworks.ImageStorerClientImpl{Directory: TestDirectory},
	)
}

type ImageFinderClientMock struct {
	mockDirectory string
}

func (imageFinderClientMock ImageFinderClientMock) Find(page int) (*goquery.Document, error) {
	data, _ := os.ReadFile(imageFinderClientMock.mockDirectory)
	query, _ := goquery.NewDocumentFromReader(strings.NewReader(string(data)))
	return query, nil
}

type ImageDownloaderClientMock struct{}

func (imageDownloaderClientMock ImageDownloaderClientMock) Download(url string) ([]byte, error) {
	return []byte{}, nil
}
