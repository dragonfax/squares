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

func (o Offset) Sub(o2 Offset) Offset {
	return Offset{x: o.x - o2.x, y: o.y - o2.y}
}

/* a dimension of -1 means "no specification" */
type Size struct {
	Width, Height float64
}

// return a new size with both dimensions able to contain s
func (s1 Size) Expand(s2 Size) Size {
	return Size{Width: math.Max(s1.Width, s2.Width), Height: math.Max(s1.Height, s2.Height)}
}

/* constraint dimensions can be max = +Inf or min = 0
 * -1's are not allowed here.
 */
type Constraints struct {
	minWidth, minHeight, maxWidth, maxHeight float64
}

func (c Constraints) loosen() Constraints {
	return Constraints{minWidth: 0, minHeight: 0, maxWidth: c.maxWidth, maxHeight: c.maxHeight}
}

func ConstraintsTight(width, height float64) Constraints {
	c := Constraints{minWidth: width, maxWidth: width, minHeight: height, maxHeight: height}
	if width == -1 {
		c.minWidth = 0
		c.maxWidth = math.Inf(1)
	}
	if height == -1 {
		c.minHeight = 0
		c.maxHeight = math.Inf(1)
	}
	return c
}

func ConstraintsUnbounded() Constraints {
	return Constraints{
		minWidth:  0,
		minHeight: 0,
		maxWidth:  math.Inf(1),
		maxHeight: math.Inf(1),
	}
}

func (c Constraints) constrain(size Size) Size {
	s := Size{
		Width:  clamp(c.minWidth, c.maxWidth, size.Width),
		Height: clamp(c.minHeight, c.maxHeight, size.Height),
	}
	if size.Width == -1 {
		s.Width = c.maxWidth
	}
	if size.Height == -1 {
		s.Height = c.maxHeight
	}
	return s
}

func (c Constraints) constrainWithRatio(size Size) Size {
	ratio := size.Width / size.Height
	// width = ratio * height
	// height = width / ratio

	// constrain by width
	wwidth := clamp(c.minWidth, c.maxWidth, size.Width)
	wheight := wwidth / ratio

	// constrain by height
	hheight := clamp(c.minHeight, c.maxHeight, size.Height)
	hwidth := ratio * hheight

	// which one fits in the constraints best?
	if wwidth < hwidth {
		return Size{Width: wwidth, Height: wheight}
	} else {
		return Size{Width: hwidth, Height: hheight}
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
