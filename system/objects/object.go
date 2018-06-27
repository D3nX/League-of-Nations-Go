package objects

import "github.com/gen2brain/raylib-go/raylib"

type Object interface {
	Update(*raylib.Camera2D)
	Draw(*raylib.Camera2D)
	IsSelected() bool
	SetSelected(bool)
	GetPosition() raylib.Vector2
}
