package usecases

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/rodrinoblega/proper_images_processor/src/domain"
)

type ImageFinderClient interface {
	Find(page int) (*goquery.Document, error)
}

type ImageFinderUseCase struct {
	ImageFinderClient ImageFinderClient
}

func NewImageFinderUseCase(imageFinderGateway ImageFinderClient) *ImageFinderUseCase {
	return &ImageFinderUseCase{
		ImageFinderClient: imageFinderGateway,
	}
}

func (imageFinder *ImageFinderUseCase) execute(imagesUrlAccumulated []domain.Image, amountOfImagesRequested int) ([]domain.Image, error) {
	//Find on first page
	var amountOfImagesFound int
	doc, err := imageFinder.ImageFinderClient.Find(1)
	imagesUrlAccumulated, amountOfImagesFound = findImagesOnTheActualPage(doc, imagesUrlAccumulated, amountOfImagesRequested)

	//Find on Following pages
	for page := 2; len(imagesUrlAccumulated) <= amountOfImagesRequested && amountOfImagesFound != 0; page++ {
		amountOfImagesFound = 0

		doc, err = imageFinder.ImageFinderClient.Find(page)

		imagesUrlAccumulated, amountOfImagesFound = findImagesOnTheActualPage(doc, imagesUrlAccumulated, amountOfImagesRequested)
	}

	if err != nil {
		return []domain.Image{}, err
	}
	return imagesUrlAccumulated, nil
}

func findImagesOnTheActualPage(doc *goquery.Document, imagesUrlAccumulated []domain.Image, amountOfImages int) ([]domain.Image, int) {
	var amountOfImagesFound int
	firstChildren := doc.Find(".mu-col-container").Children().First()

	firstChildren.Find(".resp-media").Not(".lazyload").Each(func(i int, s *goquery.Selection) {
		if len(imagesUrlAccumulated) < amountOfImages {
			val, _ := s.Attr("src")
			imagesUrlAccumulated = append(imagesUrlAccumulated, domain.Image{Url: val, Position: len(imagesUrlAccumulated) + 1})
			amountOfImagesFound = amountOfImagesFound + 1
		}
	})

	firstChildren.Find(".resp-media.lazyload").Each(func(i int, s *goquery.Selection) {
		if len(imagesUrlAccumulated) < amountOfImages {
			val, _ := s.Attr("data-src")
			imagesUrlAccumulated = append(imagesUrlAccumulated, domain.Image{Url: val, Position: len(imagesUrlAccumulated) + 1})
			amountOfImagesFound = amountOfImagesFound + 1
		}
	})

	return imagesUrlAccumulated, amountOfImagesFound
}
