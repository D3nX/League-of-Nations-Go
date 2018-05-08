package states

import (
	"../../system"
	"github.com/gen2brain/raylib-go/raylib"
)

var currentState string
var states map[string]State
var nextState string

var backAnim system.BackAnim

func Initialize() {
	states = make(map[string]State)
	states["menu"] = &MenuState{}
	states["settings"] = &SettingsState{}
	states["nation_creator"] = &NationCreatorState{}
	states["country"] = &CountryState{}
	states["loading"] = &LoadingState{}
	states["game"] = &GameState{}

	// Initializing game states
	for _, state := range states {
		state.Load()
	}

	// Setting current state to menu
	currentState = "menu"

	// Initialize the back anim struct
	backAnim = system.BackAnim{}
	backAnim.Load()
}

func Update() {

	if raylib.IsKeyDown(raylib.KeyF1) {
		SetState("menu")
	}

	states[currentState].Update()
}

func Draw() {
	states[currentState].Draw()
}

func Close() {
	backAnim.Close()

	for i := range states {
		states[i].Close()
	}
}

func SetState(input string) {
	currentState = input
	states[currentState].Reset()
}
