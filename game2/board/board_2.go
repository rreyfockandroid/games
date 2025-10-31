package board

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	el "pl.home/game2/element"
)

type Board2 struct {
	wall *el.Wall
}

func NewBoard2() *Board2 {
	return &Board2{
		wall: el.NewWall(),
	}
}

func (b *Board2) Update(ball *el.Ball) error {
	if b.wall.UpdateDirect(ball.X, ball.Y) {
		if ball.DirectRight {
			ball.DirectRight = false
		} else {
			ball.DirectRight = true
		}
	}
	return nil
}

func (b *Board2) Draw(screen *ebiten.Image) {
	vector.FillRect(screen, b.wall.X, b.wall.Y, b.wall.Width, b.wall.Height, b.wall.Color, true) // wall
}
