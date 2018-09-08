package flutter

type Divider struct {
}

type Center struct {
	Child Widget
}

func (ce *Center) GetChild() Widget {
	return ce.Child
}

func (p *Center) SetChild(c Widget) {
	p.Child = c
}

type Text struct {
	Text string
}

type Column struct {
	Children []Widget
}

func (c *Column) GetChildren() []Widget {
	return c.Children
}

func (c *Column) SetChildren(cs []Widget) {
	c.Children = cs
}
