package gltr

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

var _ StatelessWidget = &Expanded{}
var _ StatelessWidget = &Column{}
var _ StatelessWidget = &Row{}
var _ ElementWidget = &Flex{}
var _ HasChildren = &Flex{}
var _ HasChildrenElements = &FlexElement{}

type CrossAxisAlignment uint8

const (
	CrossAxisAlignmentCenter CrossAxisAlignment = iota // default
	CrossAxisAlignmentStart
)

type MainAxisAlignment uint8

const (
	MainAxisAlignmentStart MainAxisAlignment = iota // default
	MainAxisAlignmentSpaceBetween
)

type Axis uint8

const (
	Vertical Axis = iota // default
	Horizontal
)

type FlexFit uint8

const (
	FlexFitTight FlexFit = iota // default
	FlexFitLoose
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

func (ce *FlexElement) getChildCrossSize(child Element) float64 {
	switch ce.widget.(*Flex).Direction {
	case Horizontal:
		return child.getSize().height
	case Vertical:
		return child.getSize().width
	}
	return 0
}

func (ce *FlexElement) getChildMainSize(child Element) float64 {
	switch ce.widget.(*Flex).Direction {
	case Horizontal:
		return child.getSize().width
	case Vertical:
		return child.getSize().height
	}
	return 0
}

func (ce *FlexElement) layout(constraints Constraints) error {
	widget := ce.widget.(*Flex)

	var innerConstraints Constraints
	switch widget.Direction {
	case Horizontal:
		innerConstraints = Constraints{
			maxHeight: constraints.maxHeight,
			minHeight: 0,
			maxWidth:  math.MaxInt16,
			minWidth:  0,
		}
	case Vertical:
		innerConstraints = Constraints{
			maxWidth:  constraints.maxWidth,
			minWidth:  0,
			maxHeight: math.MaxInt16,
			minHeight: 0,
		}
	}

	var maxChildCrossSize float64
	var allocatedMainSize float64
	for _, child := range ce.children {

		child.layout(innerConstraints)

		maxChildCrossSize = math.Max(maxChildCrossSize, ce.getChildCrossSize(child))
		allocatedMainSize += ce.getChildMainSize(child)
	}

	idealSize := allocatedMainSize
	var actualSize float64
	var crossSize float64
	switch widget.Direction {
	case Horizontal:
		size := constraints.constrain(Size{idealSize, maxChildCrossSize})
		actualSize = size.width
		crossSize = size.height
		ce.size = size
	case Vertical:
		size := constraints.constrain(Size{maxChildCrossSize, idealSize})
		actualSize = size.height
		crossSize = size.width
		ce.size = size
	}

	actualSizeDelta := actualSize - idealSize
	remainingSpace := math.Max(0, actualSizeDelta)

	var leadingSpace float64
	var betweenSpace float64
	totalChildren := len(ce.children)

	switch widget.MainAxisAlignment {
	case MainAxisAlignmentSpaceBetween:
		leadingSpace = 0
		betweenSpace = 0
		if totalChildren > 1 {
			betweenSpace = remainingSpace / float64(totalChildren-1)
		}
	case MainAxisAlignmentStart:
		leadingSpace = 0
		betweenSpace = 0
	default:
		panic("unimplemented")
	}

	childMainPosition := leadingSpace
	for _, child := range ce.children {

		var childCrossPosition float64

		switch widget.CrossAxisAlignment {
		case CrossAxisAlignmentStart:
			childCrossPosition = crossSize - ce.getChildCrossSize(child)
		case CrossAxisAlignmentCenter:
			childCrossPosition = crossSize/2 - ce.getChildCrossSize(child)/2
		default:
			panic("unimplemented")
		}

		var childOffset Offset
		switch widget.Direction {
		case Horizontal:
			childOffset = Offset{x: childMainPosition, y: childCrossPosition}
		case Vertical:
			childOffset = Offset{x: childCrossPosition, y: childMainPosition}
		}

		childMainPosition += ce.getChildMainSize(child) + betweenSpace

		child.setOffset(childOffset)
	}

	return nil
}

func (c *FlexElement) render(offset Offset, renderer *sdl.Renderer) {
	for _, child := range c.children {
		childOffset := child.getOffset()
		child.render(offset.Add(childOffset), renderer)
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

var _ StatelessWidget = &Expanded{}
var _ ElementWidget = &Flexible{}
var _ Element = &FlexibleElement{}
var _ HasChildElement = &FlexibleElement{}

/* Expanded seems to do have no real implementation. All the magic is in Flex */
type Expanded struct {
	Child Widget
}

func (e *Expanded) Build(context BuildContext) (Widget, error) {
	return &Flexible{Fit: FlexFitTight, Child: e.Child}, nil
}

type Flexible struct {
	Child Widget
	Fit   FlexFit
}

func (f *Flexible) createElement() Element {
	fe := &FlexibleElement{}
	fe.widget = f
	return fe
}

/* FlexibleElement exists just to hold onto and respond to Flex's request for child Fit parameter */
type FlexibleElement struct {
	elementData
	childElementData
}

func (fe *FlexibleElement) getFit() FlexFit {
	return fe.widget.(*Flexible).Fit
}

func (ce *FlexibleElement) layout(constraints Constraints) error {
	return ce.getChildElement().layout(constraints)
}

func (fe *FlexibleElement) getSize() Size {
	return fe.getChildElement().getSize()
}

func (c *FlexibleElement) render(offset Offset, renderer *sdl.Renderer) {
	c.getChildElement().render(offset, renderer)
}
