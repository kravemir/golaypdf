package modifier

import (
	"github.com/kravemir/golaypdf"
)

type fontModifierContext interface {
	GetFont() (family, style string, size float64)
	SetFont(family, style string, size float64)
	SetFontStyle(style string)
	SetFontSize(size float64)
}

func measureApplyModifier(
	ctx golaypdf.MeasureContext,
	width float64,
	applyFunc func(ctx fontModifierContext) (unapply func(ctx fontModifierContext)),
	content golaypdf.FixedWidthMeasurable,
) (float64, golaypdf.Renderer) {
	unapply := applyFunc(ctx)
	height, renderer := content.Measure(ctx, width)
	unapply(ctx)

	return height, golaypdf.FuncToRenderer(func(context golaypdf.RenderContext, x, y, w, h float64) {
		unapply := applyFunc(context)
		renderer.Render(context, x, y, w, h)
		unapply(context)
	})
}
