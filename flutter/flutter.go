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
	Child() Widget
	SetChild(Widget)
}

type HasChildren interface {
	Children() []Widget
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

type Text struct {
	Text string
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
		child, err := processTree(context, p.Child())
		if err != nil {
			return nil, err
		}
		p.SetChild(child)
	} else if p, ok := w.(HasChildren); ok {
		children := p.Children()
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
