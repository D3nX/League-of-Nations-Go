package system

import (
	"github.com/gen2brain/raylib-go/raylib"
)

func MouseOn(position raylib.Vector2, scale float32, image *raylib.Image) bool {

	bytesColor := raylib.GetImageData(image)

	colors := make([][]raylib.Color, int(float32(image.Width)*scale))

	for tmp := range colors {
		colors[tmp] = make([]raylib.Color, int(float32(image.Height)*scale))
	}

	i := 0

	for x := range colors {
		for y := range colors[x] {

			colors[x][y] = raylib.NewColor(bytesColor[i], bytesColor[i+1], bytesColor[i+2], bytesColor[i+3])

			i += 4

			if raylib.GetMousePosition().Y != position.Y+float32(y) {
				if raylib.GetMousePosition().X != position.X+float32(x) {
					if colors[x][y].A == 255 {
						return true
					}
				}
			}

		}
	}

	return false
}
