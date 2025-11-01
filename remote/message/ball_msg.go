package message

import "image/color"

type BallMsg struct {
	X, Y                      float32
	DirectRight, DirectBottom bool
	Speed                     float32
	Color                     color.Color
	Stop                      bool
}
