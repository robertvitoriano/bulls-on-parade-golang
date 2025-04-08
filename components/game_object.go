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

func (gameObject *GameObject) Update() {
	gameObject.Animator.Update()
}

func (gameObject *GameObject) Draw(screen *ebiten.Image) {
	gameObject.Animator.Draw(screen, gameObject.Position)
}

func (gameObject *GameObject) CollidesWith(other GameObject) bool {
	return gameObject.GetRight() >= other.GetLeft() &&
		gameObject.GetLeft() <= other.GetRight() &&
		gameObject.GetTop() <= other.GetBottom() &&
		gameObject.GetBottom() >= other.GetTop()
}

func (gameObject *GameObject) GetLeft() float64 {
	return gameObject.Position.X
}

func (gameObject *GameObject) GetRight() float64 {
	return gameObject.Position.X + gameObject.Size.Width
}

func (gameObject *GameObject) GetTop() float64 { return gameObject.Position.Y }

func (gameObject *GameObject) GetBottom() float64 {
	return gameObject.Position.Y + gameObject.Size.Height
}
