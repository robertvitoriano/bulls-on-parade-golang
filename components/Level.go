package components

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Animation struct {
	Duration int `json:"duration,omitempty"`
	TileId   int `json:"tileid,omitempty"`
}

type Tile struct {
	Animation []Animation `json:"animation,omitempty"`
	Id        int         `json:"id,omitempty"`
}
type TileSet struct {
	Columns     int    `json:"columns,omitempty"`
	FirstGID    int    `json:"firstgid,omitempty"`
	Image       string `json:"image,omitempty"`
	ImageHeight int    `json:"imageheight,omitempty"`
	ImageWidth  int    `json:"imagewidth,omitempty"`
	Margin      int    `json:"margin,omitempty"`
	Name        string `json:"name,omitempty"`
	Spacing     int    `json:"spacing,omitempty"`
	TileCount   int    `json:"tilecount,omitempty"`
	TileHeight  int    `json:"tileheight,omitempty"`
	TileWidth   int    `json:"tilewidth,omitempty"`
	Tiles       []Tile `json:"tiles,omitempty"`
}
type PolylinePoint struct {
	X float64 `json:"x,omitempty"`
	Y float64 `json:"y,omitempty"`
}

type ObjectProperty struct {
	Name  string `json:"name,omitempty"`
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}

type Object struct {
	Id         int              `json:"id,omitempty"`
	Height     float64          `json:"height,omitempty"`
	Width      float64          `json:"width,omitempty"`
	X          float64          `json:"x,omitempty"`
	Y          float64          `json:"y,omitempty"`
	Visible    bool             `json:"visible,omitempty"`
	Name       string           `json:"name,omitempty"`
	Rotation   float64          `json:"rotation,omitempty"`
	Properties []ObjectProperty `json:"properties,omitempty"`
	Polyline   []PolylinePoint  `json:"polyline,omitempty"`
	Ellipse    bool             `json:"ellipse,omitempty"`
	Type       string           `json:"type,omitempty"`
}

type Layer struct {
	Data      []int    `json:"data,omitempty"`
	Height    int      `json:"height,omitempty"`
	Width     int      `json:"width,omitempty"`
	Id        int      `json:"id,omitempty"`
	Name      string   `json:"name,omitempty"`
	Opacity   int      `json:"opacity,omitempty"`
	Type      string   `json:"type,omitempty"`
	Visible   bool     `json:"visible,omitempty"`
	X         float64  `json:"x,omitempty"`
	Y         float64  `json:"y,omitempty"`
	DrawOrder string   `json:"draworder,omitempty"`
	Objects   []Object `json:"objects,omitempty"`
	Color     string   `json:"color,omitempty"`
}

type LevelData struct {
	Width            float64   `json:"width,omitempty"`
	Height           float64   `json:"height,omitempty"`
	Version          string    `json:"version,omitempty"`
	Type             string    `json:"type,omitempty"`
	CompressionLevel int       `json:"compressionlevel,omitempty"`
	Infinite         bool      `json:"infinite,omitempty"`
	NextLayerId      int       `json:"nextlayerid,omitempty"`
	NextObjectId     int       `json:"nextobjectid,omitempty"`
	Orientation      string    `json:"orientation,omitempty"`
	RenderOrder      string    `json:"renderorder,omitempty"`
	TiledVersion     string    `json:"tiledversion,omitempty"`
	TileHeight       int       `json:"tileheight,omitempty"`
	TileWidth        int       `json:"tilewidth,omitempty"`
	Layers           []Layer   `json:"layers,omitempty"`
	TileSets         []TileSet `json:"tilesets,omitempty"`
}
type Level struct {
	Width            float64   `json:"width,omitempty"`
	Height           float64   `json:"height,omitempty"`
	Version          string    `json:"version,omitempty"`
	Type             string    `json:"type,omitempty"`
	CompressionLevel int       `json:"compressionlevel,omitempty"`
	Infinite         bool      `json:"infinite,omitempty"`
	NextLayerId      int       `json:"nextlayerid,omitempty"`
	NextObjectId     int       `json:"nextobjectid,omitempty"`
	Orientation      string    `json:"orientation,omitempty"`
	RenderOrder      string    `json:"renderorder,omitempty"`
	TiledVersion     string    `json:"tiledversion,omitempty"`
	TileHeight       int       `json:"tileheight,omitempty"`
	TileWidth        int       `json:"tilewidth,omitempty"`
	Layers           []Layer   `json:"layers,omitempty"`
	TileSets         []TileSet `json:"tilesets,omitempty"`
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
