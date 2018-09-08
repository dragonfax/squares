package flutter

import "github.com/davecgh/go-spew/spew"

const WINDOW_WIDTH = 800
const WINDOW_HEIGHT = 600

func RunApp(w Widget) error {
	// var renderer *sdl.Renderer
	/*
		if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
			panic(err)
		}
		defer sdl.Quit()

		window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
			WINDOW_WIDTH, WINDOW_HEIGHT, sdl.WINDOW_SHOWN)
		if err != nil {
			panic(err)
		}
		defer window.Destroy()
	*/

	context := &BuildContext{}

	w, err := buildTree(context, w)
	if err != nil {
		return err
	}

	/*
		if d, ok := w.(HasRender); ok {
			d.Render(renderer)
		}
	*/

	windowConstraints := Constraints{
		minWidth:  WINDOW_WIDTH,
		minHeight: WINDOW_HEIGHT,
		maxWidth:  WINDOW_WIDTH,
		maxHeight: WINDOW_HEIGHT,
	}
	cw := w.(CoreWidget)
	err = cw.layout(windowConstraints)
	if err != nil {
		return err
	}

	spew.Dump(w)

	return nil
}

func buildTree(context *BuildContext, w Widget) (Widget, error) {

	/* Either you have a Build, or children, or nothing */

	if b, ok := w.(HasBuild); ok {
		w2, err := b.Build(context)
		if err != nil {
			return nil, err
		}
		return buildTree(context, w2)
	}

	if p, ok := w.(HasChild); ok {
		child, err := buildTree(context, p.GetChild())
		if err != nil {
			return nil, err
		}
		p.SetChild(child)
	} else if p, ok := w.(HasChildren); ok {
		children := p.GetChildren()
		newChildren := make([]Widget, len(children))
		for i, c := range children {
			nc, err := buildTree(context, c)
			if err != nil {
				return nil, err
			}
			newChildren[i] = nc
		}
		p.SetChildren(newChildren)
	}

	return w, nil
}
