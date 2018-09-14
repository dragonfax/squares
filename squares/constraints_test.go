package gltr

import (
	"math"
	"testing"
)

func TestConstraintsDeflate(t *testing.T) {
	t.Run("Constrain deflate 0 min", func(t *testing.T) {
		c := Constraints{
			minHeight: 0,
			maxHeight: 20,
			minWidth:  0,
			maxWidth:  20,
		}

		c2 := c.deflate(EdgeInsetsAll(8))

		if c2.minHeight != 0 {
			t.Fatal("wrong minHeight ", c2.maxHeight)
		}
		if c2.maxHeight != 20-8*2 {
			t.Fatal("wrong maxHeight ", c2.maxHeight)
		}

	})
}

func TestConstraintsConstrain(t *testing.T) {

	t.Run("constrain size", func(t *testing.T) {
		c := Constraints{
			minHeight: 0,
			maxHeight: 15,
			minWidth:  0,
			maxWidth:  15,
		}

		size := c.constrain(Size{20, 20})

		if size.width != 15 {
			t.Fatal("wrong width, expected 20, got ", size.width)
		}
		if size.height != 15 {
			t.Fatal("wrong height, expected 20, got ", size.width)
		}
	})

	t.Run("constrain size with +Inf", func(t *testing.T) {
		c := Constraints{
			minHeight: 0,
			maxHeight: math.MaxFloat64,
			minWidth:  0,
			maxWidth:  15,
		}

		size := c.constrain(Size{20, 20})

		if size.width != 15 {
			t.Fatal("wrong width, expected 20, got ", size.width)
		}
		if size.height != 20 {
			t.Fatal("wrong height, expected 20, got ", size.width)
		}
	})

	t.Run("constrain size with min dimension", func(t *testing.T) {
		c := Constraints{
			minHeight: 5,
			maxHeight: 15,
			minWidth:  0,
			maxWidth:  15,
		}

		size := c.constrain(Size{2, 2})

		if size.width != 2 {
			t.Fatal("wrong width, expected 20, got ", size.width)
		}
		if size.height != 5 {
			t.Fatal("wrong height, expected 20, got ", size.width)
		}
	})

}
