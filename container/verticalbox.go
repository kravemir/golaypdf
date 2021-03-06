package container

import (
	"github.com/kravemir/golaypdf"
)

type VerticalBox struct {
	Rows []golaypdf.FixedWidthMeasurable
}

func (v VerticalBox) Measure(context golaypdf.MeasureContext, width float64) (height float64, render golaypdf.Renderer) {
	type itemType struct {
		height   float64
		renderer golaypdf.Renderer
	}

	var item itemType
	items := make([]itemType, len(v.Rows))
	height = 0.0

	for idx, row := range v.Rows {
		item.height, item.renderer = row.Measure(context, width)
		height += item.height
		items[idx] = item
	}

	return height, golaypdf.FuncToRenderer(func(context golaypdf.RenderContext, x, y, w, h float64) {
		for _, row := range items {
			row.renderer.Render(context, x, y, w, row.height)
			y += row.height
		}
	})
}
