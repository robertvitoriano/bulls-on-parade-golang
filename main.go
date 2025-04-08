package main

import (
	_ "image/png"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/robertvitoriano/bulls-on-parade-golang/components/level"
	"github.com/robertvitoriano/bulls-on-parade-golang/entities"
)

const (
	screenWidth  = 320
	screenHeight = 256
	fps          = 10
)

type Game struct {
	player *entities.Player
	level  *level.Level
}

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
		g.handleUpdate()
	}
	return nil
}

func (g *Game) handleUpdate() {
	g.player.Update()
	g.level.CheckTileCollisions(g.player.GameObject)

}
func (g *Game) Draw(screen *ebiten.Image) {
	g.level.Draw(screen)
	g.player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)

	ebiten.SetWindowTitle("Animation")

	player := entities.NewPlayer()

	if player == nil {
		log.Fatal("Player is nil!")
	}

	level := level.NewLevel("content/maps/map_1.json")

	if err := ebiten.RunGame(&Game{
		player: player,
		level:  level,
	}); err != nil {
		log.Fatal(err)
	}
}
