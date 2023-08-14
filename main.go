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
	engo.Files.Load("icon/pacman_close.png",
		"icon/pacman_open_up.png",
		"icon/pacman_open_left.png",
		"icon/pacman_open_right.png",
		"icon/pacman_open_down.png")
}

// Setup is called before the main loop starts. It allows you
// to add entities and systems to your Scene.
func (*myScene) Setup(u engo.Updater) {
	engo.Input.RegisterButton("MoveRight", engo.KeyD, engo.KeyArrowRight)
	engo.Input.RegisterButton("MoveLeft", engo.KeyA, engo.KeyArrowLeft)
	engo.Input.RegisterButton("MoveUP", engo.KeyW, engo.KeyArrowUp)
	engo.Input.RegisterButton("MoveDown", engo.KeyS, engo.KeyArrowDown)
	world, _ := u.(*ecs.World)
	common.SetBackground(color.Black)
	world.AddSystem(&common.RenderSystem{})
	world.AddSystem(&systems.TileBuildingSystem{})
	world.AddSystem(&systems.FoodSystem{})
	world.AddSystem(&systems.PlayerMovementSystem{})
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
