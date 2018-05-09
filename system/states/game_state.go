package states

import (
	"fmt"

	"../gamemap"
	"github.com/gen2brain/raylib-go/raylib"
)

type GameState struct {
	Alpha uint8
	Map   *gamemap.GameMap
	Music map[string]raylib.Music
}

func (state *GameState) Load() {
	state.Alpha = 255
	state.Map = &gamemap.GameMap{}

	state.Map.Load("res/map/map.txt")

	state.Music = make(map[string]raylib.Music)

	state.Music["preparing"] = raylib.LoadMusicStream("res/musics/ost/War Thunder Soundtrack_ Menu Music 4.ogg")
	raylib.SetMusicLoopCount(state.Music["preparing"], -1)
	raylib.PlayMusicStream(state.Music["preparing"])
}

func (state *GameState) Update() {

	// Temp stuff
	// Reload map
	if raylib.IsKeyPressed(raylib.KeyR) {
		state.Map.Load("res/map/map.txt")
	}
	// End temp stuff

	// Introduction dark filter & sound uppering
	if state.Alpha-3 > 0 {
		state.Alpha -= 3
		raylib.SetMusicVolume(state.Music["preparing"], float32(255-state.Alpha)/255)
	}

	// Updating music
	// raylib.UpdateMusicStream(state.Music["preparing"])

	// Check if button clicked on one tile
	for x := range state.Map.Tiles {
		for y := range state.Map.Tiles[x] {
			if state.Map.Tiles[x][y].ClickedOn() {
				fmt.Println("Pressed on tile : x =", x, " y =", y)
			}
		}
	}
}

func (state *GameState) Draw() {

	// Drawing the map
	state.Map.Draw()

	// Drawing the filter (for un-darking screen)
	raylib.DrawRectangle(0,
		0,
		raylib.GetScreenWidth(),
		raylib.GetScreenHeight(),
		raylib.NewColor(0, 0, 0, state.Alpha))
}

func (state *GameState) Reset() {

}

func (state *GameState) Close() {

}
