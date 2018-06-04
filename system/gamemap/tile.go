package gamemap

import (
	"github.com/gen2brain/raylib-go/raylib"
)

type Tile struct {
	X     float32
	Y     float32
	Model raylib.Model
	Type  rune
}

func (tile *Tile) Draw() {
	switch tile.Type {
	case '/':
		raylib.DrawTextureRec(Tileset,
			raylib.NewRectangle(0, 0, 32, 32),
			raylib.NewVector2(tile.X, tile.Y),
			raylib.White)

	case '\\':
		raylib.DrawTextureRec(Tileset,
			raylib.NewRectangle(64, 0, 32, 32),
			raylib.NewVector2(tile.X, tile.Y),
			raylib.White)

	case '(':
		raylib.DrawTextureRec(Tileset,
			raylib.NewRectangle(0, 64, 32, 32),
			raylib.NewVector2(tile.X, tile.Y),
			raylib.White)

	case ')':
		raylib.DrawTextureRec(Tileset,
			raylib.NewRectangle(64, 64, 32, 32),
			raylib.NewVector2(tile.X, tile.Y),
			raylib.White)

	case '{':
		raylib.DrawTextureRec(Tileset,
			raylib.NewRectangle(0, 32, 32, 32),
			raylib.NewVector2(tile.X, tile.Y),
			raylib.White)

	case '}':
		raylib.DrawTextureRec(Tileset,
			raylib.NewRectangle(64, 32, 32, 32),
			raylib.NewVector2(tile.X, tile.Y),
			raylib.White)

	case '*':
		raylib.DrawTextureRec(Tileset,
			raylib.NewRectangle(32, 32, 32, 32),
			raylib.NewVector2(tile.X, tile.Y),
			raylib.White)

	case '-':
		raylib.DrawTextureRec(Tileset,
			raylib.NewRectangle(32, 0, 32, 32),
			raylib.NewVector2(tile.X, tile.Y),
			raylib.White)

	case '_':
		raylib.DrawTextureRec(Tileset,
			raylib.NewRectangle(32, 64, 32, 32),
			raylib.NewVector2(tile.X, tile.Y),
			raylib.White)
	}

	raylib.DrawModel(tile.Model, raylib.NewVector3(tile.X, tile.Y, 0), 2.0, raylib.White)
	// raylib.DrawPlane(raylib.NewVector3(tile.X-16, tile.Y-16, 0), raylib.NewVector2(32, 32), raylib.White)
}

func (tile *Tile) ClickedOn() bool {

	if raylib.IsMouseButtonPressed(raylib.MouseLeftButton) {
		if raylib.GetMouseX() >= int32(tile.X) && raylib.GetMouseX() <= int32(tile.X+32) {
			if raylib.GetMouseY() >= int32(tile.Y) && raylib.GetMouseY() <= int32(tile.Y+32) {
				return true
			}
		}
	}

	return false
}
