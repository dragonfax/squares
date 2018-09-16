package squares

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

var _ StatelessWidget = Expanded{}
var _ StatelessWidget = Column{}
var _ StatelessWidget = Row{}
var _ ElementWidget = Flex{}
var _ HasChildren = Flex{}
var _ HasChildrenElements = &FlexElement{}

type CrossAxisAlignment uint8

const (
	CrossAxisAlignmentCenter CrossAxisAlignment = iota // default
	CrossAxisAlignmentStart
	CrossAxisAlignmentEnd
	CrossAxisAlignmentStretch
)

type MainAxisAlignment uint8

const (
	MainAxisAlignmentStart MainAxisAlignment = iota // default
	MainAxisAlignmentSpaceBetween
	MainAxisAlignmentEnd
	MainAxisAlignmentCenter
	MainAxisAlignmentSpaceAround
	MainAxisAlignmentSpaceEvenly
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

func (f Flex) getChildren() []Widget {
	return f.Children
}

func (c Flex) createElement() Element {
	ce := &FlexElement{}
	ce.widget = c
	return ce
}

type FlexElement struct {
	elementData
	childrenElementsData
}

func (ce *FlexElement) getChildCrossSize(child Element) float64 {
	switch ce.widget.(Flex).Direction {
	case Horizontal:
		return child.getSize().Height
	case Vertical:
		return child.getSize().Width
	}
	return 0
}

func (ce *FlexElement) getChildMainSize(child Element) float64 {
	switch ce.widget.(Flex).Direction {
	case Horizontal:
		return child.getSize().Width
	case Vertical:
		return child.getSize().Height
	}
	return 0
}

func (ce *FlexElement) layout(constraints Constraints) error {
	widget := ce.widget.(Flex)

	var totalFlex int
	var totalChildren int
	maxMainSize := constraints.maxHeight
	if widget.Direction == Horizontal {
		maxMainSize = constraints.maxWidth
	}
	canFlex := !math.IsInf(maxMainSize, 1)

	crossSize := 0.0
	allocatedSize := 0.0

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
	var lastFlexChild *FlexibleElement
	for _, child := range ce.children {
		flexChild, ok := child.(*FlexibleElement)

		totalChildren++

		flex := 0
		if ok {
			flex = flexChild.getFlex()
		}

		if flex > 0 {
			totalFlex += flex
			lastFlexChild = flexChild
		} else {
			child.layout(innerConstraints)
			maxChildCrossSize = math.Max(maxChildCrossSize, ce.getChildCrossSize(child))
			allocatedMainSize += ce.getChildMainSize(child)
		}
	}

	t := 0.0
	if canFlex {
		t = maxMainSize
	}
	freeSpace := math.Max(0.0, t-allocatedSize)
	allocatedFlexSpace := 0.0
	if totalFlex > 0 {

		spacePerFlex := math.NaN()
		if canFlex && totalFlex > 0 {
			spacePerFlex = freeSpace / float64(totalFlex)
		}

		for _, child := range ce.children {
			flexChild, ok := child.(*FlexibleElement)

			flex := 0
			if ok {
				flex = flexChild.getFlex()
			}

			if flex > 0 {

				var maxChildExtent = math.Inf(1)
				if canFlex {
					if child == lastFlexChild {
						maxChildExtent = freeSpace - allocatedFlexSpace
					} else {
						maxChildExtent = spacePerFlex * float64(flex)
					}
				}

				var minChildExtent float64
				switch flexChild.getFit() {
				case FlexFitTight:
					minChildExtent = maxChildExtent
				case FlexFitLoose:
					minChildExtent = 0.0
				}

				if widget.CrossAxisAlignment == CrossAxisAlignmentStretch {
					switch widget.Direction {
					case Horizontal:
						innerConstraints = Constraints{
							minWidth:  minChildExtent,
							maxWidth:  maxChildExtent,
							minHeight: constraints.maxHeight,
							maxHeight: constraints.maxHeight,
						}
					case Vertical:
						innerConstraints = Constraints{
							minWidth:  constraints.maxWidth,
							maxWidth:  constraints.maxWidth,
							minHeight: minChildExtent,
							maxHeight: maxChildExtent,
						}
					}
				} else {
					switch widget.Direction {
					case Horizontal:
						innerConstraints = Constraints{
							minWidth:  minChildExtent,
							maxWidth:  maxChildExtent,
							maxHeight: constraints.maxHeight,
							minHeight: 0,
						}
					case Vertical:
						innerConstraints = Constraints{
							maxWidth:  constraints.maxWidth,
							minHeight: minChildExtent,
							maxHeight: maxChildExtent,
							minWidth:  0,
						}
					}
				}

				child.layout(innerConstraints)
				childSize := ce.getChildMainSize(child)
				allocatedSize += childSize
				allocatedFlexSpace += maxChildExtent
				crossSize = math.Max(crossSize, ce.getChildCrossSize(child))
			}
		}
	}

	idealSize := allocatedMainSize
	var actualSize float64
	switch widget.Direction {
	case Horizontal:
		size := constraints.constrain(Size{idealSize, maxChildCrossSize})
		actualSize = size.Width
		crossSize = size.Height
		ce.size = size
	case Vertical:
		size := constraints.constrain(Size{maxChildCrossSize, idealSize})
		actualSize = size.Height
		crossSize = size.Width
		ce.size = size
	}

	actualSizeDelta := actualSize - idealSize

	remainingSpace := math.Max(0, actualSizeDelta)
	var leadingSpace float64
	var betweenSpace float64

	switch widget.MainAxisAlignment {
	case MainAxisAlignmentSpaceBetween:
		leadingSpace = 0
		betweenSpace = 0
		if totalChildren > 1 {
			betweenSpace = remainingSpace / float64(totalChildren-1)
		}
	case MainAxisAlignmentEnd:
		leadingSpace = remainingSpace
		betweenSpace = 0
	case MainAxisAlignmentCenter:
		leadingSpace = remainingSpace / 2
		betweenSpace = 0
	case MainAxisAlignmentStart:
		leadingSpace = 0
		betweenSpace = 0
	case MainAxisAlignmentSpaceAround:
		betweenSpace = 0
		if totalChildren > 0 {
			betweenSpace = remainingSpace / float64(totalChildren)
		}
		leadingSpace = betweenSpace / 2
	case MainAxisAlignmentSpaceEvenly:
		betweenSpace = 0
		if totalChildren > 0 {
			betweenSpace = remainingSpace / (float64(totalChildren) + 1)
		}
		leadingSpace = betweenSpace
	default:
		panic("unimplemented")
	}

	childMainPosition := leadingSpace
	for _, child := range ce.children {

		var childCrossPosition float64

		switch widget.CrossAxisAlignment {
		case CrossAxisAlignmentEnd:
			childCrossPosition = crossSize - ce.getChildCrossSize(child)
		case CrossAxisAlignmentStart:
			childCrossPosition = 0
		case CrossAxisAlignmentCenter:
			childCrossPosition = crossSize/2 - ce.getChildCrossSize(child)/2
		case CrossAxisAlignmentStretch:
			childCrossPosition = 0
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

func (c Row) Build(context StatelessContext) (Widget, error) {
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

func (c Column) Build(context StatelessContext) (Widget, error) {
	return &Flex{
		Direction:          Vertical,
		CrossAxisAlignment: c.CrossAxisAlignment,
		MainAxisAlignment:  c.MainAxisAlignment,
		Children:           c.Children,
	}, nil
}

var _ StatelessWidget = Expanded{}
var _ ElementWidget = Flexible{}
var _ HasChild = Flexible{}
var _ Element = &FlexibleElement{}
var _ HasChildElement = &FlexibleElement{}

/* Expanded seems to do have no real implementation. All the magic is in Flex */
/* TODO widget behavior broken */
type Expanded struct {
	Child Widget
}

func (e Expanded) Build(context StatelessContext) (Widget, error) {
	return Flexible{Fit: FlexFitTight, Flex: 1, Child: e.Child}, nil
}

/* TODO you should set a flex of at lest 1, since there is no constructor here */
/* TODO widget behavior broken */
type Flexible struct {
	Child Widget
	Fit   FlexFit
	Flex  int
}

func (f Flexible) getChild() Widget {
	return f.Child
}

func (f Flexible) createElement() Element {
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
	return fe.widget.(Flexible).Fit
}

func (fe *FlexibleElement) getFlex() int {
	return fe.widget.(Flexible).Flex
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
