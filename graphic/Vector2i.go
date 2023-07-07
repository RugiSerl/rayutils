package graphic

type Vector2i struct {
	X, Y int
}

func NewVector2i(x, y int) Vector2i {
	return Vector2i{x, y}
}
