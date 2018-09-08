package flutter

type Padding struct {
	Padding EdgeInsets
	Child   Widget
	size    Size
	parentData
}

func (p *Padding) getChild() Widget {
	return p.Child
}

func (p *Padding) setChild(child Widget) {
	p.Child = child
}

func (p *Padding) getSize() Size {
	return p.size
}

func (p *Padding) layout(c constraints) error {

	c2 := c.addMargins(p.Padding)

	cw := p.Child.(coreWidget)
	cw.layout(c2)

	// multi child containers would read the sizes from the children, and position them accordingly.
	childSize := cw.getSize()
	paddedSize := childSize.addMargin(p.Padding)
	p.size = paddedSize

	// offset for Padding is easy, just offset by the padding amount.
	cw.getParentData().offset = Offset{p.Padding.All, p.Padding.All}

	return nil
}
