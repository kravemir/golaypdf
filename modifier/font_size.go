package modifier

import (
	. "github.com/kravemir/golaypdf"
)

type FontSize struct {
	Content FixedWidthMeasurable

	Size float64
}

func (f FontSize) applyFontSize(context Context) func(context Context) {
	_, _, oldSize := context.GetFont()

	context.SetFontSize(f.Size)

	return func(context Context) {
		context.SetFontSize(oldSize)
	}
}

func (f FontSize) Measure(context Context, width float64) (height float64, render Renderer) {
	return measureApplyModifier(context, width, f.applyFontSize, f.Content)
}
