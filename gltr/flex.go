package gltr

import "github.com/veandco/go-sdl2/sdl"

var _ StatelessWidget = &Expanded{}
var _ StatelessWidget = &Column{}
var _ StatelessWidget = &Row{}
var _ ElementWidget = &Flex{}
var _ HasChildren = &Flex{}
var _ HasChildrenElements = &FlexElement{}

type CrossAxisAlignment uint8

const (
	CrossAxisAlignmentStart CrossAxisAlignment = iota
)

type MainAxisAlignment uint8

const (
	MainAxisAlignmentSpaceBetween MainAxisAlignment = iota
)

type Axis uint8

const (
	Vertical Axis = iota
	Horizontal
)

type Flex struct {
	Direction          Axis
	CrossAxisAlignment CrossAxisAlignment
	MainAxisAlignment  MainAxisAlignment
	Children           []Widget
	// MainAxisSize.max
	// VerticalDirection.down
}

func (f *Flex) getChildren() []Widget {
	return f.Children
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

func (c *Row) Build(context BuildContext) (Widget, error) {
	return &Flex{
		Direction:          Horizontal,
		CrossAxisAlignment: c.CrossAxisAlignment,
		MainAxisAlignment:  c.MainAxisAlignment,
		Children:           c.Children,
	}, nil
}

type Column struct {
	CrossAxisAlignment CrossAxisAlignment
	MainAxisAlignment  MainAxisAlignment
	Children           []Widget
}

func (c *Column) Build(context BuildContext) (Widget, error) {
	return &Flex{
		Direction:          Vertical,
		CrossAxisAlignment: c.CrossAxisAlignment,
		MainAxisAlignment:  c.MainAxisAlignment,
		Children:           c.Children,
	}, nil
}

/* Expanded seems to do have no real implementation. All the magic is in Flex */
type Expanded struct {
	Child Widget
	// FlexFit.tight
}

func (e *Expanded) Build(context BuildContext) (Widget, error) {
	return e.Child, nil
}
