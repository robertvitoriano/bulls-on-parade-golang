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
	centerAX := g.Position.X + g.Size.Width/2
	centerAY := g.Position.Y + g.Size.Height/2
	centerBX := other.Position.X + other.Size.Width/2
	centerBY := other.Position.Y + other.Size.Height/2

	horizontalCenterDistance := centerBX - centerAX

	overlapX := (g.Size.Width+other.Size.Width)/2 - abs(horizontalCenterDistance)

	if overlapX <= 0 {
		return "NONE"
	}

	verticalCenterDistance := centerBY - centerAY

	overlapY := (g.Size.Height+other.Size.Height)/2 - abs(verticalCenterDistance)

	if overlapY <= 0 {
		return "NONE"
	}

	if overlapX < overlapY {
		if horizontalCenterDistance > 0 {
			return "RIGHT"
		}
		return "LEFT"
	} else {
		if verticalCenterDistance > 0 {
			return "BOTTOM"
		}
		return "TOP"
	}
}

func abs(f float64) float64 {
	if f < 0 {
		return -f
	}
	return f
}
