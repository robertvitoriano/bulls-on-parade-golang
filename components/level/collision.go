package level

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/robertvitoriano/bulls-on-parade-golang/components"
)

type Collision struct {
	GameObject components.GameObject
}

func (c *Collision) DebugDraw(screen *ebiten.Image) {
	w := int(c.GameObject.Size.Width)
	h := int(c.GameObject.Size.Height)
	x := c.GameObject.Position.X
	y := c.GameObject.Position.Y

	borderColor := color.RGBA{R: 255, G: 0, B: 0, A: 255}
	borderThickness := 1

	top := ebiten.NewImage(w, borderThickness)
	top.Fill(borderColor)
	topOpts := &ebiten.DrawImageOptions{}
	topOpts.GeoM.Translate(x, y)
	screen.DrawImage(top, topOpts)

	bottom := ebiten.NewImage(w, borderThickness)
	bottom.Fill(borderColor)
	bottomOpts := &ebiten.DrawImageOptions{}
	bottomOpts.GeoM.Translate(x, y+float64(h-borderThickness))
	screen.DrawImage(bottom, bottomOpts)

	left := ebiten.NewImage(borderThickness, h)
	left.Fill(borderColor)
	leftOpts := &ebiten.DrawImageOptions{}
	leftOpts.GeoM.Translate(x, y)
	screen.DrawImage(left, leftOpts)

	right := ebiten.NewImage(borderThickness, h)
	right.Fill(borderColor)
	rightOpts := &ebiten.DrawImageOptions{}
	rightOpts.GeoM.Translate(x+float64(w-borderThickness), y)
	screen.DrawImage(right, rightOpts)
}
