package objects

import (
	"github.com/gen2brain/raylib-go/raylib"
)

var BuildingTextures []raylib.Texture2D

type Building struct {
	Type int
	X    float32
	Y    float32
}

func (b *Building) Update() {

}

func (b *Building) Draw() {
	raylib.DrawTexture(BuildingTextures[b.Type], int32(b.X), int32(b.Y), raylib.White)
}
