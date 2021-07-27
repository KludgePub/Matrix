package actor

import (
	"math/rand"
	"time"

	"github.com/LinMAD/Snap/engine/entity"
)

// MatrixLetter which will be displayed on screen
type MatrixLetter struct {
	// Text object
	Text  *entity.TextObject
	color *entity.Color

	seededRand *rand.Rand
	maxYPos    int32
}

func NewMatrixLetter(maxY int32, t *entity.TextObject) *MatrixLetter {
	ml := &MatrixLetter{
		Text:       t,
		color:      t.DrawableInformation.Color,
		seededRand: rand.New(rand.NewSource(time.Now().UnixNano())),
		maxYPos:    maxY,
	}

	// TODO Lecture: Tip

	return ml
}

// OnUpdate update text
func (ml *MatrixLetter) OnUpdate() {

	ml.addEffectLetterDrop()
	ml.addEffectColorTransition()

	ml.Text.OnUpdate()
}

func (ml *MatrixLetter) addEffectLetterDrop() {
	// TODO Lecture: Tip
}

func (ml *MatrixLetter) addEffectColorTransition() {
	// TODO Lecture: Tip
}

// GetDrawableInformation about object
func (ml *MatrixLetter) GetDrawableInformation() *entity.DrawableInformation {
	return ml.Text.GetDrawableInformation()
}

// GetPosition in the scene
func (ml *MatrixLetter) GetPosition() *entity.Position {
	return ml.Text.GetPosition()
}
