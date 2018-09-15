package squares

/* Dummy implementation for the widgets, until I get around to building real
 * implementations */

func (d *Container) Build(context BuildContext) (Widget, error) {
	return d.Child, nil
}

func (d *Container) getChild() Widget {
	return d.Child
}

func (d *DecoratedBox) Build(context BuildContext) (Widget, error) {
	return d.Child, nil
}

func (d *DecoratedBox) getChild() Widget {
	return d.Child
}

func (d *SizedBox) Build(context BuildContext) (Widget, error) {
	return d.Child, nil
}

func (d *SizedBox) getChild() Widget {
	return d.Child
}

func (d *Stack) Build(context BuildContext) (Widget, error) {
	return &Column{Children: d.Children}, nil
}

func (d *Stack) getChildren() []Widget {
	return d.Children
}

func (d *FlexibleSpaceBar) Build(context BuildContext) (Widget, error) {
	return &Row{Children: d.getChildren()}, nil
}

func (d *FlexibleSpaceBar) getChildren() []Widget {
	return []Widget{d.Title, d.Background}
}

func (d *Icon) Build(context BuildContext) (Widget, error) {
	return &Text{Text: "Icon"}, nil
}

func (d *IconButton) Build(context BuildContext) (Widget, error) {
	return d.getChild(), nil
}

func (d *IconButton) getChild() Widget {
	return d.Icon
}

func (d *Image) Build(context BuildContext) (Widget, error) {
	return &Text{Text: "Image"}, nil
}

func (d *SnackBar) Build(context BuildContext) (Widget, error) {
	return d.getChild(), nil
}

func (d *SnackBar) getChild() Widget {
	return d.Content
}

func (d *CustomScrollView) Build(context BuildContext) (Widget, error) {
	return &Row{Children: d.getChildren()}, nil
}

func (d *CustomScrollView) getChildren() []Widget {
	return d.Slivers
}

func (d *SliverAppBar) Build(context BuildContext) (Widget, error) {
	return &Row{Children: d.getChildren()}, nil
}

func (d *SliverAppBar) getChildren() []Widget {
	return append(d.Actions, d.FlexibleSpace)
}

func (d *SliverList) Build(context BuildContext) (Widget, error) {
	return &Row{Children: d.getChildren()}, nil
}

func (d *SliverList) getChildren() []Widget {
	return d.Delegate.Children
}

func (d *MaterialApp) Build(context BuildContext) (Widget, error) {
	return d.getChild(), nil
}

func (d *MaterialApp) getChild() Widget {
	return d.Child
}

func (d *Scaffold) Build(context BuildContext) (Widget, error) {
	return d.getChild(), nil
}

func (d *Scaffold) getChild() Widget {
	return d.Body
}

func (d *PopupMenuItem) Build(context BuildContext) (Widget, error) {
	return d.getChild(), nil
}

func (d *PopupMenuItem) getChild() Widget {
	return d.Child
}

func (d *PopupMenuButton) Build(context BuildContext) (Widget, error) {
	children, err := d.ItemBuilder(context)
	if err != nil {
		return nil, err
	}
	c := make([]Widget, len(children))
	for i, child := range children {
		c[i] = child
	}
	return &Column{Children: c}, nil
}
