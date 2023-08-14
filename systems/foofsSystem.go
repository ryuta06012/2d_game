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
}

type FoodSystem struct {
	world *ecs.World
	food  *Food
	isAte bool
}

func (f *FoodSystem) New(w *ecs.World) {
	f.world = w
	fmt.Println("TileBuildingSystem was added to the Scene")
	f.generateFoodsInFields()
}

func (f *FoodSystem) generateFoodsInFields() {
	tileSize := float32(32)
	foodColor := color.RGBA{254, 184, 151, 255}
	fieldsFoods := make([]*Food, 0)
	for y, row := range Tiles {
		for x, cell := range row {
			if cell == 2 {
				food := &Food{BasicEntity: ecs.NewBasic()}
				food.SpaceComponent = common.SpaceComponent{
					Position: engo.Point{X: float32(x) * tileSize + tileSize/4, Y: float32(y) * tileSize + tileSize/4},
					Width:    tileSize,
					Height:   tileSize,
				}
				food.RenderComponent = common.RenderComponent{
					Drawable: common.Rectangle{},
					Color:    foodColor,
					Scale:    engo.Point{X: 0.3, Y: 0.3},
				}
				fieldsFoods = append(fieldsFoods, food)
			}
		}
	}
	for _, system := range f.world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			for _, v := range fieldsFoods {
				sys.Add(&v.BasicEntity, &v.RenderComponent, &v.SpaceComponent)
			}
		}
	}
}

// Update is ran every frame, with `dt` being the time
// in seconds since the last frame
func (*FoodSystem) Update(dt float32) {

}

// Update is ran every frame, with `dt` being the time
// in seconds since the last frame
func (*FoodSystem) Remove(ecs.BasicEntity) {}
