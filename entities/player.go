package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/robertvitoriano/bulls-on-parade-golang/components"
	"github.com/robertvitoriano/bulls-on-parade-golang/components/level"
)

const VELOCITY = 2

type Player struct {
	GameObject   components.GameObject
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
			Velocity: components.Velocity{
				X: VELOCITY,
				Y: VELOCITY,
			},
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
		p.GameObject.Velocity.X = VELOCITY
	}

	p.GameObject.Position.X += p.GameObject.Velocity.X

}
func (p *Player) MoveLeft() {
	p.GameObject.Animator.ChangeAnimation("walk-left")
	if p.collidedSide != "LEFT" {
		p.GameObject.Velocity.X = VELOCITY
	}
	p.GameObject.Position.X -= p.GameObject.Velocity.X

}
func (p *Player) MoveUp() {
	p.GameObject.Animator.ChangeAnimation("walk-up")
	if p.collidedSide != "TOP" {
		p.GameObject.Velocity.Y = VELOCITY
	}
	p.GameObject.Position.Y -= p.GameObject.Velocity.Y

}
func (p *Player) MoveDown() {
	p.GameObject.Animator.ChangeAnimation("walk-down")
	if p.collidedSide != "BOTTOM" {
		p.GameObject.Velocity.Y = VELOCITY
	}
	p.GameObject.Position.Y += p.GameObject.Velocity.Y

}

func (p *Player) HandleLevelCollisionsCollision(collisions []level.Collision) {
	if len(collisions) == 0 {
		p.collidedSide = "NONE"
		return
	}

	for _, collision := range collisions {

		if p.GameObject.GetRight() < collision.GameObject.GetRight() && p.collidedSide == "NONE" {
			p.collidedSide = "RIGHT"
			p.GameObject.Velocity.X = 0

		}
		if p.GameObject.GetLeft() > collision.GameObject.GetLeft() && p.collidedSide == "NONE" {
			p.collidedSide = "LEFT"
			p.GameObject.Velocity.X = 0
		}
		if p.GameObject.GetTop() > collision.GameObject.GetTop() && p.collidedSide == "NONE" {
			p.collidedSide = "TOP"
			p.GameObject.Velocity.Y = 0
		}
		if p.GameObject.GetBottom() < collision.GameObject.GetBottom() && p.collidedSide == "NONE" {
			p.collidedSide = "BOTTOM"
			p.GameObject.Velocity.Y = 0
		}
	}

}
