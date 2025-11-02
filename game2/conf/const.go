package conf

import (
	"bytes"
	_ "embed"
	"log"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

//go:embed fonts/PressStart2P-Regular.ttf
var fontData []byte

const (
	Player       = "left"
	PortListener = 8129
	PortSender   = 8128

	// Player       = "right"
	// PortListener = 8128
	// PortSender   = 8129

	PaddleSpeed = 5
	BallSpeed   = 3

	ScreenWidth  = 400
	ScreenHeight = 300

	WindowWidth = 800
	WindowHeigh = 600

	PaddleWidth  = 5
	PaddleHeight = 50

	BallRadius     = 4
	FontResultSize = 8

	StartX = ScreenWidth / 2
	StartY = ScreenHeight / 2

	ScoreTpl = "Score: %d:%d"

	start                    = float32(20)
	LeftPaddleStartPosition  = start
	RightPaddleStartPosition = ScreenHeight - PaddleHeight - start

	WallWidth = 5
	WallHeigh = 50
	WallX     = ScreenWidth/2 - WallWidth
	WallY     = ScreenHeight / 2
)

var (
	StartColor = YellowColor
	ScoreFont  = newScoreFont()
)

func newScoreFont() *text.GoTextFaceSource {
	reader := bytes.NewReader(fontData) // io.Reader
	face, err := text.NewGoTextFaceSource(reader)
	if err != nil {
		log.Fatal(err)
	}
	return face
}
