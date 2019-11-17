package modifier

import (
	. "github.com/kravemir/golaypdf"
)

type Font struct {
	Content FixedWidthMeasurable

	Family, Style string
	Size          float64
}

func (f Font) applyFont(context Context) func(context Context) {
	family, style, size := context.GetFont()

	context.SetFont(f.Family, f.Style, f.Size)

	return func(context Context) {
		context.SetFont(family, style, size)
	}
}

func (f Font) Measure(context Context, width float64) (height float64, render Renderer) {
	return measureApplyModifier(context, width, f.applyFont, f.Content)
}
