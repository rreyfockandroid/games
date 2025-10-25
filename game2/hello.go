package main

import (
	"image/color"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/hajimehoshi/ebiten/v2/vector"
)

var (
	white  = color.White
	yellow = color.RGBA{255, 255, 0, 255}
)

const (
	paddleSpeed = 5
	ballSpeed   = 3

	screenWidth  = 400
	screenHeight = 300

	paddleWidth  = 10
	paddleHeight = 50

	ballRadius = 5
)

type Game struct {
	paddleYa float32
	paddleYb float32

	ballX, ballY float32
	ballDr       bool
	ballDb       bool

	ballColor color.Color
}

func (g *Game) Reset() {
	g.ballY = 100
	g.ballX = 100
	g.ballDr = true
	time.Sleep(time.Second)
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) && g.paddleYb > 0 {
		g.paddleYb -= paddleSpeed
		// log.Println("paddleYb", g.paddleYb)
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) && g.paddleYb < screenHeight-paddleHeight {
		g.paddleYb += paddleSpeed
		// log.Println("paddleYb", g.paddleYb)
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) && g.paddleYa > 0 {
		g.paddleYa -= paddleSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) && g.paddleYa < screenHeight-paddleHeight {
		g.paddleYa += paddleSpeed
	}

	if g.ballDr {
		g.ballX += ballSpeed
	} else {
		g.ballX -= ballSpeed
	}
	if g.ballDb {
		g.ballY += ballSpeed
	} else {
		g.ballY -= ballSpeed
	}

	if g.ballX >= screenWidth-ballRadius {
		g.ballDr = false
		if g.ballY >= g.paddleYb && g.ballY <= g.paddleYb+paddleHeight {
			// log.Println("Trafiona B")
			g.ballColor = white
		} else {
			g.Reset()
		}
		// uderza w prawa
		// log.Println("right")
	} else if g.ballX <= ballRadius {
		g.ballDr = true
		if g.ballY >= g.paddleYa && g.ballY <= g.paddleYa+paddleHeight {
			// log.Println("Trafiona A")
			g.ballColor = yellow
		} else {
			g.Reset()
		}
		// uderza w lewa
		// log.Println("left")
	}

	if g.ballY >= screenHeight-ballRadius {
		g.ballDb = false
	} else if g.ballY <= ballRadius {
		g.ballDb = true
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{10, 10, 10, 255})

	vector.FillRect(screen, 0, g.paddleYa, paddleWidth, paddleHeight, white, true)

	vector.FillRect(screen, screenWidth-paddleWidth, g.paddleYb, paddleWidth, paddleHeight, yellow, true)

	vector.FillCircle(screen, g.ballX, g.ballY, ballRadius, g.ballColor, true)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	game := &Game{
		ballX:     100,
		ballY:     100,
		ballDr:    true,
		ballDb:    true,
		ballColor: yellow,
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
