package main

import (
	"fmt"
	"image/color"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"pl.home/game2/conf"

	"github.com/hajimehoshi/ebiten/v2/vector"
)

func main() {
	game := &Game{
		ballX:       conf.StartX,
		ballY:       conf.StartY,
		ballDr:      true,
		ballDb:      true,
		ballColor:   conf.StartColor,
		ballSpeed:   conf.BallSpeed,
		paddleSpeed: conf.PaddleSpeed,
		tickerReset: make(chan struct{}),
	}
	game.gameSpeedUp()

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

type Score struct {
	Left  int
	Right int
}

type Game struct {
	paddleRightY float32
	paddleLeftY  float32

	ballX, ballY float32
	ballDr       bool
	ballDb       bool

	ballColor color.Color

	ballSpeed   float32
	paddleSpeed float32

	score Score

	tickerReset chan struct{}
}

func (g *Game) gameSpeedUp() {
	tick := time.Second * 10
	ticker := time.NewTicker(tick)
	go func() {
		for {
			select {
			case <-ticker.C:
				g.ballSpeed++
				g.paddleSpeed++
				log.Println("speed up!")
			case <-g.tickerReset:
				ticker.Reset(tick)
				// log.Println("speed up reset")
			}
		}
	}()

}

func (g *Game) reset() {
	g.ballY = conf.StartX
	g.ballX = conf.StartY
	g.ballColor = conf.StartColor
	g.ballDr = true
	g.ballSpeed = conf.BallSpeed
	g.paddleSpeed = conf.PaddleSpeed
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
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) && g.paddleLeftY > 0 {
		g.paddleLeftY -= g.paddleSpeed
		// log.Println("paddleYb", g.paddleYb)
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) && g.paddleLeftY < conf.ScreenHeight-conf.PaddleHeight {
		g.paddleLeftY += g.paddleSpeed
		// log.Println("paddleYb", g.paddleYb)
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) && g.paddleRightY > 0 {
		g.paddleRightY -= g.paddleSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) && g.paddleRightY < conf.ScreenHeight-conf.PaddleHeight {
		g.paddleRightY += g.paddleSpeed
	}

	if g.ballDr {
		g.ballX += g.ballSpeed
	} else {
		g.ballX -= g.ballSpeed
	}
	if g.ballDb {
		g.ballY += g.ballSpeed
	} else {
		g.ballY -= g.ballSpeed
	}

	if g.ballX >= conf.ScreenWidth-conf.BallRadius {
		g.ballDr = false
		if g.ballY >= g.paddleLeftY && g.ballY <= g.paddleLeftY+conf.PaddleHeight {
			// log.Println("Trafiona B")
			g.ballColor = conf.WhiteColor
		} else {
			g.reset()
			g.score.Left++
		}
		// uderza w prawa
		// log.Println("right")
	} else if g.ballX <= conf.BallRadius {
		g.ballDr = true
		if g.ballY >= g.paddleRightY && g.ballY <= g.paddleRightY+conf.PaddleHeight {
			// log.Println("Trafiona A")
			g.ballColor = conf.YellowColor
		} else {
			g.reset()
			g.score.Right++
		}
		// uderza w lewa
		// log.Println("left")
	}

	if g.ballY >= conf.ScreenHeight-conf.BallRadius {
		g.ballDb = false
	} else if g.ballY <= conf.BallRadius {
		g.ballDb = true
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{10, 10, 10, 255})

	vector.FillRect(screen, 0, g.paddleRightY, conf.PaddleWidth, conf.PaddleHeight, conf.WhiteColor, true)

	vector.FillRect(screen, conf.ScreenWidth-conf.PaddleWidth, g.paddleLeftY, conf.PaddleWidth, conf.PaddleHeight, conf.YellowColor, true)

	vector.FillCircle(screen, g.ballX, g.ballY, conf.BallRadius, g.ballColor, true)

	resultMsg, face, opts := g.scoreResult()
	text.Draw(screen, resultMsg, face, opts)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return conf.ScreenWidth, conf.ScreenHeight
}
