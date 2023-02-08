package frameworks

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestImageFinderClientImpl_WithErrorInApiCall(t *testing.T) {
	//setup
	imageFinderClientImpl := ImageFinderClientImpl{Url: "asd", Client: &ImageFinderClientMockWithErrorInApiCall{}}

	//when
	_, err := imageFinderClientImpl.Find(1)

	//then
	assert.NotNil(t, err)
	assert.Equal(t, "custom error", err.Error())
}

func TestImageFinderClientImpl_WithValidGoqueryDocumentResponse(t *testing.T) {
	//setup
	imageFinderClientImpl := ImageFinderClientImpl{Url: "asd", Client: &ImageFinderClientMockWithValidGoqueryDocumentInResponse{}}

	//when
	doc, _ := imageFinderClientImpl.Find(2)

	//then
	expected, _ := createDocExpected()
	assert.Equal(t, expected, doc)
}

func createDocExpected() (*goquery.Document, error) {
	stringReader := strings.NewReader("¿%·1")
	stringReadCloser := io.NopCloser(stringReader)
	response := http.Response{Body: stringReadCloser}

	return goquery.NewDocumentFromReader(response.Body)
}

type ImageFinderClientMockWithErrorInApiCall struct{}

func (imageFinderClientMock *ImageFinderClientMockWithErrorInApiCall) Do(req *http.Request) (*http.Response, error) {
	stringReader := strings.NewReader("Test!!")
	stringReadCloser := io.NopCloser(stringReader)
	return &http.Response{Body: stringReadCloser}, errors.New("custom error")
}

type ImageFinderClientMockWithValidGoqueryDocumentInResponse struct{}

func (imageFinderClientMock *ImageFinderClientMockWithValidGoqueryDocumentInResponse) Do(req *http.Request) (*http.Response, error) {
	stringReader := strings.NewReader("¿%·1")
	stringReadCloser := io.NopCloser(stringReader)
	return &http.Response{Body: stringReadCloser}, nil
}
