package frameworks

import (
	"github.com/rodrinoblega/proper_images_processor/src/instrumentation"
	"os"
	"strconv"
)

type ImageStorerClientImpl struct {
	Directory string
}

func (imageStorerClientImpl *ImageStorerClientImpl) Store(index int, bytes []byte) {
	file, err := os.Create(imageStorerClientImpl.Directory + "image " + strconv.Itoa(index+1) + ".jpg")
	if err != nil {
		instrumentation.LogMessage("There was an error creating the directory: " + imageStorerClientImpl.Directory + "image " + strconv.Itoa(index+1) + ".jpg. Error: " + err.Error())
	}

	defer file.Close()

	_, err = file.Write(bytes)
	if err != nil {
		instrumentation.LogMessage("There was an error reading bytes: " + err.Error())
	}
}
