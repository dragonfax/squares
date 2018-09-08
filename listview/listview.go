package listview

import "github.com/dragonfax/flutter-go-example/flutter"

type ItemBuilderFunc func(*flutter.BuildContext, int) flutter.Widget

type Builder struct {
	Padding     flutter.EdgeInsets
	ItemBuilder ItemBuilderFunc
}

func (w Builder) Build(context *flutter.BuildContext) (flutter.Widget, error) {
	children := make([]flutter.Widget, 10)
	for i := 0; i < 10; i++ {
		child := w.ItemBuilder(context, i)
		children[i] = &flutter.Padding{Padding: w.Padding, Child: child}
	}
	return &flutter.Column{Children: children}, nil
}

type ListTile struct {
	Title flutter.Widget
}

func (w ListTile) Build(bc *flutter.BuildContext) (flutter.Widget, error) {
	return w.Title, nil
}
