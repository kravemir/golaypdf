package modifier

import (
	. "github.com/kravemir/golaypdf"
)

func measureApplyModifier(context Context, width float64, applyFunc func(context Context) (unapply func(context Context)), content FixedWidthMeasurable) (float64, Renderer) {
	unapply := applyFunc(context)
	height, renderer := content.Measure(context, width)
	unapply(context)

	return height, FuncToRenderer(func(context Context, x, y, w, h float64) {
		unapply := applyFunc(context)
		renderer.Render(context, x, y, w, h)
		unapply(context)
	})
}
