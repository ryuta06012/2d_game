package systems

import (
	"fmt"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

type Player struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

type PlayerMovementSystem struct {
	world            *ecs.World
	playerEntity     *Player
	field            [][]int
	fieldWidth       int
	fieldHeight      int
	currentPositionX int
	currentPositionY int
}

// New is the initialisation of the System
func (pms *PlayerMovementSystem) New(w *ecs.World) {
	pms.world = w
	fmt.Println("TileBuildingSystem was added to the Scene")
	player := Player{BasicEntity: ecs.NewBasic()}
	positionX := 32 * 4
	positionY := 32 * 1
	player.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{X: float32(positionX), Y: float32(positionY)},
		Width:    32,
		Height:   32,
	}
	texture, err := common.LoadedSprite("icon/pacman_open_up.png")
	if err != nil {
		fmt.Println("Unable to load texture: " + err.Error())
	}
	player.RenderComponent = common.RenderComponent{
		Drawable: texture,
		Scale:    engo.Point{X: 0.9, Y: 0.9},
	}
	pms.playerEntity = &player
	pms.currentPositionX = positionX / 32
	pms.currentPositionY = positionY / 32
	pms.field = Tiles
	pms.fieldHeight = len(Tiles)
	pms.fieldWidth = len(Tiles[0])
	for _, system := range pms.world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&player.BasicEntity, &player.RenderComponent, &player.SpaceComponent)
		}
	}
}

// Update is ran every frame, with `dt` being the time
// in seconds since the last frame
func (pms *PlayerMovementSystem) Update(dt float32) {
	// fmt.Println("###################")
	velocity := float32(10)
	if engo.Input.Button("MoveRight").Down() {
		fmt.Println("----------MoveRight--------")
		if pms.currentPositionX < pms.fieldWidth && pms.field[pms.currentPositionY][pms.currentPositionX+1] == 2 {
			pms.playerEntity.SpaceComponent.Position.X += 32 - velocity
			pms.currentPositionX = int(pms.playerEntity.SpaceComponent.Position.X) / 32
		}
		texture, err := common.LoadedSprite("icon/pacman_open_right.png")
		if err != nil {
			fmt.Println("Unable to load texture: " + err.Error())
		}
		pms.playerEntity.RenderComponent.Drawable = texture
	}
	if engo.Input.Button("MoveLeft").Down() {
		fmt.Println("------------MoveLeft-------")
		if pms.currentPositionX > 0 && pms.field[pms.currentPositionY][pms.currentPositionX-1] == 2 {
			pms.playerEntity.SpaceComponent.Position.X -= 32 - velocity
			pms.currentPositionX = int(pms.playerEntity.SpaceComponent.Position.X) / 32
		}
		texture, err := common.LoadedSprite("icon/pacman_open_left.png")
		if err != nil {
			fmt.Println("Unable to load texture: " + err.Error())
		}
		pms.playerEntity.RenderComponent.Drawable = texture

	}
	if engo.Input.Button("MoveDown").Down() {
		fmt.Println("--------MoveDown----------")

		if pms.currentPositionY < pms.fieldHeight && pms.field[pms.currentPositionY+1][pms.currentPositionX] == 2 {
			pms.playerEntity.SpaceComponent.Position.Y += 32 - velocity
			pms.currentPositionY = int(pms.playerEntity.SpaceComponent.Position.Y) / 32
		}
		texture, err := common.LoadedSprite("icon/pacman_open_down.png")
		if err != nil {
			fmt.Println("Unable to load texture: " + err.Error())
		}
		pms.playerEntity.RenderComponent.Drawable = texture

	}
	if engo.Input.Button("MoveUP").Down() {
		fmt.Println("---------------MoveUP--------------")
		if pms.currentPositionY > 0 && pms.field[pms.currentPositionY-1][pms.currentPositionX] == 2 {
			pms.playerEntity.SpaceComponent.Position.Y -= 32 - velocity
			pms.currentPositionY = int(pms.playerEntity.SpaceComponent.Position.Y) / 32
		}
		texture, err := common.LoadedSprite("icon/pacman_open_up.png")
		if err != nil {
			fmt.Println("Unable to load texture: " + err.Error())
		}
		pms.playerEntity.RenderComponent.Drawable = texture

	}
	/* for _, system := range pms.world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&pms.playerEntity.BasicEntity, &pms.playerEntity.RenderComponent, &pms.playerEntity.SpaceComponent)
		}
	} */
}

// Update is ran every frame, with `dt` being the time
// in seconds since the last frame
func (pms *PlayerMovementSystem) Remove(ecs.BasicEntity) {}
