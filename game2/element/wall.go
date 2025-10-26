package element

import (
	"image/color"

	"pl.home/game2/conf"
)

type Wall struct {
	X, Y          float32
	Width, Height float32
	Color         color.Color
}

func NewWall() *Wall {
	return &Wall{
		X:      conf.WallX,
		Y:      conf.WallY,
		Width:  conf.WallWidth,
		Height: conf.WallHeigh,
		Color:  conf.WhiteColor,
	}
}

func (w *Wall) UpdateDirect(x, y float32) bool {
	if x <= conf.WallX+3 && x >= conf.WallX-3 &&
		y >= conf.WallY && y <= conf.WallHeigh+conf.WallY {
		return true
	}
	return false
}
