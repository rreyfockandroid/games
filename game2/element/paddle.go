package element

import "pl.home/game2/conf"

type Paddle struct {
	X, Y  float32
	Speed float32
}

func NewPaddle() *Paddle {
	return &Paddle{
		Speed: conf.PaddleSpeed,
		Y:     conf.PaddleStartPosition,
	}
}

func (p *Paddle) Reset() {
	p.Speed = conf.PaddleSpeed
	p.Y = conf.PaddleStartPosition
}
