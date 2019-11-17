package element

import . "github.com/kravemir/golaypdf"

type HorizontaLine struct {
}

func (l HorizontaLine) Measure(context Context, width float64) (height float64, render Renderer) {
	// TODO: line sizing
	return 0.2, FuncToRenderer(func(context Context, x, y, w, h float64) {
		context.PDF().Line(x, y+h/2, x+w, y+h/2)
	})
}
