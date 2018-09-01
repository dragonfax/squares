package listview

import "github.com/dragonfax/flutter-go-example/flutter"

type ItemBuilderFunc func(flutter.BuildContext, int) flutter.Widget

type Builder struct {
	Padding     flutter.EdgeInsets
	ItemBuilder ItemBuilderFunc
}

type ListTile struct {
	Title flutter.Widget
}
