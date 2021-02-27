package element

import (
	"github.com/kravemir/golaypdf"
)

type Empty struct {
	Width, Height float64
}

func (e Empty) Measure(_ golaypdf.MeasureContext, _ float64) (height float64, render golaypdf.Renderer) {
	return e.Height, golaypdf.FuncToRenderer(func(_ golaypdf.RenderContext, x, y, w, h float64) {})
}
