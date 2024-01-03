package systems

import (
	"fmt"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

type GoastEntity struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
	common.AnimationComponent
	fieldWidth       int
	fieldHeight      int
	speed            float32
	direction        int // 1: up, 2: down, 3: right, 4: left
	nextDirection    int
	currentPositionX int
	currentPositionY int
	animationPattern map[string][]common.Drawable
	name             string
}

type GoastSystem struct {
	world      *ecs.World
	goastEnity []*GoastEntity
	field      [][]int
}

func (gs *GoastSystem) New(w *ecs.World) {
	gs.world = w
	fmt.Println("GoastSystem was added to the Scene")
	gs.field = Tiles
	gs.createGoasts()
	for _, system := range gs.world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			for _, goastEnity := range gs.goastEnity {
				sys.Add(&goastEnity.BasicEntity, &goastEnity.RenderComponent, &goastEnity.SpaceComponent)
			}
		case *common.AnimationSystem:
			for _, goastEnity := range gs.goastEnity {
				println("###############")
				sys.Add(&goastEnity.BasicEntity, &goastEnity.AnimationComponent, &goastEnity.RenderComponent)
			}
		}
	}

}

var AkabeAnimationPatternUp []string = []string{"icon/goast/akabe/akabe_up_pattern1.png", "icon/goast/akabe/akabe_up_pattern2.png"}
var AkabeAnimationPatternDown []string = []string{"icon/goast/akabe/akabe_down_pattern1.png", "icon/goast/akabe/akabe_down_pattern2.png"}
var AkabeAnimationPatternLeft []string = []string{"icon/goast/akabe/akabe_left_pattern1.png", "icon/goast/akabe/akabe_left_pattern2.png"}
var AkabeAnimationPatternRight []string = []string{"icon/goast/akabe/akabe_right_pattern1.png", "icon/goast/akabe/akabe_right_pattern2.png"}

func (gs *GoastSystem) createGoasts() {
	count := 0
	for y, row := range gs.field {
		for x, cell := range row {
			if count < 1 {
				if cell == 3 {
					positionX := 32 * x
					positionY := 32 * y
					goast := &GoastEntity{BasicEntity: ecs.NewBasic()}
					goast.name = "AKABE"
					goast.SpaceComponent = common.SpaceComponent{
						Position: engo.Point{X: float32(positionX), Y: float32(positionY)},
						Width:    32,
						Height:   32,
					}
					goast.animationPattern = make(map[string][]common.Drawable)
					drawablesUpPattern1, err := common.LoadedSprite("icon/goast/akabe/akabe_up_pattern1.png")
					if err != nil {
						fmt.Println("Unable to load texture: " + err.Error())
					}
					goast.RenderComponent = common.RenderComponent{
						Drawable: drawablesUpPattern1,
						Scale:    engo.Point{X: 2, Y: 2},
					}
					drawablesUpPattern2, err := common.LoadedSprite("icon/goast/akabe/akabe_up_pattern2.png")
					if err != nil {
						fmt.Println("Unable to load texture: " + err.Error())
					}
					drawablesDownPattern1, err := common.LoadedSprite("icon/goast/akabe/akabe_down_pattern1.png")
					if err != nil {
						fmt.Println("Unable to load texture: " + err.Error())
					}
					drawablesDownPattern2, err := common.LoadedSprite("icon/goast/akabe/akabe_down_pattern2.png")
					if err != nil {
						fmt.Println("Unable to load texture: " + err.Error())
					}
					drawablesLeftPattern1, err := common.LoadedSprite("icon/goast/akabe/akabe_left_pattern1.png")
					if err != nil {
						fmt.Println("Unable to load texture: " + err.Error())
					}
					drawablesLeftPattern2, err := common.LoadedSprite("icon/goast/akabe/akabe_left_pattern2.png")
					if err != nil {
						fmt.Println("Unable to load texture: " + err.Error())
					}
					drawablesRightPattern1, err := common.LoadedSprite("icon/goast/akabe/akabe_right_pattern1.png")
					if err != nil {
						fmt.Println("Unable to load texture: " + err.Error())
					}
					drawablesRightPattern2, err := common.LoadedSprite("icon/goast/akabe/akabe_right_pattern2.png")
					if err != nil {
						fmt.Println("Unable to load texture: " + err.Error())
					}
					drawablesUp := []common.Drawable{drawablesUpPattern1, drawablesUpPattern2}
					drawablesDown := []common.Drawable{drawablesDownPattern1, drawablesDownPattern2}
					drawablesLeft := []common.Drawable{drawablesLeftPattern1, drawablesLeftPattern2}
					drawablesRight := []common.Drawable{drawablesRightPattern1, drawablesRightPattern2}
					goast.animationPattern["UP"] = drawablesUp
					goast.animationPattern["DOWN"] = drawablesDown
					goast.animationPattern["LEFT"] = drawablesLeft
					goast.animationPattern["RIGHT"] = drawablesRight

					// 2. Define animations for each direction
					animationUp := &common.Animation{Name: "UP", Frames: []int{0, 1}, Loop: true}
					animationDown := &common.Animation{Name: "DOWN", Frames: []int{0, 1}, Loop: true}
					animationLeft := &common.Animation{Name: "LEFT", Frames: []int{0, 1}, Loop: true}
					animationRight := &common.Animation{Name: "RIGHT", Frames: []int{0, 1}, Loop: true}

					// 3. Create AnimationComponents for the ghost
					animationComponentGhost := common.NewAnimationComponent(drawablesUp, 0.3) // Initial direction is up

					// 4. Add animations to the AnimationComponent
					animationComponentGhost.AddAnimation(animationUp)
					animationComponentGhost.AddAnimation(animationDown)
					animationComponentGhost.AddAnimation(animationLeft)
					animationComponentGhost.AddAnimation(animationRight)

					// 5. Select the animation for the current direction
					currentDirection := "UP" // This should be updated whenever the ghost changes direction
					animationComponentGhost.SelectAnimationByName(currentDirection)
					goast.AnimationComponent = animationComponentGhost
					gs.goastEnity = append(gs.goastEnity, goast)
					count++
				}
			}
		}
	}
}

