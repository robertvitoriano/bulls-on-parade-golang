package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/robertvitoriano/bulls-on-parade-golang/components"
	"github.com/robertvitoriano/bulls-on-parade-golang/physics"
	"github.com/robertvitoriano/bulls-on-parade-golang/utils"
)

const SPEED = 2.5
const GRAVITY float64 = 20
const GRAVITY_SMOTHING float64 = .4541561
const JUMP_SPEED float64 = 40
const FALL_SPEED float64 = GRAVITY / 2

type Player struct {
	GameObject       components.GameObject
	collidedSide     utils.CollisionSide
	isJumping        bool
	XMovementEnabled bool
}

func NewPlayer() *Player {
	player := &Player{
		GameObject: components.GameObject{
			Animator: components.Animator{},
			Position: utils.Vector2{
				X: 0,
				Y: 0,
			},
			Size: utils.Size{
				Width:  16,
				Height: 16,
			},
			Velocity: utils.Vector2{
				X: 0,
				Y: 0,
			},
		},
		collidedSide:     utils.CollisionNone,
		XMovementEnabled: true,
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
	p.handleGravity()
}

func (p *Player) handleGravity() {
	if p.collidedSide != utils.CollisionBottom {
		if p.isJumping {
			p.GameObject.Position.Y += FALL_SPEED * GRAVITY_SMOTHING
			return
		}
		p.isJumping = false
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
	} else {
		p.GameObject.Velocity.X = 0
	}
	if ebiten.IsKeyPressed(ebiten.KeySpace) && p.collidedSide == utils.CollisionBottom {
		p.Jump()
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	p.GameObject.Draw(screen)
}

func (p *Player) MoveRight() {
	p.GameObject.Animator.ChangeAnimation("walk-right")

	if p.collidedSide != utils.CollisionRight {
		p.GameObject.Velocity.X = SPEED
	}
	if !p.XMovementEnabled {
		return
	}
	p.GameObject.Position.X += p.GameObject.Velocity.X

}
func (p *Player) Jump() {
	p.GameObject.Position.Y -= 35
	p.isJumping = true

}
func (p *Player) MoveLeft() {
	p.GameObject.Animator.ChangeAnimation("walk-left")

	if p.collidedSide != utils.CollisionLeft {
		p.GameObject.Velocity.X = -SPEED
	}
	if !p.XMovementEnabled {
		return
	}
	p.GameObject.Position.X += p.GameObject.Velocity.X

}
func (p *Player) MoveUp() {
	p.GameObject.Animator.ChangeAnimation("walk-up")
	if p.collidedSide != utils.CollisionTop {
		p.GameObject.Velocity.Y = SPEED
	}

	p.GameObject.Position.Y -= p.GameObject.Velocity.Y

}
func (p *Player) MoveDown() {
	p.GameObject.Animator.ChangeAnimation("walk-down")
	if p.collidedSide != utils.CollisionBottom {
		p.GameObject.Velocity.Y = SPEED
	}

	p.GameObject.Position.Y += p.GameObject.Velocity.Y

}

func (p *Player) HandleLevelCollisions(collisions []*physics.Collision) {
	p.collidedSide = utils.CollisionNone

	for _, collision := range collisions {

		switch p.GameObject.GetCollisionSide(collision.GameObject) {
		case utils.CollisionRight:
			p.collidedSide = utils.CollisionRight
			p.GameObject.Position.X = collision.GameObject.GetLeft() - p.GameObject.Size.Width
			p.GameObject.Velocity.X = 0

		case utils.CollisionLeft:
			p.collidedSide = utils.CollisionLeft
			p.GameObject.Position.X = collision.GameObject.GetRight()
			p.GameObject.Velocity.X = 0

		case utils.CollisionTop:
			p.collidedSide = utils.CollisionTop
			p.GameObject.Position.Y = collision.GameObject.GetBottom()
			p.GameObject.Velocity.Y = 0

		case utils.CollisionBottom:
			p.collidedSide = utils.CollisionBottom
			p.GameObject.Position.Y = collision.GameObject.GetTop() - p.GameObject.Size.Height
			p.GameObject.Velocity.Y = 0
		}

	}
}
