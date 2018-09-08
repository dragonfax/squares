package flutter

type Padding struct {
	Padding    EdgeInsets
	Child      Widget
	size       Size
	parentData ParentData
}

func (p *Padding) GetChild() Widget {
	return p.Child
}

func (p *Padding) SetChild(child Widget) {
	p.Child = child
}

func (p *Padding) getParentData() *ParentData {
	return &p.parentData
}

func (p *Padding) getSize() Size {
	return p.size
}

func (p *Padding) layout(c Constraints) error {

	c2 := c.addMargins(p.Padding)

	cw := p.Child.(CoreWidget)
	cw.layout(c2)

	// multi child containers would read the sizes from the children, and position them accordingly.
	childSize := cw.getSize()
	paddedSize := childSize.addMargin(p.Padding)
	p.size = paddedSize

	// offset for Padding is easy, just offset by the padding amount.
	cw.getParentData().offset = Offset{p.Padding.All, p.Padding.All}

	return nil
}
