package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/robertvitoriano/bulls-on-parade-golang/components"
)

type Player struct {
	gameObject components.GameObject
}

func NewPlayer() *Player {
	player := &Player{
		gameObject: components.GameObject{
			Animator: components.Animator{},
		},
	}

	player.gameObject.Animator.AddAnimation("walk-left", "character.png", 0, 0, "horizontal", components.FrameProperties{
		Width:  16,
		Height: 16,
		Count:  5,
	})

	player.gameObject.Animator.AddAnimation("walk-right", "character.png", 1, 0, "horizontal", components.FrameProperties{
		Width:  16,
		Height: 16,
		Count:  5,
	})

	player.gameObject.Animator.ChangeAnimation("walk-right")

	return player
}

func (p *Player) Update() {
	p.gameObject.Update()
}

func (p *Player) Draw(screen *ebiten.Image) {
	p.gameObject.Draw(screen)
}
