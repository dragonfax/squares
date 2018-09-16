package squares

/* TODO
 * Forces the child to match its size specification, even if they are 0.
 *
 * * needs to be implemented to allow for unspecified constraint dimensions.
 * * needs to reimplemented based on a ConstrainedBox class
 */

import "github.com/veandco/go-sdl2/sdl"

var _ ElementWidget = SizedBox{}
var _ HasChild = SizedBox{}
var _ Element = &SizedBoxElement{}
var _ HasChildElement = &SizedBoxElement{}

type SizedBox struct {
	Size  Size
	Child Widget
}

func (sb SizedBox) createElement() Element {
	sbe := &SizedBoxElement{}
	sbe.widget = sb
	return sbe
}

type SizedBoxElement struct {
	elementData
	childElementData
}

func (sbe *SizedBoxElement) layout(constraints Constraints) error {
	widget := sbe.widget.(SizedBox)

	if sbe.child != nil {
		// "tight" constraints to child,
		// -1 constraints are honored, though. child adopts parent size
		c := ConstraintsTight(widget.Size.Width, widget.Size.Height)
		sbe.child.layout(c)
		sbe.size = sbe.child.getSize()
	} else {
		// no child: use given size but treat -1's as zeros
		sbe.size = widget.Size
		if sbe.size.Width == -1 {
			sbe.size.Width = 0
		}
		if sbe.size.Height == -1 {
			sbe.size.Height = 0
		}
	}

	return nil
}

func (sbe *SizedBoxElement) render(offset Offset, renderer *sdl.Renderer) {
	if sbe.child != nil {
		sbe.child.render(offset, renderer)
	}
}
