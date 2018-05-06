package states

import (
	"math/rand"

	"../../system"
	"github.com/gen2brain/raylib-go/raygui"
	"github.com/gen2brain/raylib-go/raylib"
)

type NationCreatorState struct {
	// Buttons
	BackButtonPressed bool
	PlayButtonPressed bool
	// Flag stuff
	GenNewFlag    bool
	FlagDirChoice int
	Flag          raylib.Texture2D
	// Leader stuff
	CurrentLeader int
	Leaders       []raylib.Texture2D
	// Country stuff
	CountryName       string
	CountryLeader     string
	LastIdeology      string
	EnableCommunism   bool
	EnableNationalism bool
	EnableDemocracy   bool
	AnthemPlaying     bool
	CurrentAnthem     int
	CountrySize       int
}

func (state *NationCreatorState) Load() {

	// Initialize variables
	state.BackButtonPressed = false

	state.FlagDirChoice = 0

	state.GenNewFlag = false

	state.CurrentLeader = 1

	state.CountryName = ""

	state.CountryLeader = ""

	state.LastIdeology = "communism"

	state.CurrentAnthem = 1

	state.AnthemPlaying = false

	state.EnableCommunism = true
	state.EnableDemocracy = false
	state.EnableNationalism = false

	state.CountrySize = 0

	state.PlayButtonPressed = false

	dir := false
	if state.FlagDirChoice == 0 {
		dir = true
	}
	image := system.GetNewFlag(dir, 1+rand.Intn(4), state.LastIdeology)
	state.Flag = raylib.LoadTextureFromImage(image)
	raylib.UnloadImage(image)

	// Copying leader image and resizing it
	state.Leaders = make([]raylib.Texture2D, 5)

	for i := 0; i < len(state.Leaders); i++ {
		image := raylib.ImageCopy(system.Leaders[i])
		raylib.ImageResize(image, 346, 386)

		state.Leaders[i] = raylib.LoadTextureFromImage(image)

		raylib.UnloadImage(image)
	}

}

func (state *NationCreatorState) Update() {

	// Update the background animation
	backAnim.Update()

	// Go to main menu if back button pressed
	if state.BackButtonPressed {
		SetState("menu")
	}

	// Generate a new flag if needed
	if state.GenNewFlag {
		dir := false
		if state.FlagDirChoice == 0 {
			dir = true
		}
		image := system.GetNewFlag(dir, 1+rand.Intn(4), state.LastIdeology)
		raylib.UnloadTexture(state.Flag)
		state.Flag = raylib.LoadTextureFromImage(image)
		raylib.UnloadImage(image)
	}

	// Leader stuff
	if state.CurrentLeader > len(system.Leaders) {
		state.CurrentLeader = 1
	}

	if state.CurrentLeader < 1 {
		state.CurrentLeader = len(system.Leaders)
	}

	// Anthem stuff
	if state.AnthemPlaying {
		raylib.UpdateMusicStream(system.Anthems[state.CurrentAnthem-1])
	}

	// Ideology stuff
	state.EnableCommunism = false
	state.EnableDemocracy = false
	state.EnableNationalism = false

	switch state.LastIdeology {
	case "communism":
		state.EnableCommunism = true

	case "nationalism":
		state.EnableNationalism = true

	case "democracy":
		state.EnableDemocracy = true
	}

	// Country button stuff
	if state.PlayButtonPressed {
		state.AnthemPlaying = false
		system.CurrentFlag = state.Flag
		raylib.StopMusicStream(system.Anthems[state.CurrentAnthem-1])
		SetState("game")
	}
}

