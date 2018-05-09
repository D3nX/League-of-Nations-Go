package objects

import (
	"github.com/gen2brain/raylib-go/raylib"
)

func Initialize() {
	BuildingTextures = make([]raylib.Texture2D, 1)

	BuildingTextures[0] = raylib.LoadTexture("res/sprites/building_0.png")

	TankTextures = make([]raylib.Texture2D, 1)

	TankTextures[0] = raylib.LoadTexture("res/sprites/tank_0.png")
}

func Close() {
	for i := range BuildingTextures {
		raylib.UnloadTexture(BuildingTextures[i])
	}
}

func NewBuilding(buildingType int, x, y float32) *Building {
	building := &Building{}

	building.X = x
	building.Y = y
	building.Type = buildingType

	return building
}

func NewTank(tankType int, x, y, angle float32) *Tank {
	tank := &Tank{}

	tank.X = x
	tank.Y = y
	tank.Angle = angle
	tank.Type = tankType

	return tank
}
