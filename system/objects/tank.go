package objects

import (
	"github.com/gen2brain/raylib-go/raylib"
)

var TankTextures []raylib.Texture2D

type Tank struct {
	Type     int
	X        float32
	Y        float32
	Angle    float32
	Selected bool
}

func (t *Tank) Update() {
	switch t.Angle {
	case 0.0:
		fallthrough
	case 90.0:
		fallthrough
	case 180.0:
		fallthrough
	case 270.0:
		if raylib.IsMouseButtonPressed(raylib.MouseLeftButton) {
			if t.Selected {
				t.Selected = false
			} else {
				t.Selected = true
			}
		}
	}
}

func (t *Tank) Draw() {
	/*raylib.DrawTextureEx(TankTextures[t.Type],
	raylib.NewVector2(t.X, t.Y),
	t.Angle,
	1.0,
	raylib.White)*/

	if t.Selected {
		raylib.DrawTexturePro(TankTextures[t.Type],
			raylib.NewRectangle(0, 0, float32(TankTextures[t.Type].Width), float32(TankTextures[t.Type].Height)),
			raylib.NewRectangle(t.X, t.Y, float32(TankTextures[t.Type].Width), float32(TankTextures[t.Type].Height)),
			raylib.NewVector2(float32(TankTextures[t.Type].Width)/2.0, float32(TankTextures[t.Type].Height)/2.0),
			t.Angle,
			raylib.White)
	} else {
		raylib.DrawTexturePro(TankTextures[t.Type],
			raylib.NewRectangle(0, 0, float32(TankTextures[t.Type].Width), float32(TankTextures[t.Type].Height)),
			raylib.NewRectangle(t.X, t.Y, float32(TankTextures[t.Type].Width), float32(TankTextures[t.Type].Height)),
			raylib.NewVector2(float32(TankTextures[t.Type].Width)/2.0, float32(TankTextures[t.Type].Height)/2.0),
			t.Angle,
			raylib.NewColor(255, 24, 24, 255))
	}
}
