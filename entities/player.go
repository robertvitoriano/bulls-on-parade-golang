package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/robertvitoriano/bulls-on-parade-golang/components"
	"github.com/robertvitoriano/bulls-on-parade-golang/components/level"
)

const VELOCITY = 2
const GRAVITY float64 = 20
const GRAVITY_SMOTHING float64 = .4541561

type Player struct {
	GameObject   components.GameObject
	collidedSide components.CollisionSide
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
		collidedSide: components.CollisionNone,
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
	if p.collidedSide != components.CollisionBottom {
		p.GameObject.Position.Y += GRAVITY * GRAVITY_SMOTHING
	}
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
	if ebiten.IsKeyPressed(ebiten.KeySpace) && p.collidedSide == components.CollisionBottom {
		p.Jump()
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	p.GameObject.Draw(screen)
}

func (p *Player) MoveRight() {
	p.GameObject.Animator.ChangeAnimation("walk-right")

	if p.collidedSide != components.CollisionRight {
		p.GameObject.Velocity.X = VELOCITY
	}

	p.GameObject.Position.X += p.GameObject.Velocity.X

}
func (p *Player) Jump() {
	p.GameObject.Position.Y -= 35

}
func (p *Player) MoveLeft() {
	p.GameObject.Animator.ChangeAnimation("walk-left")
	if p.collidedSide != components.CollisionLeft {
		p.GameObject.Velocity.X = VELOCITY
	}
	p.GameObject.Position.X -= p.GameObject.Velocity.X

}
func (p *Player) MoveUp() {
	p.GameObject.Animator.ChangeAnimation("walk-up")
	if p.collidedSide != components.CollisionTop {
		p.GameObject.Velocity.Y = VELOCITY
	}
	p.GameObject.Position.Y -= p.GameObject.Velocity.Y

}
func (p *Player) MoveDown() {
	p.GameObject.Animator.ChangeAnimation("walk-down")
	if p.collidedSide != components.CollisionBottom {
		p.GameObject.Velocity.Y = VELOCITY
	}
	p.GameObject.Position.Y += p.GameObject.Velocity.Y

}

func (p *Player) HandleLevelCollisionsCollision(collisions []level.Collision) {
	p.collidedSide = components.CollisionNone

	for _, collision := range collisions {

		switch p.GameObject.GetCollisionSide(collision.GameObject) {
		case components.CollisionRight:
			p.collidedSide = components.CollisionRight
			p.GameObject.Position.X = collision.GameObject.GetLeft() - p.GameObject.Size.Width
			p.GameObject.Velocity.X = 0

		case components.CollisionLeft:
			p.collidedSide = components.CollisionLeft
			p.GameObject.Position.X = collision.GameObject.GetRight()
			p.GameObject.Velocity.X = 0

		case components.CollisionTop:
			p.collidedSide = components.CollisionTop
			p.GameObject.Position.Y = collision.GameObject.GetBottom()
			p.GameObject.Velocity.Y = 0

		case components.CollisionBottom:
			p.collidedSide = components.CollisionBottom
			p.GameObject.Position.Y = collision.GameObject.GetTop() - p.GameObject.Size.Height
			p.GameObject.Velocity.Y = 0
		}

	}
}
