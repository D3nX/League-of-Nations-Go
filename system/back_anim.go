package system

import (
	"fmt"
	"time"

	"github.com/gen2brain/raylib-go/raylib"
)

type BackAnim struct {
	backgrounds       []raylib.Texture2D
	alpha             int16
	currentBackground uint8
	action            int
}

func (ba *BackAnim) Load() {
	ba.backgrounds = make([]raylib.Texture2D, 14)

	ba.currentBackground = 0

	ba.action = 1

	ba.alpha = 255

	for i := 0; i < len(ba.backgrounds); i++ {
		image := raylib.LoadImage(fmt.Sprint("res/backgrounds/bg_", i+1, ".png"))

		raylib.ImageResize(image, raylib.GetScreenWidth(), raylib.GetScreenHeight())

		ba.backgrounds[i] = raylib.LoadTextureFromImage(image)

		raylib.UnloadImage(image)
	}

	// Launch a routine that will count time
	go func(input *int) {
		for true {
			time.Sleep(time.Second) // Wait 1 sec
			*input = 0
			time.Sleep(time.Second * 4) // Wait 2 sec
			*input = 1
			time.Sleep(time.Second) // Wait 1 sec again
			*input = 2
		}
	}(&ba.action)
}

func (ba *BackAnim) Update() {

	// fmt.Println("action -> ", ba.action)

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
