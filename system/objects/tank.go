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
	Speed     float32

	Direction string

	ButtonRectangles map[string]*raylib.Rectangle

	Rectangle raylib.Rectangle

	CanRotateVertical, CanRotateHorizontal bool // Affect the direction
	CanRotate                              bool // Let the rotation or not, but doesn't affect the direction

	Selected      bool
	buttonPressed bool
	canMove       bool
	Rotating      bool
}

func (t *Tank) Update(cam *raylib.Camera2D, pickable bool) {

	if pickable {

		if raylib.IsMouseButtonPressed(raylib.MouseLeftButton) {

			mPos := raylib.GetMousePosition()

			mPos.X -= cam.Offset.X
			mPos.Y -= cam.Offset.Y

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

			mPos = raylib.GetMousePosition()

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

	}

	// Angle manipulations
	t.Rotating = false

	if t.CanRotate {
		if int(t.AngleToGo) != -1 {
			t.Rotating = true
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
	}

	// Move if needed
	if t.canMove {
		if int(t.AngleToGo) == -1 || !t.CanRotate {
			switch t.Direction {
			case "right":
				t.X += t.Speed

			case "left":
				t.X -= t.Speed

			case "up":
				t.Y -= t.Speed

			case "down":
				t.Y += t.Speed
			}
		}
	}

	// Update the collision rectangle position
	t.UpdateRectanglePosition()

	// Enable moving
	t.canMove = true
}

func (t *Tank) Draw(cam *raylib.Camera2D) {
	/*raylib.DrawTextureEx(TankTextures[t.Type],
	raylib.NewVector2(t.X, t.Y),
	t.Angle,
	1.0,
	raylib.White)*/

	t.buttonPressed = false

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
			t.buttonPressed = true

			if t.CanRotateHorizontal {
				if int(t.Angle) != 270 {
					if t.Angle > 90.0 {
						t.AngleToGo = 270
					} else {
						t.AngleToGo = -90.0
					}
				}

				t.X += 1
				t.Direction = "right"
			}
		}

		// LEFT button
		t.ButtonRectangles["left"].X = (float32(raylib.GetScreenWidth()/2) - t.ButtonRectangles["left"].Width*2) - t.Rectangle.Width/2
		t.ButtonRectangles["left"].Y = (float32(raylib.GetScreenHeight()/2) + (t.Rectangle.Height-t.ButtonRectangles["left"].Height)/2) - t.Rectangle.Height/2
		if raygui.Button(*t.ButtonRectangles["left"], "<") {
			t.buttonPressed = true

			if t.CanRotateHorizontal {
				if int(t.Angle) != 90 {
					t.AngleToGo = 90.0
				}

				t.X -= 1
				t.Direction = "left"
			}
		}

		// UP button
		t.ButtonRectangles["up"].X = (float32(raylib.GetScreenWidth()/2) + (t.Rectangle.Width-t.ButtonRectangles["up"].Width)/2) - t.Rectangle.Width/2
		t.ButtonRectangles["up"].Y = (float32(raylib.GetScreenHeight()/2) - t.ButtonRectangles["up"].Height*2) - t.Rectangle.Height/2 + 20
		if raygui.Button(*t.ButtonRectangles["up"], "^") {
			t.buttonPressed = true

			if t.CanRotateVertical {
				if int(t.Angle) != 180 {
					t.AngleToGo = 180.0
				}

				t.Y -= 1
				t.Direction = "up"
			}
		}

		// DOWN button
		t.ButtonRectangles["down"].X = (float32(raylib.GetScreenWidth()/2) + (t.Rectangle.Width-t.ButtonRectangles["up"].Width)/2) - t.Rectangle.Width/2
		t.ButtonRectangles["down"].Y = (float32(raylib.GetScreenHeight()/2) + t.Rectangle.Height + 44) - t.Rectangle.Height/2 - 35
		if raygui.Button(*t.ButtonRectangles["down"], "v") {
			t.buttonPressed = true

			if t.CanRotateVertical {
				if int(t.Angle) != 360 {
					if t.Angle > 180 {
						t.AngleToGo = 360.0
					} else {
						t.AngleToGo = 0.0
					}
				}

				t.Y += 1
				t.Direction = "down"
			}
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

func (t *Tank) SetSelected(selected bool) {
	t.Selected = selected
}

func (t *Tank) GetPosition() raylib.Vector2 {
	return raylib.NewVector2(t.X, t.Y)
}

func (t *Tank) Collides(rect raylib.Rectangle) bool {

	if raylib.CheckCollisionRecs(rect, t.Rectangle) {
		return true
	}

	return false

}

func (t *Tank) StopMoving(direction string) {
	/*if t.Direction == direction {
		t.Direction = ""
	}*/
	t.canMove = false
}

func (t Tank) CanMove() bool {
	return t.canMove
}

func TankCollides(t1, t2 *Tank) bool {
	return t1.Collides(raylib.NewRectangle(t2.X-t2.Rectangle.Width/2, t2.Y-t2.Rectangle.Height/2, t2.Rectangle.Width, t2.Rectangle.Height))
}

func TankWillCollides(t1, t2 *Tank, distance float32) bool {
	switch t2.Direction {
	case "right":
		return t1.Collides(raylib.NewRectangle(t2.X-t2.Rectangle.Width/2+distance, t2.Y-t2.Rectangle.Height/2, t2.Rectangle.Width, t2.Rectangle.Height))

	case "left":
		return t1.Collides(raylib.NewRectangle(t2.X-t2.Rectangle.Width/2-distance, t2.Y-t2.Rectangle.Height/2, t2.Rectangle.Width, t2.Rectangle.Height))

	case "up":
		return t1.Collides(raylib.NewRectangle(t2.X-t2.Rectangle.Width/2, t2.Y-t2.Rectangle.Height/2-distance, t2.Rectangle.Width, t2.Rectangle.Height))

	case "down":
		return t1.Collides(raylib.NewRectangle(t2.X-t2.Rectangle.Width/2, t2.Y-t2.Rectangle.Height/2+distance, t2.Rectangle.Width, t2.Rectangle.Height))
	}

	return false
}

// Check the collision with tank 1 & tank 2
// if the tank is rotated
func TankCollidesRotation(t1, t2 *Tank) bool {
	var rectangle raylib.Rectangle
	switch int(t1.Angle) {
	case 0:
		fallthrough
	case 180:
		fallthrough
	case -180:
		rectangle.Width = float32(TankTextures[t1.Type].Height)
		rectangle.Height = float32(TankTextures[t1.Type].Width)
		rectangle.X = t1.X - float32(TankTextures[t1.Type].Height)/2
		rectangle.Y = t1.Y - float32(TankTextures[t1.Type].Width)/2
	case 90:
		fallthrough
	case -90:
		fallthrough
	case 270:
		rectangle.Width = float32(TankTextures[t1.Type].Width)
		rectangle.Height = float32(TankTextures[t1.Type].Height)
		rectangle.X = t1.X - float32(TankTextures[t1.Type].Width)/2
		rectangle.Y = t1.Y - float32(TankTextures[t1.Type].Height)/2
	}

	return raylib.CheckCollisionRecs(rectangle, raylib.NewRectangle(t2.X-t2.Rectangle.Width/2, t2.Y-t2.Rectangle.Height/2, t2.Rectangle.Width, t2.Rectangle.Height))
}

func TankCollidesRotationRec(t *Tank, rect raylib.Rectangle) bool {
	var rectangle raylib.Rectangle
	switch int(t.Angle) {
	case 0:
		fallthrough
	case 180:
		fallthrough
	case -180:
		rectangle.Width = float32(TankTextures[t.Type].Height)
		rectangle.Height = float32(TankTextures[t.Type].Width)
		rectangle.X = t.X - float32(TankTextures[t.Type].Height)/2
		rectangle.Y = t.Y - float32(TankTextures[t.Type].Width)/2
	case 90:
		fallthrough
	case -90:
		fallthrough
	case 270:
		rectangle.Width = float32(TankTextures[t.Type].Width)
		rectangle.Height = float32(TankTextures[t.Type].Height)
		rectangle.X = t.X - float32(TankTextures[t.Type].Width)/2
		rectangle.Y = t.Y - float32(TankTextures[t.Type].Height)/2
	}

	return raylib.CheckCollisionRecs(rectangle, rect)
}
