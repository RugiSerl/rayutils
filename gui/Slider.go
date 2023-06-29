package gui

// Par Raphaël

import (
	"github.com/RugiSerl/rayutils/graphic"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// le slider, ou "curseur" est une barre verticale qui permet d'ajuster une valeur en déplaçant le rond qui est devant
type Slider struct {
	value *float32
	min   float32
	max   float32

	HoverState bool
	position   graphic.Vector2
	anchor     graphic.Anchor

	size graphic.Vector2
}

// file scope var to get the on-screen position of the object
var sliderRect graphic.Rect

// Initialisation du slider
func NewSlider(position graphic.Vector2, anchor graphic.Anchor) *Slider {

	s := new(Slider)

	s.anchor = anchor

	s.position = position

	s.size = graphic.NewVector2(70, 20)

	return s

}

func (s *Slider) SetValue(value *float32, min float32, max float32) {
	s.value, s.min, s.max = value, min, max

}

func (s *Slider) GetSize() graphic.Vector2 {
	return s.size
}

// Fonction de mise à jour du slider
func (s *Slider) Update(containingRect graphic.Rect) {
	sliderRect = graphic.NewRectFromVector(graphic.GetRectCoordinatesWithAnchor(s.position, s.size, s.anchor, containingRect), s.size)

	s.handleInput()
	s.render()

}

// Fonction permettant de gérer les inputs du slider
func (s *Slider) handleInput() {
	s.HoverState = false
	if graphic.DetectRectCollision(sliderRect, graphic.GetMouseRect()) {
		s.HoverState = true
		if rl.IsMouseButtonDown(rl.MouseLeftButton) {
			*s.value = s.min + (s.max-s.min)*(rl.GetMousePosition().X-sliderRect.X)/sliderRect.Width
		}

	}

}

// Fonction d'affichage du slider
func (s *Slider) render() {

	bar := graphic.GetInnerHorizontalrect(sliderRect, sliderRect.Height/3)

	ballXPosition := (*s.value-s.min)/(s.max-s.min)*sliderRect.Width + sliderRect.X
	ball := graphic.NewCircle(5, ballXPosition, sliderRect.Y+sliderRect.Height/2)
	ball.Fill(rl.Black)

	bar.Fill(rl.Black, 0)

}
