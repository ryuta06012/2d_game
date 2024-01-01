package systems

import (
	"log"

	"github.com/EngoEngine/ecs"
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

type Sound struct {
	ecs.BasicEntity
	common.SpaceComponent
	common.RenderComponent
}

type SoundSystem struct {
	text Score
}

func (ss *SoundSystem) New(w *ecs.World) {
	audio := &Audio{BasicEntity: ecs.NewBasic()}
	player, err := common.LoadedPlayer(AudioFile[i])
	if err != nil {
		log.Println(err)
	}
	player.SetVolume(0.5)
	audio.AudioComponent = common.AudioComponent{Player: player}
}

func (ss *SoundSystem) Update(dt float32) {}

func (ss *SoundSystem) Remove(basic ecs.BasicEntity) {}
