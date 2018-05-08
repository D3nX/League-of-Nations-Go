package gamemap

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"../../system"
	"../objects"

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
	Tiles   [][]Tile
	Objects []objects.Object
	Width   int
	Height  int
}

func (gm *GameMap) Load(path string) {

	// Load the file and get the cotent
	f, err := ioutil.ReadFile("res/map/map.txt")

	if err != nil {
		panic(err)
	}

	content := string(f)

	// Now parsing tiles & objects
	var x, y float32 = 0, 0

	parseObject := false

	gm.Tiles = make([][]Tile, 0)

	for _, line := range strings.Split(content, "\n") {
		if line[0] == '#' {
			continue
		}

		if strings.HasPrefix(line, "objects:") {
			parseObject = true
			continue
		}

		if !parseObject {
			for _, char := range line {
				gm.Tiles = append(gm.Tiles, []Tile{})
				gm.Tiles[int(x/32)] = append(gm.Tiles[int(x/32)], Tile{X: x, Y: y, Type: char})

				x += 32
			}

			y += 32
			x = 0
		} else {
			keywords := strings.Split(line, " ")

			if len(keywords) < 4 {
				system.Log(fmt.Sprint("Error: Too less arguments ! (Only ", len(keywords), ")"))
				continue
			}

			switch keywords[0] {
			case "building":
				buildingType, _ := strconv.Atoi(keywords[1])
				x, _ := strconv.Atoi(keywords[2])
				y, _ := strconv.Atoi(keywords[3])

				gm.Objects = append(gm.Objects, objects.NewBuilding(buildingType, float32(x)*32, float32(y)*32))
			}
		}
	}

}

func (gm *GameMap) Update() {
	for _, object := range gm.Objects {
		object.Update()
	}
}

func (gm *GameMap) Draw() {
	for x := range gm.Tiles {
		for y := range gm.Tiles[x] {
			gm.Tiles[x][y].Draw()
		}
	}

	for _, object := range gm.Objects {
		object.Draw()
	}
}
