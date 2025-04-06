package main

import "github.com/hajimehoshi/ebiten/v2"

type GameObject struct {
	animator *Animator
}

func (g *GameObject) Update() {
	g.animator.Update()
}

func (g *GameObject) Draw(screen *ebiten.Image) {
	g.animator.Draw(screen)
}
