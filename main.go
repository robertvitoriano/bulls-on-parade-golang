package main

import (
	_ "image/png"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 640
	screenHeight = 480
	fps          = 10
	frameCount   = 5
	frameWidth   = 16
	frameHeight  = 16
)

type Game struct {
	player GameObject
}

var currentFrameIndex int
var timeToUpdate int64

func (g *Game) Update() error {
	durationPerFrame := time.Second / time.Duration(fps)
	durationPerFrameMs := durationPerFrame.Milliseconds()

	now := time.Now().UnixMilli()

	if timeToUpdate == 0 {
		timeToUpdate = now + durationPerFrameMs
	}

	if now >= timeToUpdate {
		nextTimeToUpdate := now + durationPerFrameMs
		timeToUpdate = nextTimeToUpdate
		currentFrameIndex = (currentFrameIndex + 1) % frameCount
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(40, 30)

	if len(g.player.animator.animations) > 0 {
		screen.DrawImage(g.player.animator.animations["walk"][currentFrameIndex], op)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)

	ebiten.SetWindowTitle("Animation")

	player := GameObject{
		animator: &Animator{},
	}

	player.animator.AddAnimation("walk", "character.png", "horizontal", FrameProperties{
		width:  frameWidth,
		height: frameHeight,
		count:  frameCount,
	})

	if err := ebiten.RunGame(&Game{player: player}); err != nil {
		log.Fatal(err)
	}
}
