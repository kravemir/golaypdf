package element

import . "github.com/kravemir/golaypdf"

type Empty struct {
	Width, Height float64
}

func (e Empty) Measure(context Context, width float64) (height float64, render Renderer) {
	return e.Height, FuncToRenderer(func(context Context, x, y, w, h float64) {})
}
