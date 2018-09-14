package squares

import "github.com/veandco/go-sdl2/sdl"

var _ HasChild = &Composite{}
var _ ElementWidget = &Composite{}
var _ Element = &CompositeElement{}
var _ HasChildElement = &CompositeElement{}

/* caches the render of its children in a texture */
type Composite struct {
	Child Widget
}

func (cw *Composite) createElement() Element {
	e := &CompositeElement{}
	e.widget = cw
	return e
}

type CompositeElement struct {
	elementData
	childElementData
	rendered        bool
	renderedSize    Size
	renderedTexture *sdl.Texture
}

func (se *CompositeElement) getSize() Size {
	return se.child.getSize()
}

func (se *CompositeElement) layout(c Constraints) error {
	return se.child.layout(c)
}

func (se *CompositeElement) updateWidget(widget Widget) {
	se.rendered = false
	se.elementData.updateWidget(widget)
}

func (se *CompositeElement) render(o Offset, r *sdl.Renderer) {

	size := se.child.getSize()

	if size.height == 0 || size.width == 0 {
		panic("can't composite a zero size child")
	}

	if !se.rendered {
		// create a new composite

		if se.renderedSize != size {
			// reclaim the texture and create a new one of the right size
			if se.renderedTexture != nil {
				se.renderedTexture.Destroy()
			}
			t, err := r.CreateTexture(sdl.PIXELFORMAT_RGBA8888, sdl.TEXTUREACCESS_TARGET, int32(size.width), int32(size.height))
			if err != nil {
				panic(err)
			}
			se.renderedTexture = t
		}

		// render to the texture.
		prevTarget := r.GetRenderTarget()
		r.SetRenderTarget(se.renderedTexture)
		r.SetDrawColor(0, 0, 0, 255)
		r.Clear()
		se.child.render(Offset{0, 0}, r)
		// r.Present()
		r.SetRenderTarget(prevTarget)

		se.renderedSize = size
		se.rendered = true
	}

	// use the composite
	srcRect := &sdl.Rect{X: 0, Y: 0, W: int32(size.width), H: int32(size.height)}
	dstRect := &sdl.Rect{X: int32(o.x), Y: int32(o.y), W: int32(size.width), H: int32(size.height)}
	r.Copy(se.renderedTexture, srcRect, dstRect)
}
