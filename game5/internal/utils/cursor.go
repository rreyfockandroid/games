package utils

import (
	"fmt"
	"image/color"
	"sync"

	"github.com/ebitengine/debugui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"pl.home/game5/internal/cfg"
)

type Cursor struct {
	pointerImage *ebiten.Image
	x            int
	y            int
	mouseX       int
	mouseY       int
	initOnce     sync.Once
}

func NewCursor() *Cursor {
	c := &Cursor{}
	c.x = 100
	c.y = 100
	c.pointerImage = ebiten.NewImage(8, 8)
	c.pointerImage.Fill(color.RGBA{0xff, 0, 0, 0xff})
	return c
}

func (c *Cursor) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(c.x), float64(c.y))
	screen.DrawImage(c.pointerImage, op)
}

func (c *Cursor) pointer() {
	cursorX, cursorY := ebiten.CursorPosition()
	deltaX, deltaY := cursorX-c.mouseX, cursorY-c.mouseY
	c.mouseX, c.mouseY = cursorX, cursorY

	if deltaX != 0 {
		c.x += deltaX
	}

	if deltaY != 0 {
		c.y += deltaY
	}

	// Constrain red dot within screen view.
	if c.x < 0 {
		c.x = 0
	} else if c.x > cfg.WindowWidth-c.pointerImage.Bounds().Dx() {
		c.x = cfg.WindowWidth - c.pointerImage.Bounds().Dx()
	}

	if c.y < 0 {
		c.y = 0
	} else if c.y > cfg.WindowHeigh-c.pointerImage.Bounds().Dy() {
		c.y = cfg.WindowHeigh - c.pointerImage.Bounds().Dy()
	}
}

func (c *Cursor) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyControl) && inpututil.IsKeyJustPressed(ebiten.KeyM) {
		if !c.isOn() {
			ebiten.SetCursorMode(ebiten.CursorModeCaptured)
			c.mouseX, c.mouseY = ebiten.CursorPosition()
		} else {
			ebiten.SetCursorMode(ebiten.CursorModeVisible)
		}
	}
	if c.isOn() {
		c.pointer()
	}

	return nil
}

func (c *Cursor) Append(ctx *debugui.Context) {
	on := c.isOn()
	ctx.Header("Cursor info", true, func() {
		ctx.Checkbox(&on, "Cursor position [ctrl+m]").On(func() {
			if on {
				ebiten.SetCursorMode(ebiten.CursorModeCaptured)
			}
			fmt.Println("pressed")
		})
	})
}

func (c *Cursor) isOn() bool {
	return ebiten.CursorMode() == ebiten.CursorModeCaptured
}
