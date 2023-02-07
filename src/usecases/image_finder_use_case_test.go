package usecases

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/rodrinoblega/proper_images_processor/src/domain"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

func TestImageFinderUseCase_ErrorInFinderClient(t *testing.T) {

	//SETUP
	imageFinderUseCase := ImageFinderUseCase{ImageFinderClient: ImageFinderClientMock{}}

	//WHEN
	_, err := imageFinderUseCase.execute([]domain.Image{{Position: 1, Url: "asd.com"}}, 1)

	//then
	assert.NotNil(t, err)
	assert.Equal(t, "custom error", err.Error())
}

type ImageFinderClientMock struct{}

func (imageFinderClientMock ImageFinderClientMock) Find(page int) (*goquery.Document, error) {
	data, _ := os.ReadFile("images_test/")
	query, _ := goquery.NewDocumentFromReader(strings.NewReader(string(data)))
	return query, errors.New("custom error")
}
