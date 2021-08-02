package actor

import (
	"math/rand"
	"time"

	"github.com/LinMAD/Snap/engine/entity"
)

// MatrixLetter which will be displayed on screen
type MatrixLetter struct {
	Text  *entity.TextObject
	color *entity.Color

	seededRand *rand.Rand
	maxYPos    int32
}

// NewMatrixLetter to be shown on the screen
func NewMatrixLetter(maxY int32, t *entity.TextObject) *MatrixLetter {
	return &MatrixLetter{
		Text:       t,
		color:      t.DrawableInformation.Color,
		seededRand: rand.New(rand.NewSource(time.Now().UnixNano())),
		maxYPos:    maxY,
	}
}

// OnUpdate update text
func (ml *MatrixLetter) OnUpdate() {

	// TODO Implement rainfall effect
	// TODO Implement color transition

	ml.Text.OnUpdate()
}

// GetDrawableInformation about object
func (ml *MatrixLetter) GetDrawableInformation() *entity.DrawableInformation {
	return ml.Text.GetDrawableInformation()
}

// GetPosition in the scene
func (ml *MatrixLetter) GetPosition() *entity.Position {
	return ml.Text.GetPosition()
}
