package states

import (
	"fmt"

	"../../system"
	"github.com/gen2brain/raylib-go/raygui"
	"github.com/gen2brain/raylib-go/raylib"
)

type SettingsState struct {
	// Button stuff
	BackButtonPressed bool
}

func (state *SettingsState) Load() {

}

func (state *SettingsState) Update() {
	// Update the background animation
	backAnim.Update()

	// Check button pressed
	if state.BackButtonPressed {
		SetState("menu")
	}
}

func (state *SettingsState) Draw() {
	// Draw the background animation
	backAnim.Draw()

	// Drawing "Settings" text
	raylib.DrawTextEx(system.FontKremlin,
		"Settings",
		raylib.Vector2{X: float32((raylib.GetScreenWidth() - raylib.MeasureText("Settings", 96)) / 2),
			Y: 5},
		96,
		2,
		raylib.Yellow)

	// Draw the panel
	raylib.DrawRectangleLines(10,
		100,
		raylib.GetScreenWidth()-20,
		raylib.GetScreenHeight()-200,
		raygui.LinesColor())

	raylib.DrawRectangle(11,
		101,
		raylib.GetScreenWidth()-22,
		raylib.GetScreenHeight()-202,
		raylib.NewColor(128, 128, 128, 128))

	// The categories
	raylib.DrawTextEx(system.FontKremlin,
		"Sound and Music :",
		raylib.Vector2{X: 20, Y: 110},
		20,
		2,
		raylib.Yellow)

	// Music Volume stuff
	raylib.DrawTextEx(system.FontKremlin,
		"Music Volume :",
		raylib.Vector2{X: 20, Y: 140},
		20,
		2,
		raylib.White)

	raylib.DrawTextEx(system.FontKremlin,
		fmt.Sprint(int(system.GameSettings.MusicVolume), "%"), // Font to change (-> not supporting "%" symbole)
		raylib.Vector2{X: 395, Y: 140},
		20,
		2,
		raylib.White)

	system.GameSettings.MusicVolume = raygui.SliderBar(
		raylib.NewRectangle(40+float32(raylib.MeasureText("Music Volume :", 20)), 140, 200, 20),
		system.GameSettings.MusicVolume,
		0,
		100)

	// Sound Volume stuff
	raylib.DrawTextEx(system.FontKremlin,
		"SFX Volume :",
		raylib.Vector2{X: 20, Y: 170},
		20,
		2,
		raylib.White)

	raylib.DrawTextEx(system.FontKremlin,
		fmt.Sprint(int(system.GameSettings.SFXVolume), "%"), // Font to change (-> not supporting "%" symbole)
		raylib.Vector2{X: 395, Y: 170},
		20,
		2,
		raylib.White)

	system.GameSettings.SFXVolume = raygui.SliderBar(
		raylib.NewRectangle(40+float32(raylib.MeasureText("Music Volume :", 20)), 170, 200, 20),
		system.GameSettings.SFXVolume,
		0,
		100)

	// Drawing the buttons

	// Back button
	state.BackButtonPressed = raygui.Button(raylib.NewRectangle(5,
		float32(raylib.GetScreenHeight()-47),
		200,
		45),
		"Back")
}

func (state *SettingsState) Reset() {

}

func (state *SettingsState) Close() {

}
