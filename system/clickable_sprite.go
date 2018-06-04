package system

import (
	"github.com/gen2brain/raylib-go/raylib"
)

type ClickableSprite struct {
	Position raylib.Vector2
	Scale    float32
	Color    raylib.Color
	Texture  raylib.Texture2D
	Pixels   [][]raylib.Color
}

func NewClickableSprite(path string) ClickableSprite {
	tex := raylib.LoadTexture(path)
	return NewClickableSpriteFromTexture(&tex)
}

func NewClickableSpriteFromTexture(texture *raylib.Texture2D) ClickableSprite {
	cp := ClickableSprite{}

	cp.Texture = *texture
	cp.Position = raylib.NewVector2(0, 0)
	cp.Scale = 1.0
	cp.Color = raylib.White

	// Read pixels

	imageData := raylib.GetImageData(raylib.GetTextureData(cp.Texture))

	cp.Pixels = make([][]raylib.Color, cp.Texture.Width)

	for tmp := range cp.Pixels {
		cp.Pixels[tmp] = make([]raylib.Color, cp.Texture.Height)
	}

	i := 0

	for y := int32(0); y < cp.Texture.Height; y++ {
		for x := int32(0); x < cp.Texture.Width; x++ {
			cp.Pixels[x][y] = imageData[i]
			i++
		}
	}

	return cp
}

func (cp *ClickableSprite) Hover() bool {

	px := raylib.GetMouseX() - int32(cp.Position.X)
	py := raylib.GetMouseY() - int32(cp.Position.Y)

	if px >= 0 && px < cp.Texture.Width {
		if py >= 0 && py < cp.Texture.Height {
			if cp.Pixels[px][py].A != 0 {
				return true
			}
		}
	}

	return false
}

func (cp *ClickableSprite) Draw() {
	raylib.DrawTexture(cp.Texture, int32(cp.Position.X), int32(cp.Position.Y), cp.Color)
}
