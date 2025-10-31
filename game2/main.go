package main

import (
	"fmt"
	"image/color"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"pl.home/game2/board"
	"pl.home/game2/conf"
	el "pl.home/game2/element"

	"github.com/hajimehoshi/ebiten/v2/vector"
)

func main() {
	game := &Game{
		ball:        el.NewBall(),
		paddleLeft:  el.NewPaddle(conf.LeftPaddleStartPosition),
		paddleRight: el.NewPaddle(conf.RightPaddleStartPosition),
		board:       board.NewBoard1(),

		tickerReset: make(chan struct{}),
	}
	game.gameSpeedUp()

	ebiten.SetWindowSize(conf.WindowWidth, conf.WindowHeigh)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

type Game struct {
	ball *el.Ball

	paddleLeft  *el.Paddle
	paddleRight *el.Paddle

	score el.Score
	board board.Board

	tickerReset chan struct{}
}

func (g *Game) gameSpeedUp() {
	tick := time.Second * 10
	ticker := time.NewTicker(tick)
	go func() {
		for {
			select {
			case <-ticker.C:
				g.ball.Speed++
				g.paddleLeft.Speed++
				g.paddleRight.Speed++
				log.Println("speed up!")
			case <-g.tickerReset:
				ticker.Reset(tick)
				// log.Println("speed up reset")
			}
		}
	}()

}

func (g *Game) reset() {
	g.ball.Reset()
	g.paddleRight.Reset()
	g.paddleLeft.Reset()
	g.tickerReset <- struct{}{}
	time.Sleep(time.Second)
}

func (g *Game) scoreResult() (resultMsg string, face *text.GoTextFace, opts *text.DrawOptions) {
	opts = &text.DrawOptions{}
	opts.GeoM.Translate(conf.ScreenWidth/2-float64(len(conf.ScoreTpl)*5), conf.ScreenHeight-conf.FontResultSize)
	opts.ColorScale.ScaleWithColor(conf.WhiteColor)

	resultMsg = fmt.Sprintf(conf.ScoreTpl, g.score.Left, g.score.Right)

	face = &text.GoTextFace{
		Source: conf.ScoreFont,
		Size:   conf.FontResultSize,
	}
	return
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) && g.paddleRight.Y > 0 {
		g.paddleRight.Y -= g.paddleRight.Speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) && g.paddleRight.Y < conf.ScreenHeight-conf.PaddleHeight {
		g.paddleRight.Y += g.paddleRight.Speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) && g.paddleLeft.Y > 0 {
		g.paddleLeft.Y -= g.paddleLeft.Speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) && g.paddleLeft.Y < conf.ScreenHeight-conf.PaddleHeight {
		g.paddleLeft.Y += g.paddleLeft.Speed
	}
	if g.ball.DirectRight {
		g.ball.X += g.ball.Speed
	} else {
		g.ball.X -= g.ball.Speed
	}
	if g.ball.DirectBottom {
		g.ball.Y += g.ball.Speed
	} else {
		g.ball.Y -= g.ball.Speed
	}

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

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{10, 10, 10, 255})

	vector.FillRect(screen, 0, g.paddleLeft.Y, conf.PaddleWidth, conf.PaddleHeight, conf.WhiteColor, true) // left paddle

	vector.FillRect(screen, conf.ScreenWidth-conf.PaddleWidth, g.paddleRight.Y, conf.PaddleWidth, conf.PaddleHeight, conf.YellowColor, true) // right paddle

	vector.FillCircle(screen, g.ball.X, g.ball.Y, conf.BallRadius, g.ball.Color, true) // ball

	g.board.Draw(screen)

	resultMsg, face, opts := g.scoreResult()
	text.Draw(screen, resultMsg, face, opts)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return conf.ScreenWidth, conf.ScreenHeight
}
