package components

import (
	"encoding/json"
	"image"
	"io"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Animation struct {
	Duration int `json:"duration"`
	TileId   int `json:"tileid"`
}

type RawTile struct {
	Animation []Animation `json:"animation"`
	Id        int         `json:"id"`
}
type TileSet struct {
	Columns     int       `json:"columns"`
	FirstGID    int       `json:"firstgid"`
	Image       string    `json:"image"`
	ImageHeight int       `json:"imageheight"`
	ImageWidth  int       `json:"imagewidth"`
	Margin      int       `json:"margin"`
	Name        string    `json:"name"`
	Spacing     int       `json:"spacing"`
	TileCount   int       `json:"tilecount"`
	TileHeight  int       `json:"tileheight"`
	TileWidth   int       `json:"tilewidth"`
	Tiles       []RawTile `json:"tiles"`
}
type PolylinePoint struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type ObjectProperty struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

type Object struct {
	Id         int              `json:"id"`
	Height     float64          `json:"height"`
	Width      float64          `json:"width"`
	X          float64          `json:"x"`
	Y          float64          `json:"y"`
	Visible    bool             `json:"visible"`
	Name       string           `json:"name"`
	Rotation   float64          `json:"rotation"`
	Properties []ObjectProperty `json:"properties"`
	Polyline   []PolylinePoint  `json:"polyline"`
	Ellipse    bool             `json:"ellipse"`
	Type       string           `json:"type"`
}

type Layer struct {
	Data      []int    `json:"data"`
	Height    int      `json:"height"`
	Width     int      `json:"width"`
	Id        int      `json:"id"`
	Name      string   `json:"name"`
	Opacity   int      `json:"opacity"`
	Type      string   `json:"type"`
	Visible   bool     `json:"visible"`
	X         float64  `json:"x"`
	Y         float64  `json:"y"`
	DrawOrder string   `json:"draworder"`
	Objects   []Object `json:"objects"`
	Color     string   `json:"color"`
}

type LevelData struct {
	Width            float64   `json:"width"`
	Height           float64   `json:"height"`
	Version          string    `json:"version"`
	Type             string    `json:"type"`
	CompressionLevel int       `json:"compressionlevel"`
	Infinite         bool      `json:"infinite"`
	NextLayerId      int       `json:"nextlayerid"`
	NextObjectId     int       `json:"nextobjectid"`
	Orientation      string    `json:"orientation"`
	RenderOrder      string    `json:"renderorder"`
	TiledVersion     string    `json:"tiledversion"`
	TileHeight       int       `json:"tileheight"`
	TileWidth        int       `json:"tilewidth"`
	Layers           []Layer   `json:"layers"`
	TileSets         []TileSet `json:"tilesets"`
}
type Level struct {
	Width         float64
	Height        float64
	TileWidth     float64
	TileHeight    float64
	X             float64
	Y             float64
	Layers        []Layer
	TileSets      []TileSet
	TilesetImages map[int]*ebiten.Image
	tiles         []Tile
}

type Tile struct {
	X     float64
	Y     float64
	image *ebiten.Image
}

func NewLevel(filePath string) *Level {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	jsonFileData, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	var parsedData LevelData
	err = json.Unmarshal(jsonFileData, &parsedData)
	if err != nil {
		log.Fatal(err)
	}

	level := &Level{
		TilesetImages: make(map[int]*ebiten.Image),
	}
	level.loadMap(parsedData)

	return level
}

func (l *Level) loadMap(levelData LevelData) {
	l.Width = levelData.Width
	l.Height = levelData.Height
	l.TileWidth = float64(levelData.TileWidth)
	l.TileHeight = float64(levelData.TileHeight)
	l.Layers = levelData.Layers
	l.TileSets = levelData.TileSets

	// Load tile images
	for _, tileset := range levelData.TileSets {
		img, _, err := ebitenutil.NewImageFromFile(tileset.Image)
		if err != nil {
			log.Fatal(err)
		}

		// Store the image with its firstgid as the key
		l.TilesetImages[tileset.FirstGID] = img
	}
	for _, layer := range l.Layers {
		if layer.Type != "tilelayer" {
			continue
		}

		rows := layer.Height
		columns := layer.Width

		for row := 0; row < rows; row++ {
			for column := 0; column < columns; column++ {
				tileIndex := row*layer.Width + column

				if tileIndex >= len(layer.Data) {
					continue
				}

				tileGID := layer.Data[tileIndex]

				if tileGID == 0 {
					continue
				}

				var currentTileset TileSet
				var tilesetGID int
				for _, tileset := range l.TileSets {
					if tileGID >= tileset.FirstGID {
						currentTileset = tileset
						tilesetGID = tileset.FirstGID
					}
				}

				if currentTileset.FirstGID == 0 {
					continue
				}

				localID := tileGID - tilesetGID
				tilesetX := (localID % currentTileset.Columns) * currentTileset.TileWidth
				tilesetY := (localID / currentTileset.Columns) * currentTileset.TileHeight

				tileImg := l.TilesetImages[tilesetGID].SubImage(image.Rect(
					tilesetX,
					tilesetY,
					tilesetX+currentTileset.TileWidth,
					tilesetY+currentTileset.TileHeight,
				)).(*ebiten.Image)

				l.tiles = append(l.tiles, Tile{
					X:     float64(column) * l.TileWidth,
					Y:     float64(row) * l.TileHeight,
					image: tileImg,
				})
			}
		}
	}
}

func (l *Level) Draw(screen *ebiten.Image) {

	for _, tile := range l.tiles {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(
			tile.X,
			tile.Y,
		)
		screen.DrawImage(tile.image, op)
	}
}
