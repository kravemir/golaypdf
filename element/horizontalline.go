package element

import (
	"github.com/kravemir/golaypdf"
)

type HorizontaLine struct {
}

func (l HorizontaLine) Measure(context golaypdf.MeasureContext, width float64) (height float64, render golaypdf.Renderer) {
	// TODO: line sizing
	return 0.2, golaypdf.FuncToRenderer(func(context golaypdf.RenderContext, x, y, w, h float64) {
		context.PDF().Line(x, y+h/2, x+w, y+h/2)
	})
}
