package modifier

import (
	. "github.com/kravemir/golaypdf"
)

type FontStyle struct {
	Content FixedWidthMeasurable

	Style string
}

func (f FontStyle) applyFontSize(context Context) func(context Context) {
	_, oldStyle, _ := context.GetFont()

	context.SetFontStyle(f.Style)

	return func(context Context) {
		context.SetFontStyle(oldStyle)
	}
}

func (f FontStyle) Measure(context Context, width float64) (height float64, render Renderer) {
	return measureApplyModifier(context, width, f.applyFontSize, f.Content)
}
