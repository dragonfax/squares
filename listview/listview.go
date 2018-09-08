package listview

import "github.com/dragonfax/flutter-go-example/flutter"

type ItemBuilderFunc func(flutter.BuildContext, int) flutter.Widget

type Builder struct {
	Padding     flutter.EdgeInsets
	ItemBuilder ItemBuilderFunc
}

func (w Builder) Build(bc *flutter.BuildContext) (flutter.Widget, error) {
	return nil, nil
}

type ListTile struct {
	Title flutter.Widget
}

func (w ListTile) Build(bc *flutter.BuildContext) (flutter.Widget, error) {
	return nil, nil
}
