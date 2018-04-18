/*
 * League of Nations
 * ©DreamVelopper 2018
 */

// The main package
package main

// Importing all the stuff
import (
	"./system"
	"./system/states"
	"github.com/gen2brain/raylib-go/raylib"
)

func main() {
	system.Log("Initializing game...")

	raylib.InitWindow(1280, 720, "League of Nations")
	defer raylib.CloseWindow()

	raylib.InitAudioDevice()
	defer raylib.CloseAudioDevice()

	raylib.SetTargetFPS(60)

	// Initialize global stuff
	system.Initialize()

	// Initialize states manager
	states.Initialize()

	// Log it
	system.Log("Done !")

	for !raylib.WindowShouldClose() {
		raylib.BeginDrawing()
		raylib.ClearBackground(raylib.Black)

		// States manager
		states.Update()
		states.Draw()

		// Check if close window was requested while Update / Draw
		if system.Closed {
			break
		}

		raylib.EndDrawing()
	}

	// Close everything
	system.Close()
}
