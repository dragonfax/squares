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

	if err := sdl.Init(sdl.INIT_VIDEO | sdl.INIT_EVENTS | sdl.INIT_TIMER); err != nil {
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

	var rootElement Element

	running := true
	for running {

		var err error
		rootElement, err = buildElementTree(app, rootElement)
		if err != nil {
			return err
		}

		err = rootElement.layout(windowConstraints)
		if err != nil {
			return err
		}

		gfx.FramerateDelay(fps)

		render(rootElement)

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
			case *sdl.MouseWheelEvent:
				if mouseWheelCallback != nil {
					if event.Y > 0 {
						mouseWheelCallback(MOUSEWHEEL_UP)
					} else if event.Y < 0 {
						mouseWheelCallback(MOUSEWHEEL_DOWN)
					}
				}
			}
		}
	}

	return nil
}

func elementFromStatelessWidget(sw StatelessWidget, currentElement Element) (Element, error) {
	builtWidget, err := sw.Build()
	if err != nil {
		return nil, err
	}
	// we don't keep an element for stateless widgets.
	return buildElementTree(builtWidget, currentElement)
}

func elementFromStatefulWidget(sw StatefulWidget, currentElement Element) (Element, error) {

	// reuse an existing state?
	var state State
	ce, ok := currentElement.(*StatefulElement)
	if ok && ce.state != nil && reflect.TypeOf(sw) == reflect.TypeOf(ce.widget) {
		state = ce.state
	} else if !ok {
		println("creating state, currentElement was not a StatefulElement")
		state = sw.CreateState()
	} else if ce.state == nil {
		println("creating state, state was nil")
		state = sw.CreateState()
	} else {
		println(fmt.Sprintf("creating state, types didn't match %T vs %T", sw, ce.widget))
		state = sw.CreateState()
	}

	e := &StatefulElement{widget: sw, state: state}
	childElement, err := buildElementTree(state, e)
	if err != nil {
		return nil, err
	}
	e.child = childElement
	return e, nil

}

func elementFromElementWidget(ew ElementWidget, currentElement Element) (Element, error) {
	// you're a concrete widget, you may (or may not) have children.

	// TODO wasteful if we won't use it.
	newElement := ew.createElement()

	if parentWidget, ok := ew.(HasChild); ok {

		// get the next child to pass down.
		var childElement Element
		if currentElement != nil && reflect.TypeOf(newElement) == reflect.TypeOf(currentElement) {
			childElement = currentElement.(HasChildElement).getChildElement()
		}

		childWidget := parentWidget.getChild()
		newChild, err := buildElementTree(childWidget, childElement)
		if err != nil {
			return nil, err
		}
		newElement.(HasChildElement).setChildElement(newChild)
	} else if parentWidget, ok := ew.(HasChildren); ok {

		var oldChildrenElements []Element
		if currentElement != nil && reflect.TypeOf(newElement) == reflect.TypeOf(currentElement) {
			oldChildrenElements = currentElement.(HasChildrenElements).getChildrenElements()
		}

		children := parentWidget.getChildren()
		newChildrenElements := make([]Element, len(children))
		for i, c := range children {

			var oldChildElement Element
			if oldChildrenElements != nil {
				oldChildElement = oldChildrenElements[i]
			}

			nc, err := buildElementTree(c, oldChildElement)
			if err != nil {
				return nil, err
			}
			newChildrenElements[i] = nc
		}
		newElement.(HasChildrenElements).setChildrenElements(newChildrenElements)
	}
	return newElement, nil

}

func buildElementTree(w Widget, currentElement Element) (Element, error) {

	if w == nil {
		return nil, errors.New("widget was nil.")
	}

	if reflect.ValueOf(w).Kind() != reflect.Ptr {
		return nil, errors.New(fmt.Sprintf("widget in tree is not a pointer, type %T, value %v", w, w))
	}

	if b, ok := w.(StatelessWidget); ok {
		return elementFromStatelessWidget(b, currentElement)
	} else if sw, ok := w.(StatefulWidget); ok {
		return elementFromStatefulWidget(sw, currentElement)
	} else if ew, ok := w.(ElementWidget); ok {
		return elementFromElementWidget(ew, currentElement)
	} else {
		return nil, errors.New(fmt.Sprintf("unknown widget type in tree, type %T, value %v", w, w))
	}
}
