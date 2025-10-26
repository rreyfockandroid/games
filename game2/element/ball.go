package element

import (
	"image/color"

	"pl.home/game2/conf"
)

type Ball struct {
	X, Y                      float32
	DirectRight, DirectBottom bool
	Speed                     float32
	Color                     color.Color
	Stop                      bool
}

func NewBall() *Ball {
	return &Ball{
		X:            conf.StartX,
		Y:            conf.StartY,
		DirectRight:  true,
		DirectBottom: true,
		Color:        conf.StartColor,
		Speed:        conf.BallSpeed,
	}
}

func (b *Ball) Reset() {
	b.X = conf.StartX
	b.Y = conf.StartY
	b.Color = conf.StartColor
	if b.DirectRight {
		b.DirectRight = false
		b.Color = conf.WhiteColor
	} else {
		b.Color = conf.YellowColor
		b.DirectRight = true
	}
	b.Speed = conf.BallSpeed
}
