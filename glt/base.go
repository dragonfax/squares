package glt

import (
	"math"
)

type EdgeInsets struct {
	All uint16
}

type Offset struct {
	x, y uint16
}

type Size struct {
	width, height uint16
}

type sizeData struct {
	size Size
}

func (sd sizeData) getSize() Size {
	return sd.size
}

type parentData struct {
	offset Offset
}

// let an embedded struct return itself in order to match interfaces (interfaces for struct elements)
func (ce parentData) getParentData() *parentData {
	return &ce
}

// use MaxUint32 for +Inf during layout
type constraints struct {
	minWidth, minHeight, maxWidth, maxHeight uint16
}

func (s Size) addMargin(in EdgeInsets) Size {
	return Size{width: s.width + in.All, height: s.height + in.All}
}

func (c constraints) addMargins(in EdgeInsets) constraints {
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
