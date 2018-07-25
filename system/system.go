package system

import (
	"fmt"

	"github.com/gen2brain/raylib-go/raygui"
	"github.com/gen2brain/raylib-go/raylib"
)

// The settings structure
type Settings struct {
	MusicVolume float32
	SFXVolume   float32
}

// Global variables
const Title = "League of Nations"

var Closed bool
var FontKremlin raylib.Font
var GameSettings Settings
var Leaders []*raylib.Image
var Logos []*raylib.Image
var Anthems []raylib.Music
var Screenshot raylib.Texture2D

// Global and current game variables
var CurrentFlag raylib.Texture2D

// Global functions
func Initialize() {
	Closed = false
	FontKremlin = raylib.LoadFontEx("res/fonts/kremlin.ttf", 96, 0, nil)
	GameSettings = Settings{
		MusicVolume: 100.0,
		SFXVolume:   100.0,
	}

	Leaders = make([]*raylib.Image, 5)
	for i := 0; i < len(Leaders); i++ {
		Leaders[i] = raylib.LoadImage(fmt.Sprint("res/leaders/leader_", i+1, ".png"))
	}

	// Screenshot = nil

	// Logo stuff
	Logos = make([]*raylib.Image, 7)

	// Fascism / Nationalism
	Logos[0] = raylib.LoadImage("res/logos/nationalism.png")
	Logos[1] = raylib.LoadImage("res/logos/fascism.png")
	Logos[2] = raylib.LoadImage("res/logos/imperialism.png")

	// Communsim / Socialism
	Logos[3] = raylib.LoadImage("res/logos/communism.png")
	Logos[4] = raylib.LoadImage("res/logos/socialism.png")
	Logos[5] = raylib.LoadImage("res/logos/red_star.png")

	// Democracy
	Logos[6] = raylib.LoadImage("res/logos/star.png")

	// Anthems
	Anthems = make([]raylib.Music, 2)
	Anthems[0] = raylib.LoadMusicStream("res/musics/anthems/ussr_anthem.ogg")
	Anthems[1] = raylib.LoadMusicStream("res/musics/anthems/ddr_anthem.ogg")

	// Load style
	raygui.LoadGuiStyle("res/gui.style")
}

func Close() {
	raylib.UnloadFont(FontKremlin)

	for i := 0; i < len(Leaders); i++ {
		raylib.UnloadImage(Leaders[i])
	}

	for i := 0; i < len(Logos); i++ {
		raylib.UnloadImage(Logos[i])
	}
}

func Log(input string) {
	fmt.Println("[League of Nations][Build 01][Log]:", input)
}

func LogError(input string) {
	fmt.Println("[League of Nations][Build 01][ERROR]:", input)
}
