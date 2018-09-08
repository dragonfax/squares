package flutter

type Divider struct {
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
