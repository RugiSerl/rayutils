package graphic

// Anchor help to attach buttons, rect, or texture to a certain part of the screen.
// Anchor comes very handy when the window is resized.

// Object representing a set of two anchor: one horizontal and vertical.
// Those anchor types are defined as enum.
// if X or Y are mixed, the resulting transormation will ignore the anchor.
type Anchor struct {
	X, Y int
}

// Type of anchor
const (
	//horizontal anchor

	ANCHOR_LEFT = iota
	ANCHOR_RIGHT
	ANCHOR_HORIZONTAL_MiDDLE

	//vertical anchor

	ANCHOR_TOP
	ANCHOR_BOTTOM
	ANCHOR_VERTICAL_MiDDLE
)

// Get the position of a rect given its set of anchors.
// Position and size : define the position and size of the rectangle.
// bounds : the rectangle which define the bounds where the rect is contained. Often the rect of the window
func GetRectCoordinatesWithAnchor(position Vector2, size Vector2, anchor Anchor, bounds Rect) Vector2 {
	var DestVector Vector2

	if anchor.X == ANCHOR_LEFT {
		DestVector.X = position.X + bounds.X
	} else if anchor.X == ANCHOR_RIGHT {
		DestVector.X = bounds.Width + bounds.X - position.X - size.X
	} else if anchor.X == ANCHOR_HORIZONTAL_MiDDLE {
		DestVector.X = bounds.X + bounds.Width/2 - position.X - size.X/2
	}

	if anchor.Y == ANCHOR_TOP {
		DestVector.Y = position.Y + bounds.Y
	} else if anchor.Y == ANCHOR_BOTTOM {
		DestVector.Y = bounds.Height + bounds.Y - position.Y - size.Y
	} else if anchor.Y == ANCHOR_VERTICAL_MiDDLE {
		DestVector.Y = bounds.Y + bounds.Height/2 - position.Y - size.Y/2
	}
	return DestVector

}
