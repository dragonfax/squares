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

	rootElement, err := buildElementTree(app, nil)
	if err != nil {
		return err
	}

	running := true
	for running {

		err = rebuildDirty(rootElement)
		if err != nil {
			return err
		}

		err = rootElement.layout(windowConstraints)
		if err != nil {
			return err
		}

		floatUpRendered(rootElement)

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

/* recurse the element tree, rebuilding any subtree that is marked dirty */
func rebuildDirty(element Element) error {
	statefulElement, ok := element.(*StatefulElement)
	if ok && !statefulElement.built {
		// this follows the children implicitely
		newElement, err := buildElementTree(element.GetWidget(), element)
		if err != nil {
			return err
		}
		if statefulElement != newElement.(*StatefulElement) {
			panic("statefulelement changed element during dirty rebuild")
		}
		statefulElement.rendered = false
	} else {
		// just follow the children.
		var children = getElementChildren(element)
		for _, child := range children {
			err := rebuildDirty(child)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func floatUpRendered(element Element) bool {
	compositeElement, ok := element.(*CompositeElement)

	var children = getElementChildren(element)
	var rendered = true
	for _, child := range children {
		r := floatUpRendered(child)
		if !r {
			if ok {
				compositeElement.rendered = false
			}
			rendered = false
		}
	}
	if rendered == false {
		return rendered
	}

	statefulElement, ok2 := element.(*StatefulElement)
	if ok && !compositeElement.rendered {
		return false
	} else if ok2 && !statefulElement.rendered {
		statefulElement.rendered = true
		return false
	} else {
		return true
	}
}

func elementFromStatelessWidget(sw StatelessWidget, oldElement Element) (Element, error) {
	oldStatelessElement, ok := oldElement.(*StatelessElement)

	// reusing the existing element
	var element *StatelessElement
	if ok && sameType(sw, oldStatelessElement.GetWidget()) {
		element = oldStatelessElement
		element.updateWidget(sw)
	} else {
		element = NewStatelessElement(sw)
	}

	// rebuild the widget
	builtWidget, err := sw.Build(element)
	if err != nil {
		return nil, err
	}

	// check on the children
	var oldChildElement Element
	if oldStatelessElement != nil {
		oldChildElement = oldStatelessElement.child
	}

	childElement, err := buildElementTree(builtWidget, oldChildElement)
	if err != nil {
		return nil, err
	}
	element.child = childElement

	return element, nil
}

func elementFromStatefulWidget(widget StatefulWidget, oldElement Element) (Element, error) {
	oldStatefulElement, ok := oldElement.(*StatefulElement)

	// reusing the existing element & state?
	var element *StatefulElement
	if ok && sameType(widget, oldStatefulElement.GetWidget()) {
		element = oldStatefulElement
		element.updateWidget(widget)
	} else {
		println("new element, creating new state")
		state := widget.CreateState()
		element = NewStatefulElement(widget, state)
	}

	// build the state
	childWidget, err := element.state.Build(element)
	if err != nil {
		return nil, err
	}

	// build the child subtree
	var oldChildElement Element
	if oldStatefulElement != nil {
		oldChildElement = oldStatefulElement.child
	}
	childElement, err := buildElementTree(childWidget, oldChildElement)
	if err != nil {
		return nil, err
	}
	element.child = childElement
	element.built = true
	return element, nil
}

func sameType(a, b interface{}) bool {
	return reflect.TypeOf(a) == reflect.TypeOf(b)
}

func getWidgetChildren(ew ElementWidget) []Widget {
	if parent, ok := ew.(HasChild); ok {
		return []Widget{parent.getChild()}
	} else if parent, ok := ew.(HasChildren); ok {
		return parent.getChildren()
	}
	return []Widget{}
}

func setElementChildren(e Element, ec []Element) {
	if parent, ok := e.(HasChildElement); ok {
		if len(ec) > 1 {
			panic("unhandleable")
		}
		if len(ec) == 1 {
			parent.setChildElement(e, ec[0])
		}
	} else if parent, ok := e.(HasChildrenElements); ok {
		parent.setChildrenElements(e, ec)
	}
}

func getElementChildren(e Element) []Element {
	if parent, ok := e.(HasChildElement); ok {
		return []Element{parent.getChildElement()}
	} else if parent, ok := e.(HasChildrenElements); ok {
		return parent.getChildrenElements()
	}
	return []Element{}
}

func processElementChildren(widget ElementWidget, newElement Element, oldElement Element) error {

	widgetChildren := getWidgetChildren(widget)
	newElementChildren := make([]Element, len(widgetChildren))
	oldElementChildren := getElementChildren(oldElement)

	for i, widgetChild := range widgetChildren {

		// check if we have an old element to reuse (for keeping state)
		var oldChildElement Element
		if len(oldElementChildren) > i && sameType(widgetChild, oldElementChildren[i].GetWidget()) {
			oldChildElement = oldElementChildren[i]
		}

		newChildElement, err := buildElementTree(widgetChild, oldChildElement)
		if err != nil {
			return err
		}
		newElementChildren[i] = newChildElement
	}
	setElementChildren(newElement, newElementChildren)
	return nil
}

func elementFromElementWidget(ew ElementWidget, oldElement Element) (Element, error) {

	// reusing the existing element & state?
	var element Element
	if oldElement != nil && sameType(ew, oldElement.GetWidget()) {
		element = oldElement
		element.updateWidget(ew)
	} else {
		element = ew.createElement()
	}

	err := processElementChildren(ew, element, oldElement)
	return element, err
}

func buildElementTree(w Widget, oldElement Element) (Element, error) {

	if w == nil {
		return nil, errors.New("widget was nil.")
	}

	if reflect.ValueOf(w).Kind() != reflect.Ptr {
		return nil, errors.New(fmt.Sprintf("widget in tree is not a pointer, type %T, value %v", w, w))
	}

	if b, ok := w.(StatelessWidget); ok {
		return elementFromStatelessWidget(b, oldElement)
	} else if sw, ok := w.(StatefulWidget); ok {
		return elementFromStatefulWidget(sw, oldElement)
	} else if ew, ok := w.(ElementWidget); ok {
		return elementFromElementWidget(ew, oldElement)
	} else {
		return nil, errors.New(fmt.Sprintf("unknown widget type in tree, type %T, value %v", w, w))
	}
}

func ContextOf(context BuildContext, typeOf interface{}) BuildContext {
	element := context.(Element)
	for parentContext := element.getParentElement(); parentContext != nil; parentContext = element.getParentElement() {
		widget := parentContext.GetWidget()
		if sameType(widget, typeOf) {
			return parentContext
		}
	}
	return nil
}
