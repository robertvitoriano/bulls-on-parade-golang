package components

import (
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/robertvitoriano/bulls-on-parade-golang/utils"
)

type FrameProperties struct {
	Width  int
	Height int
	Count  int
}
type Animator struct {
	animations        map[string][]*ebiten.Image
	currentAnimation  string
	currentFrameIndex int
}

func (a *Animator) AddAnimation(animationName string, spriteSheetPath string, initialRow int, initialColumn int, orientation string, frameProperties FrameProperties) {
	charImageAnimationSprite, err := utils.ReadImageFile(spriteSheetPath)

	if err != nil {
		log.Fatal(err)
	}

	if a.animations == nil {
		a.animations = make(map[string][]*ebiten.Image)
		a.currentAnimation = animationName
	}

	images := make([]*ebiten.Image, 0, frameProperties.Count)

	if orientation == "horizontal" {
		startX := initialColumn * frameProperties.Width
		startY := initialRow * frameProperties.Height

		for i := 0; i < frameProperties.Count; i++ {
			xOffset := startX + (frameProperties.Width * i)
			yOffset := startY
			frame := charImageAnimationSprite.SubImage(image.Rect(
				xOffset,
				yOffset,
				xOffset+frameProperties.Width,
				yOffset+frameProperties.Height)).(*ebiten.Image)
			images = append(images, frame)
		}
	}

	if orientation == "vertical" {
		startX := initialColumn * frameProperties.Width
		startY := initialRow * frameProperties.Height

		for i := 0; i < frameProperties.Count; i++ {
			xOffset := startX
			yOffset := startY + (frameProperties.Height * i)
			frame := charImageAnimationSprite.SubImage(image.Rect(
				xOffset,
				yOffset,
				xOffset+frameProperties.Width,
				yOffset+frameProperties.Height)).(*ebiten.Image)
			images = append(images, frame)
		}
	}

	a.animations[animationName] = images
}
func (a *Animator) ChangeAnimation(animationName string) {
	a.currentAnimation = animationName
}

func (a *Animator) Update() {
	a.currentFrameIndex = (a.currentFrameIndex + 1) % len(a.animations[a.currentAnimation])
}

func (a *Animator) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(40, 30)

	if len(a.animations) > 0 {
		screen.DrawImage(a.animations[a.currentAnimation][a.currentFrameIndex], op)
	}
}
