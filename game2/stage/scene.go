package stage

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"pl.home/game2/board"
	"pl.home/game2/conf"
	el "pl.home/game2/element"
)

type Scene struct {
	ball *el.Ball

	paddleLeft  *el.Paddle
	paddleRight *el.Paddle

	score el.Score
	board board.Board

	motion *MotionController

	started bool
}

func NewScene() *Scene {
	paddleLeft := el.NewPaddle()
	paddleRight := el.NewPaddle()
	ball := el.NewBall()
	scene := &Scene{
		ball:        ball,
		paddleLeft:  paddleLeft,
		paddleRight: paddleRight,
		board:       board.NewBoard1(),
		motion:      NewMotionController(paddleLeft, paddleRight, ball),
	}
	return scene
}

func (g *Scene) Start() {
	g.motion.Start()
	g.started = true
}

func (g *Scene) Stop() {
	g.motion.Stop()
	g.started = false

	g.ball.Reset()
	g.paddleRight.Reset()
	g.paddleLeft.Reset()
}

func (g *Scene) reset() {
	g.ball.Reset()
	g.paddleRight.Reset()
	g.paddleLeft.Reset()
	g.motion.Reset()
	// time.Sleep(time.Second)
}

func (g *Scene) Update() error {
	if !g.started {
		return nil
	}

	g.motion.Update(ebiten.IsKeyPressed(ebiten.KeyArrowUp), ebiten.IsKeyPressed(ebiten.KeyArrowDown), ebiten.IsKeyPressed(ebiten.KeyW), ebiten.IsKeyPressed(ebiten.KeyS))
	g.board.Update(g.ball)

	if g.ball.X >= conf.ScreenWidth-conf.BallRadius {
		if g.ball.Y >= g.paddleRight.Y && g.ball.Y <= g.paddleRight.Y+conf.PaddleHeight {
			// log.Println("Trafiona B")
			g.ball.Color = conf.WhiteColor
			g.ball.DirectRight = false
		} else {
			g.reset()
			g.score.Left++
		}
		// uderza w prawa
		// log.Println("right")
	} else if g.ball.X <= conf.BallRadius {
		if g.ball.Y >= g.paddleLeft.Y && g.ball.Y <= g.paddleLeft.Y+conf.PaddleHeight {
			// log.Println("Trafiona A")
			g.ball.Color = conf.YellowColor
			g.ball.DirectRight = true
		} else {
			g.reset()
			g.score.Right++
		}
		// uderza w lewa
		// log.Println("left")
	}

	if g.ball.Y >= conf.ScreenHeight-conf.BallRadius {
		g.ball.DirectBottom = false
	} else if g.ball.Y <= conf.BallRadius {
		g.ball.DirectBottom = true
	}

	return nil
}

func (g *Scene) Draw(screen *ebiten.Image) {
	if !g.started {
		return
	}

	screen.Fill(color.RGBA{10, 10, 10, 255})

	DrawLeftPadle(screen, g.paddleLeft.Y)
	DrawRightPadle(screen, g.paddleRight.Y)
	DrawBall(screen, g.ball.X, g.ball.Y, g.ball.Color)

	g.board.Draw(screen)

	DrawResult(screen, fmt.Sprintf(conf.ScoreTpl, g.score.Left, g.score.Right))
}
