package golaypdf

import (
	"fmt"
	"io"

	"github.com/jung-kurt/gofpdf"
)

type MeasureContext interface {
	// TODO: expose TextMeasurer instead of mutable PDF here
	PDF() *gofpdf.Fpdf

	GetFont() (family, style string, size float64)
	SetFont(family, style string, size float64)
	SetFontStyle(style string)
	SetFontSize(size float64)
}

type FixedWidthMeasurable interface {
	Measure(context MeasureContext, width float64) (height float64, render Renderer)
}

type PdfBuilder struct {
	pdf *gofpdf.Fpdf

	fontBytesProvider FontLoader
	loadedFonts       map[string]struct{}

	fontFamily, fontStyle string
	fontSize              float64
}

func NewPdfBuilder(
	orientationStr, unitStr, sizeStr string,
	left, top, right float64,
	fontBytesProvider FontLoader,
) *PdfBuilder {
	pdf := gofpdf.New(orientationStr, unitStr, sizeStr, "")
	pdf.SetMargins(left, top, right)
	pdf.AddPage()

	return &PdfBuilder{
		pdf:               pdf,
		fontBytesProvider: fontBytesProvider,
		loadedFonts:       map[string]struct{}{},
		fontFamily:        "",
		fontStyle:         "",
		fontSize:          0,
	}
}

func (builder *PdfBuilder) GetFont() (family, style string, size float64) {
	return builder.fontFamily, builder.fontStyle, builder.fontSize
}

func (builder *PdfBuilder) SetFont(family, style string, size float64) {
	err := builder.ensureFontLoaded(family, style)
	if err != nil {
		builder.PDF().SetError(err)
		return
	}

	builder.fontFamily, builder.fontStyle, builder.fontSize = family, style, size
	builder.setFontToPDF()
}

func (builder *PdfBuilder) SetFontStyle(style string) {
	err := builder.ensureFontLoaded(builder.fontFamily, style)
	if err != nil {
		builder.PDF().SetError(err)
		return
	}

	builder.fontStyle = style
	builder.setFontToPDF()
}

func (builder *PdfBuilder) ensureFontLoaded(family string, style string) error {
	key := fontKey(family, style)

	if _, ok := builder.loadedFonts[key]; !ok {
		// TODO: support fonts aliases with same identity, no duplicate embedding
		// identity, fontFile, err := builder.fontBytesProvider.OpenFont(family, style)

		_, fontFile, err := builder.fontBytesProvider.OpenFont(family, style)
		if err != nil {
			return fmt.Errorf("get font: %w", err)
		}
		defer fontFile.Close()

		fontBytes, err := io.ReadAll(fontFile)
		if err != nil {
			return fmt.Errorf("read font: %w", err)
		}

		builder.PDF().AddUTF8FontFromBytes(key, "", fontBytes)
	}

	return nil
}

func (builder *PdfBuilder) setFontToPDF() {
	builder.pdf.SetFont(fontKey(builder.fontFamily, builder.fontStyle), "", builder.fontSize)
}

func (builder *PdfBuilder) SetFontSize(size float64) {
	builder.fontSize = size
	builder.pdf.SetFontSize(size)
}

func (builder *PdfBuilder) PDF() *gofpdf.Fpdf {
	return builder.pdf
}

func (builder *PdfBuilder) Render(block FixedWidthMeasurable) {
	leftMargin, _, rightMargin, _ := builder.pdf.GetMargins()

	width, _ := builder.pdf.GetPageSize()
	width = width - leftMargin - rightMargin

	height, renderer := block.Measure(builder, width)

	x := leftMargin
	y := builder.pdf.GetY()
	renderer.Render(builder, x, y, width, height)

	builder.pdf.SetY(y + height)
}

func fontKey(family string, style string) string {
	return family + "___" + style
}
