package systems

import (
	"fmt"
	"math"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

type Player struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
	common.AnimationComponent
}

type PlayerMovementSystem struct {
	world            *ecs.World
	playerEntity     *Player
	field            [][]int
	fieldWidth       int
	fieldHeight      int
	speed            float32
	direction        int // 1: up, 2: down, 3: right, 4: left
	nextDirection    int
	currentPositionX int
	currentPositionY int
	suspend          bool
}

const (
	DIRECTION_UP    = 1
	DIRECTION_DOWN  = 2
	DIRECTION_RIGHT = 3
	DIRECTION_LEFT  = 4
)

var AnimationPatternUp []string = []string{"icon/pacman_up_open.png", "icon/pacman_up_middle.png", "icon/pacman_close.png"}
var AnimationPatternDown []string = []string{"icon/pacman_down_open.png", "icon/pacman_down_middle.png", "icon/pacman_close.png"}
var AnimationPatternLeft []string = []string{"icon/pacman_left_open.png", "icon/pacman_left_middle.png", "icon/pacman_close.png"}
var AnimationPatternRight []string = []string{"icon/pacman_right_open.png", "icon/pacman_right_middle.png", "icon/pacman_close.png"}
var AnimationFrameStop []string = []string{"icon/pacman_up_open.png", "icon/pacman_down_open.png", "icon/pacman_right_open.png", "icon/pacman_left_open.png"}

// New is the initialisation of the System
func (pms *PlayerMovementSystem) New(w *ecs.World) {
	pms.world = w
	fmt.Println("PlayerMovementSystem was added to the Scene")
	player := Player{BasicEntity: ecs.NewBasic()}
	positionX := 32 * 10
	positionY := 32 * 21
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
		Scale:    engo.Point{X: 1, Y: 1},
	}
	pms.playerEntity = &player
	pms.currentPositionX = positionX / 32
	pms.currentPositionY = positionY / 32
	pms.field = Tiles
	pms.fieldHeight = len(Tiles)
	pms.fieldWidth = len(Tiles[0])
	pms.speed = 32 / 8
	pms.direction = 4
	pms.nextDirection = 4
	pms.suspend = true
	for _, system := range pms.world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&player.BasicEntity, &player.RenderComponent, &player.SpaceComponent)
		}
	}
	engo.Mailbox.Listen(HUDOpeningMessageType, func(m engo.Message) {
		pms.suspend = false
	})
}

// Update is ran every frame, with `dt` being the time
// in seconds since the last frame
func (pms *PlayerMovementSystem) Update(dt float32) {
	fmt.Printf("pms.suspend: %v\n", pms.suspend)
	if !pms.suspend {
		pms.nextDirectionPlayer()
		pms.changeDirectionIfPossible()
		pms.changeDirectionTexture()
		pms.moveAdd()
		if pms.checkCollisions() {
			pms.moveDecrease()
			texture, err := common.LoadedSprite(AnimationFrameStop[pms.direction-1])
			if err != nil {
				fmt.Println("Unable to load texture: " + err.Error())
			}
			pms.playerEntity.RenderComponent.Drawable = texture
		}
	}
}

func (pms *PlayerMovementSystem) checkCollisions() bool {
	isCollided := false
	if pms.field[int(math.Floor(float64(pms.playerEntity.SpaceComponent.Position.Y/32)))][int(math.Floor(float64(pms.playerEntity.SpaceComponent.Position.X/32)))] == 1 ||
		pms.field[int(math.Floor(float64(pms.playerEntity.SpaceComponent.Position.Y/32+0.9999)))][int(math.Floor(float64(pms.playerEntity.SpaceComponent.Position.X/32)))] == 1 ||
		pms.field[int(math.Floor(float64(pms.playerEntity.SpaceComponent.Position.Y/32)))][int(math.Floor(float64(pms.playerEntity.SpaceComponent.Position.X/32+0.9999)))] == 1 ||
		pms.field[int(math.Floor(float64(pms.playerEntity.SpaceComponent.Position.Y/32+0.9999)))][int(math.Floor(float64(pms.playerEntity.SpaceComponent.Position.X/32+0.9999)))] == 1 {
		isCollided = true
	}
	return isCollided
}

func (pms *PlayerMovementSystem) moveAdd() {
	switch pms.direction {
	case DIRECTION_RIGHT: // Right
		pms.playerEntity.SpaceComponent.Position.X += pms.speed
		break
	case DIRECTION_LEFT: // Left
		pms.playerEntity.SpaceComponent.Position.X -= pms.speed
		break
	case DIRECTION_UP: // Up
		pms.playerEntity.SpaceComponent.Position.Y -= pms.speed
		break
	case DIRECTION_DOWN: // Bottom
		pms.playerEntity.SpaceComponent.Position.Y += pms.speed
		break
	}
}

