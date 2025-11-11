package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"pl.home/game5/internal/utils"
)

type Game struct {
	debugWindow *utils.DebugWindow
}

func NewGame() *Game {
	return &Game{
		debugWindow: utils.NewDebugWindow(),
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func (g *Game) Update() error {
	if err := g.debugWindow.Update(); err != nil {
		return err
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.debugWindow.Draw(screen)
}
