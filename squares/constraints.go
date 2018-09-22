package squares

import (
	"math"
)

var Inf float64 = math.MaxFloat64

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

// return a new size able to contain both s1 and s2
func (s1 Size) Union(s2 Size) Size {
	return Size{Width: math.Max(s1.Width, s2.Width), Height: math.Max(s1.Height, s2.Height)}
}

func (s Size) addMargin(in EdgeInsets) Size {
	return Size{Width: in.Left + s.Width + in.Right, Height: in.Top + s.Height + in.Bottom}
}