func (pms *PlayerMovementSystem) moveDecrease() {
	switch pms.direction {
	case DIRECTION_RIGHT: // Right
		pms.playerEntity.SpaceComponent.Position.X -= pms.speed
		break
	case DIRECTION_LEFT: // Left
		pms.playerEntity.SpaceComponent.Position.X += pms.speed
		break
	case DIRECTION_UP: // Up
		pms.playerEntity.SpaceComponent.Position.Y += pms.speed
		break
	case DIRECTION_DOWN: // Bottom
		pms.playerEntity.SpaceComponent.Position.Y -= pms.speed
		break
	}
}

var frame uint64 = 0
var animation int = 0
var animationFrame uint64 = 5

func (pms *PlayerMovementSystem) changeDirectionTexture() {
	switch pms.direction {
	case DIRECTION_RIGHT: // Right
		if frame%animationFrame == 0 {
			texture, err := common.LoadedSprite(AnimationPatternRight[animation])
			if err != nil {
				fmt.Println("Unable to load texture: " + err.Error())
			}
			pms.playerEntity.RenderComponent.Drawable = texture
			pms.AddAnimationFrame()
			break
		}
	case DIRECTION_LEFT: // Left
		if frame%animationFrame == 0 {
			// fmt.Println("##############")
			texture, err := common.LoadedSprite(AnimationPatternLeft[animation])
			if err != nil {
				fmt.Println("Unable to load texture: " + err.Error())
			}
			pms.playerEntity.RenderComponent.Drawable = texture
			pms.AddAnimationFrame()
			break
		}
	case DIRECTION_UP: // Up
		if frame%animationFrame == 0 {
			texture, err := common.LoadedSprite(AnimationPatternUp[animation])
			if err != nil {
				fmt.Println("Unable to load texture: " + err.Error())
			}
			pms.playerEntity.RenderComponent.Drawable = texture
			pms.AddAnimationFrame()
			break
		}
	case DIRECTION_DOWN: // Bottom
		if frame%animationFrame == 0 {
			texture, err := common.LoadedSprite(AnimationPatternDown[animation])
			if err != nil {
				fmt.Println("Unable to load texture: " + err.Error())
			}
			pms.playerEntity.RenderComponent.Drawable = texture
			pms.AddAnimationFrame()
			break
		}
	}
	frame++
}

func (pms *PlayerMovementSystem) AddAnimationFrame() {
	if animation >= 2 {
		animation = 0
	} else {
		animation++
	}
}

func (pms *PlayerMovementSystem) changeDirectionIfPossible() {
	if pms.direction == pms.nextDirection {
		return
	}
	tempDirection := pms.direction
	pms.direction = pms.nextDirection
	pms.moveAdd()
	if pms.checkCollisions() {
		pms.moveDecrease()
		pms.direction = tempDirection
	} else {
		pms.moveDecrease()
	}
}

func (pms *PlayerMovementSystem) nextDirectionPlayer() {
	if engo.Input.Button("MoveRight").Down() {
		pms.nextDirection = DIRECTION_RIGHT
	}
	if engo.Input.Button("MoveLeft").Down() {
		pms.nextDirection = DIRECTION_LEFT
	}
	if engo.Input.Button("MoveDown").Down() {
		pms.nextDirection = DIRECTION_DOWN
	}
	if engo.Input.Button("MoveUP").Down() {
		pms.nextDirection = DIRECTION_UP
	}
}

func (pms *PlayerMovementSystem) animationPattern() {
	switch pms.direction {
	case DIRECTION_RIGHT: // Right
		pms.playerEntity.SpaceComponent.Position.X -= pms.speed
		break
	case DIRECTION_LEFT: // Left
		pms.playerEntity.SpaceComponent.Position.X += pms.speed
		break
	case DIRECTION_UP: // Up
		pms.playerEntity.SpaceComponent.Position.Y += pms.speed
		break
	case DIRECTION_DOWN: // Bottom
		pms.playerEntity.SpaceComponent.Position.Y -= pms.speed
		break
	}
}

// Update is ran every frame, with `dt` being the time
// in seconds since the last frame
func (pms *PlayerMovementSystem) Remove(ecs.BasicEntity) {}

func (pms *PlayerMovementSystem) GetMapX() int {
	return int(math.Floor(float64(pms.playerEntity.SpaceComponent.Position.X / 32)))
}

func (pms *PlayerMovementSystem) GetMapY() int {
	return int(math.Floor(float64(pms.playerEntity.SpaceComponent.Position.Y / 32)))
}
