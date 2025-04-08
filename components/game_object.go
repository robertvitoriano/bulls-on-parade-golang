package components

import "github.com/hajimehoshi/ebiten/v2"

type Position struct {
	X float64
	Y float64
}
type Size struct {
	Width  float64
	Height float64
}
type Velocity struct {
	X float64
	Y float64
}
type GameObject struct {
	Animator Animator
	Position Position
	Size     Size
	Velocity Velocity
}

func (g *GameObject) Update() {
	g.Animator.Update()
}

func (g *GameObject) Draw(screen *ebiten.Image) {
	g.Animator.Draw(screen, g.Position)
}

func (g *GameObject) CollidesWith(other GameObject) bool {
	return g.GetRight() >= other.GetLeft() &&
		g.GetLeft() <= other.GetRight() &&
		g.GetTop() <= other.GetBottom() &&
		g.GetBottom() >= other.GetTop()
}

func (g *GameObject) GetLeft() float64 {
	return g.Position.X
}

func (g *GameObject) GetRight() float64 {
	return g.Position.X + g.Size.Width
}

func (g *GameObject) GetTop() float64 { return g.Position.Y }

func (g *GameObject) GetBottom() float64 {
	return g.Position.Y + g.Size.Height
}

func (g *GameObject) GetCollisionSide(other GameObject) string {

	collisionSide := "NONE"

	if g.GetRight() < other.GetRight() && g.GetBottom() >= other.GetTop() {
		collisionSide = "RIGHT"
	} else if g.GetLeft() > other.GetLeft() {
		collisionSide = "LEFT"
	} else if g.GetTop() > other.GetTop() && g.GetBottom() >= other.GetTop() {
		collisionSide = "TOP"
	} else if g.GetBottom() < other.GetBottom() {
		collisionSide = "BOTTOM"
	}
	return collisionSide
}
