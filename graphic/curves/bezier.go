package curves

import "github.com/RugiSerl/rayutils/graphic"

// ret
func BezierCurves(t float32, points []graphic.Vector2) graphic.Vector2 {

	if len(points) == 2 {
		return LinearInterpolation(t, points[0], points[1])
	}

	newList := []graphic.Vector2{}
	for i := 0; i+i < len(points); i++ {
		newList = append(newList, LinearInterpolation(t, points[i], points[i+1]))
	}

	return BezierCurves(t, newList)

}
