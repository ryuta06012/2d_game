package systems

import (
	"fmt"
	"image/color"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

// HUDMoneyMessage updates the HUD text when changes are made to the amount of
// money available to the player
type HUDScoreMessage struct {
	Score int
}

// HUDTextMessageType is the type for an HUDTextMessage
const HUDScoreMessageType string = "HUDTextMessage"

// Type implements the engo.Message Interface
func (HUDScoreMessage) Type() string {
	return HUDScoreMessageType
}

type Score struct {
	ecs.BasicEntity
	common.SpaceComponent
	common.RenderComponent
}

type ScoreSystem struct {
	text        Score
	scoreUpdate bool
	score       int
}

func (st *ScoreSystem) New(w *ecs.World) {
	fnt := &common.Font{
		URL:  "go.ttf",
		FG:   color.White,
		Size: 32,
	}
	fnt.CreatePreloaded()
	st.text = Score{BasicEntity: ecs.NewBasic()}
	st.text.RenderComponent.Drawable = common.Text{
		Font: fnt,
		Text: "0",
	}
	st.text.SetShader(common.TextHUDShader)
	st.text.RenderComponent.SetZIndex(1001)
	st.text.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{X: 0, Y: engo.WindowHeight() - 32*3},
	}
	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&st.text.BasicEntity, &st.text.RenderComponent, &st.text.SpaceComponent)
		}
	}
	engo.Mailbox.Listen(HUDScoreMessageType, func(m engo.Message) {
		st.score++
		sm, ok := m.(HUDScoreMessage)
		if !ok {
			return
		}
		fmt.Printf("#####################sm.Score: %v\n", sm.Score)
		txt := st.text.RenderComponent.Drawable.(common.Text)
		txt.Text = fmt.Sprintf("SCORE: %v", st.score)
		st.text.RenderComponent.Drawable = common.Text{
			Font: fnt,
			Text: txt.Text,
		}
	})
}

func (st *ScoreSystem) Update(dt float32) {

}

// Remove takes an enitty out of the system.
func (st *ScoreSystem) Remove(basic ecs.BasicEntity) {}
