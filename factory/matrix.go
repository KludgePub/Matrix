package factory

import (
	"fmt"

	"github.com/LinMAD/Matrix/actor"
	"github.com/LinMAD/Snap/engine/entity"
	"github.com/LinMAD/Snap/engine/graphics/data"
	"github.com/LinMAD/Snap/engine/platform"
)

const fontPath = "assets/matrix.ttf"

// CreateMatrix must generate slice of actors for scene
func CreateMatrix(sc *platform.ScreenConfiguration) []entity.SceneObject {
	actors := make([]entity.SceneObject, 0)

	// TODO Lecture: Tip

	return actors
}

// createNewLetterActor to be placed in the scene
func createNewLetterActor(screenConfig *platform.ScreenConfiguration, x, y int32) *actor.MatrixLetter {
	letter := entity.NewTextObject(
		&entity.Position{X: x, Y: y},
		&data.FontData{FontFilePath: fontPath, Size: 30},
		new(entity.Color),
	)

	letter.DrawableInformation.FontData.ID = fmt.Sprintf("%p", &letter)

	return actor.NewMatrixLetter(int32(screenConfig.Height), letter)
}
