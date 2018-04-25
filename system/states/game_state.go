package states

import (
	"fmt"

	"../../system"
	"github.com/gen2brain/raylib-go/raylib"
)

type GameState struct {
	WaterAnimation raylib.Texture2D
	CurrentFrame   float32

	// Test
	Map raylib.Texture2D
}

func (state *GameState) Load() {
	// Load the water animation
	state.WaterAnimation = raylib.LoadTexture("res/sprites/water_animation.png")
	state.CurrentFrame = 0
	state.Map = raylib.LoadTexture("res/test/square.png")
}

func (state *GameState) Update() {
	raylib.ClearBackground(raylib.NewColor(0, 128, 255, 255))

	state.CurrentFrame += 0.5

	if state.WaterAnimation.Width/64 == int32(state.CurrentFrame) {
		state.CurrentFrame = 0.0
	}

	if system.MouseOn(raylib.NewVector2(0, 0), 1.0, raylib.GetTextureData(state.Map)) {
		fmt.Println("It's amazing ! It works !")
	}
}

func (state *GameState) Draw() {
	// Water
	for x := float32(0); x < float32(raylib.GetScreenWidth()); x += 64 {
		for y := float32(0); y < float32(raylib.GetScreenHeight()); y += 64 {
			raylib.DrawTextureRec(state.WaterAnimation,
				raylib.NewRectangle(int32(state.CurrentFrame)*64, 0, 64, 64),
				raylib.NewVector2(x, y),
				raylib.White)
		}
	}

	// Map
	raylib.DrawTexture(state.Map, 0, 0, raylib.White)

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
