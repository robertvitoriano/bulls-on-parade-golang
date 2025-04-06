package components

import "github.com/hajimehoshi/ebiten/v2"

type GameObject struct {
	Animator Animator
}

func (g *GameObject) Update() {
	g.Animator.Update()
}

func (g *GameObject) Draw(screen *ebiten.Image) {
	g.Animator.Draw(screen)
}
