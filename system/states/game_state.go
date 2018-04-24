package states

import "github.com/gen2brain/raylib-go/raylib"

type GameState struct {
	WaterAnimation raylib.Texture2D
	CurrentFrame   float32
}

func (state *GameState) Load() {
	// Load the water animation
	state.WaterAnimation = raylib.LoadTexture("res/sprites/water_animation.png")
	state.CurrentFrame = 0
}

func (state *GameState) Update() {
	raylib.ClearBackground(raylib.NewColor(0, 128, 255, 255))

	state.CurrentFrame += 0.5

	if state.WaterAnimation.Width/64 == int32(state.CurrentFrame) {
		state.CurrentFrame = 0.0
	}
}

func (state *GameState) Draw() {
	for x := float32(0); x < float32(raylib.GetScreenWidth()); x += 64 {
		for y := float32(0); y < float32(raylib.GetScreenHeight()); y += 64 {
			raylib.DrawTextureRec(state.WaterAnimation,
				raylib.NewRectangle(int32(state.CurrentFrame)*64, 0, 64, 64),
				raylib.NewVector2(x, y),
				raylib.White)
		}
	}
}

func (state *GameState) Reset() {

}

func (state *GameState) Close() {
	raylib.UnloadTexture(state.WaterAnimation)
}
