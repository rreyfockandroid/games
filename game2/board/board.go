package board

import (
	"github.com/hajimehoshi/ebiten/v2"
	el "pl.home/game2/element"
)

type Board interface {
	Update(ball *el.Ball) error
	Draw(screen *ebiten.Image)
}

type Board1 struct {
}

func NewBoard1() *Board1 {
	return &Board1{}
}

func (b *Board1) Update(ball *el.Ball) error {
	return nil
}

func (b *Board1) Draw(screen *ebiten.Image) {
}
