package stage

import (
	"errors"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"pl.home/game2/conf"
	"pl.home/game5/pkg/utils"
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

	scene *Scene

	debugWindow *utils.DebugWindow
}

func NewGame() *Game {
	return &Game{
		debugWindow: utils.NewDebugWindow(conf.ScreenWidth, conf.ScreenHeight, false),
		state:       GameStateMenu,
		menuIndex:   0,
		mainMenu:    []string{"Start Game", "Options", "Remote", "Exit"},
		scene:       NewScene(),
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func (g *Game) Update() error {
	if err := g.debugWindow.Update(); err != nil {
		return err
	}
	g.updatePauseEsc()
	switch g.state {
	case GameStatePlaying:
		if err := g.scene.Update(); err != nil {
			return err
		}
	case GameStatePause:

	case GameStateMenu:
		g.updateMenu()
	case GameStateExit:
		return errors.New("exit game")
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.debugWindow.Draw(screen)
	switch g.state {
	case GameStatePlaying:
		g.scene.Draw(screen)
	case GameStateMenu:
		g.drawMenu(screen)
	case GameStatePause:
		g.drawPause(screen)
	case GameStateOptions:
		g.drawOptions(screen)
	}
}

func (g *Game) drawOptions(screen *ebiten.Image) {
	g.state = GameStateOptions
}

func (g *Game) drawPause(screen *ebiten.Image) {
	DrawPause(screen)
}

func (g *Game) drawMenu(screen *ebiten.Image) {
	for i := 0; i < len(g.mainMenu); i++ {
		DrawMainMenu(screen, i == g.menuIndex, g.mainMenu[i], i)
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
		if g.state == GameStatePlaying {
			g.scene.Stop()
		}
		g.state = GameStateMenu
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
			g.scene.Start()
		case 1:
			g.state = GameStateOptions
		case 2:

		case 3:
			g.state = GameStateExit
		}
	}

}
