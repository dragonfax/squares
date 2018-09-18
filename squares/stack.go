package squares

import "github.com/veandco/go-sdl2/sdl"

var _ ElementWidget = Stack{}
var _ HasChildren = Stack{}
var _ Element = &StackElement{}
var _ HasChildrenElements = &StackElement{}

type StackFit uint8

const (
	StackFitExpand StackFit = iota
)

type Stack struct {
	Fit      StackFit
	Children []Widget
}

func (d Stack) getChildren() []Widget {
	return d.Children
}

func (s Stack) createElement() Element {
	se := &StackElement{}
	se.widget = s
	return se
}

type StackElement struct {
	elementData
	childrenElementsData
}

func (se *StackElement) layout(c Constraints) error {

	size := c.constrain(Size{})
	for _, childElement := range se.children {
		err := childElement.layout(c)
		if err != nil {
			return err
		}
		childSize := childElement.getSize()
		size = size.Expand(childSize)
	}
	se.size = size

	return nil
}

func (se *StackElement) render(o Offset, r *sdl.Renderer) {
	for _, childElement := range se.children {
		childElement.render(o, r)
	}

}
