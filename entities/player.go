package entities

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/robertvitoriano/bulls-on-parade-golang/components"
)

type Player struct {
	GameObject components.GameObject
}

func NewPlayer() *Player {
	player := &Player{
		GameObject: components.GameObject{
			Animator: components.Animator{},
			Position: components.Position{},
			Size: components.Size{
				Width:  16,
				Height: 16,
			},
		},
	}

	player.GameObject.Animator.AddAnimation("walk-left", "character.png", 0, 0, "horizontal", components.FrameProperties{
		Width:  16,
		Height: 16,
		Count:  5,
	})

	player.GameObject.Animator.AddAnimation("walk-right", "character.png", 1, 0, "horizontal", components.FrameProperties{
		Width:  16,
		Height: 16,
		Count:  5,
	})

	player.GameObject.Animator.AddAnimation("walk-up", "character.png", 0, 7, "vertical", components.FrameProperties{
		Width:  16,
		Height: 16,
		Count:  2,
	})

	player.GameObject.Animator.AddAnimation("walk-down", "character.png", 0, 5, "vertical", components.FrameProperties{
		Width:  16,
		Height: 16,
		Count:  2,
	})

	player.GameObject.Animator.ChangeAnimation("walk-right")

	return player
}

func (p *Player) Update() {
	p.GameObject.Update()
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
	p.GameObject.Draw(screen)
}

func (p *Player) MoveRight() {
	p.GameObject.Animator.ChangeAnimation("walk-right")
	p.GameObject.Position.X++

}
func (p *Player) MoveLeft() {
	p.GameObject.Animator.ChangeAnimation("walk-left")
	p.GameObject.Position.X--

}
func (p *Player) MoveUp() {
	p.GameObject.Animator.ChangeAnimation("walk-up")
	p.GameObject.Position.Y--

}
func (p *Player) MoveDown() {
	p.GameObject.Animator.ChangeAnimation("walk-down")
	p.GameObject.Position.Y++

}

func (p *Player) HandleTileCollision() {
	fmt.Println("Collided with player")
}
