package glt

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
	return Size{width: in.Left + s.width + in.Right, height: in.Top + s.height + in.Bottom}
}

func (c Constraints) horizontal() uint16 {
	return c.Left + c.Right
}

func (c Constraints) vertical() uint16 {
	return c.Top + c.Bottom
}

func (c Constraints) deflate(in EdgeInsets) Constraints {
	deflatedMinWidth := MaxUint16(0.0, c.minWidth-c.horizontal())
	deflatedMinHeight := MaxUint16(0.0, c.minHeight-c.vertical())
	return Constraints{
		minWidth:  deflatedMinWidth,
		maxWidth:  math.max(deflatedMinWidth, c.maxWidth-c.horizontal()),
		minHeight: deflatedMinHeight,
		maxHeight: math.max(deflatedMinHeight, c.maxHeight-c.vertical()),
	}
}
