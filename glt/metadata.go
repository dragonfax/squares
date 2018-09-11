package glt

import (
	"math"
)

type EdgeInsets struct {
	All      uint16
	Vertical uint16
}

type Offset struct {
	x, y uint16
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
		width:  clampUint32(c.minWidth, c.maxWidth, size.width),
		height: clampUint32(c.minHeight, c.maxHeight, size.height),
	}
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
