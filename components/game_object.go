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
type GameObject struct {
	Animator Animator
	Position Position
	Size     Size
}

func (gameObject *GameObject) Update() {
	gameObject.Animator.Update()
}

func (gameObject *GameObject) Draw(screen *ebiten.Image) {
	gameObject.Animator.Draw(screen, gameObject.Position)
}

func (gameObject *GameObject) CollidesWith(other GameObject) bool {
	return gameObject.getRight() >= other.getLeft() &&
		gameObject.getLeft() <= other.getRight() &&
		gameObject.getTop() <= other.getBottom() &&
		gameObject.getBottom() >= other.getTop()
}

func (gameObject *GameObject) getLeft() float64 {
	return gameObject.Position.X
}

func (gameObject *GameObject) getRight() float64 {
	return gameObject.Position.X + gameObject.Size.Width
}

func (gameObject *GameObject) getTop() float64 { return gameObject.Position.Y }

func (gameObject *GameObject) getBottom() float64 {
	return gameObject.Position.Y + gameObject.Size.Height
}
