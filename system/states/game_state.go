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
	DEBUG_MUSIC = false
)

type GameState struct {
	Alpha        uint8
	Map          *gamemap.GameMap
	Music        map[string]raylib.Music
	Camera       raylib.Camera2D
	CameraTarget raylib.Vector2

	// Private stuff
	panelButtons []ui.PanelButton
	pickedObject string
	selectedId   int
	pickable     bool
}

func (state *GameState) Load() {

	// Initialize the alpha
	state.Alpha = 255

	// Load the map
	state.Map = &gamemap.GameMap{}

	state.Map.Load("res/map/map.txt")

	// Load & initialize the music
	state.Music = make(map[string]raylib.Music)

	state.Music["preparing"] = raylib.LoadMusicStream("res/musics/ost/War Thunder Soundtrack_ Menu Music 4.ogg")
	raylib.SetMusicLoopCount(state.Music["preparing"], -1)
	raylib.PlayMusicStream(state.Music["preparing"])

	// Create the camera and its target
	state.Camera = raylib.NewCamera2D(raylib.NewVector2(0, 0),
		raylib.NewVector2(float32(raylib.GetScreenWidth())/2, float32(raylib.GetScreenHeight())/2),
		0.0,
		1.0)

	state.CameraTarget = raylib.NewVector2(0, 0)

	// For know which object is picked for being placed
	state.pickedObject = ""

	// Set the selected object to -1, since no any are selected
	state.selectedId = -1

	// Initialize panel button
	state.panelButtons = make([]ui.PanelButton, 1)
	state.panelButtons[0] = ui.NewPanelButton(5, float32(raylib.GetScreenHeight())-175, 170, 170, "12000$", &objects.TankTextures[0])
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
	/*if state.pickedObject == "" {
		for x := range state.Map.Tiles {
			for y := range state.Map.Tiles[x] {
				if state.Map.Tiles[x][y].ClickedOn() {
					// fmt.Println("Pressed on tile : x =", x, " y =", y)
				}
			}
		}
	}*/

	// Iterate over objects
	state.selectedId = -1
	for i, obj := range state.Map.Objects {
		if obj.IsSelected() {
			state.selectedId = i
		}
	}

	if state.selectedId != -1 {
		for i, obj := range state.Map.Objects {
			if i != state.selectedId {
				obj.SetSelected(false)
			}
		}
	}

	// Move camera depending mouse position

	state.pickable = true // A pickable variable to know if we can pick any object

	if raylib.IsKeyDown(raylib.KeyLeft) {
		state.Camera.Offset.X += 5
		state.CameraTarget.X += 5

		if state.selectedId != -1 {
			state.Map.Objects[state.selectedId].SetSelected(false)
			state.selectedId = -1
		}

		state.pickable = false
	}

	if raylib.IsKeyDown(raylib.KeyRight) {
		state.Camera.Offset.X -= 5
		state.CameraTarget.X -= 5

		if state.selectedId != -1 {
			state.Map.Objects[state.selectedId].SetSelected(false)
			state.selectedId = -1
		}

		state.pickable = false
	}

	if raylib.IsKeyDown(raylib.KeyUp) {
		state.Camera.Offset.Y += 5
		state.CameraTarget.Y += 5

		if state.selectedId != -1 {
			state.Map.Objects[state.selectedId].SetSelected(false)
			state.selectedId = -1
		}

		state.pickable = false
	}

	if raylib.IsKeyDown(raylib.KeyDown) {
		state.Camera.Offset.Y -= 5
		state.CameraTarget.Y -= 5

		if state.selectedId != -1 {
			state.Map.Objects[state.selectedId].SetSelected(false)
			state.selectedId = -1
		}

		state.pickable = false
	}

	if state.selectedId != -1 {
		state.CameraTarget = state.Map.Objects[state.selectedId].GetPosition()

		state.CameraTarget.X = -state.CameraTarget.X
		state.CameraTarget.Y = -state.CameraTarget.Y

		state.CameraTarget.X += float32(raylib.GetScreenWidth()) / 2
		state.CameraTarget.Y += float32(raylib.GetScreenHeight()) / 2
	}

	if state.Camera.Offset.X > state.CameraTarget.X {
		state.Camera.Offset.X += float32(float64(state.CameraTarget.X-state.Camera.Offset.X) * 0.1)

		if state.Camera.Offset.X <= state.CameraTarget.X {
			state.Camera.Offset.X = state.CameraTarget.X
		}
	} else if state.Camera.Offset.X < state.CameraTarget.X {
		state.Camera.Offset.X += float32(float64(state.CameraTarget.X-state.Camera.Offset.X) * 0.1)

		if state.Camera.Offset.X >= state.CameraTarget.X {
			state.Camera.Offset.X = state.CameraTarget.X
		}
	}

	if state.Camera.Offset.Y > state.CameraTarget.Y {
		state.Camera.Offset.Y += float32(float64(state.CameraTarget.Y-state.Camera.Offset.Y) * 0.1)

		if state.Camera.Offset.Y <= state.CameraTarget.Y {
			state.Camera.Offset.Y = state.CameraTarget.Y
		}
	} else if state.Camera.Offset.Y < state.CameraTarget.Y {
		state.Camera.Offset.Y += float32(float64(state.CameraTarget.Y-state.Camera.Offset.Y) * 0.1)

		if state.Camera.Offset.Y >= state.CameraTarget.Y {
			state.Camera.Offset.Y = state.CameraTarget.Y
		}
	}

	// Cursor stuff
	if state.pickedObject != "" && raylib.IsMouseButtonPressed(raylib.MouseRightButton) {
		state.pickedObject = ""
	}

	if state.pickedObject != "" {

		pos := raylib.NewVector2(float32(int32(float32(raylib.GetMouseX()-objects.TankTextures[0].Width/2)/32)*32),
			float32(int32(float32(raylib.GetMouseY()-objects.TankTextures[0].Height/2)/32)*32))

		switch state.pickedObject {
		case "tank:0":
			if raylib.IsMouseButtonPressed(raylib.MouseLeftButton) {
				for _, object := range state.Map.Objects {
					switch object.(type) {
					case *objects.Tank:

						if object.Collides(raylib.NewRectangle(pos.X-state.Camera.Offset.X,
							pos.Y-state.Camera.Offset.Y,
							float32(objects.TankTextures[0].Width),
							float32(objects.TankTextures[0].Height))) {
							goto end
						}
					}
				}
				if tile := state.Map.GetTile(int32(pos.X-state.Camera.Offset.X+float32(objects.TankTextures[0].Width/2)), int32(pos.Y-state.Camera.Offset.Y+float32(objects.TankTextures[0].Height/2))); tile != nil {
					if !state.Map.IsObjectAt(tile.X, tile.Y) {
						state.Map.AddObject(objects.NewTank(0, tile.X, tile.Y, 0.0))
					}
				}
			}
		}
	}

end:

	// Update the game map
	if state.pickedObject != "" || raylib.CheckCollisionRecs(raylib.NewRectangle(0, float32(raylib.GetScreenHeight())-180, float32(raylib.GetScreenWidth()), 200),
		raylib.NewRectangle(float32(raylib.GetMouseX()), float32(raylib.GetMouseY()), 2, 2)) {
		state.pickable = false
	}
}

