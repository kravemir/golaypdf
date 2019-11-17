package golaypdf

import "github.com/jung-kurt/gofpdf"

type Context interface {
	PDF() *gofpdf.Fpdf

	GetFont() (family, style string, size float64)
	SetFont(family, style string, size float64)
	SetFontSize(size float64)
}

type FixedWidthMeasurable interface {
	Measure(context Context, width float64) (height float64, render func(context Context, x, y, w, h float64))
}

type BorderBox struct {
	Content                      FixedWidthMeasurable
	PTop, PRight, PLeft, PBottom float64
}

func (b BorderBox) Measure(context Context, width float64) (height float64, render func(context Context, x, y, w, h float64)) {
	contentWidth := width - b.PLeft - b.PRight
	contentHeight, contentRender := b.Content.Measure(context, contentWidth)

	return contentHeight + b.PTop + b.PBottom, func(context Context, x, y, w, h float64) {
		context.PDF().Rect(x, y, w, h, "D")

		contentRender(context, x+b.PLeft, y+b.PTop, contentWidth, contentHeight)
	}
}

type Text struct {
	Text string

	LineHt float64
	Align  string
}

func (c Text) Measure(context Context, width float64) (height float64, render func(context Context, x, y, w, h float64)) {
	textList := context.PDF().SplitText(c.Text, width)
	height = float64(len(textList)) * c.LineHt

	return height, func(context Context, x, y, w, h float64) {
		cellY := y + (h-height)/2

		for _, textLine := range textList {
			context.PDF().SetXY(x, cellY)
			context.PDF().CellFormat(w, c.LineHt, textLine, "", 0, c.Align, false, 0, "")
			cellY += c.LineHt
		}
	}
}

type PdfBuilder struct {
	Pdf     *gofpdf.Fpdf
	MarginH float64

	fontFamily, fontStyle string
	fontSize              float64
}

func (builder *PdfBuilder) GetFont() (family, style string, size float64) {
	return builder.fontFamily, builder.fontStyle, builder.fontSize
}

func (builder *PdfBuilder) SetFont(family, style string, size float64) {
	builder.fontFamily, builder.fontStyle, builder.fontSize = family, style, size
	builder.Pdf.SetFont(family, style, size)
}

func (builder *PdfBuilder) SetFontSize(size float64) {
	builder.fontSize = size
	builder.Pdf.SetFontSize(size)
}

func (builder *PdfBuilder) PDF() *gofpdf.Fpdf {
	return builder.Pdf
}

func (builder *PdfBuilder) Render(block FixedWidthMeasurable) {
	width, _ := builder.Pdf.GetPageSize()
	width = width - builder.MarginH - builder.MarginH

	height, render := block.Measure(builder, width)

	x := builder.MarginH
	y := builder.Pdf.GetY()
	render(builder, x, y, width, height)

	builder.Pdf.SetY(y + height)
}

type HorizontalBoxItem struct {
	Content FixedWidthMeasurable
	Width   float64
	Grow    float64
}

type HorizontalBox struct {
	Columns []HorizontalBoxItem
}

func (h HorizontalBox) Measure(context Context, width float64) (height float64, render func(context Context, x, y, w, h float64)) {
	type cellType struct {
		width, height float64
		render        func(context Context, x, y, w, h float64)
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
		cell.height, cell.render = column.Content.Measure(context, cell.width)
		if cell.height > maxHeight {
			maxHeight = cell.height
		}
		cellList[idx] = cell
	}

	return maxHeight, func(context Context, x, y, w, h float64) {
		for _, cell := range cellList {
			cell.render(context, x, y, cell.width, maxHeight)
			x += cell.width
		}
	}
}

type VerticalBox struct {
	Rows []FixedWidthMeasurable
}

func (v VerticalBox) Measure(context Context, width float64) (height float64, render func(context Context, x, y, w, h float64)) {
	type itemType struct {
		height float64
		render func(context Context, x, y, w, h float64)
	}

	var item itemType
	items := make([]itemType, len(v.Rows))
	height = 0.0

	for idx, row := range v.Rows {
		item.height, item.render = row.Measure(context, width)
		height += item.height
		items[idx] = item
	}

	return height, func(context Context, x, y, w, h float64) {
		for _, row := range items {
			row.render(context, x, y, w, row.height)
			y += row.height
		}
	}
}

type Empty struct {
	Width, Height float64
}

func (e Empty) Measure(context Context, width float64) (height float64, render func(context Context, x, y, w, h float64)) {
	return e.Height, func(context Context, x, y, w, h float64) {}
}
