package element

import . "github.com/kravemir/golaypdf"

type Text struct {
	Text string

	LineHt float64
	Align  string
}

func (c Text) Measure(context Context, width float64) (height float64, render Renderer) {
	textList := context.PDF().SplitText(c.Text, width)
	height = float64(len(textList)) * c.LineHt

	return height, FuncToRenderer(func(context Context, x, y, w, h float64) {
		cellY := y + (h-height)/2

		for _, textLine := range textList {
			context.PDF().SetXY(x, cellY)
			context.PDF().CellFormat(w, c.LineHt, textLine, "", 0, c.Align, false, 0, "")
			cellY += c.LineHt
		}
	})
}
