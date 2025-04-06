package components

import "github.com/hajimehoshi/ebiten/v2"

type Position struct {
	X float64
	Y float64
}
type GameObject struct {
	Animator Animator
	Position Position
}

func (g *GameObject) Update() {
	g.Animator.Update()
}

func (g *GameObject) Draw(screen *ebiten.Image) {
	g.Animator.Draw(screen, g.Position)
}
