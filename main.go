package main

import (
	"fmt"

	"image/color"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"github.com/ryuta06012/2d_game/systems"
)

type myScene struct{}

// Type uniquely defines your game type
func (*myScene) Type() string { return "myGame" }

// Preload is called before loading any assets from the disk,
// to allow you to register / queue them
func (*myScene) Preload() {
	engo.Files.Load("icon/city.png")
}

// Setup is called before the main loop starts. It allows you
// to add entities and systems to your Scene.
func (*myScene) Setup(u engo.Updater) {
	world, _ := u.(*ecs.World)
	common.SetBackground(color.Black)
	world.AddSystem(&common.RenderSystem{})
	world.AddSystem(&systems.TileBuildingSystem{})
}

func main() {
	fmt.Println("hello world!")
	opts := engo.RunOptions{
		Title:  "Pacman Field",
		Width:  21 * 32,
		Height: 23 * 32,
	}
	engo.Run(opts, &myScene{})
}
