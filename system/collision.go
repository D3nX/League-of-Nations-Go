package system

import (
	"github.com/gen2brain/raylib-go/raylib"
)

func MouseOn(position raylib.Vector2, scale float32, image *raylib.Image) bool {

	imageColors := raylib.GetImageData(image)

	i := 0

	for y := int32(0); y < image.Height; y++ {
		for x := int32(0); x < image.Width; x++ {
			if raylib.GetMousePosition().X == position.X+float32(x) {
				if raylib.GetMousePosition().Y == position.Y+float32(y) {
					if imageColors[i].A != 0 {
						return true
					}
				}
			}

			i++
		}
	}

	return false
}

func MouseOnEx(position raylib.Vector2, scale float32, image *raylib.Image, imageColors []raylib.Color) bool {

	i := 0

	for y := int32(0); y < image.Height; y++ {
		for x := int32(0); x < image.Width; x++ {
			if raylib.GetMousePosition().X == position.X+float32(x) {
				if raylib.GetMousePosition().Y == position.Y+float32(y) {
					if imageColors[i].A != 0 {
						return true
					}
				}
			}

			i++
		}
	}

	return false
}
