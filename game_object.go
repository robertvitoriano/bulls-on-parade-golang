package main

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/robertvitoriano/bulls-on-parade-golang/utils"
)

type FrameProperties struct {
	width  int
	height int
	count  int
}

type GameObject struct {
	animations map[string][]*ebiten.Image
}

func (g *GameObject) AddAnimation(animationName string, spriteSheetPath string, orientation string, frameProperties FrameProperties) {
	charImageAnimationSprite, err := utils.ReadImage(spriteSheetPath)

	if err != nil {
		log.Fatal(err)
	}

	if g.animations == nil {
		g.animations = make(map[string][]*ebiten.Image)
	}

	images := make([]*ebiten.Image, 0, frameProperties.count)

	if orientation == "horizontal" {
		for i := 0; i < frameProperties.count; i++ {
			xOffset := frameProperties.width * i
			frame := charImageAnimationSprite.SubImage(image.Rect(xOffset, 0, xOffset+frameProperties.width, frameProperties.height)).(*ebiten.Image)
			images = append(images, frame)
		}
	}

	if orientation == "vertical" {
		for i := 0; i < frameProperties.count; i++ {
			yOffset := frameProperties.height * i
			frame := charImageAnimationSprite.SubImage(image.Rect(0, yOffset, frameProperties.width, frameProperties.height+yOffset)).(*ebiten.Image)
			images = append(images, frame)
		}
	}
	g.animations[animationName] = images
}

func (g *GameObject) PlayAnimation(animationName string) {

}
