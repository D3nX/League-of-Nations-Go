package states

import (
	"../../system"
	"github.com/gen2brain/raylib-go/raylib"
)

type GameState struct {
	WaterAnimation raylib.Texture2D
	CurrentFrame   float32

	// Test
	Colors    []raylib.Color
	Countries []raylib.Texture2D
}

func (state *GameState) Load() {
	// Load the water animation
	state.WaterAnimation = raylib.LoadTexture("res/sprites/water_animation.png")
	state.CurrentFrame = 0

	state.Countries = append(state.Countries, raylib.LoadTexture("res/test/map/russia.png"))
	state.Countries = append(state.Countries, raylib.LoadTexture("res/test/map/ukraine.png"))
	state.Countries = append(state.Countries, raylib.LoadTexture("res/test/map/baltic_states.png"))
	state.Countries = append(state.Countries, raylib.LoadTexture("res/test/map/poland.png"))
	state.Countries = append(state.Countries, raylib.LoadTexture("res/test/map/kazakhstan.png"))
	state.Countries = append(state.Countries, raylib.LoadTexture("res/test/map/belarus.png"))
	state.Countries = append(state.Countries, raylib.LoadTexture("res/test/map/czech_republic.png"))
	state.Countries = append(state.Countries, raylib.LoadTexture("res/test/map/balkan.png"))
	state.Countries = append(state.Countries, raylib.LoadTexture("res/test/map/slovakia.png"))
	state.Countries = append(state.Countries, raylib.LoadTexture("res/test/map/hungary.png"))
	state.Countries = append(state.Countries, raylib.LoadTexture("res/test/map/united_kingdom.png"))
	state.Countries = append(state.Countries, raylib.LoadTexture("res/test/map/germany.png"))
	state.Countries = append(state.Countries, raylib.LoadTexture("res/test/map/austria.png"))
	state.Countries = append(state.Countries, raylib.LoadTexture("res/test/map/france.png"))
	state.Countries = append(state.Countries, raylib.LoadTexture("res/test/map/italy.png"))
	state.Countries = append(state.Countries, raylib.LoadTexture("res/test/map/slovenia.png"))
	state.Countries = append(state.Countries, raylib.LoadTexture("res/test/map/finland.png"))
	state.Countries = append(state.Countries, raylib.LoadTexture("res/test/map/norway.png"))
	state.Countries = append(state.Countries, raylib.LoadTexture("res/test/map/sweden.png"))
	state.Countries = append(state.Countries, raylib.LoadTexture("res/test/map/denmark.png"))
	state.Countries = append(state.Countries, raylib.LoadTexture("res/test/map/switzerland.png"))
	state.Countries = append(state.Countries, raylib.LoadTexture("res/test/map/spain.png"))
	state.Countries = append(state.Countries, raylib.LoadTexture("res/test/map/portugal.png"))
	state.Countries = append(state.Countries, raylib.LoadTexture("res/test/map/belgium.png"))
	state.Countries = append(state.Countries, raylib.LoadTexture("res/test/map/netherlands.png"))
	state.Countries = append(state.Countries, raylib.LoadTexture("res/test/map/luxembourg.png"))
	state.Countries = append(state.Countries, raylib.LoadTexture("res/test/map/ireland.png"))
	state.Countries = append(state.Countries, raylib.LoadTexture("res/test/map/turkey.png"))
	state.Countries = append(state.Countries, raylib.LoadTexture("res/test/map/greece.png"))

	state.Colors = make([]raylib.Color, len(state.Countries))
}

func (state *GameState) Update() {
	raylib.ClearBackground(raylib.NewColor(0, 128, 255, 255))

	state.CurrentFrame += 0.5

	if state.WaterAnimation.Width/64 == int32(state.CurrentFrame) {
		state.CurrentFrame = 0.0
	}

	/*for i, country := range state.Countries {
		if system.MouseOn(raylib.NewVector2(0, 0), 1.0, raylib.GetTextureData(country)) {
			state.Colors[i] = raylib.Red
		} else {
			state.Colors[i] = raylib.White
		}
	}*/
}

func (state *GameState) Draw() {
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
	for i, country := range state.Countries {
		state.Colors[i] = raylib.White
		raylib.DrawTexture(country, 0, 0, state.Colors[i])
	}

	// Gui
	raylib.DrawRectangle(3,
		3,
		int32(float32(system.CurrentFlag.Width)*0.15+4),
		int32(float32(system.CurrentFlag.Height)*0.15+4),
		raylib.Black) // Border

	raylib.DrawTextureEx(system.CurrentFlag, raylib.NewVector2(5, 5), 0.0, 0.15, raylib.White)
}

func (state *GameState) Reset() {

}

func (state *GameState) Close() {
	raylib.UnloadTexture(state.WaterAnimation)
}
