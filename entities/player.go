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

	player.gameObject.Animator.AddAnimation("walk-up", "character.png", 0, 7, "vertical", components.FrameProperties{
		Width:  16,
		Height: 16,
		Count:  2,
	})

	player.gameObject.Animator.AddAnimation("walk-down", "character.png", 0, 5, "vertical", components.FrameProperties{
		Width:  16,
		Height: 16,
		Count:  2,
	})

	player.gameObject.Animator.ChangeAnimation("walk-right")

	return player
}

func (p *Player) Update() {
	p.gameObject.Update()
	p.Move()
}

func (p *Player) Move() {
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.MoveRight()
	} else if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.MoveLeft()
	} else if ebiten.IsKeyPressed(ebiten.KeyUp) {
		p.MoveUp()
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
		p.MoveDown()
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	p.gameObject.Draw(screen)
}

func (p *Player) MoveRight() {
	p.gameObject.Animator.ChangeAnimation("walk-right")
	p.gameObject.Position.X++

}
func (p *Player) MoveLeft() {
	p.gameObject.Animator.ChangeAnimation("walk-left")
	p.gameObject.Position.X--

}
func (p *Player) MoveUp() {
	p.gameObject.Animator.ChangeAnimation("walk-up")
	p.gameObject.Position.Y--

}
func (p *Player) MoveDown() {
	p.gameObject.Animator.ChangeAnimation("walk-down")
	p.gameObject.Position.Y++

}
