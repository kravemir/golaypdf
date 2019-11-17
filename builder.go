package golaypdf

import (
	"github.com/jung-kurt/gofpdf"
)

type Context interface {
	PDF() *gofpdf.Fpdf

	GetFont() (family, style string, size float64)
	SetFont(family, style string, size float64)
	SetFontSize(size float64)
}

type FixedWidthMeasurable interface {
	Measure(context Context, width float64) (height float64, render Renderer)
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

	height, renderer := block.Measure(builder, width)

	x := builder.MarginH
	y := builder.Pdf.GetY()
	renderer.Render(builder, x, y, width, height)

	builder.Pdf.SetY(y + height)
}
