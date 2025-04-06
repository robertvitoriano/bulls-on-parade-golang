package components

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Animation struct {
	Duration int `json:"duration"`
	TileId   int `json:"tileid"`
}

type Tile struct {
	Animation []Animation `json:"animation"`
	Id        int         `json:"id"`
}
type TileSet struct {
	Columns     int    `json:"columns"`
	FirstGID    int    `json:"firstgid"`
	Image       string `json:"image"`
	ImageHeight int    `json:"imageheight"`
	ImageWidth  int    `json:"imagewidth"`
	Margin      int    `json:"margin"`
	Name        string `json:"name"`
	Spacing     int    `json:"spacing"`
	TileCount   int    `json:"tilecount"`
	TileHeight  int    `json:"tileheight"`
	TileWidth   int    `json:"tilewidth"`
	Tiles       []Tile `json:"tiles"`
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
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
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
	for _, layer := range parsedData.Layers {
		fmt.Println(layer)
	}
	fmt.Println(parsedData)
	return &Level{}
}
