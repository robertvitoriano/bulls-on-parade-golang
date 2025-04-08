package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/robertvitoriano/bulls-on-parade-golang/components"
	"github.com/robertvitoriano/bulls-on-parade-golang/components/level"
)

const VELOCITY = 2

type Velocity struct {
	x float64
	y float64
}
type Player struct {
	GameObject   components.GameObject
	velocity     Velocity
	collidedSide string
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
		velocity: Velocity{
			x: VELOCITY,
			y: VELOCITY,
		},
		collidedSide: "NONE",
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

	if p.collidedSide != "RIGHT" {
		p.velocity.x = VELOCITY
	}

	p.GameObject.Position.X += p.velocity.x

}
func (p *Player) MoveLeft() {
	p.GameObject.Animator.ChangeAnimation("walk-left")
	if p.collidedSide != "LEFT" {
		p.velocity.x = VELOCITY
	}
	p.GameObject.Position.X -= p.velocity.x

}
func (p *Player) MoveUp() {
	p.GameObject.Animator.ChangeAnimation("walk-up")
	if p.collidedSide != "BOTTOM" {
		p.velocity.x = VELOCITY
	}
	p.GameObject.Position.Y -= p.velocity.y

}
func (p *Player) MoveDown() {
	p.GameObject.Animator.ChangeAnimation("walk-down")
	if p.collidedSide != "TOP" {
		p.velocity.x = VELOCITY
	}
	p.GameObject.Position.Y += p.velocity.y

}

func (p *Player) HandleLevelCollisionsCollision(collisions []level.Collision) {
	for _, collision := range collisions {

		if p.GameObject.GetRight() < collision.GameObject.GetRight() {
			p.collidedSide = "RIGHT"
			p.velocity.x = 0
			return
		} else if p.GameObject.GetLeft() < collision.GameObject.GetLeft() {
			p.collidedSide = "LEFT"
			p.velocity.x = 0
			return
		}
	}
	p.collidedSide = "NONE"
}
