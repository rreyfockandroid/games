package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 640
	screenHeight = 480
	paddleWidth  = 10
	paddleHeight = 80
	ballSize     = 10
)

type Game struct {
	paddleY                      float64
	ballX, ballY, ballVX, ballVY float64
}

func (g *Game) Update() error {
	// Sterowanie graczem (strzałki góra/dół)
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) && g.paddleY > 0 {
		g.paddleY -= 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) && g.paddleY < screenHeight-paddleHeight {
		g.paddleY += 5
	}

	// Ruch piłki
	g.ballX += g.ballVX
	g.ballY += g.ballVY

	// Odbicie od ścian
	if g.ballY <= 0 || g.ballY >= screenHeight-ballSize {
		g.ballVY *= -1
	}

	// Odbicie od paletki
	if g.ballX <= paddleWidth &&
		g.ballY+ballSize >= g.paddleY &&
		g.ballY <= g.paddleY+paddleHeight {
		g.ballVX *= -1
	}

	// Odbicie od prawej ściany
	if g.ballX >= screenWidth-ballSize {
		g.ballVX *= -1
	}

	// Reset piłki, jeśli wpadnie za paletkę
	if g.ballX < 0 {
		g.ballX = screenWidth / 2
		g.ballY = screenHeight / 2
		g.ballVX = 4
		g.ballVY = 4
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 255}) // tło czarne

	// Paletka
	ebitenutil.DrawRect(screen, 0, g.paddleY, paddleWidth, paddleHeight, color.White)

	// Piłka
	ebitenutil.DrawRect(screen, g.ballX, g.ballY, ballSize, ballSize, color.White)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	game := &Game{
		paddleY: screenHeight/2 - paddleHeight/2,
		ballX:   screenWidth / 2,
		ballY:   screenHeight / 2,
		ballVX:  4,
		ballVY:  4,
	}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Simple Pong in Go")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
