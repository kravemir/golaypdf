package container

import (
	"github.com/kravemir/golaypdf"
)

type TableCell struct {
	Content golaypdf.FixedWidthMeasurable
}

type TableRow struct {
	Cells []TableCell
}

type TableColumn struct {
	Width float64
	Grow  float64
}

type Table struct {
	Columns []TableColumn
	Rows    []TableRow
}

func (t Table) Measure(context golaypdf.MeasureContext, width float64) (height float64, render golaypdf.Renderer) {
	type MeasuredColumn struct {
		width  float64
		height float64
	}
	type MeasuredCell struct {
		height   float64
		renderer golaypdf.Renderer
	}
	type MeasuredRow struct {
		height float64
		cells  []MeasuredCell
	}

	measuredColumns := make([]MeasuredColumn, len(t.Columns), len(t.Columns))
	measuredRows := make([]MeasuredRow, len(t.Rows), len(t.Rows))

	for i := 0; i < len(t.Rows); i++ {
		measuredRows[i] = MeasuredRow{
			height: 0,
			cells:  make([]MeasuredCell, len(t.Columns), len(t.Columns)),
		}
	}

	totalWidth := 0.0
	totalWeight := 0.000001
	for _, column := range t.Columns {
		totalWidth += column.Width
		totalWeight += column.Grow
	}

	maxHeight := 0.0
	baseGrowSize := (width - totalWidth) / totalWeight
	for columnIdx, column := range t.Columns {
		columnWidth := column.Width + baseGrowSize*column.Grow
		columnHeight := 0.0

		for rowIdx, row := range t.Rows {
			cellHeight, cellRender := row.Cells[columnIdx].Content.Measure(context, columnWidth)

			columnHeight += cellHeight

			if cellHeight > measuredRows[rowIdx].height {
				measuredRows[rowIdx].height = cellHeight
			}

			measuredRows[rowIdx].cells[columnIdx] = MeasuredCell{
				height:   cellHeight,
				renderer: cellRender,
			}
		}

		if columnHeight > maxHeight {
			maxHeight = columnHeight
		}

		measuredColumns[columnIdx] = MeasuredColumn{
			width:  columnWidth,
			height: columnHeight,
		}
	}

	return maxHeight, golaypdf.FuncToRenderer(func(context golaypdf.RenderContext, base_x, y, w, h float64) {
		for _, measuredRow := range measuredRows {
			x := base_x
			h := measuredRow.height

			for columnIdx, columnCell := range measuredRow.cells {
				w := measuredColumns[columnIdx].width

				columnCell.renderer.Render(context, x, y, w, h)

				x = x + w
			}

			y = y + h
		}
	})
}
