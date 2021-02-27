package container

import (
	"github.com/kravemir/golaypdf"
)

type BorderBox struct {
	Content                      golaypdf.FixedWidthMeasurable
	PTop, PRight, PLeft, PBottom float64
}

func (b BorderBox) Measure(context golaypdf.MeasureContext, width float64) (height float64, render golaypdf.Renderer) {
	contentWidth := width - b.PLeft - b.PRight
	contentHeight, contentRender := b.Content.Measure(context, contentWidth)

	return contentHeight + b.PTop + b.PBottom, golaypdf.FuncToRenderer(func(context golaypdf.RenderContext, x, y, w, h float64) {
		context.PDF().Rect(x, y, w, h, "D")

		contentRender.Render(context, x+b.PLeft, y+b.PTop, contentWidth, contentHeight)
	})
}
