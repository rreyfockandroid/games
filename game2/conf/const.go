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
	PaddleSpeed = 5
	BallSpeed   = 3

	ScreenWidth  = 400
	ScreenHeight = 300

	PaddleWidth  = 10
	PaddleHeight = 50

	BallRadius     = 5
	FontResultSize = 8

	StartX = 100
	StartY = 100

	ScoreTpl = "Score: %d:%d"
)

var (
	StartColor = YellowColor
	ScoreFont  = newScoreFont()
)

func newScoreFont() *text.GoTextFaceSource {
	reader := bytes.NewReader(fontData) // ✅ io.Reader
	face, err := text.NewGoTextFaceSource(reader)
	if err != nil {
		log.Fatal(err)
	}
	return face
}
