package factory

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/KludgePub/Matrix/actor"
)

var letterRunes []rune
var seededRand *rand.Rand

func init() {
	letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// PrintSymbol loads new letter to entity.TextObject
func setSymbol(ml *actor.MatrixLetter) {
	symbol := []rune{
		letterRunes[seededRand.Intn(len(letterRunes))],
	}

	ml.Text.SetTextField(fmt.Sprintf("%s", string(symbol)))
}
