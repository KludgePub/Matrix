package main

import (
	"github.com/LinMAD/Snap/engine"
	"github.com/LinMAD/Snap/engine/platform"

	"github.com/LinMAD/Matrix/factory"
)

func main() {
	// Set screen properties
	screenConfig := &platform.ScreenConfiguration{
		Title:  "Enter The Matrix",
		Width:  800,
		Height: 600,
	}

	// Create actors for scene
	actors := factory.CreateMatrix(screenConfig)

	// Init engine instance
	snapEngine := engine.New(screenConfig, false)
	snapEngine.LoadSceneObjects(actors)

	// Execute scene window
	if err := snapEngine.Run(); err != nil {
		panic(err.Error())
	}
}
