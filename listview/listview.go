package listview

type ItemBuilderFunc func(*glt.BuildContext, int) glt.Widget

type Builder struct {
	Padding     glt.EdgeInsets
	ItemBuilder ItemBuilderFunc
}

func (w Builder) Build(context *glt.BuildContext) (glt.Widget, error) {
	children := make([]glt.Widget, 10)
	for i := 0; i < 10; i++ {
		child := w.ItemBuilder(context, i)
		children[i] = &glt.Padding{Padding: w.Padding, Child: child}
	}
	return &glt.Column{Children: children}, nil
}

type ListTile struct {
	Title glt.Widget
}

func (w ListTile) Build(bc *glt.BuildContext) (glt.Widget, error) {
	return w.Title, nil
}
