package glt

import "github.com/veandco/go-sdl2/sdl"

type StatefulElement struct {
	elementData
	state State
	childElementData
	built bool

	// for compositing
	rendered        bool
	renderedSize    Size
	renderedTexture *sdl.Texture
}

var _ Element = &StatefulElement{}
var _ StatefulContext = &StatefulElement{}

func NewStatefulElement(widget Widget, state State) *StatefulElement {
	se := &StatefulElement{}
	se.widget = widget
	se.state = state
	return se
}

func (se *StatefulElement) getSize() Size {
	return se.child.getSize()
}

func (se *StatefulElement) layout(c Constraints) error {
	return se.child.layout(c)
}

func (se *StatefulElement) updateWidget(widget Widget) {
	se.rendered = false
	se.elementData.updateWidget(widget)
}

func (se *StatefulElement) render(o Offset, r *sdl.Renderer) {

	size := se.getSize()

	if !se.rendered || se.renderedSize != size {
		// create a new composite

		println("rendering to texture")

		if se.renderedSize != size {
			// reclaim the texture and create a new one of the right size
			if se.renderedTexture != nil {
				println("destroying texture")
				se.renderedTexture.Destroy()
			}
			println("creating new texture")
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
		println("rendering")
		se.child.render(Offset{0, 0}, r)
		// r.Present()
		r.SetRenderTarget(prevTarget)

		se.renderedSize = size
		se.rendered = true
	}

	// use the composite
	println("compositing texture")
	srcRect := &sdl.Rect{X: 0, Y: 0, W: int32(size.width), H: int32(size.height)}
	dstRect := &sdl.Rect{X: int32(o.x), Y: int32(o.y), W: int32(size.width), H: int32(size.height)}
	r.Copy(se.renderedTexture, srcRect, dstRect)
}

func (se *StatefulElement) SetState(callback SetStateFunc) {
	callback()
	se.built = false
	se.rendered = false
}
