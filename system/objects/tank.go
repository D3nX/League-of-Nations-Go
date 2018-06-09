package objects

import (
	"math"

	"github.com/gen2brain/raylib-go/raygui"

	"github.com/gen2brain/raylib-go/raylib"
)

var TankTextures []raylib.Texture2D

type Tank struct {
	Type int

	X         float32
	Y         float32
	Angle     float32
	AngleToGo float32

	Direction string

	Selected bool

	ButtonRectangles map[string]*raylib.Rectangle

	Rectangle raylib.Rectangle
}

func (t *Tank) Update() {
	if raylib.IsMouseButtonPressed(raylib.MouseLeftButton) {

		mPos := raylib.GetMousePosition()

		if mPos.X >= t.Rectangle.X && mPos.X <= t.Rectangle.X+t.Rectangle.Width {
			if mPos.Y >= t.Rectangle.Y && mPos.Y <= t.Rectangle.Y+t.Rectangle.Height {
				if t.Selected {
					t.Selected = false
				} else {
					t.Selected = true
				}
				goto end
			}
		}

		if t.Selected {
			for _, c := range t.ButtonRectangles {
				if mPos.X >= c.X && mPos.X <= c.X+c.Width {
					if mPos.Y >= c.Y && mPos.Y <= c.Y+c.Height {
						goto end
					}
				}
			}

			t.Selected = false
		}

	end:
	}

	// Angle manipulations
	if int(t.AngleToGo) != -1 {
		if t.Angle < t.AngleToGo {
			t.Angle += 5.0

			if t.Angle > t.AngleToGo {
				t.Angle = float32(int(t.AngleToGo) % 360)
				t.AngleToGo = -1.0
			}
		} else if t.Angle > t.AngleToGo {
			t.Angle -= 5.0

			if t.Angle < t.AngleToGo {
				t.Angle = float32(int(t.AngleToGo) % 360)
				t.AngleToGo = -1.0
			}
		} else {
			if t.Angle >= 0.0 {
				t.Angle = t.AngleToGo
				if t.Angle == 360 {
					t.Angle = 0
				}
			} else {
				t.Angle = 360 - float32(math.Abs(float64(t.AngleToGo)))
			}
			t.AngleToGo = -1.0
		}
	}

	// Move if needed
	if int(t.AngleToGo) == -1 {
		switch t.Direction {
		case "right":
			t.X += 1

		case "left":
			t.X -= 1

		case "up":
			t.Y -= 1

		case "down":
			t.Y += 1
		}
	}

	// Update the collision rectangle position
	t.UpdateRectanglePosition()
}

