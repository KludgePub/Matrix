package factory

import (
	"fmt"
	"github.com/LinMAD/Snap/engine/entity"
	"github.com/LinMAD/Snap/engine/graphics/data"
	"github.com/LinMAD/Snap/engine/platform"

	"github.com/KludgePub/Matrix/actor"
)

const fontPath = "assets/matrix.ttf"

// CreateMatrix must generate slice of actors for scene
func CreateMatrix(sc *platform.ScreenConfiguration) []entity.SceneObject {
	actors := make([]entity.SceneObject, 1)

	// TODO Implement an algorithm to create structure of actors

	// Add actor with coordinates on the screen
	letter := createNewLetterActor(sc, int32(sc.Width>>1), int32(sc.Height>>2))

	// Set random symbol for the letter
	setSymbol(letter)

	actors[0] = letter

	return actors
}

// createNewLetterActor to be placed in the scene
func createNewLetterActor(screenConfig *platform.ScreenConfiguration, x, y int32) *actor.MatrixLetter {
	letter := entity.NewTextObject(
		&entity.Position{X: x, Y: y},
		&data.FontData{FontFilePath: fontPath, Size: 30},
		&entity.Color{Green: 255},
	)

	letter.DrawableInformation.FontData.ID = fmt.Sprintf("%p", &letter)

	return actor.NewMatrixLetter(int32(screenConfig.Height), letter)
}
