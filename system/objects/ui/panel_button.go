package ui

import (
	"github.com/gen2brain/raylib-go/raygui"
	"github.com/gen2brain/raylib-go/raylib"
)

type PanelButton struct {
	Rectangle raylib.Rectangle
	Text      string
	Texture   *raylib.Texture2D
}

func (pb PanelButton) Draw() bool {
	pressed := raygui.Button(pb.Rectangle, "")

	scale := float32(1.0)
	texPos := raylib.NewVector2(pb.Rectangle.X, pb.Rectangle.Y)

	if pb.Texture != nil {
		scale = float32(pb.Rectangle.Height/float32(pb.Texture.Height)) - 0.2

		texPos = raylib.NewVector2(pb.Rectangle.X+float32((pb.Rectangle.Width-float32(pb.Texture.Width)*scale)/2),
			pb.Rectangle.Y+float32((pb.Rectangle.Height-float32(pb.Texture.Height)*scale))/2)

		// DrawTextureEx(Texture2D texture, Vector2 position, float rotation, float scale, Color tint)
		raylib.DrawTextureEx(*pb.Texture,
			texPos,
			0.0,
			scale,
			raylib.White)
	}

	raylib.DrawText(pb.Text, int32(texPos.X), int32(texPos.Y)+int32(float32(pb.Texture.Height)*scale), 25, raylib.Black)

	return pressed
}

func NewPanelButton(x, y, width, height float32, text string, texture *raylib.Texture2D) PanelButton {
	pb := PanelButton{
		Rectangle: raylib.NewRectangle(x, y, width, height),
		Text:      text,
		Texture:   texture,
	}

	return pb
}
