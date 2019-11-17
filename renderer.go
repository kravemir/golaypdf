package golaypdf

type Renderer interface {
	Render(context Context, x, y, w, h float64)
}

type RendererFunc func(context Context, x, y, w, h float64)

func (r RendererFunc) Render(context Context, x, y, w, h float64) {
	r(context, x, y, w, h)
}

func FuncToRenderer(f RendererFunc) Renderer {
	return f
}
