package utils

import (
	"github.com/ebitengine/debugui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Screen struct {
}

func (c Screen) Append(ctx *debugui.Context) {
	on := ebiten.IsFullscreen()
	ctx.Checkbox(&on, "Full screen [ctrl+f]").On(func() {
		c.fullScreen()
	})
}

func (c Screen) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyControl) && inpututil.IsKeyJustPressed(ebiten.KeyF) {
		c.fullScreen()
	}
	return nil
}

func (c Screen) fullScreen() {
	ebiten.SetFullscreen(!ebiten.IsFullscreen())
}
