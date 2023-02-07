package usecases

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestImagesProcessorUseCase_ErrorInImageFinderClient(t *testing.T) {
	//setup
	imageProcessorUseCase := NewImageProcessorUseCase(NewImageFinderUseCase(ImageFinderClientMock{}), nil)

	//when
	err := imageProcessorUseCase.Execute(10, 5)

	//then
	assert.NotNil(t, err)
	assert.Equal(t, "custom error", err.Error())
}
