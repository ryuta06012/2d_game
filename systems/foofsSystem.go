package systems

import (
	"fmt"
	"image/color"
	"log"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

type Food struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
	mapX  int
	mapY  int
	isEat bool
}

type Audio struct {
	ecs.BasicEntity
	common.AudioComponent
}

type FoodSystem struct {
	world       *ecs.World
	foodEntity  []*Food
	foodColor   color.RGBA
	audioEntity []*Audio
}

var AudioFile []string = []string{"sounds/chewing_sound_pattern1.mp3", "sounds/chewing_sound_pattern2.mp3"}

func (f *FoodSystem) New(w *ecs.World) {
	f.world = w
	for i := 0; i < 2; i++ {
		audio := &Audio{BasicEntity: ecs.NewBasic()}
		player, err := common.LoadedPlayer(AudioFile[i])
		if err != nil {
			log.Println(err)
		}
		player.SetVolume(0.5)
		audio.AudioComponent = common.AudioComponent{Player: player}
		f.audioEntity = append(f.audioEntity, audio)
	}
	fmt.Println("FoodSystemSystem was added to the Scene")
	f.generateFoodsInFields()
}

var wallOffset float32

func (f *FoodSystem) generateFoodsInFields() {
	tileSize := float32(32)
	wallSpaceWidth := tileSize / 6
	wallOffset = (tileSize - wallSpaceWidth) / 2
	//wallInnerColor := color.RGBA{0, 0, 0, 255}
	f.foodColor = color.RGBA{254, 184, 151, 255}

	for y, row := range Tiles {
		for x, cell := range row {
			if cell == 2 {
				food := &Food{BasicEntity: ecs.NewBasic()}
				food.SpaceComponent = common.SpaceComponent{
					Position: engo.Point{X: float32(x)*tileSize + wallOffset, Y: float32(y)*tileSize + wallOffset},
					Width:    wallSpaceWidth,
					Height:   wallSpaceWidth,
				}
				food.RenderComponent = common.RenderComponent{
					Drawable: common.Rectangle{},
					Color:    color.RGBA{254, 184, 151, 255},
					Scale:    engo.Point{X: 1, Y: 1},
				}
				food.mapX = x
				food.mapY = y
				f.foodEntity = append(f.foodEntity, food)
			}
		}
	}
	for _, system := range f.world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			for _, v := range f.foodEntity {
				sys.Add(&v.BasicEntity, &v.RenderComponent, &v.SpaceComponent)
			}
		case *common.AudioSystem:
			for _, v := range f.audioEntity {
				sys.Add(&v.BasicEntity, &v.AudioComponent)
			}
		}
	}
}

// Update is ran every frame, with `dt` being the time
// in seconds since the last frame
func (f *FoodSystem) Update(dt float32) {
	var playerMapX int
	var playerMapY int
	var playerDirection int

	for _, system := range f.world.Systems() {
		switch sys := system.(type) {
		case *PlayerMovementSystem:
			playerMapX = sys.GetMapX()
			playerMapY = sys.GetMapY()
			playerDirection = sys.direction
		}
	}
	var count int = 0
	for _, food := range f.foodEntity {
		if food.isEat == false && food.mapX == playerMapX && food.mapY == playerMapY {
			println(food.mapX % 2)
			println(food.mapY % 2)
			f.selectAudioPlayerByDirection(f.audioEntity, playerDirection, food.mapX, food.mapY)
			food.isEat = true
			food.RenderComponent.Color = color.RGBA{0, 0, 0, 255}
			engo.Mailbox.Dispatch(HUDScoreMessage{
				Score: 100,
			})
			break
		}
		count++
	}
}

func (f *FoodSystem) selectAudioPlayerByDirection(audioEntity []*Audio, direction, x, y int) {
	if direction == 1 || direction == 2 {
		audioEntity[y%2].Player.Play()
		audioEntity[y%2].Player.Rewind()
	} else {
		audioEntity[x%2].Player.Play()
		audioEntity[x%2].Player.Rewind()
	}
}

// Update is ran every frame, with `dt` being the time
// in seconds since the last frame
func (*FoodSystem) Remove(ecs.BasicEntity) {}
