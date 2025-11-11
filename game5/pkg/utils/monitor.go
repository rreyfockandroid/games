package utils

import (
	"fmt"

	"github.com/ebitengine/debugui"
	"github.com/hajimehoshi/ebiten/v2"
)

type Monitor struct {
	monitors []*ebiten.MonitorType
}

func NewMonitor() *Monitor {
	m := &Monitor{}
	m.monitors = ebiten.AppendMonitors(nil)
	targetMonitor := m.monitors[0]
	ebiten.SetMonitor(targetMonitor)
	return m
}

func (m *Monitor) Append(ctx *debugui.Context) {
	m.monitors = ebiten.AppendMonitors(m.monitors[:0])

	activeMonitor := ebiten.Monitor()
	index := -1
	for i, m := range m.monitors {
		if m == activeMonitor {
			index = i
			break
		}
	}

	ctx.Header("Monitors", false, func() {
		ctx.TreeNode("Active Monitor Info", func() {
			ctx.SetGridLayout([]int{-1, -2}, nil)
			ctx.Text("Index")
			ctx.Text(fmt.Sprintf("%d", index))
			ctx.Text("Name")
			ctx.Text(activeMonitor.Name())
			ctx.Text("Size")
			w, h := activeMonitor.Size()
			ctx.Text(fmt.Sprintf("%d x %d", w, h))
			ctx.Text("Device Scale Factor")
			ctx.Text(fmt.Sprintf("%0.2f", activeMonitor.DeviceScaleFactor()))
		})
		ctx.TreeNode("Details", func() {
			for i, m := range m.monitors {
				ctx.IDScope(fmt.Sprintf("%d", i), func() {
					name := fmt.Sprintf("%d: %s", i, m.Name())
					if i == index {
						name += " (Active)"
					}
					ctx.TreeNode(name, func() {
						if index != i {
							ctx.Button("Activate").On(func() {
								ebiten.SetMonitor(m)
							})
						}
						ctx.SetGridLayout([]int{-1, -2}, nil)
						ctx.Text("Name")
						ctx.Text(m.Name())
						ctx.Text("Size")
						w, h := m.Size()
						ctx.Text(fmt.Sprintf("%d x %d", w, h))
						ctx.Text("Device Scale Factor")
						ctx.Text(fmt.Sprintf("%0.2f", m.DeviceScaleFactor()))
					})
				})
			}
		})
	})
}

func (m *Monitor) Update() error {
	return nil
}

func (m *Monitor) Draw(screen *ebiten.Image) {}
