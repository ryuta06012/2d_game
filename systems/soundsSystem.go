package systems

import (
	"image/color"
	"log"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

// HUDMoneyMessage updates the HUD text when changes are made to the amount of
// money available to the player
type HUDOpeningMessage struct {
	finishOpening bool
}

// HUDOpeningMessageType is the type for an HUDOpeningMessage
const HUDOpeningMessageType string = "HUDOpeningMessage"

// Type implements the engo.Message Interface
func (HUDOpeningMessage) Type() string {
	return HUDOpeningMessageType
}

type SoundEnity struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
	common.AudioComponent
}

type SoundSystem struct {
	sound *SoundEnity
}

func (ss *SoundSystem) New(w *ecs.World) {
	audio := &SoundEnity{BasicEntity: ecs.NewBasic()}
	player, err := common.LoadedPlayer("sounds/opening_audio.mp3")
	if err != nil {
		log.Println(err)
	}
	player.SetVolume(0.5)
	audio.AudioComponent = common.AudioComponent{Player: player}
	audio.Player.Rewind()
	audio.Player.Play()
	fnt := &common.Font{
		URL:  "emulogic.ttf",
		FG:   color.RGBA{255, 255, 0, 255},
		Size: 28,
	}
	fnt.CreatePreloaded()
	audio.RenderComponent = common.RenderComponent{
		Drawable: common.Text{
			Font: fnt,
			Text: "READY!",
		},
		Scale: engo.Point{X: 0.8, Y: 0.8},
	}
	audio.SetShader(common.TextHUDShader)
	audio.RenderComponent.SetZIndex(1002)
	audio.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{X: 32 * 8.5, Y: 32 * 12},
	}
	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.AudioSystem:
			sys.Add(&audio.BasicEntity, &audio.AudioComponent)
		case *common.RenderSystem:
			sys.Add(&audio.BasicEntity, &audio.RenderComponent, &audio.SpaceComponent)
		}
	}
	audio.Player.Current()
	ss.sound = audio
}

func (ss *SoundSystem) Update(dt float32) {
	if !ss.sound.Player.IsPlaying() {
		ss.sound.RenderComponent.Hidden = true
		engo.Mailbox.Dispatch(HUDOpeningMessage{
			finishOpening: true,
		})
	}
}

func (ss *SoundSystem) Remove(basic ecs.BasicEntity) {}
