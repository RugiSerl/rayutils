package gui

// Par Raphaël

import (
	"github.com/RugiSerl/rayutils/graphic"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// la CheckBox est une case à cocher
type CheckBox struct {
	value      *bool
	HoverState bool
	position   graphic.Vector2
	anchor     graphic.Anchor
}

var (
	checkBoxRect graphic.Rect
)

const (
	CHECKBOX_SIZE = 100
)

// Initialisation de la checkbox
func NewCheckBox(position graphic.Vector2, anchor graphic.Anchor) *CheckBox {

	c := new(CheckBox)

	c.anchor = anchor

	c.position = position

	return c

}

func (c *CheckBox) SetValue(value *bool) {
	c.value = value

}

// Fonction qui met à jour la checkbox
func (c *CheckBox) Update(containingRect graphic.Rect) {
	checkBoxRect = graphic.NewRectFromVector(graphic.GetRectCoordinatesWithAnchor(c.position, graphic.NewVector2(CHECKBOX_SIZE*interfaceScale, CHECKBOX_SIZE*interfaceScale), c.anchor, containingRect), graphic.NewVector2(CHECKBOX_SIZE*interfaceScale, CHECKBOX_SIZE*interfaceScale))

	c.handleInput()
	c.render()

}

// Fonction permettant de gérer les inputs de la checkbox
func (c *CheckBox) handleInput() {

	c.HoverState = false

	if graphic.DetectRectCollision(checkBoxRect, graphic.GetMouseRect()) {
		c.HoverState = true
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {

			*c.value = !*c.value

		}

	}

}

// Fonction d'affichage de la Checkbox
func (c *CheckBox) render() {

	innerRect := graphic.GetInnerRect(checkBoxRect, 2)

	checkBoxRect.Fill(rl.Black, 0)
	innerRect.Fill(rl.White, 0)

	if *c.value {
		innerRectConfirmation := graphic.GetInnerRect(innerRect, 3)
		innerRectConfirmation.Fill(rl.Black, 0)
	}

}
