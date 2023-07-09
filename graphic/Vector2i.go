package graphic

type Vector2i struct {
	X, Y int
}

func NewVector2i(x, y int) Vector2i {
	return Vector2i{x, y}
}

func (v Vector2i) ToVector2() Vector2 {
	return NewVector2(float32(v.X), float32(v.Y))
}
