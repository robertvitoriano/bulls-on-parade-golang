package level

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/robertvitoriano/bulls-on-parade-golang/components"
)

type Tile struct {
	GameObject components.GameObject
	image      *ebiten.Image
}

func (t *Tile) draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(
		t.GameObject.Position.X,
		t.GameObject.Position.Y,
	)
	screen.DrawImage(t.image, op)
}
