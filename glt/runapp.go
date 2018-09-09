package glt

import (
	"errors"
	"fmt"
	"path"
	"reflect"
	"runtime"

	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

const WINDOW_WIDTH = 800
const WINDOW_HEIGHT = 600

var renderer *sdl.Renderer
var font *ttf.Font

var windowConstraints = Constraints{
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

func render(rootElement Element) {

	renderer.SetDrawColor(0, 0, 0, 255)
	renderer.Clear()

	rootElement.render(Offset{0, 0}, renderer)

	renderer.Present()
}

func RunApp(app Widget) error {

	initRender()

	fps := &gfx.FPSmanager{}
	gfx.InitFramerate(fps)
	gfx.SetFramerate(fps, 60)

	running := true
	for running {

		elementTree, err := buildElementTree(app, nil)
		if err != nil {
			return err
		}

		err = elementTree.layout(windowConstraints)
		if err != nil {
			return err
		}

		gfx.FramerateDelay(fps)

		render(elementTree)

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

func buildElementTree(w Widget, currentElement Element) (Element, error) {

	/* Concrete widget (possibly with children) or non-concrete widget (you have a Build method) */

	if b, ok := w.(StatelessWidget); ok {
		w2, err := b.Build()
		if err != nil {
			return nil, err
		}
		// we don't keep an element for stateless widgets.
		return buildElementTree(w2, currentElement)
	} else if sw, ok := w.(StatefulWidget); ok {

		// reuse an existing state?
		var state State
		ce, ok := currentElement.(StatefulElement)
		if ok && ce.state != nil && reflect.TypeOf(w) == reflect.TypeOf(ce.widget) {
			state = ce.state
		} else {
			state = sw.CreateState()
		}

		e := StatefulElement{widget: w, state: state}
		childElement, err := buildElementTree(state, e)
		if err != nil {
			return nil, err
		}
		e.child = childElement
		return e, nil
	} else if cw, ok := w.(ElementWidget); ok {
		// you're a concrete widget, you may (or may not) have children.

		e := cw.createElement()

		if p, ok := cw.(HasChild); ok {
			pe := e.(HasChildElement)
			child, err := buildElementTree(p.getChild(), e)
			if err != nil {
				return nil, err
			}
			pe.setChildElement(child)
		} else if p, ok := w.(HasChildren); ok {
			pe := e.(HasChildrenElements)
			children := p.getChildren()
			newChildrenElements := make([]Element, len(children))
			for i, c := range children {
				nc, err := buildElementTree(c, e)
				if err != nil {
					return nil, err
				}
				newChildrenElements[i] = nc
			}
			pe.setChildrenElements(newChildrenElements)
		}
		return e, nil
	} else {
		return nil, errors.New(fmt.Sprintf("unknown widget type in tree, type %T, value %v", w, w))
	}
}
