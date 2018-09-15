package squares

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
	Width, Height float64
}

type Constraints struct {
	minWidth, minHeight, maxWidth, maxHeight float64
}

func (c Constraints) loosen() Constraints {
	return Constraints{minWidth: 0, minHeight: 0, maxWidth: c.maxWidth, maxHeight: c.maxHeight}
}

func ConstraintsAbsolute(width, height float64) Constraints {
	return Constraints{minWidth: width, maxWidth: width, minHeight: height, maxHeight: height}
}

func (c Constraints) constrain(size Size) Size {
	return Size{
		Width:  clamp(c.minWidth, c.maxWidth, size.Width),
		Height: clamp(c.minHeight, c.maxHeight, size.Height),
	}
}

func (s Size) addMargin(in EdgeInsets) Size {
	return Size{Width: in.Left + s.Width + in.Right, Height: in.Top + s.Height + in.Bottom}
}

func (c Constraints) deflate(in EdgeInsets) Constraints {
	deflatedMinWidth := math.Max(0.0, c.minWidth-in.horizontal())
	deflatedMinHeight := math.Max(0.0, c.minHeight-in.vertical())
	deflatedMaxWidth := math.Max(deflatedMinWidth, c.maxWidth-in.horizontal())
	if math.IsInf(c.maxWidth, 1) {
		deflatedMaxWidth = math.Inf(1)
	}
	deflatedMaxHeight := math.Max(deflatedMinHeight, c.maxHeight-in.vertical())
	if math.IsInf(c.maxHeight, 1) {
		deflatedMaxHeight = math.Inf(1)
	}
	return Constraints{
		minWidth:  deflatedMinWidth,
		maxWidth:  deflatedMaxWidth,
		minHeight: deflatedMinHeight,
		maxHeight: deflatedMaxHeight,
	}
}
