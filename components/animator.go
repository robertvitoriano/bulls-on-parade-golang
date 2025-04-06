package components

import (
	"image"
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

func (a *Animator) AddAnimation(animationName string, spriteSheetPath string, orientation string, frameProperties FrameProperties) {
	charImageAnimationSprite, err := utils.ReadImage(spriteSheetPath)

	if err != nil {
		log.Fatal(err)
	}

	if a.animations == nil {
		a.animations = make(map[string][]*ebiten.Image)
		a.currentAnimation = animationName
	}

	images := make([]*ebiten.Image, 0, frameProperties.Count)

	if orientation == "horizontal" {
		for i := 0; i < frameProperties.Count; i++ {
			xOffset := frameProperties.Width * i
			frame := charImageAnimationSprite.SubImage(image.Rect(xOffset, 0, xOffset+frameProperties.Width, frameProperties.Height)).(*ebiten.Image)
			images = append(images, frame)
		}
	}

	if orientation == "vertical" {
		for i := 0; i < frameProperties.Count; i++ {
			yOffset := frameProperties.Height * i
			frame := charImageAnimationSprite.SubImage(image.Rect(0, yOffset, frameProperties.Width, frameProperties.Height+yOffset)).(*ebiten.Image)
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
