package states

import (
	"../../system"
	"github.com/gen2brain/raylib-go/raygui"
	"github.com/gen2brain/raylib-go/raylib"
)

type LoadingState struct {
	LoadingTexture    raylib.Texture2D
	LoadingTexturePos raylib.Vector2
	CurrentText       string
	Progress          float32
	Alpha             uint8
}

func (state *LoadingState) Load() {
	state.LoadingTexture = raylib.LoadTexture("res/backgrounds/loading_screen.png")
	state.LoadingTexturePos = raylib.NewVector2(float32(raylib.GetScreenWidth()), 0)
	state.CurrentText = "Preparing the best army..."

	state.Alpha = 0
}

func (state *LoadingState) Update() {

	// Updating texture position (while introduction)
	if state.LoadingTexturePos.X > 0 {
		state.LoadingTexturePos.X -= 15

		if state.LoadingTexturePos.X < 0 {
			state.LoadingTexturePos.X = 0
		}
	}

	// Updating progess bar
	if int32(state.LoadingTexturePos.X) <= 0 {
		if state.Progress < 1.0 {
			state.Progress += 0.005
		} else {
			state.Progress = 1.0
		}
	}

	// Darking the screen when quitting state
	if int32(state.Progress) >= 1.0 {

		if state.Alpha+3 < 255 {
			state.Alpha += 3
		} else {
			SetState("game")
		}

	}

}

func (state *LoadingState) Draw() {

	// Drawing the screenshot
	if int32(state.LoadingTexturePos.X) > 0 {
		raylib.DrawTexture(system.Screenshot,
			int32(state.LoadingTexturePos.X)-raylib.GetScreenWidth(),
			0,
			raylib.White)
	}

	// Drawing the texture
	raylib.DrawTexture(state.LoadingTexture,
		int32(state.LoadingTexturePos.X),
		0,
		raylib.White)

	// Draw the ui

	// Panel
	raylib.DrawRectangle(int32(state.LoadingTexturePos.X), 0, raylib.GetScreenWidth(), 95, raylib.White)
	raylib.DrawRectangle(int32(state.LoadingTexturePos.X)+4, 4, raylib.GetScreenWidth()-8, 87, raylib.Black)

	// Text
	raylib.DrawTextEx(system.FontKremlin,
		state.CurrentText,
		raylib.NewVector2(state.LoadingTexturePos.X+10, 23),
		50,
		0,
		raylib.White)

	// Progress bar
	raygui.SetStyleProperty(raygui.ProgressbarProgressColor, 0xff0000ff)
	raygui.SetStyleProperty(raygui.ProgressbarInsideColor, 0xffffffff)
	raygui.ProgressBar(raylib.NewRectangle(state.LoadingTexturePos.X+700, 18, 550, 48), state.Progress)

	// Drawing the filter (for darking screen)
	raylib.DrawRectangle(0,
		0,
		raylib.GetScreenWidth(),
		raylib.GetScreenHeight(),
		raylib.NewColor(0, 0, 0, state.Alpha))
}

func (state *LoadingState) Reset() {

}

func (state *LoadingState) Close() {

}
