package usecases

import (
	"bytes"
	"errors"
	"github.com/rodrinoblega/proper_images_processor/src/domain"
	"github.com/rodrinoblega/proper_images_processor/src/frameworks"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestImageDownloaderUseCase_ErrorInImageDownloaderClient(t *testing.T) {
	//set up
	imageFinderUseCase := NewImageDownloaderAndStorerUseCase(ImageDownloaderClientMockWithError{}, &frameworks.ImageStorerClientImpl{Directory: "asd/"})
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()

	//when
	imageFinderUseCase.execute(domain.Image{Url: "https://asd.com"}, 1)

	//then
	assert.Contains(t, buf.String(), "There was an error downloading an image. URL:")
}

func TestImageDownloaderUseCase_ErrorCreatingDirectory(t *testing.T) {
	//set up
	imageFinderUseCase := NewImageDownloaderAndStorerUseCase(ImageDownloaderClientMock{}, &frameworks.ImageStorerClientImpl{Directory: "asd/"})
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()

	//when
	imageFinderUseCase.execute(domain.Image{Url: "https://asd.com"}, 1)

	//then
	assert.Contains(t, buf.String(), "There was an error creating the directory: asd/image 2.jpg. Error: open asd/image 2.jpg: no such file or directory")
}

func TestImageDownloaderUseCase_ErrorReadingResponseBody(t *testing.T) {
	//set up
	imageFinderUseCase := NewImageDownloaderAndStorerUseCase(ImageDownloaderClientMock{}, &frameworks.ImageStorerClientImpl{Directory: "asd/"})
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()

	//when
	imageFinderUseCase.execute(domain.Image{Url: "https://asd.com"}, 1)

	//then
	assert.Contains(t, buf.String(), "There was an error reading bytes: invalid argument")
}

type ImageDownloaderClientMock struct{}

func (imageDownloaderClientMock ImageDownloaderClientMock) Download(url string) ([]byte, error) {
	return []byte{}, nil
}

type ImageDownloaderClientMockWithError struct{}

func (imageDownloaderClientMockWithError ImageDownloaderClientMockWithError) Download(url string) ([]byte, error) {
	return []byte{}, errors.New("custom error")
}
