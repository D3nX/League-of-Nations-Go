package system

import (
	"fmt"

	"github.com/gen2brain/raylib-go/raylib"
)

type BackAnim struct {
	backgrounds       []raylib.Texture2D
	alpha             int16
	currentBackground uint8
	action            int
	time              int
}

func (ba *BackAnim) Load() {
	ba.backgrounds = make([]raylib.Texture2D, 14)

	ba.currentBackground = 0

	ba.action = 1

	ba.alpha = 255

	ba.time = 0

	for i := 0; i < len(ba.backgrounds); i++ {
		image := raylib.LoadImage(fmt.Sprint("res/backgrounds/bg_", i+1, ".png"))

		raylib.ImageResize(image, raylib.GetScreenWidth(), raylib.GetScreenHeight())

		ba.backgrounds[i] = raylib.LoadTextureFromImage(image)

		raylib.UnloadImage(image)
	}
}

func (ba *BackAnim) Update() {

	// fmt.Println("action -> ", ba.action)

	ba.time += 1

	if ba.action != 0 {
		if ba.time%61 == 60 {
			ba.action++
		}
	} else {
		if ba.time%(60*6+1) == 60*6 {
			ba.action++
		}
	}

	switch ba.action {
	case 0:
		ba.alpha -= 5
		if ba.alpha < 0 {
			ba.alpha = 0
		}
		break
	case 1:
		ba.alpha = 0
		break
	case 2:
		ba.alpha += 5
		if ba.alpha > 255 {
			ba.alpha = 255
			ba.action = 0
		}

		if ba.alpha == 255 {
			ba.currentBackground += 1

			if ba.currentBackground >= uint8(len(ba.backgrounds)) {
				ba.currentBackground = 0
			}
		}
		break
	}
}

func (ba *BackAnim) Draw() {
	raylib.DrawTexture(ba.backgrounds[ba.currentBackground], 0, 0, raylib.White)
	raylib.DrawRectangle(0, 0, raylib.GetScreenWidth(), raylib.GetScreenHeight(),
		raylib.Color{R: 0, G: 0, B: 0, A: uint8(ba.alpha)})
}

func (ba *BackAnim) Close() {
	for _, texture := range ba.backgrounds {
		raylib.UnloadTexture(texture)
	}
}
