package states

import (
	"../gamemap"
)

type GameState struct {
	Alpha uint8
	Map   *gamemap.GameMap
}

func (state *GameState) Load() {
	state.Alpha = 255
	state.Map = &gamemap.GameMap{}

	state.Map.Load("res/map/map.txt")
}

func (state *GameState) Update() {

	// Introduction dark filter
	/*if state.Alpha-3 > 0 {
		state.Alpha -= 3
	}*/
}

func (state *GameState) Draw() {

	// Drawing the map
	state.Map.Draw()

	// Drawing the filter (for un-darking screen)
	/*raylib.DrawRectangle(0,
	0,
	raylib.GetScreenWidth(),
	raylib.GetScreenHeight(),
	raylib.NewColor(0, 0, 0, state.Alpha))*/
}

func (state *GameState) Reset() {

}

func (state *GameState) Close() {

}
