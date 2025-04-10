package level

import (
	"encoding/json"
	"fmt"
	"image"
	"io"
	"log"
	"os"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/robertvitoriano/bulls-on-parade-golang/components"
	"github.com/robertvitoriano/bulls-on-parade-golang/entities"
	"github.com/robertvitoriano/bulls-on-parade-golang/physics"
	"github.com/robertvitoriano/bulls-on-parade-golang/utils"
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
	ImageHeight float64   `json:"imageheight"`
	ImageWidth  float64   `json:"imagewidth"`
	Margin      int       `json:"margin"`
	Name        string    `json:"name"`
	Spacing     int       `json:"spacing"`
	TileCount   int       `json:"tilecount"`
	TileHeight  float64   `json:"tileheight"`
	TileWidth   float64   `json:"tilewidth"`
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
	Height    float64  `json:"height"`
	Width     float64  `json:"width"`
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
	TileHeight       float64   `json:"tileheight"`
	TileWidth        float64   `json:"tilewidth"`
	Layers           []Layer   `json:"layers"`
	TileSets         []TileSet `json:"tilesets"`
}
type Level struct {
	Width               float64
	Height              float64
	TileWidth           float64
	TileHeight          float64
	X                   float64
	Y                   float64
	Layers              []Layer
	TileSets            []TileSet
	TilesetImages       map[int]*ebiten.Image
	tiles               []*Tile
	collisions          []*physics.Collision
	PlayerSpawnPosition utils.Position
	Player              *entities.Player
}

func NewLevel(filePath string, player *entities.Player) *Level {
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
		Player:        player,
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

	for _, tileset := range levelData.TileSets {

		tileSetImagePrefix := "content/tilesets/"
		fileNameChunks := strings.Split(tileset.Image, "/")
		fileName := fileNameChunks[len(fileNameChunks)-1]
		filePath := fmt.Sprintf("%s/%s", tileSetImagePrefix, fileName)

		img, _, err := ebitenutil.NewImageFromFile(filePath)

		if err != nil {
			log.Fatal(err)
		}

		l.TilesetImages[tileset.FirstGID] = img
	}
	for _, layer := range l.Layers {
		if layer.Type == "tilelayer" {
			l.addTilesToLevel(layer)
		} else if layer.Type == "objectgroup" {
			l.addObjectsToLevel(layer)
		}
	}
}
func (l *Level) addObjectsToLevel(layer Layer) {
	for _, object := range layer.Objects {
		if layer.Name == "collisions" {
			l.collisions = append(l.collisions, &physics.Collision{
				GameObject: components.GameObject{
					Position: utils.Position{
						X: object.X,
						Y: object.Y,
					},
					Size: utils.Size{
						Width:  object.Width,
						Height: object.Height,
					},
				},
			})
		} else if layer.Name == "spawn points" {
			l.PlayerSpawnPosition = utils.Position{
				X: object.X,
				Y: object.Y,
			}
		}
	}
}
func (l *Level) addTilesToLevel(layer Layer) {
	rows := int(layer.Height)
	columns := int(layer.Width)

	for row := 0; row < rows; row++ {
		for column := 0; column < columns; column++ {
			tileIndex := row*columns + column

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
			tilesetX := (localID % currentTileset.Columns) * int(currentTileset.TileWidth)
			tilesetY := (localID / currentTileset.Columns) * int(currentTileset.TileHeight)

			tileImg := l.TilesetImages[tilesetGID].SubImage(image.Rect(
				tilesetX,
				tilesetY,
				tilesetX+int(currentTileset.TileWidth),
				tilesetY+int(currentTileset.TileHeight),
			)).(*ebiten.Image)

			l.tiles = append(l.tiles, &Tile{

				GameObject: components.GameObject{
					Position: utils.Position{
						X: float64(column) * l.TileWidth,
						Y: float64(row) * l.TileHeight,
					},
					Size: utils.Size{
						Width:  currentTileset.TileWidth,
						Height: currentTileset.TileHeight,
					},
				},
				image: tileImg,
			})
		}
	}
}

func (l *Level) Update() {
	offset := &utils.Position{
		X: 0,
	}
	for _, tile := range l.tiles {
		tile.GameObject.SetOffset(utils.Position{
			X: offset.X,
			Y: 0.00,
		})
	}
	for _, collision := range l.collisions {
		collision.GameObject.SetOffset(utils.Position{
			X: offset.X,
			Y: 0.00,
		})
	}

}

func (l *Level) Draw(screen *ebiten.Image) {
	fmt.Println("TILE X", l.tiles[0].GameObject.Position.X)
	for _, tile := range l.tiles {
		tile.draw(screen)
	}
	// for _, collision := range l.collisions {
	// 	collision.DebugDraw(screen)
	// }
}
func (l *Level) GetLevelCollisions(other components.GameObject) []*physics.Collision {

	collisions := []*physics.Collision{}

	for _, collision := range l.collisions {
		if collision.GameObject.CollidesWith(other) {
			collisions = append(collisions, collision)
		}
	}
	return collisions
}
