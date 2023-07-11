package graphic

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func DrawAdjustedTexture(image rl.Texture2D) {
	//draw image ajusted to the window

	var imagex, imagey, scale, imageW, imageH float32

	winW, winH := float32(rl.GetScreenWidth()), float32(rl.GetScreenHeight())
	winRatio := winW / winH

	imgW, imgH := float32(image.Width), float32(image.Height)
	imgRatio := imgW / imgH

	if imgRatio > winRatio { //change Height
		imageW = winW
		imageH = imageW / imgRatio

		imagex = 0
		imagey = winH/2 - imageH/2

		scale = imageH / imgH

	} else if imgRatio < winRatio {
		imageH = winH
		imageW = imageH * imgRatio

		imagey = 0
		imagex = winW/2 - imageW/2

		scale = imageW / imgW

	} else {
		imagex = 0
		imagey = 0
		scale = winW / imgW
	}

	rl.DrawTextureEx(image, rl.NewVector2(imagex, imagey), 0, scale, rl.White)

}

func DrawTextureFromCenter(texture rl.Texture2D, position Vector2, scale float32, tint color.RGBA) {
	//draw an image from its center

	position = position.Substract(NewVector2(float32(texture.Width)/2, float32(texture.Height)/2).Scale(scale))

	rl.DrawTextureEx(texture, rl.Vector2(position), 0, scale, tint)

}
