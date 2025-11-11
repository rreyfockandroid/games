package utils

import (
	"image"

	"github.com/ebitengine/debugui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type DebugWindow struct {
	monitor *Monitor
	cursor  Cursor
	screen  Screen
	debugui debugui.DebugUI

	visible bool
}

func NewDebugWindow(windowWidth, windowHeigh int, visible bool) *DebugWindow {
	debug := &DebugWindow{
		monitor: NewMonitor(),
		visible: visible,
		cursor:  *NewCursor(windowWidth, windowHeigh),
	}
	return debug
}

func (d *DebugWindow) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyControl) && inpututil.IsKeyJustPressed(ebiten.KeyD) {
		d.visible = !d.visible
	}
	if !d.show() {
		return nil
	}
	if err := d.cursor.Update(); err != nil {
		return err
	}
	if err := d.screen.Update(); err != nil {
		return err
	}
	if _, err := d.debugui.Update(func(ctx *debugui.Context) error {
		ctx.Window("Debug [ctrl+d]", image.Rect(10, 10, 200, 200), func(layout debugui.ContainerLayout) {
			d.screen.Append(ctx)
			d.cursor.Append(ctx)
			d.monitor.Append(ctx)
		})
		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (d *DebugWindow) Draw(screen *ebiten.Image) {
	if !d.show() {
		return
	}
	d.debugui.Draw(screen)
	d.cursor.Draw(screen)
}

func (d *DebugWindow) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func (d *DebugWindow) show() bool {
	return d.visible
}
