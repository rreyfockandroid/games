package utils

import (
	"image"

	"github.com/ebitengine/debugui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type DebugComponent interface {
	Update() error
	Draw(screen *ebiten.Image)
	Append(ctx *debugui.Context)
}

type DebugWindow struct {
	debugComponents []DebugComponent

	debugui debugui.DebugUI

	visible bool
}

func NewDebugWindow(windowWidth, windowHeigh int, visible bool) *DebugWindow {
	debugComponents := []DebugComponent{
		NewScreen(),
		NewCursor(windowWidth, windowHeigh),
		NewMonitor(),
	}

	debug := &DebugWindow{
		debugComponents: debugComponents,
		visible:         visible,
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
	for _, c := range d.debugComponents {
		if err := c.Update(); err != nil {
			return nil
		}
	}
	if _, err := d.debugui.Update(func(ctx *debugui.Context) error {
		ctx.Window("Debug [ctrl+d]", image.Rect(10, 10, 200, 200), func(layout debugui.ContainerLayout) {
			for _, c := range d.debugComponents {
				c.Append(ctx)
			}
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
	for _, c := range d.debugComponents {
		c.Draw(screen)
	}
}

func (d *DebugWindow) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func (d *DebugWindow) show() bool {
	return d.visible
}
