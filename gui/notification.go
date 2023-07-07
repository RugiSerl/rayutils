package gui

import (
	"github.com/RugiSerl/rayutils/graphic"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const NOTIFICATION_DURATION = 1
const NOTIFICATION_FONT_SIZE = 10

var (
	time             float32 = -1
	text             string
	textSize         graphic.Vector2
	font             rl.Font
	notificationInit bool = false
)

func InitNotification() {
	font = rl.LoadFontEx("assets/VarelaRound-Regular.ttf", int32(NOTIFICATION_DURATION), []rune("'é!èabcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ:0123456789.- ()"))
	rl.SetTextureFilter(font.Texture, rl.FilterBilinear)
	notificationInit = true
}

func NewNotificationText(notifText string) {
	text = notifText
	time = 0
	textSize = graphic.Vector2(rl.MeasureTextEx(font, text, NOTIFICATION_FONT_SIZE, 0))

}

func UpdateNotification() {
	if !notificationInit {
		panic("the notification is not initialized")
	}
	time += rl.GetFrameTime()
	if time > 0 && time < NOTIFICATION_DURATION {
		rect := graphic.NewRect(0, 0, textSize.X, textSize.Y)
		rect.Fill(rl.Black, 0.3)
		rl.DrawTextEx(font, text, rl.NewVector2(0, 0), NOTIFICATION_FONT_SIZE, 0, rl.White)
	}

}
