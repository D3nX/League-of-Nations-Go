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
	LastProgess       float32
	Alpha             uint8
	Music             raylib.Music
}

func (state *LoadingState) Load() {
	state.LoadingTexture = raylib.LoadTexture("res/backgrounds/loading_screen.png")
	state.LoadingTexturePos = raylib.NewVector2(float32(raylib.GetScreenWidth()), 0)
	state.CurrentText = ""

	state.Alpha = 0

	state.Music = raylib.LoadMusicStream("res/musics/ost/Hearts of Iron IV Soundtrack_ Retribution.ogg")
	raylib.SetMusicLoopCount(state.Music, -1)
	raylib.PlayMusicStream(state.Music)
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
			state.Progress += 0.002
		} else {
			state.Progress = 1.0
		}

		// And the text
		if int32(state.Progress*100)%25 == 1 && int32(state.LastProgess*100) != int32(state.Progress*100) {

			state.LastProgess = state.Progress

			switch state.CurrentText {
			case "":
				state.CurrentText = "Preparing the best army..."

			case "Preparing the best army...":
				state.CurrentText = "Launching plane..."
				break

			case "Launching plane...":
				state.CurrentText = "Creating next-gen Panzer..."
				break

			case "Creating next-gen Panzer...":
				state.CurrentText = "Scaring the ennemy..."
				break
			}
		}
	}

	// Darking the screen & make music lower when quitting state
	if int32(state.Progress) >= 1.0 {

		if state.Alpha+3 < 255 {
			state.Alpha += 3
			raylib.SetMusicVolume(state.Music, float32(255-state.Alpha)/255)
		} else {
			SetState("game")
		}

	}

	// Updating music
	raylib.UpdateMusicStream(state.Music)

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
	raylib.StopMusicStream(state.Music)
	raylib.UnloadMusicStream(state.Music)
}
