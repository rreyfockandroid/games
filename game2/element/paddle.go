package element

import "pl.home/game2/conf"

type Paddle struct {
	X, Y  float32
	Speed float32
}

func NewPaddle(y float32) *Paddle {
	return &Paddle{
		Speed: conf.PaddleSpeed,
		Y:     y,
	}
}

func (p *Paddle) Reset() {
	p.Speed = conf.PaddleSpeed
}
