package utils

import (
	"bytes"
	"image"
	"io"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

func ReadImageFile(path string) (*ebiten.Image, error) {
	charAnimationFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer charAnimationFile.Close()

	charAnimationFileData, err := io.ReadAll(charAnimationFile)
	if err != nil {
		log.Fatal(err)
	}

	charImageDecoded, _, err := image.Decode(bytes.NewReader(charAnimationFileData))
	if err != nil {
		log.Fatal(err)
	}

	return ebiten.NewImageFromImage(charImageDecoded), nil
}
