package states

import (
	"../../system"
	"github.com/gen2brain/raylib-go/raygui"
	"github.com/gen2brain/raylib-go/raylib"
)

type MenuState struct {
	ButtonsClicked []bool
}

func (state *MenuState) Load() {

	// Button stuff
	state.ButtonsClicked = make([]bool, 3)
}

func (state *MenuState) Update() {

	// Update the background animation
	backAnim.Update()
	// fmt.Println("alpha = ", state.BlackScreenAlpha)

	// Pressed on the new game button
	if state.ButtonsClicked[0] {
		SetState("nation_creator")
	}

	// Pressed on the settings button
	if state.ButtonsClicked[1] {
		SetState("settings")
	}

	// Pressed on the exit button
	if state.ButtonsClicked[2] {
		system.Closed = true
	}
}

func (state *MenuState) Draw() {
	// Drawing the background
	backAnim.Draw()

	// Drawing the title
	raylib.DrawTextEx(system.FontKremlin,
		system.Title,
		raylib.Vector2{X: float32((raylib.GetScreenWidth() - raylib.MeasureText(system.Title, 96)) / 2),
			Y: 5},
		96,
		2,
		raylib.Yellow)

	// Drawing the little panel
	width := int32(300)
	height := int32(350)
	raylib.DrawRectangle((raylib.GetScreenWidth()-width)/2,
		(raylib.GetScreenHeight()-height)/2,
		width,
		height,
		raylib.NewColor(32, 32, 32, 255))

	raylib.DrawRectangle((raylib.GetScreenWidth()-width)/2+5,
		(raylib.GetScreenHeight()-height)/2+5,
		width-10,
		height-10,
		raylib.NewColor(64, 64, 64, 255))

	// Drawing the button
	state.ButtonsClicked[0] = raygui.Button(raylib.NewRectangle((raylib.GetScreenWidth()-200)/2,
		250,
		200,
		45),
		"New game")

	state.ButtonsClicked[1] = raygui.Button(raylib.NewRectangle((raylib.GetScreenWidth()-200)/2,
		320,
		200,
		45),
		"Settings")

	state.ButtonsClicked[2] = raygui.Button(raylib.NewRectangle((raylib.GetScreenWidth()-200)/2,
		390,
		200,
		45),
		"Exit")
}

func (state *MenuState) Reset() {

}

func (state *MenuState) Close() {

}
