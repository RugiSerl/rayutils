package graphic

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// A rect is a set of 4 numbers, which represent its position and its size.
type Rect struct {
	X      float32
	Y      float32
	Width  float32
	Height float32
}

// Returns a new rect.
func NewRect(x float32, y float32, width float32, height float32) Rect {
	r := Rect{}
	r.X, r.Y, r.Width, r.Height = x, y, width, height
	return r
}

// Returns a new rect, using vectors.
func NewRectFromVector(position Vector2, size Vector2) Rect {
	return Rect{X: position.X, Y: position.Y, Width: size.X, Height: size.Y}
}

// Get the position of the rect.
func (r Rect) GetPosition() Vector2 {
	return NewVector2(r.X, r.Y)
}

// Get the size of the rect.
func (r Rect) GetSize() Vector2 {
	return NewVector2(r.Width, r.Height)
}

// Get the rect representing the window, with position (0, 0) and size (wWidth, wHeight).
func GetWindowRect() Rect {
	return NewRect(0, 0, float32(rl.GetScreenWidth()), float32(rl.GetScreenHeight()))
}

// Get the rect representing the mouse, with position (mX, mY) and size (1, 1).
func GetMouseRect() Rect {
	return NewRect(rl.GetMousePosition().X, rl.GetMousePosition().Y, 1, 1)
}

// Get the rect inside another, with a certain padding.
func GetInnerRect(sourceRect Rect, padding float32) Rect {
	sourceRect.X += padding
	sourceRect.Y += padding
	sourceRect.Width -= padding * 2
	sourceRect.Height -= padding * 2
	return sourceRect
}

// GetInnerRect() but only horizontally.
func GetInnerHorizontalrect(sourceRect Rect, padding float32) Rect {
	sourceRect.Y += padding
	sourceRect.Height -= padding * 2

	return sourceRect
}

// Get the position of the center of the rect.
func (r Rect) GetCenter() Vector2 {
	return r.GetPosition().Add(r.GetSize().Scale(0.5))
}

//---------------------
// Collision methods

// Detect if two rectangles overlap.
func DetectRectCollision(rect1 Rect, rect2 Rect) bool {
	return (rect1.X+rect1.Width >= rect2.X && rect1.X <= rect2.X+rect2.Width && rect1.Y+rect1.Height >= rect2.Y && rect1.Y <= rect2.Y+rect2.Height)
}

// Detect if a point is inside the area of the rectangle.
func DetectPointRectCollision(point Vector2, rect Rect) bool {
	return DetectRectCollision(rect, NewRectFromVector(point, NewVector2(1, 1)))
}

//---------------------
// Draw methods

// Draw a filled rectangle with a color using the coordinates of the rectangle.
func (r Rect) Fill(color color.RGBA, roundness float32) {
	rectangle := rl.NewRectangle(r.X, r.Y, r.Width, r.Height)

	rl.DrawRectangleRounded(rectangle, roundness, 5, color)
}

// Draw the lines of the rectangle.
func (r *Rect) DrawLines(color color.RGBA, roundness float32, thickness float32) {
	rectangle := rl.NewRectangle(r.X, r.Y, r.Width, r.Height)

	rl.DrawRectangleRoundedLines(rectangle, roundness, 5, thickness, color)
}
