package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"

	"image/color"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"github.com/ryuta06012/2d_game/systems"
	"golang.org/x/image/font/gofont/gosmallcaps"
)

type myScene struct{}

// Type uniquely defines your game type
func (*myScene) Type() string { return "myGame" }

// Preload is called before loading any assets from the disk,
// to allow you to register / queue them
func (*myScene) Preload() {
	engo.Files.Load(
		"icon/pacman_close.png",
		"icon/pacman_open_up.png",
		"icon/pacman_open_left.png",
		"icon/pacman_open_right.png",
		"icon/pacman_open_down.png",
		"icon/pacman_up_open.png",
		"icon/pacman_up_middle.png",
		"icon/pacman_down_open.png",
		"icon/pacman_down_middle.png",
		"icon/pacman_left_open.png",
		"icon/pacman_left_middle.png",
		"icon/pacman_right_open.png",
		"icon/pacman_right_middle.png",
		"icon/goast/akabe/akabe_up_pattern1.png",
		"icon/goast/akabe/akabe_up_pattern2.png",
		"icon/goast/akabe/akabe_down_pattern1.png",
		"icon/goast/akabe/akabe_down_pattern2.png",
		"icon/goast/akabe/akabe_left_pattern1.png",
		"icon/goast/akabe/akabe_left_pattern2.png",
		"icon/goast/akabe/akabe_right_pattern1.png",
		"icon/goast/akabe/akabe_right_pattern2.png",
		"sounds/chewing_sound_pattern1.mp3",
		"sounds/chewing_sound_pattern2.mp3",
		"sounds/opening_audio.mp3",
	)
	fontData, err := ioutil.ReadFile("assets/font/emulogic.ttf")
	if err != nil {
		log.Fatal(err)
	}

	engo.Files.LoadReaderData("go.ttf", bytes.NewReader(gosmallcaps.TTF))
	engo.Files.LoadReaderData("emulogic.ttf", bytes.NewReader(fontData))
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
	world.AddSystem(&common.AudioSystem{})
	world.AddSystem(&common.AnimationSystem{})
	world.AddSystem(&systems.TileBuildingSystem{})
	world.AddSystem(&systems.FoodSystem{})
	world.AddSystem(&systems.PlayerMovementSystem{})
	world.AddSystem(&systems.ScoreSystem{})
	world.AddSystem(&systems.SoundSystem{})
	world.AddSystem(&systems.GoastSystem{})
}

func main() {
	fmt.Println("hello world!")
	opts := engo.RunOptions{
		Title: "Pacman Field",
		// Width:  21 * 32,
		// Height: 23 * 32,
		Width:  21 * 32,
		Height: 26 * 32,
	}
	engo.Run(opts, &myScene{})
}
