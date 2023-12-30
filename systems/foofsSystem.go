package systems

import (
	"fmt"
	"image/color"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

type Food struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
	isEat bool
}

type FoodSystem struct {
	world      *ecs.World
	foodEntity []*Food
	isAte      bool
	foodColor  color.RGBA
}

func (f *FoodSystem) New(w *ecs.World) {
	f.world = w
	fmt.Println("TileBuildingSystem was added to the Scene")
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
				fmt.Printf("food.SpaceComponent.Position.X: %v\n", food.SpaceComponent.Position.X)
				fmt.Printf("food.SpaceComponent.Position.Y: %v\n", food.SpaceComponent.Position.Y)
				food.RenderComponent = common.RenderComponent{
					Drawable: common.Rectangle{},
					Color:    color.RGBA{254, 184, 151, 255},
					Scale:    engo.Point{X: 1, Y: 1},
				}
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
		}
	}
}

// Update is ran every frame, with `dt` being the time
// in seconds since the last frame
func (f *FoodSystem) Update(dt float32) {
	// var playerPositionX float32
	// var playerPositionY float32
	var playerMapX int
	var playerMapY int
	for _, system := range f.world.Systems() {
		switch sys := system.(type) {
		case *PlayerMovementSystem:
			// playerPositionX = sys.playerEntity.SpaceComponent.Position.X
			// playerPositionY = sys.playerEntity.SpaceComponent.Position.Y
			playerMapX = sys.GetMapX()
			playerMapY = sys.GetMapY()
		}
	}
	var count int=0
	// for y, tile := range Tiles {
	// 	for x, cell := range tile {
	// 		if cell == 2 && playerMapX == x && playerMapY == y {

	// 		}
	// 	}
	// }
	for i, food := range f.foodEntity {
		fmt.Printf("i: %v\n", i)
		fmt.Printf("count: %v\n", count)
		fmt.Printf("food.SpaceComponent.Position.X: %v\n", food.SpaceComponent.Position.X)
		fmt.Printf("food.SpaceComponent.Position.Y: %v\n", food.SpaceComponent.Position.Y)
		// fmt.Printf("food.SpaceComponent.Position.X+wallOffset/2: %v\n", food.SpaceComponent.Position.X+wallOffset/2)
		// fmt.Printf("food.SpaceComponent.Position.Y+wallOffset/2: %v\n", food.SpaceComponent.Position.Y+wallOffset/2)
		// fmt.Printf("playerPositionX: %v\n", playerPositionX+16)
		// fmt.Printf("playerPosition: %v\n", playerPositionY+16)
		fmt.Printf("int((food.SpaceComponent.Position.X - wallOffset) / 32): %v\n", int((food.SpaceComponent.Position.X-wallOffset)/32))
		fmt.Printf("int((food.SpaceComponent.Position.Y - wallOffset) / 32): %v\n", int((food.SpaceComponent.Position.Y-wallOffset)/32))
		fmt.Printf("playerMapX: %v\n", playerMapX)
		fmt.Printf("playerMapY: %v\n", playerMapY)
		fmt.Printf("wallOffset: %v\n", wallOffset)
		if food.isEat != true && int((food.SpaceComponent.Position.X - wallOffset) / 32) == playerMapX && int((food.SpaceComponent.Position.Y - wallOffset) / 32) == playerMapY {
			println("#################################")
			food.isEat = true
			food.RenderComponent.Color = color.RGBA{0, 0, 0, 255}
			break
		}
		count++
	}
}

// Update is ran every frame, with `dt` being the time
// in seconds since the last frame
func (*FoodSystem) Remove(ecs.BasicEntity) {}
