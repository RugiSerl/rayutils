package curves

import (
	"github.com/RugiSerl/rayutils/graphic"
)

// Linear interpolation between two points using normalized t.
func LinearInterpolation(t float32, controlPoint1 graphic.Vector2, controlPoint2 graphic.Vector2) graphic.Vector2 {
	if t > 1 || t < 0 {
		panic("t must be between 0 and 1")
	}

	return controlPoint2.Scale(t).Add(controlPoint1.Scale(1 - t))
}
