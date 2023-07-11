package graphic

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// A Vector2 is a set of 2 numbers, often representing a 2d coordinate/movement.
type Vector2 struct {
	X float32
	Y float32
}

// Returns a new Vector.
func NewVector2(x float32, y float32) Vector2 {
	return Vector2{X: x, Y: y}
}

func NewVectorFromAngle(angle float64) Vector2 {
	cos := float32(math.Cos(angle))
	sin := float32(math.Sin(angle))
	return NewVector2(cos, sin)
}

// Get the length of the Vector.
func (v Vector2) GetNorm() float32 {
	return float32(math.Sqrt(math.Pow((float64(v.X)), 2) + math.Pow((float64(v.Y)), 2)))
}

// return the addition of the two vectors.
func (v Vector2) Add(otherVector Vector2) Vector2 {
	return NewVector2(v.X+otherVector.X, v.Y+otherVector.Y)
}

// Return the substraction of the two vectors.
func (v Vector2) Substract(otherVector Vector2) Vector2 {
	return NewVector2(v.X-otherVector.X, v.Y-otherVector.Y)
}

// Scale the Vector by multiplying numbers by a certain value.
func (v Vector2) Scale(scale float32) Vector2 {
	return NewVector2(v.X*scale, v.Y*scale)
}

// Scale the vector to get the desired norm. Doesn't work with Vector (0, 0) (division by 0)
func (v Vector2) ScaleToNorm(norm float32) Vector2 {
	return v.Scale(norm / v.GetNorm())
}

// Return the angle in radians between the vector and (1, 0)
func (v Vector2) GetAngle() float64 {
	if v.Y > 0 {
		return math.Acos(float64(v.ScaleToNorm(1).X))
	} else {
		return -math.Acos(float64(v.ScaleToNorm(1).X))
	}
}

// Rotate the Vector to a certain angle. Better way to do this smh (https://matthew-brett.github.io/teaching/rotation_2d.html).
func (v Vector2) Rotate(angle float64) Vector2 {
	return NewVectorFromAngle(v.GetAngle() + angle).ScaleToNorm(v.GetNorm())
}

// Weird function I invented and can't explain. Probably something stupid.
func (v Vector2) FlattenToLine(lineAngle float64) Vector2 {
	return NewVector2(v.Rotate(lineAngle).X, 0).Rotate(-lineAngle)
}

// Conversion to raylib Vector2.
// Useless since it possible to do rl.Vector2() of the object Vector2
func (v Vector2) ToRaylibVector2() rl.Vector2 {
	return rl.NewVector2(v.X, v.Y)
}

// Returns (v1.x*v2.x, v.1.y*v.2.y)
func (v Vector2) Multiply(otherVector Vector2) Vector2 {
	return NewVector2(v.X*otherVector.X, v.Y*otherVector.Y)
}

// vector inverse -> no null vector
func (v Vector2) Inverse() Vector2 {
	return NewVector2(1/v.X, 1/v.Y)
}

// Normalize vector to another. So (v1.x/v2.x, v.1.y/v.2.y)
func (v Vector2) NormalizeToVector(otherVector Vector2) Vector2 {
	return v.Multiply(otherVector.Inverse())
}

// Normalize vector to originRect and multiply it by the size of the destRect
func (v Vector2) NormalizeToRect(originRect Rect, destRect Rect) Vector2 {
	return v.NormalizeToVector(originRect.GetSize()).Multiply(destRect.GetSize()).Add(destRect.GetPosition())
}
