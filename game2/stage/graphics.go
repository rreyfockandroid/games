package stage

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"pl.home/game2/conf"
)

func DrawLeftPadle(screen *ebiten.Image, y float32) {
	vector.FillRect(screen, 0, y, conf.PaddleWidth, conf.PaddleHeight, conf.WhiteColor, true) // left paddle
}

func DrawRightPadle(screen *ebiten.Image, y float32) {
	vector.FillRect(screen, conf.ScreenWidth-conf.PaddleWidth, y, conf.PaddleWidth, conf.PaddleHeight, conf.YellowColor, true) // right paddle
}

func DrawBall(screen *ebiten.Image, x, y float32, color color.Color) {
	vector.FillCircle(screen, x, y, conf.BallRadius, color, true) // ball
}

func DrawResult(screen *ebiten.Image, resultMsg string) {
	opts := &text.DrawOptions{}
	opts.GeoM.Translate(conf.ScreenWidth/2-float64(len(conf.ScoreTpl)*5), conf.ScreenHeight-conf.FontResultSize)
	opts.ColorScale.ScaleWithColor(conf.WhiteColor)
	text.Draw(screen, resultMsg, face, opts)
}

var face = &text.GoTextFace{
	Source: conf.ScoreFont,
	Size:   conf.FontResultSize,
}

func DrawMainMenu(screen *ebiten.Image, menu bool, msg string, idx int) {
	opts := &text.DrawOptions{}
	opts.ColorScale.ScaleWithColor(conf.WhiteColor)
	opts.GeoM.Translate(conf.ScreenWidth/2-50, conf.ScreenHeight/2+float64(2*idx*conf.FontResultSize)-50)

	if menu {
		opts.ColorScale.ScaleWithColor(conf.YellowColor)
	}

	text.Draw(screen, msg, face, opts)
}

func DrawPause(screen *ebiten.Image) {
	opts := &text.DrawOptions{}
	opts.ColorScale.ScaleWithColor(conf.WhiteColor)
	opts.GeoM.Translate(conf.ScreenWidth/2-50, conf.ScreenHeight/2-10)

	text.Draw(screen, "PAUSE", face, opts)
}
