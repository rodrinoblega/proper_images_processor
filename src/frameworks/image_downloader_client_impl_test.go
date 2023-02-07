package frameworks

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestImageDownloaderClientImpl_WithErrorInApiCall(t *testing.T) {
	imageDownloaderClientImpl := ImageDownloaderClientImpl{Client: &ImageDownloaderClientMockWithErrorInApiCall{}}

	_, err := imageDownloaderClientImpl.Download("https://asd.com")

	//then
	assert.NotNil(t, err)
	assert.Equal(t, "custom error", err.Error())
}

func TestImageDownloaderClientImpl_WithoutErrors(t *testing.T) {
	imageDownloaderClientImpl := ImageDownloaderClientImpl{Client: &ImageDownloaderClientMockWithoutErrors{}}

	_, err := imageDownloaderClientImpl.Download("https://asd.com")

	//then
	assert.Nil(t, err)
}

type ImageDownloaderClientMockWithErrorInApiCall struct{}

func (imageDownloaderClientMock *ImageDownloaderClientMockWithErrorInApiCall) Do(req *http.Request) (*http.Response, error) {
	stringReader := strings.NewReader("Test!!")
	stringReadCloser := io.NopCloser(stringReader)
	return &http.Response{Body: stringReadCloser}, errors.New("custom error")
}

type ImageDownloaderClientMockWithoutErrors struct{}

func (imageDownloaderClientMockWithoutError *ImageDownloaderClientMockWithoutErrors) Do(req *http.Request) (*http.Response, error) {
	stringReader := strings.NewReader("Test!!")
	stringReadCloser := io.NopCloser(stringReader)
	return &http.Response{Body: stringReadCloser}, nil
}
