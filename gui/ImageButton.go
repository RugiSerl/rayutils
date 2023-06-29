package gui

import (
	"github.com/RugiSerl/rayutils/graphic"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// l'objet ImageButton désigne un bouton qui a pour hitbox et pour visuel une texture
type ImageButton struct {
	position     graphic.Vector2
	size         graphic.Vector2
	img          rl.Texture2D
	HoverState   bool
	PressedState bool
	anchor       graphic.Anchor
}

var (
	imageButtonPhysicalPosition graphic.Vector2
)

// Initialisation du bouton
func NewImageButton(position graphic.Vector2, texture rl.Texture2D, anchor graphic.Anchor) *ImageButton {
	b := new(ImageButton)

	b.img = texture

	b.position = position
	b.size = graphic.NewVector2(float32(b.img.Width), float32(b.img.Height))

	b.anchor = anchor

	b.HoverState = false
	b.PressedState = false

	return b
}

// Fonction de mise à jour du bouton
func (b *ImageButton) Update(containingRect graphic.Rect) {
	imageButtonPhysicalPosition = graphic.GetRectCoordinatesWithAnchor(b.position, b.anchor, b.size.Scale(interfaceScale), containingRect)
	b.handleInput()
	b.render()
}

// Fonction qui permet de gérer les inputs du bouton
func (b *ImageButton) handleInput() {
	b.HoverState, b.PressedState = false, false
	if graphic.DetectRectCollision(graphic.GetMouseRect(), graphic.NewRectFromVector(imageButtonPhysicalPosition, b.size.Scale(interfaceScale))) {
		b.HoverState = true
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			b.PressedState = true
		}
	}

}

// Fonction d'affichage du bouton
func (b *ImageButton) render() {

	if b.HoverState {
		rl.DrawTextureEx(b.img, rl.Vector2(imageButtonPhysicalPosition), 0, interfaceScale, rl.White)
	} else {
		rl.DrawTextureEx(b.img, rl.Vector2(imageButtonPhysicalPosition), 0, interfaceScale, rl.NewColor(255, 255, 255, 120))
	}

}
