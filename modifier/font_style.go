package modifier

import (
	"github.com/kravemir/golaypdf"
)

type FontStyle struct {
	Content golaypdf.FixedWidthMeasurable

	Style string
}

func (f FontStyle) applyFontSize(context fontModifierContext) func(context fontModifierContext) {
	_, oldStyle, _ := context.GetFont()

	context.SetFontStyle(f.Style)

	return func(context fontModifierContext) {
		context.SetFontStyle(oldStyle)
	}
}

func (f FontStyle) Measure(context golaypdf.MeasureContext, width float64) (height float64, render golaypdf.Renderer) {
	return measureApplyModifier(context, width, f.applyFontSize, f.Content)
}
