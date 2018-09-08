package flutter

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/veandco/go-sdl2/sdl"
)

type BuildContext struct {
}

type Widget interface {
}

type HasBuild interface { // StatelessWidget
	Build(*BuildContext) (Widget, error)
}

type HasChild interface { // Container
	GetChild() Widget
	SetChild(Widget)
}

type HasChildren interface {
	GetChildren() []Widget
	SetChildren([]Widget)
}

type HasRender interface { // RenderObject
	Render(*sdl.Renderer) error
}

type EdgeInsets struct {
	All int
}

type Divider struct {
}

type Center struct {
	Child Widget
}

func (ce *Center) GetChild() Widget {
	return ce.Child
}

func (p *Center) SetChild(c Widget) {
	p.Child = c
}

type Text struct {
	Text string
}

type Column struct {
	Children []Widget
}

func (c *Column) GetChildren() []Widget {
	return c.Children
}

func (c *Column) SetChildren(cs []Widget) {
	c.Children = cs
}

type Padding struct {
	Padding EdgeInsets
	Child   Widget
}

func (p *Padding) GetChild() Widget {
	return p.Child
}

func (p *Padding) SetChild(c Widget) {
	p.Child = c
}

func RunApp(w Widget) error {
	// var renderer *sdl.Renderer
	/*
		if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
			panic(err)
		}
		defer sdl.Quit()

		window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
			800, 600, sdl.WINDOW_SHOWN)
		if err != nil {
			panic(err)
		}
		defer window.Destroy()
	*/

	context := &BuildContext{}

	w, err := processTree(context, w)
	if err != nil {
		return err
	}

	/*
		if d, ok := w.(HasRender); ok {
			d.Render(renderer)
		}
	*/

	spew.Dump(w)

	return nil
}

func processTree(context *BuildContext, w Widget) (Widget, error) {

	/* Either you have a Build, or children, or nothing */

	if b, ok := w.(HasBuild); ok {
		w2, err := b.Build(context)
		if err != nil {
			return nil, err
		}
		return processTree(context, w2)
	}

	if p, ok := w.(HasChild); ok {
		child, err := processTree(context, p.GetChild())
		if err != nil {
			return nil, err
		}
		p.SetChild(child)
	} else if p, ok := w.(HasChildren); ok {
		children := p.GetChildren()
		newChildren := make([]Widget, len(children))
		for i, c := range children {
			nc, err := processTree(context, c)
			if err != nil {
				return nil, err
			}
			newChildren[i] = nc
		}
		p.SetChildren(newChildren)
	}

	return w, nil
}
