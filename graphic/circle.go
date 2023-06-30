package graphic

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// object representing a circle
type Circle struct {
	CenterPosition Vector2
	Radius         float32
}

// Return a new circle.
func NewCircle(radius float32, x float32, y float32) Circle {
	return Circle{Radius: radius, CenterPosition: NewVector2(x, y)}

}

//---------------------
// Collision methods

// Detect if the two circles are overlapping.
func (c *Circle) DetectCircleCollision(otherCircle Circle) bool {

	return (c.CenterPosition.Substract(otherCircle.CenterPosition).GetNorm() <= c.Radius+otherCircle.Radius)
}

// Returns if the position of the mouse is inside the area of the circle.
func (c *Circle) DetectMouseCollision() bool {
	return c.DetectPointCollision(Vector2(rl.GetMousePosition()))
}

// Returns if a point is inside the area of the circle.
func (c *Circle) DetectPointCollision(position Vector2) bool {
	return (position.Substract(c.CenterPosition).GetNorm() <= c.Radius)
}

// Returns if the circle and a given rectangle are overlapping.
func (c *Circle) DetectRectCollision(rect Rect) bool {

	// NOTE: instead of detecting if the circle and the rect are overlapping, the rect is made larger by the radius of the circle.
	// This method is simple but not perfect, especially in the corners
	// see answer : https://stackoverflow.com/questions/401847/circle-rectangle-collision-detection-intersection
	// might improve it if I'm not lazy
	rect = GetInnerRect(rect, -c.Radius)

	return DetectPointRectCollision(c.CenterPosition, rect)

}

//---------------------
// Draw methods

// Draw a circle filled with a certain color, using the coordinates of the circle.
func (c *Circle) Fill(color color.RGBA) {
	rl.DrawCircleV(rl.Vector2(c.CenterPosition), c.Radius, color)

}

// Draw the border line of the cirle. For some reason it also draw an additionnal line, whose angle can be changed by lineAngle (in degrees).
func (c *Circle) DrawLines(color color.RGBA, LineAngle float32) {
	//rl.DrawCircleLines(int32(c.CenterPosition.X), int32(c.CenterPosition.Y), c.Radius, color) //problème: convertit les coordonnées en entier
	rl.DrawCircleSectorLines(rl.Vector2(c.CenterPosition), c.Radius, LineAngle+90, LineAngle+450, 100, color) //problème: laisse un trait de couleur entre le centre et le cercle
}

// draw the border line of the circle, and a cross inside
func (c *Circle) DrawCross(color color.RGBA) {
	c.DrawLines(color, 45)
	c.DrawLines(color, 135)
	c.DrawLines(color, 225)
	c.DrawLines(color, 315)

}
