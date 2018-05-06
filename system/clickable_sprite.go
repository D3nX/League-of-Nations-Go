package system

import (
	"github.com/gen2brain/raylib-go/raylib"
)

type ClickableSprite struct {
	Position  raylib.Vector2
	Scale     float32
	Image     *raylib.Image
	ImageData []raylib.Color
}

func (ct *ClickableSprite) Clicked() bool {

	if MouseOnEx(ct.Position, ct.Scale, ct.Image, ct.ImageData) {
		return true
	}

	return false
}

func (ct *ClickableSprite) Draw() {

}

func NewClickeableSprite(path string) ClickableSprite {
	return ClickableSprite{}
}
