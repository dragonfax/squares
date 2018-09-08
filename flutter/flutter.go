package flutter

import (
	"math"

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
	All uint16
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

type Offset struct {
	x, y uint16
}

type Size struct {
	width, height uint16
}

type ParentData struct {
	offset Offset
}

type CoreWidget interface {
	layout(c Constraints) error
	getParentData() *ParentData
	getSize() Size
}

// use MaxUint32 for +Inf during layout
type Constraints struct {
	minWidth, minHeight, maxWidth, maxHeight uint16
}

func (s Size) addMargin(in EdgeInsets) Size {
	return Size{width: s.width + in.All, height: s.height + in.All}
}

func (c Constraints) addMargins(in EdgeInsets) Constraints {
	// TODO fix the math here
	if c.minWidth > in.All {
		c.minWidth -= in.All
	} else {
		c.minWidth = 0
	}

	if c.minHeight > in.All {
		c.minHeight -= in.All
	} else {
		c.minHeight = 0
	}

	if c.maxWidth != math.MaxUint16 {
		c.maxWidth -= in.All
	}
	if c.maxHeight != math.MaxUint16 {
		c.maxHeight -= in.All
	}
	return c
}

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

func layoutTree(root Widget) {

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
