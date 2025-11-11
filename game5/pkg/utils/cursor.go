package utils

import (
	"fmt"
	"image/color"

	"github.com/ebitengine/debugui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Cursor struct {
	windowWidth, windowHeigh int

	pointerImage   *ebiten.Image
	x, y           int
	mouseX, mouseY int

	collapsed bool
}

func NewCursor(windowWidth, windowHeigh int) *Cursor {
	c := &Cursor{
		windowWidth: windowWidth,
		windowHeigh: windowHeigh,
	}
	c.pointerImage = ebiten.NewImage(4, 4)
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

	maxW := c.windowWidth
	maxH := c.windowHeigh
	if ebiten.IsFullscreen() {
		maxW, maxH = ebiten.Monitor().Size()
	}

	// Constrain red dot within screen view.
	if c.x < 0 {
		c.x = 0
	} else if c.x > maxW-c.pointerImage.Bounds().Dx() {
		c.x = maxW - c.pointerImage.Bounds().Dx()
	}

	if c.y < 0 {
		c.y = 0
	} else if c.y > maxH-c.pointerImage.Bounds().Dy() {
		c.y = maxH - c.pointerImage.Bounds().Dy()
	}
}

func (c *Cursor) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyControl) && inpututil.IsKeyJustPressed(ebiten.KeyM) {
		if !c.isOn() {
			c.collapsed = true
			ebiten.SetCursorMode(ebiten.CursorModeCaptured)
			c.mouseX, c.mouseY = ebiten.CursorPosition()
		} else {
			c.collapsed = false
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
	ctx.Header("Cursor info", c.collapsed, func() {
		ctx.Checkbox(&on, "Cursor position [ctrl+m]").On(func() {
			if on {
				ebiten.SetCursorMode(ebiten.CursorModeCaptured)
			}
		})
		msg := fmt.Sprintf("x: %d\ny: %d", c.x, c.y)
		ctx.Text(msg)

	})
}

func (c *Cursor) isOn() bool {
	return ebiten.CursorMode() == ebiten.CursorModeCaptured
}
