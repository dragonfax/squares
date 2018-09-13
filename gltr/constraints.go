package gltr

import (
	"math"
)

type Offset struct {
	x, y float64
}

func (o Offset) Add(o2 Offset) Offset {
	return Offset{x: o.x + o2.x, y: o.y + o2.y}
}

type Size struct {
	width, height float64
}

type Constraints struct {
	minWidth, minHeight, maxWidth, maxHeight float64
}

func (c Constraints) loosen() Constraints {
	return Constraints{minWidth: 0, minHeight: 0, maxWidth: c.maxWidth, maxHeight: c.maxHeight}
}

func (c Constraints) constrain(size Size) Size {
	return Size{
		width:  clamp(c.minWidth, c.maxWidth, size.width),
		height: clamp(c.minHeight, c.maxHeight, size.height),
	}
}

func (s Size) addMargin(in EdgeInsets) Size {
	return Size{width: in.Left + s.width + in.Right, height: in.Top + s.height + in.Bottom}
}

func (c Constraints) deflate(in EdgeInsets) Constraints {
	deflatedMinWidth := math.Max(0.0, c.minWidth-in.horizontal())
	deflatedMinHeight := math.Max(0.0, c.minHeight-in.vertical())
	deflatedMaxWidth := math.Max(deflatedMinWidth, c.maxWidth-in.horizontal())
	if c.maxWidth == math.MaxFloat64 {
		deflatedMaxWidth = math.MaxFloat64
	}
	deflatedMaxHeight := math.Max(deflatedMinHeight, c.maxHeight-in.vertical())
	if c.maxHeight == math.MaxFloat64 {
		deflatedMaxHeight = math.MaxFloat64
	}
	return Constraints{
		minWidth:  deflatedMinWidth,
		maxWidth:  deflatedMaxWidth,
		minHeight: deflatedMinHeight,
		maxHeight: deflatedMaxHeight,
	}
}
