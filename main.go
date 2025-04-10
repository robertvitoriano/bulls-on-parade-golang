package main

import (
	_ "image/png"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/robertvitoriano/bulls-on-parade-golang/entities"
	"github.com/robertvitoriano/bulls-on-parade-golang/level"
	"github.com/robertvitoriano/bulls-on-parade-golang/utils"
)

const (
	ScreenWidth  = 320
	ScreenHeight = 256
	FPS          = 15
)

type Game struct {
	player *entities.Player
	level  *level.Level
}

var timeToUpdate int64

func (g *Game) Update() error {

	durationPerFrame := time.Second / time.Duration(FPS)
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
	collisions := g.level.GetLevelCollisions(g.player.GameObject)
	g.player.HandleLevelCollisions(collisions)
	g.level.Update()

}
func (g *Game) Draw(screen *ebiten.Image) {
	g.level.Draw(screen)
	g.player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func main() {
	ebiten.SetWindowSize(ScreenWidth*utils.SCALE, ScreenHeight*utils.SCALE)

	ebiten.SetWindowTitle("Animation")

	player := entities.NewPlayer()

	if player == nil {
		log.Fatal("Player is nil!")
	}

	level := level.NewLevel("content/maps/map_1.json", player)

	player.GameObject.Position = level.PlayerSpawnPosition

	if err := ebiten.RunGame(&Game{
		player: player,
		level:  level,
	}); err != nil {
		log.Fatal(err)
	}
}
