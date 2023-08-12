package systems

import (
	"fmt"
	"image/color"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

var tiles = [][]int{
	{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	{1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1},
	{1, 2, 1, 1, 1, 2, 1, 1, 1, 2, 1, 2, 1, 1, 1, 2, 1, 1, 1, 2, 1},
	{1, 2, 1, 1, 1, 2, 1, 1, 1, 2, 1, 2, 1, 1, 1, 2, 1, 1, 1, 2, 1},
	{1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1},
	{1, 2, 1, 1, 1, 2, 1, 2, 1, 1, 1, 1, 1, 2, 1, 2, 1, 1, 1, 2, 1},
	{1, 2, 2, 2, 2, 2, 1, 2, 2, 2, 1, 2, 2, 2, 1, 2, 2, 2, 2, 2, 1},
	{1, 1, 1, 1, 1, 2, 1, 1, 1, 2, 1, 2, 1, 1, 1, 2, 1, 1, 1, 1, 1},
	{0, 0, 0, 0, 1, 2, 1, 2, 2, 2, 2, 2, 2, 2, 1, 2, 1, 0, 0, 0, 0},
	{1, 1, 1, 1, 1, 2, 1, 2, 1, 1, 2, 1, 1, 2, 1, 2, 1, 1, 1, 1, 1},
	{2, 2, 2, 2, 2, 2, 2, 2, 1, 2, 2, 2, 1, 2, 2, 2, 2, 2, 2, 2, 2},
	{1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 2, 2, 1, 2, 1, 2, 1, 1, 1, 1, 1},
	{0, 0, 0, 0, 1, 2, 1, 2, 1, 1, 1, 1, 1, 2, 1, 2, 1, 0, 0, 0, 0},
	{0, 0, 0, 0, 1, 2, 1, 2, 2, 2, 2, 2, 2, 2, 1, 2, 1, 0, 0, 0, 0},
	{1, 1, 1, 1, 1, 2, 2, 2, 1, 1, 1, 1, 1, 2, 2, 2, 1, 1, 1, 1, 1},
	{1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1},
	{1, 2, 1, 1, 1, 2, 1, 1, 1, 2, 1, 2, 1, 1, 1, 2, 1, 1, 1, 2, 1},
	{1, 2, 2, 2, 1, 2, 2, 2, 2, 2, 1, 2, 2, 2, 2, 2, 1, 2, 2, 2, 1},
	{1, 1, 2, 2, 1, 2, 1, 2, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 2, 1, 1},
	{1, 2, 2, 2, 2, 2, 1, 2, 2, 2, 1, 2, 2, 2, 1, 2, 2, 2, 2, 2, 1},
	{1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1},
	{1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1},
	{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
}

// Tileはマップを表示するためのtile
type Tile struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

// マップを作成するためのsystem
type TileBuildingSystem struct {
	world *ecs.World
	tile  *Tile
}

// New is the initialisation of the System
func (cb *TileBuildingSystem) New(w *ecs.World) {
	cb.world = w
	fmt.Println("TileBuildingSystem was added to the Scene")
	cb.generateFields()
}

// Update is ran every frame, with `dt` being the time
// in seconds since the last frame
func (cb *TileBuildingSystem) generateFields() {
	tileSize := float32(32)
	wallSpaceWidth := tileSize / 1.6
	wallOffset := (tileSize - wallSpaceWidth) / 2
	wallInnerColor := color.RGBA{0, 0, 0, 255}
	fieldsTiles := make([]*Tile, 0)

	for y, row := range tiles {
		for x, cell := range row {
			if cell == 1 {
				tile := &Tile{BasicEntity: ecs.NewBasic()}
				tile.SpaceComponent = common.SpaceComponent{
					Position: engo.Point{X: float32(x) * tileSize, Y: float32(y) * tileSize},
					Width:    tileSize,
					Height:   tileSize,
				}
				tile.RenderComponent = common.RenderComponent{
					Drawable: common.Rectangle{},
					Color:    color.RGBA{0, 0, 255, 255},
					Scale:    engo.Point{X: 1, Y: 1},
				}
				fieldsTiles = append(fieldsTiles, tile)
				if x > 0 && tiles[y][x-1] == 1 {
					blackTile := &Tile{BasicEntity: ecs.NewBasic()}
					blackTile.SpaceComponent = common.SpaceComponent{
						Position: engo.Point{X: float32(x) * tileSize, Y: float32(y)*tileSize + wallOffset},
						Width:    wallSpaceWidth + wallOffset,
						Height:   wallSpaceWidth,
					}
					blackTile.RenderComponent = common.RenderComponent{
						Drawable: common.Rectangle{},
						Color:    wallInnerColor,
						Scale:    engo.Point{X: 1, Y: 1},
					}
					fieldsTiles = append(fieldsTiles, blackTile)
				}
				if x < len(tiles[0])-1 && tiles[y][x+1] == 1 {
					blackTile := &Tile{BasicEntity: ecs.NewBasic()}
					blackTile.SpaceComponent = common.SpaceComponent{
						Position: engo.Point{X: float32(x)*tileSize + wallOffset, Y: float32(y)*tileSize + wallOffset},
						Width:    wallSpaceWidth + wallOffset,
						Height:   wallSpaceWidth,
					}
					blackTile.RenderComponent = common.RenderComponent{
						Drawable: common.Rectangle{},
						Color:    wallInnerColor,
						Scale:    engo.Point{X: 1, Y: 1},
					}
					fieldsTiles = append(fieldsTiles, blackTile)
				}
				if y < len(tiles)-1 && tiles[y+1][x] == 1 {
					blackTile := &Tile{BasicEntity: ecs.NewBasic()}
					blackTile.SpaceComponent = common.SpaceComponent{
						Position: engo.Point{X: float32(x)*tileSize + wallOffset, Y: float32(y)*tileSize + wallOffset},
						Width:    wallSpaceWidth,
						Height:   wallSpaceWidth + wallOffset,
					}
					blackTile.RenderComponent = common.RenderComponent{
						Drawable: common.Rectangle{},
						Color:    wallInnerColor,
						Scale:    engo.Point{X: 1, Y: 1},
					}
					fieldsTiles = append(fieldsTiles, blackTile)
				}

				if y > 0 && tiles[y-1][x] == 1 {
					blackTile := &Tile{BasicEntity: ecs.NewBasic()}
					blackTile.SpaceComponent = common.SpaceComponent{
						Position: engo.Point{X: float32(x)*tileSize + wallOffset, Y: float32(y) * tileSize},
						Width:    wallSpaceWidth,
						Height:   wallSpaceWidth + wallOffset,
					}
					blackTile.RenderComponent = common.RenderComponent{
						Drawable: common.Rectangle{},
						Color:    wallInnerColor,
						Scale:    engo.Point{X: 1, Y: 1},
					}
					fieldsTiles = append(fieldsTiles, blackTile)
				}
			} else if cell == 2 {
				continue
			}
		}
	}

	for _, system := range cb.world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			for _, v := range fieldsTiles {
				sys.Add(&v.BasicEntity, &v.RenderComponent, &v.SpaceComponent)
			}
		}
	}
}

// Update is ran every frame, with `dt` being the time
// in seconds since the last frame
func (cb *TileBuildingSystem) Update(dt float32) {}

// Update is ran every frame, with `dt` being the time
// in seconds since the last frame
func (*TileBuildingSystem) Remove(ecs.BasicEntity) {}
