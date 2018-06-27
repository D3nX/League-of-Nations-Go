package states

import (
	"fmt"

	"../gamemap"
	"../objects"
	"../objects/ui"
	"github.com/gen2brain/raylib-go/raylib"
)

const (
	DEBUG_UI    = true
	DEBUG_MUSIC = true
)

type GameState struct {
	Alpha        uint8
	Map          *gamemap.GameMap
	Music        map[string]raylib.Music
	Camera       raylib.Camera2D
	panelButtons []ui.PanelButton
}

func (state *GameState) Load() {
	state.Alpha = 255
	state.Map = &gamemap.GameMap{}

	state.Map.Load("res/map/map.txt")

	state.Music = make(map[string]raylib.Music)

	state.Music["preparing"] = raylib.LoadMusicStream("res/musics/ost/War Thunder Soundtrack_ Menu Music 4.ogg")
	raylib.SetMusicLoopCount(state.Music["preparing"], -1)
	raylib.PlayMusicStream(state.Music["preparing"])

	state.Camera = raylib.NewCamera2D(raylib.NewVector2(0, 0),
		raylib.NewVector2(float32(raylib.GetScreenWidth())/2, float32(raylib.GetScreenHeight())/2),
		0.0,
		1.0)

	// Initialize panel button
	state.panelButtons = make([]ui.PanelButton, 1)
	state.panelButtons[0] = ui.NewPanelButton(5, float32(raylib.GetScreenHeight())-195, 190, 190, "12000$", &objects.TankTextures[0])
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
	if DEBUG_MUSIC {
		raylib.UpdateMusicStream(state.Music["preparing"])
	}

	// Check if button clicked on one tile
	for x := range state.Map.Tiles {
		for y := range state.Map.Tiles[x] {
			if state.Map.Tiles[x][y].ClickedOn() {
				// fmt.Println("Pressed on tile : x =", x, " y =", y)
			}
		}
	}

	// Iterate over objects
	selectedId := int(-1)
	for i, obj := range state.Map.Objects {
		if obj.IsSelected() {
			selectedId = i
		}
	}

	if selectedId != -1 {
		for i, obj := range state.Map.Objects {
			if i != selectedId {
				obj.SetSelected(false)
			}
		}
	}

	// Move camera depending mouse position

	if raylib.IsKeyDown(raylib.KeyLeft) {
		state.Camera.Offset.X += 5

		if selectedId != -1 {
			state.Map.Objects[selectedId].SetSelected(false)
			selectedId = -1
		}
	}

	if raylib.IsKeyDown(raylib.KeyRight) {
		state.Camera.Offset.X -= 5

		if selectedId != -1 {
			state.Map.Objects[selectedId].SetSelected(false)
			selectedId = -1
		}
	}

	if raylib.IsKeyDown(raylib.KeyUp) {
		state.Camera.Offset.Y += 5

		if selectedId != -1 {
			state.Map.Objects[selectedId].SetSelected(false)
			selectedId = -1
		}
	}

	if raylib.IsKeyDown(raylib.KeyDown) {
		state.Camera.Offset.Y -= 5

		if selectedId != -1 {
			state.Map.Objects[selectedId].SetSelected(false)
			selectedId = -1
		}
	}

	if selectedId != -1 {
		state.Camera.Offset = state.Map.Objects[selectedId].GetPosition()

		state.Camera.Offset.X = -state.Camera.Offset.X
		state.Camera.Offset.Y = -state.Camera.Offset.Y

		state.Camera.Offset.X += float32(raylib.GetScreenWidth()) / 2
		state.Camera.Offset.Y += float32(raylib.GetScreenHeight()) / 2
	}

	// Update the game map
	state.Map.Update(&state.Camera)
}

func (state *GameState) Draw() {

	// Drawing the map
	state.Map.Draw(&state.Camera)

	// Helper
	raylib.DrawText(fmt.Sprint("Camera\nX : ", state.Camera.Offset.X, "\nY : ", state.Camera.Offset.Y), 5, 0, 30, raylib.White)
	// End helper stuff

	// The UI

	// Drawing the panel
	if DEBUG_UI {
		raylib.DrawRectangle(0,
			raylib.GetScreenHeight()-200,
			raylib.GetScreenWidth(),
			200,
			raylib.White)

		raylib.DrawRectangleLines(3,
			raylib.GetScreenHeight()-197,
			raylib.GetScreenWidth()-6,
			194,
			raylib.Black)

		// Draw the panel buttons
		for _, b := range state.panelButtons {
			if b.Draw() {
				fmt.Println("It finally works comrade !")
			}
		}

		// Drawing the filter (for un-darking screen)
		raylib.DrawRectangle(0,
			0,
			raylib.GetScreenWidth(),
			raylib.GetScreenHeight(),
			raylib.NewColor(0, 0, 0, state.Alpha))
	}

	// DO NOT ADD CODE UNDER
}

func (state *GameState) Reset() {

}

func (state *GameState) Close() {

}
