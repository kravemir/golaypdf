package modifier

import (
	"github.com/kravemir/golaypdf"
)

type FontSize struct {
	Content golaypdf.FixedWidthMeasurable

	Size float64
}

func (f FontSize) applyFontSize(context golaypdf.Context) func(context golaypdf.Context) {
	_, _, oldSize := context.GetFont()

	context.SetFontSize(f.Size)

	return func(context golaypdf.Context) {
		context.SetFontSize(oldSize)
	}
}

func (f FontSize) Measure(context golaypdf.Context, width float64) (float64, func(context golaypdf.Context, x, y, w, h float64)) {
	return measureApplyModifier(context, width, f.applyFontSize, f.Content)
}
