package element

import (
	"math"

	"github.com/kravemir/golaypdf"
)

const float64EqualityThreshold = 1e-9

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}

type Text struct {
	Text string

	LineHt          float64
	HorizontalAlign string
}

func (c Text) Measure(context golaypdf.MeasureContext, width float64) (height float64, render golaypdf.Renderer) {
	lineHt := c.LineHt

	if almostEqual(lineHt, 0.0) {
		_, unitSize := context.PDF().GetFontSize()
		lineHt = unitSize * 1.4
	}

	descriptor := context.PDF().GetFontDesc("", "")
	baseline := float64(descriptor.Ascent) / float64(descriptor.Ascent-descriptor.Descent)

	textList := context.PDF().SplitText(c.Text, width)
	height = float64(len(textList)) * lineHt

	return height, golaypdf.FuncToRenderer(func(context golaypdf.RenderContext, x, y, w, h float64) {
		lineY := y + lineHt*baseline

		for _, textLine := range textList {
			var lineX float64

			switch c.HorizontalAlign {
			default:
				fallthrough
			case "L":
				lineX = x
			case "C":
				lineX = x + w/2 - context.PDF().GetStringWidth(textLine)/2
			case "R":
				lineX = x + w - context.PDF().GetStringWidth(textLine)
			}

			context.PDF().Text(lineX, lineY, textLine)
			lineY += lineHt
		}
	})
}
