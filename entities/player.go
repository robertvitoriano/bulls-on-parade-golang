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
	player.gameObject.Animator.AddAnimation("walk", "character.png", "horizontal", components.FrameProperties{
		Width:  16,
		Height: 16,
		Count:  5,
	})
	player.gameObject.Animator.ChangeAnimation("walk")

	return player
}

func (p *Player) Update() {
	p.gameObject.Update()
}

func (p *Player) Draw(screen *ebiten.Image) {
	p.gameObject.Draw(screen)
}
