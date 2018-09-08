package flutter

type Center struct {
	Child Widget
	sizeData
	parentData
}

func (ce *Center) getChild() Widget {
	return ce.Child
}

func (ce *Center) setChild(c Widget) {
	ce.Child = c
}

func (ce *Center) layout(c constraints) error {

	cw := ce.Child.(coreWidget)
	cw.layout(c)

	ce.size = Size{width: c.maxWidth, height: c.maxHeight}

	// multi child containers would read the sizes from the children, and position them accordingly.
	childSize := cw.getSize()

	// offset for Padding is easy, just offset by the padding amount.
	cw.getParentData().offset = Offset{
		x: (ce.size.width - childSize.width) / 2,
		y: (ce.size.height - childSize.height) / 2,
	}

	return nil
}
