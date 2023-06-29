package graphic

// anchor help to attach buttons, rect, or texture to a certain part of the screen.
//

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

// Get the position of a rect ()
func GetRectCoordinatesWithAnchor(position Vector2, anchor Anchor, size Vector2, surfaceRect Rect) Vector2 {
	var DestVector Vector2

	if anchor.X == ANCHOR_LEFT {
		DestVector.X = position.X + surfaceRect.X
	} else if anchor.X == ANCHOR_RIGHT {
		DestVector.X = surfaceRect.Width + surfaceRect.X - position.X - size.X
	} else if anchor.X == ANCHOR_HORIZONTAL_MiDDLE {
		DestVector.X = surfaceRect.X + surfaceRect.Width/2 - position.X - size.X/2
	}

	if anchor.Y == ANCHOR_TOP {
		DestVector.Y = position.Y + surfaceRect.Y
	} else if anchor.Y == ANCHOR_BOTTOM {
		DestVector.Y = surfaceRect.Height + surfaceRect.Y - position.Y - size.Y
	} else if anchor.Y == ANCHOR_VERTICAL_MiDDLE {
		DestVector.Y = surfaceRect.Y + surfaceRect.Height/2 - position.Y - size.Y/2
	}
	return DestVector

}
