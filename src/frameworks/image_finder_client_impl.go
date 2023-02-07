package frameworks

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strconv"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type ImageFinderClientImpl struct {
	Client HTTPClient
	Url    string
}

func NewImageFinderClientImpl(url string, client HTTPClient) *ImageFinderClientImpl {
	return &ImageFinderClientImpl{Url: url, Client: client}
}

func (imageFinderGateway ImageFinderClientImpl) Find(page int) (*goquery.Document, error) {
	var url string

	if page == 1 {
		url = imageFinderGateway.Url
	} else {
		url = imageFinderGateway.Url + "page/" + strconv.Itoa(page)
	}
	page = page + 1

	request, err := http.NewRequest(http.MethodGet, url, nil)
	resp, err := imageFinderGateway.Client.Do(request)

	if err != nil {
		return &goquery.Document{}, err
	}
	defer resp.Body.Close()

	return goquery.NewDocumentFromReader(resp.Body)

}
