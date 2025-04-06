package main

import (
	"image"
	_ "image/png"
	"log"

	"github.com/robertvitoriano/bulls-on-parade-golang/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 640
	screenHeight = 480

	frameWidth     = 16
	frameHeight    = 16
	frameCount     = 5
	animationSpeed = 5
	fps            = 30
)

type Game struct {
	animation    []*ebiten.Image
	currentFrame int
	tick         int
}

func (g *Game) Update() error {
	g.tick++

	if g.tick%animationSpeed == 0 {
		g.currentFrame = (g.currentFrame + 1) % frameCount
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(frameWidth)/2, -float64(frameHeight)/2)
	op.GeoM.Translate(screenWidth/2, screenHeight/2)

	screen.DrawImage(g.animation[g.currentFrame], op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)

	ebiten.SetWindowTitle("Animation")

	charImageAnimationSprite, err := utils.ReadImage("character.png")

	if err != nil {
		log.Fatal(err)
	}

	animation := make([]*ebiten.Image, frameCount)

	for i := 0; i < frameCount; i++ {
		xOffset := frameWidth * i
		sub := charImageAnimationSprite.SubImage(image.Rect(xOffset, 0, xOffset+frameWidth, frameHeight)).(*ebiten.Image)
		animation[i] = sub
	}

	if err := ebiten.RunGame(&Game{
		animation: animation,
	}); err != nil {
		log.Fatal(err)
	}
}
