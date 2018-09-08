package flutter

type Column struct {
	Children []Widget
	sizeData
	parentData
}

func (c *Column) getChildren() []Widget {
	return c.Children
}

func (c *Column) setChildren(cs []Widget) {
	c.Children = cs
}

func (ce *Column) layout(c constraints) error {

	ce.size = Size{0, 0}

	numChildren := uint16(len(ce.Children))

	for _, child := range ce.Children {

		cw := child.(coreWidget)

		// TODO not sure about this.
		// might need to do them one at a time. and see whats left for the others.
		con := constraints{
			minWidth:  c.maxWidth,
			minHeight: c.minHeight / numChildren,
			maxWidth:  c.maxWidth,
			maxHeight: c.maxHeight / numChildren,
		}

		cw.layout(con)

		childSize := cw.getSize()
		ce.size.width = MaxUint16(childSize.width, ce.size.width)
		offsetHeight := ce.size.height
		ce.size.height += childSize.height

		cw.getParentData().offset = Offset{
			x: 0,
			y: offsetHeight,
		}

	}

	return nil
}