func (t *Tank) Draw(cam *raylib.Camera2D) {
	/*raylib.DrawTextureEx(TankTextures[t.Type],
	raylib.NewVector2(t.X, t.Y),
	t.Angle,
	1.0,
	raylib.White)*/

	if t.Selected {
		// Begin rendering camera
		raylib.BeginMode2D(*cam)

		// Draw Texture
		raylib.DrawTexturePro(TankTextures[t.Type],
			raylib.NewRectangle(0, 0, float32(TankTextures[t.Type].Width), float32(TankTextures[t.Type].Height)),
			raylib.NewRectangle(t.X, t.Y, float32(TankTextures[t.Type].Width), float32(TankTextures[t.Type].Height)),
			raylib.NewVector2(float32(TankTextures[t.Type].Width)/2.0, float32(TankTextures[t.Type].Height)/2.0),
			t.Angle,
			raylib.NewColor(255, 24, 24, 255))

		// End rendering camera
		raylib.EndMode2D()

		// Draw buttons

		// RIGHT button
		t.ButtonRectangles["right"].X = (float32(float32(raylib.GetScreenWidth()/2) + t.Rectangle.Width + 20)) - t.Rectangle.Width/2
		t.ButtonRectangles["right"].Y = (float32(raylib.GetScreenHeight()/2) + (t.Rectangle.Height-t.ButtonRectangles["right"].Height)/2) - t.Rectangle.Height/2
		if raygui.Button(*t.ButtonRectangles["right"], ">") {
			if int(t.Angle) != 270 {
				if t.Angle > 90.0 {
					t.AngleToGo = 270
				} else {
					t.AngleToGo = -90.0
				}
			}

			t.Direction = "right"
		}

		// LEFT button
		t.ButtonRectangles["left"].X = (float32(raylib.GetScreenWidth()/2) - t.ButtonRectangles["left"].Width*2) - t.Rectangle.Width/2
		t.ButtonRectangles["left"].Y = (float32(raylib.GetScreenHeight()/2) + (t.Rectangle.Height-t.ButtonRectangles["left"].Height)/2) - t.Rectangle.Height/2
		if raygui.Button(*t.ButtonRectangles["left"], "<") {
			if int(t.Angle) != 90 {
				t.AngleToGo = 90.0
			}

			t.Direction = "left"
		}

		// UP button
		t.ButtonRectangles["up"].X = (float32(raylib.GetScreenWidth()/2) + (t.Rectangle.Width-t.ButtonRectangles["up"].Width)/2) - t.Rectangle.Width/2
		t.ButtonRectangles["up"].Y = (float32(raylib.GetScreenHeight()/2) - t.ButtonRectangles["up"].Height*2) - t.Rectangle.Height/2
		if raygui.Button(*t.ButtonRectangles["up"], "^") {
			if int(t.Angle) != 180 {
				t.AngleToGo = 180.0
			}

			t.Direction = "up"
		}

		// DOWN button
		t.ButtonRectangles["down"].X = (float32(raylib.GetScreenWidth()/2) + (t.Rectangle.Width-t.ButtonRectangles["up"].Width)/2) - t.Rectangle.Width/2
		t.ButtonRectangles["down"].Y = (float32(raylib.GetScreenHeight()/2) + t.Rectangle.Height + 44) - t.Rectangle.Height/2
		if raygui.Button(*t.ButtonRectangles["down"], "v") {
			if int(t.Angle) != 360 {
				if t.Angle > 180 {
					t.AngleToGo = 360.0
				} else {
					t.AngleToGo = 0.0
				}
			}
			t.Direction = "down"
		}
	} else {
		// Begin rendering camera
		raylib.BeginMode2D(*cam)

		// Draw texture
		raylib.DrawTexturePro(TankTextures[t.Type],
			raylib.NewRectangle(0, 0, float32(TankTextures[t.Type].Width), float32(TankTextures[t.Type].Height)),
			raylib.NewRectangle(t.X, t.Y, float32(TankTextures[t.Type].Width), float32(TankTextures[t.Type].Height)),
			raylib.NewVector2(float32(TankTextures[t.Type].Width)/2.0, float32(TankTextures[t.Type].Height)/2.0),
			t.Angle,
			raylib.White)

		// End rendering camera
		raylib.EndMode2D()
	}

	// raylib.DrawRectangleRec(t.Rectangle, raylib.White)
}

func (t *Tank) UpdateRectanglePosition() {
	switch int(t.Angle) {
	case 0:
		fallthrough
	case 180:
		fallthrough
	case -180:
		t.Rectangle.Width = float32(TankTextures[t.Type].Width)
		t.Rectangle.Height = float32(TankTextures[t.Type].Height)
		t.Rectangle.X = t.X - float32(TankTextures[t.Type].Width)/2
		t.Rectangle.Y = t.Y - float32(TankTextures[t.Type].Height)/2
	case 90:
		fallthrough
	case -90:
		fallthrough
	case 270:
		t.Rectangle.Width = float32(TankTextures[t.Type].Height)
		t.Rectangle.Height = float32(TankTextures[t.Type].Width)
		t.Rectangle.X = t.X - float32(TankTextures[t.Type].Height)/2
		t.Rectangle.Y = t.Y - float32(TankTextures[t.Type].Width)/2
	}
}

func (t *Tank) IsSelected() bool {
	return t.Selected
}

func (t *Tank) GetPosition() raylib.Vector2 {
	return raylib.NewVector2(t.X, t.Y)
}