func (state *NationCreatorState) Draw() {
	// Drawing the background
	backAnim.Draw()

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

	// Draw the state title
	raylib.DrawTextEx(system.FontKremlin,
		"Nation creator",
		raylib.NewVector2(float32((raylib.GetScreenWidth()-raylib.MeasureText("Nation creator", 96))/2), 5),
		96,
		2,
		raylib.Yellow)

	// Draw buttons

	// Back button
	state.BackButtonPressed = raygui.Button(raylib.NewRectangle(5,
		float32(raylib.GetScreenHeight()-47),
		200,
		45),
		"Back")

	// Play button !
	state.PlayButtonPressed = raygui.Button(raylib.NewRectangle(float32(raylib.GetScreenWidth()-205),
		float32(raylib.GetScreenHeight()-47),
		200,
		45),
		"Play")

	// Draw the flag stuff
	raylib.DrawTextEx(system.FontKremlin,
		"Flag editor :",
		raylib.NewVector2(20, 110),
		20,
		2,
		raylib.Yellow)

	raylib.DrawRectangle(20,
		143,
		int32(float32(state.Flag.Width)*0.35+4),
		int32(float32(state.Flag.Height)*0.35+4),
		raylib.Black) // Border

	raylib.DrawTextureEx(state.Flag, raylib.NewVector2(22, 145), 0.0, 0.35, raylib.White)

	// Vertical of Horizontal ?
	state.FlagDirChoice = raygui.ToggleGroup(raylib.NewRectangle(22, 147+float32(float32(state.Flag.Height)*0.36), 80, 25),
		[]string{"Vertical", "Horizontal"},
		state.FlagDirChoice)

	// Generate ?
	state.GenNewFlag = raygui.Button(raylib.NewRectangle(22,
		350,
		180,
		45),
		"Generate")

	// Country stuff

	// Draw the category title
	raylib.DrawTextEx(system.FontKremlin,
		"Country settings :",
		raylib.NewVector2(400, 110),
		20,
		2,
		raylib.Yellow)

	// Get country name
	raylib.DrawTextEx(system.FontKremlin,
		"Country name :",
		raylib.NewVector2(400, 145),
		20,
		2,
		raylib.White)

	state.CountryName = raygui.TextBox(raylib.NewRectangle(400, 170, 350, 32), state.CountryName)

	// Get the leader name
	raylib.DrawTextEx(system.FontKremlin,
		"Leader name :",
		raylib.NewVector2(400, 207),
		20,
		2,
		raylib.White)

	state.CountryLeader = raygui.TextBox(raylib.NewRectangle(400, 232, 350, 32), state.CountryLeader)

	// Get the country regime
	currentIdeology := state.LastIdeology
	if raygui.CheckBox(raylib.NewRectangle(550, 269, 20, 20), state.EnableCommunism) {
		currentIdeology = "communism"
		state.EnableDemocracy = false
		state.EnableNationalism = false
	}
	if raygui.CheckBox(raylib.NewRectangle(550, 292, 20, 20), state.EnableNationalism) {
		currentIdeology = "nationalism"
		state.EnableCommunism = false
		state.EnableDemocracy = false
	}
	if raygui.CheckBox(raylib.NewRectangle(550, 316, 20, 20), state.EnableDemocracy) {
		currentIdeology = "democracy"
		state.EnableCommunism = false
		state.EnableNationalism = false
	}

	if state.LastIdeology != currentIdeology {
		state.LastIdeology = currentIdeology
		state.GenNewFlag = true
	}

	raylib.DrawTextEx(system.FontKremlin,
		"Communism",
		raylib.NewVector2(400, 269),
		20,
		2,
		raylib.White)

	raylib.DrawTextEx(system.FontKremlin,
		"Nationalism",
		raylib.NewVector2(400, 292),
		20,
		2,
		raylib.White)

	raylib.DrawTextEx(system.FontKremlin,
		"Democracy",
		raylib.NewVector2(400, 316),
		20,
		2,
		raylib.White)

	// Add the country size selector
	raylib.DrawTextEx(system.FontKremlin,
		"Country size :",
		raylib.NewVector2(400, 346),
		20,
		2,
		raylib.White)

	state.CountrySize = raygui.ToggleGroup(raylib.NewRectangle(400, 370, 170, 32),
		[]string{"Small", "Vast"},
		state.CountrySize)

	// Add the national anthem spinner
	lastAnthem := state.CurrentAnthem
	state.CurrentAnthem = raygui.Spinner(raylib.NewRectangle(400, 508, 350, 32), state.CurrentAnthem, 1, len(system.Anthems))

	if lastAnthem != state.CurrentAnthem {
		raylib.StopMusicStream(system.Anthems[lastAnthem-1])
	}

	raylib.DrawTextEx(system.FontKremlin,
		"Country Anthem :",
		raylib.NewVector2(400, 484),
		20,
		2,
		raylib.White)

	if raygui.Button(raylib.NewRectangle(400,
		550,
		128,
		45),
		"Play") {
		state.AnthemPlaying = true
		raylib.PlayMusicStream(system.Anthems[state.CurrentAnthem-1])
	}

	if raygui.Button(raylib.NewRectangle(622,
		550,
		128,
		45),
		"Stop") {
		state.AnthemPlaying = false
		raylib.StopMusicStream(system.Anthems[state.CurrentAnthem-1])
	}

	// Leader stuff
	raylib.DrawRectangle(900,
		145,
		350,
		390,
		raylib.Black) // Back

	raylib.DrawTexture(state.Leaders[state.CurrentLeader-1], 902, 147, raylib.White) // Leader

	state.CurrentLeader = raygui.Spinner(raylib.NewRectangle(900, 550, 350, 32), state.CurrentLeader, 0, 100)
}

func (state *NationCreatorState) Reset() {
	for _, music := range system.Anthems {
		raylib.StopMusicStream(music)
	}

	state.BackButtonPressed = false

	state.FlagDirChoice = 0

	state.GenNewFlag = false

	state.CurrentLeader = 1

	state.CountryName = ""

	state.CountryLeader = ""

	state.LastIdeology = ""

	state.CurrentAnthem = 1

	state.AnthemPlaying = false

	state.EnableCommunism = true
	state.EnableDemocracy = false
	state.EnableNationalism = false

	state.CountrySize = 0

	state.PlayButtonPressed = false
}

func (state *NationCreatorState) Close() {

}
