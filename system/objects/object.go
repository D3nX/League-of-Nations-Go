package objects

import "github.com/gen2brain/raylib-go/raylib"

type Object interface {
	Update()
	Draw(*raylib.Camera2D)
	IsSelected() bool
	GetPosition() raylib.Vector2
}
