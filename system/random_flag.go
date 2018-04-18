package system

import (
	"math/rand"

	"github.com/gen2brain/raylib-go/raylib"
)

func GetNewFlag(vertical bool, linesCount int, ideology string) *raylib.Image {

	width := 720
	height := 480

	flagImage := raylib.GenImageGradientH(width, height, raylib.White, raylib.White)

	// Logo for the following ideology
	posX := int32(0)
	posY := int32(0)
	var logo *raylib.Image = nil

	switch ideology {

	case "communism":
		if rand.Intn(100) < 50 {
			logo = raylib.ImageCopy(Logos[3+rand.Intn(3)])
			raylib.ImageResize(logo, 128, 128)
		} else {
			logo = nil
		}

	case "nationalism":
		if rand.Intn(100) < 50 {
			logo = raylib.ImageCopy(Logos[rand.Intn(3)])
			raylib.ImageResize(logo, 128, 128)
		} else {
			logo = nil
		}

	case "democracy":
		if rand.Intn(100) < 50 {
			logo = raylib.ImageCopy(Logos[6 /*+ rand.Intn(3)*/])
			raylib.ImageResize(logo, 128, 128)
		} else {
			logo = nil
		}
	}

	// Now, generate line
	if vertical {
		x := float32(0)
		addX := float32(width / linesCount)
		for i := 0; i < linesCount; i++ {
			var color raylib.Color
			switch int32(rand.Intn(6)) {
			case 0:
				color = raylib.Blue
				break
			case 1:
				color = raylib.Red
				break
			case 2:
				color = raylib.Yellow
				break
			case 3:
				color = raylib.White
				break
			case 4:
				color = raylib.Brown
				break
			case 5:
				color = raylib.Green
				break
			}

			image := raylib.GenImageGradientH(width/linesCount, height, color, color)
			raylib.ImageDraw(flagImage,
				image,
				raylib.NewRectangle(int32(x), 0, image.Width, image.Height),
				raylib.NewRectangle(int32(x), 0, image.Width, image.Height))
			x += addX
		}
	} else {
		y := float32(0)
		addY := float32(height / linesCount)
		for i := 0; i < linesCount; i++ {
			var color raylib.Color
			switch int32(rand.Intn(6)) {
			case 0:
				color = raylib.Blue
				break
			case 1:
				color = raylib.Red
				break
			case 2:
				color = raylib.Yellow
				break
			case 3:
				color = raylib.White
				break
			case 4:
				color = raylib.Brown
				break
			case 5:
				color = raylib.Green
				break
			}
			image := raylib.GenImageGradientH(width, height/linesCount, color, color)
			raylib.ImageDraw(flagImage,
				image,
				raylib.NewRectangle(0, int32(y), image.Width, image.Height),
				raylib.NewRectangle(0, int32(y), image.Width, image.Height))
			y += addY
		}
	}

	// Add the logo, if needed
	if logo != nil {
		raylib.ImageDraw(flagImage,
			logo,
			raylib.NewRectangle(posX, posY, logo.Width, logo.Height),
			raylib.NewRectangle(posX, posY, logo.Width, logo.Height))
	}

	return flagImage
}
