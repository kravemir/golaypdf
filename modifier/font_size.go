package modifier

import (
	"github.com/kravemir/golaypdf"
)

type FontSize struct {
	Content golaypdf.FixedWidthMeasurable

	Size float64
}

func (f FontSize) applyFontSize(context fontModifierContext) func(context fontModifierContext) {
	_, _, oldSize := context.GetFont()

	context.SetFontSize(f.Size)

	return func(context fontModifierContext) {
		context.SetFontSize(oldSize)
	}
}

func (f FontSize) Measure(context golaypdf.MeasureContext, width float64) (height float64, render golaypdf.Renderer) {
	return measureApplyModifier(context, width, f.applyFontSize, f.Content)
}
