package modifier

import (
	"github.com/kravemir/golaypdf"
)

func measureApplyModifier(context golaypdf.Context, width float64, applyFunc func(context golaypdf.Context) (unapply func(context golaypdf.Context)), content golaypdf.FixedWidthMeasurable) (float64, func(context golaypdf.Context, x float64, y float64, w float64, h float64)) {
	unapply := applyFunc(context)
	height, render := content.Measure(context, width)
	unapply(context)

	return height, func(context golaypdf.Context, x, y, w, h float64) {
		unapply := applyFunc(context)
		render(context, x, y, w, h)
		unapply(context)
	}
}
