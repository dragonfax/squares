package glt

import (
	"path"
	"runtime"

	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

const WINDOW_WIDTH = 800
const WINDOW_HEIGHT = 600

var renderer *sdl.Renderer
var font *ttf.Font

var windowConstraints = constraints{
	minWidth:  WINDOW_WIDTH,
	minHeight: WINDOW_HEIGHT,
	maxWidth:  WINDOW_WIDTH,
	maxHeight: WINDOW_HEIGHT,
}

func initRender() {

	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		panic(err)
	}
	// defer sdl.Quit()

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		WINDOW_WIDTH, WINDOW_HEIGHT, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}

	renderer, err = sdl.CreateRenderer(window, -1, 0)
	if err != nil {
		panic(err)
	}
	//defer window.Destroy()

	err = ttf.Init()
	if err != nil {
		panic(err)
	}
	_, file, _, _ := runtime.Caller(0)
	font, err = ttf.OpenFont(path.Dir(file)+"/fonts/OpenSans-Regular.ttf", 12)
	if err != nil {
		panic(err)
	}
}

func render(rootWidget coreWidget) {

	renderer.SetDrawColor(0, 0, 0, 255)
	renderer.Clear()

	rootWidget.render(Offset{0, 0}, renderer)

	renderer.Present()
}

func RunApp(app Widget) error {

	initRender()

	fps := &gfx.FPSmanager{}
	gfx.InitFramerate(fps)
	gfx.SetFramerate(fps, 60)

	running := true
	for running {

		w, err := buildRenderTree(app)
		if err != nil {
			return err
		}

		cw := w.(coreWidget)
		err = cw.layout(windowConstraints)
		if err != nil {
			return err
		}

		gfx.FramerateDelay(fps)

		render(cw)

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event := event.(type) {
			case *sdl.QuitEvent:
				running = false
				break
			case *sdl.KeyboardEvent:
				if event.Type == sdl.KEYDOWN && event.Keysym.Sym == sdl.K_q {
					running = false
					break
				}
			}
		}
	}

	return nil
}

func buildRenderTree(w Widget) (Widget, error) {

	/* Either you have a Build, or children, or nothing */

	if b, ok := w.(hasBuild); ok {
		w2, err := b.Build()
		if err != nil {
			return nil, err
		}
		return buildRenderTree(w2)
	}

	if p, ok := w.(hasChild); ok {
		child, err := buildRenderTree(p.getChild())
		if err != nil {
			return nil, err
		}
		p.setChild(child)
	} else if p, ok := w.(hasChildren); ok {
		children := p.getChildren()
		newChildren := make([]Widget, len(children))
		for i, c := range children {
			nc, err := buildRenderTree(c)
			if err != nil {
				return nil, err
			}
			newChildren[i] = nc
		}
		p.setChildren(newChildren)
	}

	return w, nil
}
