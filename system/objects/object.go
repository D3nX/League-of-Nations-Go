package objects

import "github.com/gen2brain/raylib-go/raylib"

type Object interface {
	Update(*raylib.Camera2D, bool)
	Draw(*raylib.Camera2D)
	IsSelected() bool
	SetSelected(bool)
	GetPosition() raylib.Vector2
	Collides(raylib.Rectangle) bool
	StopMoving(string)
	CanMove() bool
}
