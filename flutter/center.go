package flutter

type Center struct {
	Child      Widget
	size       Size
	parentData ParentData
}

func (ce *Center) GetChild() Widget {
	return ce.Child
}

func (ce *Center) SetChild(c Widget) {
	ce.Child = c
}

func (ce *Center) getParentData() *ParentData {
	return &ce.parentData
}

func (ce *Center) getSize() Size {
	return ce.size
}

func (ce *Center) layout(c Constraints) error {

	cw := ce.Child.(CoreWidget)
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
