package golaypdf

import "github.com/jung-kurt/gofpdf"

type RenderContext interface {
	PDF() *gofpdf.Fpdf

	GetFont() (family, style string, size float64)
	SetFont(family, style string, size float64)
	SetFontStyle(style string)
	SetFontSize(size float64)
}

type Renderer interface {
	Render(ctx RenderContext, x, y, w, h float64)
}

type RendererFunc func(context RenderContext, x, y, w, h float64)

func (r RendererFunc) Render(ctx RenderContext, x, y, w, h float64) {
	r(ctx, x, y, w, h)
}

func FuncToRenderer(f RendererFunc) Renderer {
	return f
}
