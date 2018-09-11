package glt

import "github.com/veandco/go-sdl2/sdl"

var _ StatelessWidget = &Column{}
var _ HasChildren = &Column{}
var _ StatelessWidget = &Row{}
var _ HasChildren = &Row{}
var _ ElementWidget = &Expanded{}
var _ HasChild = &Expanded{}
var _ ElementWidget = &Flex{}
var _ HasChildren = &Flex{}

type CrossAxisAlignment uint8

const (
	CrossAxisAlignmentStart CrossAxisAlignment = iota
)

type MainAxisAlignment uint8

const (
	MainAxisAlignmentSpaceBetween MainAxisAlignment = iota
)

type AxisDirection uint8

const (
	Vertical AxisDirection = iota
	Horizontal
)

type Flex struct {
	Axis               AxisDirection
	CrossAxisAlignment CrossAxisAlignment
	MainAxisAlignment  MainAxisAlignment
	Children           []Widget
	// MainAxisSize.max
	// VerticalDirection.down
}

func (c *Flex) createElement() Element {
	ce := &FlexElement{}
	ce.widget = c
	return ce
}

type FlexElement struct {
	elementData
	childrenElementsData
}

func (ce *FlexElement) layout(c Constraints) error {

	ce.size = Size{0, 0}

	numChildren := uint16(len(ce.children))

	for _, child := range ce.children {

		// TODO not sure about this.
		// might need to do them one at a time. and see whats left for the others.
		con := Constraints{
			minWidth:  c.maxWidth,
			minHeight: c.minHeight / numChildren,
			maxWidth:  c.maxWidth,
			maxHeight: c.maxHeight / numChildren,
		}

		child.layout(con)

		childSize := child.getSize()
		ce.size.width = MaxUint16(childSize.width, ce.size.width)
		offsetHeight := ce.size.height
		ce.size.height += childSize.height

		child.setOffset(Offset{
			x: 0,
			y: offsetHeight,
		})

	}

	return nil
}

func (c *FlexElement) render(offset Offset, renderer *sdl.Renderer) {

	for _, child := range c.children {
		child.render(offset, renderer)
		offset.y += child.getSize().height
	}
}

type Row struct {
	CrossAxisAlignment CrossAxisAlignment
	MainAxisAlignment  MainAxisAlignment
	Children           []Widget
}

type Column struct {
	CrossAxisAlignment CrossAxisAlignment
	MainAxisAlignment  MainAxisAlignment
	Children           []Widget
}
type Expanded struct {
	Child Widget
	// FlexFit.tight
}