func (state *GameState) Draw() {

	// Drawing the map & update objects
	state.Map.Draw(&state.Camera, func(id int, object *objects.Object) {
		if state.selectedId == -1 {
			(*object).Update(&state.Camera, state.pickable)
			(*object).Draw(&state.Camera)
		} else {
			if state.selectedId != id {
				/*for _, object := range state.Map.Objects {
					object.Update(&state.Camera, false)
				}*/
				(*object).Update(&state.Camera, false)
				(*object).Draw(&state.Camera)
			}
		}
	})

	// We draw the selected object first
	if state.selectedId != -1 {
		state.Map.Objects[state.selectedId].Update(&state.Camera, true)
		state.Map.Objects[state.selectedId].Draw(&state.Camera)
	}

	// Helper
	raylib.DrawText(fmt.Sprint("Camera\nX : ", state.Camera.Offset.X, "\nY : ", state.Camera.Offset.Y), 5, 0, 30, raylib.White)

	// Cursor stuff
	if state.pickedObject != "" {
		switch state.pickedObject {
		case "tank:0":
			pos := raylib.NewVector2(float32(int32(float32(raylib.GetMouseX()-objects.TankTextures[0].Width/2)/32)*32),
				float32(int32(float32(raylib.GetMouseY()-objects.TankTextures[0].Height/2)/32)*32))

			color := raylib.NewColor(255, 255, 255, 128)
			for _, object := range state.Map.Objects {
				switch object.(type) {
				case *objects.Tank:

					if object.Collides(raylib.NewRectangle(pos.X-state.Camera.Offset.X, pos.Y-state.Camera.Offset.Y, float32(objects.TankTextures[0].Width), float32(objects.TankTextures[0].Height))) {
						color = raylib.Red
						color.A = 128
						goto drawing
					}
				}
			}

		drawing:
			raylib.DrawTexture(objects.TankTextures[0],
				int32(pos.X),
				int32(pos.Y),
				color)
		}
	}

	// The UI
	if DEBUG_UI {

		// Drawing the panel title
		text := "Army"

		raylib.DrawRectangle(0,
			raylib.GetScreenHeight()-200,
			raylib.MeasureText(text, 15)+9,
			23,
			raylib.White)

		raylib.DrawRectangleLines(3,
			raylib.GetScreenHeight()-197,
			raylib.MeasureText(text, 15)+3,
			17,
			raylib.Black)

		raylib.DrawText(text, 5, raylib.GetScreenHeight()-196, 15, raylib.Black)

		// Drawing the panel
		raylib.DrawRectangle(0,
			raylib.GetScreenHeight()-180,
			raylib.GetScreenWidth(),
			200,
			raylib.White)

		raylib.DrawRectangleLines(3,
			raylib.GetScreenHeight()-177,
			raylib.GetScreenWidth()-6,
			174,
			raylib.Black)

		// Draw the panel buttons
		if state.panelButtons[0].Draw() {
			state.pickedObject = "tank:0"
		}

		// Drawing the filter (for un-darking screen)
		raylib.DrawRectangle(0,
			0,
			raylib.GetScreenWidth(),
			raylib.GetScreenHeight(),
			raylib.NewColor(0, 0, 0, state.Alpha))
	}

	// DO NOT ADD CODE UNDER IF IT'S NOT UI
}

func (state *GameState) Reset() {

}

func (state *GameState) Close() {

}
