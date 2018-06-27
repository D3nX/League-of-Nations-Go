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

func (b *Building) Update(cam *raylib.Camera2D) {

}

func (b *Building) Draw(cam *raylib.Camera2D) {
	// Begin rendering camera
	raylib.BeginMode2D(*cam)
	// Render the texture
	raylib.DrawTexture(BuildingTextures[b.Type], int32(b.X), int32(b.Y), raylib.White)
	// End rendering camera
	raylib.EndMode2D()
}

func (b *Building) IsSelected() bool {
	return false
}

func (t *Building) SetSelected(selected bool) {

}

func (t *Building) GetPosition() raylib.Vector2 {
	return raylib.NewVector2(t.X, t.Y)
}
