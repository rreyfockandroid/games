package board

import (
	"errors"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"pl.home/game2/conf"
)

type GameState int

const (
	GameStateMenu GameState = iota
	GameStatePlaying
	GameStateOptions
	GameStatePause
	GameStateExit
)

type Game struct {
	state     GameState
	menuIndex int
	mainMenu  []string
}

func NewGame() *Game {
	return &Game{
		state:     GameStateMenu,
		menuIndex: 0,
		mainMenu:  []string{"Start Game", "Options", "Exit"},
	}
}

func (g *Game) IsPlaying() bool {
	return g.state == GameStatePlaying
}

func (g *Game) Update() error {
	g.updatePauseEsc()
	switch g.state {
	case GameStatePlaying:
	case GameStatePause:

	case GameStateMenu:
		g.updateMenu()
	case GameStateExit:
		return errors.New("exit game")
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch g.state {
	case GameStateMenu:
		g.drawMenu(screen)
	case GameStatePause:
		g.drawPause(screen)
	}
}

func (g *Game) updatePauseEsc() {
	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		if g.state == GameStatePause {
			g.state = GameStatePlaying
		} else {
			g.state = GameStatePause
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		g.state = GameStateMenu
	}
}

func (g *Game) drawPause(screen *ebiten.Image) {
	face := &text.GoTextFace{
		Source: conf.ScoreFont,
		Size:   conf.FontResultSize,
	}
	opts := &text.DrawOptions{}
	opts.ColorScale.ScaleWithColor(conf.WhiteColor)
	opts.GeoM.Translate(conf.ScreenWidth/2-50, conf.ScreenHeight/2-10)

	text.Draw(screen, "PAUSE", face, opts)
}

func (g *Game) drawMenu(screen *ebiten.Image) {
	face := &text.GoTextFace{
		Source: conf.ScoreFont,
		Size:   conf.FontResultSize,
	}
	for i := 0; i < len(g.mainMenu); i++ {
		opts := &text.DrawOptions{}
		opts.ColorScale.ScaleWithColor(conf.WhiteColor)
		opts.GeoM.Translate(conf.ScreenWidth/2-50, conf.ScreenHeight/2+float64(2*i*conf.FontResultSize)-50)

		msg := g.mainMenu[i]
		if i == g.menuIndex {
			opts.ColorScale.ScaleWithColor(conf.YellowColor)
		}

		text.Draw(screen, msg, face, opts)
	}

}

func (g *Game) updateMenu() {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		g.menuIndex--
		if g.menuIndex < 0 {
			g.menuIndex = len(g.mainMenu) - 1
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		g.menuIndex++
		if g.menuIndex >= len(g.mainMenu) {
			g.menuIndex = 0
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		switch g.menuIndex {
		case 0:
			g.state = GameStatePlaying
		case 1:
			g.state = GameStateOptions
		case 2:
			g.state = GameStateExit
		}
	}

}
