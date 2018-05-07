package gamemap

import (
	"io/ioutil"
	"strings"

	"github.com/gen2brain/raylib-go/raylib"
)

var Tileset raylib.Texture2D

func Initialize() {
	Tileset = raylib.LoadTexture("res/sprites/grass_tileset.png")
}

func Close() {
	raylib.UnloadTexture(Tileset)
}

type GameMap struct {
	Tiles  []Tile
	Width  int
	Height int
}

func (gm *GameMap) Load(path string) {

	f, err := ioutil.ReadFile("res/map/map.txt")

	if err != nil {
		panic(err)
	}

	content := string(f)

	var x, y float32 = 0, 0

	for _, line := range strings.Split(content, "\n") {
		for _, char := range line {
			gm.Tiles = append(gm.Tiles, Tile{X: x, Y: y, Type: char})

			x += 32
		}
		y += 32
		x = 0
	}

}

func (gm *GameMap) Draw() {
	for _, tile := range gm.Tiles {
		tile.Draw()
	}
}
