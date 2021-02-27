package modifier

import (
	"github.com/kravemir/golaypdf"
)

type Font struct {
	Content golaypdf.FixedWidthMeasurable

	Family, Style string
	Size          float64
}

func (f Font) applyFont(context fontModifierContext) func(context fontModifierContext) {
	family, style, size := context.GetFont()

	context.SetFont(f.Family, f.Style, f.Size)

	return func(context fontModifierContext) {
		context.SetFont(family, style, size)
	}
}

func (f Font) Measure(context golaypdf.MeasureContext, width float64) (height float64, render golaypdf.Renderer) {
	return measureApplyModifier(context, width, f.applyFont, f.Content)
}