// func (gs *GoastSystem) CreateGoastAnimationComponent() common.AnimationComponent {
// 	drawablesUpPattern1, err := common.LoadedSprite("icon/goast/akabe/akabe_up_pattern1.png")
// 	if err != nil {
// 		fmt.Println("Unable to load texture: " + err.Error())
// 	}
// 	drawablesUpPattern2, err := common.LoadedSprite("icon/goast/akabe/akabe_up_pattern2.png")
// 	if err != nil {
// 		fmt.Println("Unable to load texture: " + err.Error())
// 	}
// 	drawablesUp := []common.Drawable{drawablesUpPattern1, drawablesUpPattern2}

// 	// 2. Define animations for each direction
// 	animationUp := &common.Animation{Name: "UP", Frames: []int{0, 1}, Loop: true}
// 	animationDown := &common.Animation{Name: "DOWN", Frames: []int{0, 1}, Loop: true}
// 	animationLeft := &common.Animation{Name: "LEFT", Frames: []int{0, 1}, Loop: true}
// 	animationRight := &common.Animation{Name: "RIGHT", Frames: []int{0, 1}, Loop: true}

// 	// 3. Create AnimationComponents for the ghost
// 	animationComponentGhost := common.NewAnimationComponent(drawablesUp, 0.1) // Initial direction is up

// 	// 4. Add animations to the AnimationComponent
// 	animationComponentGhost.AddAnimation(animationUp)
// 	animationComponentGhost.AddAnimation(animationDown)
// 	animationComponentGhost.AddAnimation(animationLeft)
// 	animationComponentGhost.AddAnimation(animationRight)

// 	// 5. Select the animation for the current direction
// 	currentDirection := "UP" // This should be updated whenever the ghost changes direction
// 	animationComponentGhost.SelectAnimationByName(currentDirection)
// 	return animationComponentGhost
// }

var array []string = []string{"UP", "DOWN", "LEFT", "RIGHT"}
var count int = 0
var index int = 0

func (gs *GoastSystem) Update(dt float32) {
	if count%100 == 0 {
		currentDirection := array[index] // This should be updated whenever the ghost changes direction
		for _, v := range gs.goastEnity {
			if v.name == "AKABE" {
				v.AnimationComponent.Drawables = v.animationPattern[currentDirection]
				v.AnimationComponent.SelectAnimationByName(currentDirection)
			}
		}
		index++
		if index == 4 {
			index = 0
		}
	}
	count++
}

// Remove takes an enitty out of the system.
func (st *GoastSystem) Remove(basic ecs.BasicEntity) {}
