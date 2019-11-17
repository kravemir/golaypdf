package modifier

import (
	"github.com/kravemir/golaypdf"
)

type Font struct {
	Content golaypdf.FixedWidthMeasurable

	Family, Style string
	Size          float64
}

func (f Font) applyFont(context golaypdf.Context) func(context golaypdf.Context) {
	family, style, size := context.GetFont()

	context.SetFont(f.Family, f.Style, f.Size)

	return func(context golaypdf.Context) {
		context.SetFont(family, style, size)
	}
}

func (f Font) Measure(context golaypdf.Context, width float64) (float64, func(context golaypdf.Context, x, y, w, h float64)) {
	return measureApplyModifier(context, width, f.applyFont, f.Content)
}
