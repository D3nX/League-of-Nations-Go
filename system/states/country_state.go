package states

import (
	"../../system"
	"github.com/gen2brain/raylib-go/raylib"
)

type CountryState struct {
	WaterAnimation raylib.Texture2D
	CurrentFrame   float32

	Countries      []system.ClickableSprite
	Camera         raylib.Camera2D
	ZoomOn         bool
	TakeScreenshot int
}

func (state *CountryState) Load() {
	// Load the water animation
	state.WaterAnimation = raylib.LoadTexture("res/sprites/water_animation.png")
	state.CurrentFrame = 0

	state.Camera = raylib.NewCamera2D(raylib.NewVector2(0, 0), raylib.NewVector2(0, 0), 0.0, 1.0)
	state.ZoomOn = false

	state.TakeScreenshot = 0

	state.Countries = append(state.Countries, system.NewClickableSprite("res/test/map/russia.png"))
	state.Countries = append(state.Countries, system.NewClickableSprite("res/test/map/ukraine.png"))
	state.Countries = append(state.Countries, system.NewClickableSprite("res/test/map/baltic_states.png"))
	state.Countries = append(state.Countries, system.NewClickableSprite("res/test/map/poland.png"))
	state.Countries = append(state.Countries, system.NewClickableSprite("res/test/map/kazakhstan.png"))
	state.Countries = append(state.Countries, system.NewClickableSprite("res/test/map/belarus.png"))
	state.Countries = append(state.Countries, system.NewClickableSprite("res/test/map/czech_republic.png"))
	state.Countries = append(state.Countries, system.NewClickableSprite("res/test/map/balkan.png"))
	state.Countries = append(state.Countries, system.NewClickableSprite("res/test/map/slovakia.png"))
	state.Countries = append(state.Countries, system.NewClickableSprite("res/test/map/hungary.png"))
	state.Countries = append(state.Countries, system.NewClickableSprite("res/test/map/united_kingdom.png"))
	state.Countries = append(state.Countries, system.NewClickableSprite("res/test/map/germany.png"))
	state.Countries = append(state.Countries, system.NewClickableSprite("res/test/map/austria.png"))
	state.Countries = append(state.Countries, system.NewClickableSprite("res/test/map/france.png"))
	state.Countries = append(state.Countries, system.NewClickableSprite("res/test/map/italy.png"))
	state.Countries = append(state.Countries, system.NewClickableSprite("res/test/map/slovenia.png"))
	state.Countries = append(state.Countries, system.NewClickableSprite("res/test/map/finland.png"))
	state.Countries = append(state.Countries, system.NewClickableSprite("res/test/map/norway.png"))
	state.Countries = append(state.Countries, system.NewClickableSprite("res/test/map/sweden.png"))
	state.Countries = append(state.Countries, system.NewClickableSprite("res/test/map/denmark.png"))
	state.Countries = append(state.Countries, system.NewClickableSprite("res/test/map/switzerland.png"))
	state.Countries = append(state.Countries, system.NewClickableSprite("res/test/map/spain.png"))
	state.Countries = append(state.Countries, system.NewClickableSprite("res/test/map/portugal.png"))
	state.Countries = append(state.Countries, system.NewClickableSprite("res/test/map/belgium.png"))
	state.Countries = append(state.Countries, system.NewClickableSprite("res/test/map/netherlands.png"))
	state.Countries = append(state.Countries, system.NewClickableSprite("res/test/map/luxembourg.png"))
	state.Countries = append(state.Countries, system.NewClickableSprite("res/test/map/ireland.png"))
	state.Countries = append(state.Countries, system.NewClickableSprite("res/test/map/turkey.png"))
	state.Countries = append(state.Countries, system.NewClickableSprite("res/test/map/greece.png"))

	for i := range state.Countries {
		state.Countries[i].Position.X = float32(raylib.GetScreenWidth() - state.Countries[i].Texture.Width)
	}
}

func (state *CountryState) Update() {
	raylib.ClearBackground(raylib.NewColor(0, 128, 255, 255))

	state.CurrentFrame += 0.5

	if state.WaterAnimation.Width/64 == int32(state.CurrentFrame) {
		state.CurrentFrame = 0.0
	}

	if state.ZoomOn {
		if state.Camera.Zoom < 8.0 {
			state.Camera.Zoom += 0.1
		} else {
			state.TakeScreenshot++

			if state.TakeScreenshot >= 2 {
				system.Screenshot = raylib.LoadTexture("res/tmp/screenshot.png")
				SetState("loading")
			}
		}
	}

}

func (state *CountryState) Draw() {

	// Begin camera
	raylib.BeginMode2D(state.Camera)

	// Water
	for x := float32(0); x < float32(raylib.GetScreenWidth()); x += 64 {
		for y := float32(0); y < float32(raylib.GetScreenHeight()); y += 64 {
			raylib.DrawTextureRec(state.WaterAnimation,
				raylib.NewRectangle(float32(int32(state.CurrentFrame)*64), 0, 64, 64),
				raylib.NewVector2(x, y),
				raylib.White)
		}
	}

	// Map
	for i := range state.Countries {

		if !state.ZoomOn {
			if state.Countries[i].Hover() {
				state.Countries[i].Color = raylib.NewColor(255, 0, 0, 255)

				if raylib.IsMouseButtonDown(raylib.MouseLeftButton) {
					state.ZoomOn = true
					state.Camera.Target = raylib.GetMousePosition()
				}
			} else {
				state.Countries[i].Color = raylib.White
			}
		}

		state.Countries[i].Draw()
	}

	// Gui
	raylib.DrawRectangle(3,
		3,
		int32(float32(system.CurrentFlag.Width)*0.15+4),
		int32(float32(system.CurrentFlag.Height)*0.15+4),
		raylib.Black) // Border

	raylib.DrawTextureEx(system.CurrentFlag, raylib.NewVector2(5, 5), 0.0, 0.15, raylib.White)

	// End camera
	raylib.EndMode2D()

	// Take screenshot, if needed
	if state.TakeScreenshot == 1 {
		state.TakeScreenshot++
		raylib.TakeScreenshot("res/tmp/screenshot.png")
	}
}

func (state *CountryState) Reset() {

}

func (state *CountryState) Close() {
	raylib.UnloadTexture(state.WaterAnimation)
}
