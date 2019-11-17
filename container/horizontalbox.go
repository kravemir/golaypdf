package container

import . "github.com/kravemir/golaypdf"

type HorizontalBoxItem struct {
	Content FixedWidthMeasurable
	Width   float64
	Grow    float64
}

type HorizontalBox struct {
	Columns []HorizontalBoxItem
}

func (h HorizontalBox) Measure(context Context, width float64) (height float64, render Renderer) {
	type cellType struct {
		width, height float64
		renderer      Renderer
	}

	colCount := len(h.Columns)

	var (
		cellList = make([]cellType, colCount, colCount)
		cell     cellType

		maxHeight float64 = 0
	)

	totalFixed := 0.0
	totalWeight := 0.0000001
	for _, column := range h.Columns {
		totalFixed += column.Width
		totalWeight += column.Grow
	}

	baseExtra := (width - totalFixed) / totalWeight

	for idx, column := range h.Columns {
		cell.width = column.Width + baseExtra*column.Grow
		cell.height, cell.renderer = column.Content.Measure(context, cell.width)
		if cell.height > maxHeight {
			maxHeight = cell.height
		}
		cellList[idx] = cell
	}

	return maxHeight, FuncToRenderer(func(context Context, x, y, w, h float64) {
		for _, cell := range cellList {
			cell.renderer.Render(context, x, y, cell.width, maxHeight)
			x += cell.width
		}
	})
}
