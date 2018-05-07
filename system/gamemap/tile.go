package gamemap

import (
	"github.com/gen2brain/raylib-go/raylib"
)

type Tile struct {
	X    float32
	Y    float32
	Type rune
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
}
