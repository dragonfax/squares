package gltr

import (
	"math"
)

type Offset struct {
	x, y uint16
}

func (o Offset) Add(o2 Offset) Offset {
	return Offset{x: o.x + o2.x, y: o.y + o2.y}
}

type Size struct {
	width, height uint16
}

// use MaxUint32 for +Inf during layout
type Constraints struct {
	minWidth, minHeight, maxWidth, maxHeight uint16
}

func (c Constraints) loosen() Constraints {
	return Constraints{minWidth: 0, minHeight: 0, maxWidth: c.maxWidth, maxHeight: c.maxHeight}
}

func (c Constraints) constrain(size Size) Size {
	return Size{
		width:  clampUint16(c.minWidth, c.maxWidth, size.width),
		height: clampUint16(c.minHeight, c.maxHeight, size.height),
	}
}

func (s Size) addMargin(in EdgeInsets) Size {
	return Size{width: in.Left + s.width + in.Right, height: in.Top + s.height + in.Bottom}
}

func (c Constraints) deflate(in EdgeInsets) Constraints {
	deflatedMinWidth := MaxUint16(0.0, c.minWidth-in.horizontal())
	deflatedMinHeight := MaxUint16(0.0, c.minHeight-in.vertical())
	deflatedMaxWidth := MaxUint16(deflatedMinWidth, c.maxWidth-in.horizontal())
	if c.maxWidth == math.MaxUint16 {
		deflatedMaxWidth = math.MaxUint16
	}
	deflatedMaxHeight := MaxUint16(deflatedMinHeight, c.maxHeight-in.vertical())
	if c.maxHeight == math.MaxUint16 {
		deflatedMaxHeight = math.MaxUint16
	}
	return Constraints{
		minWidth:  deflatedMinWidth,
		maxWidth:  deflatedMaxWidth,
		minHeight: deflatedMinHeight,
		maxHeight: deflatedMaxHeight,
	}
}
