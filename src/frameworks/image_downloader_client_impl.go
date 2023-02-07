package frameworks

import (
	"io"
	"net/http"
)

type ImageDownloaderClientImpl struct {
	Client HTTPClient
}

func NewImageDownloaderClientImpl(client HTTPClient) *ImageDownloaderClientImpl {
	return &ImageDownloaderClientImpl{Client: client}
}

func (imageDownloaderGateway ImageDownloaderClientImpl) Download(url string) ([]byte, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	response, err := imageDownloaderGateway.Client.Do(request)

	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return io.ReadAll(response.Body)
}
