package squares

var _ StatefulWidget = &Window{}
var _ HasChild = &Window{}

/* This widget isn't created by the user, but by the framework to dole out
   Window resize events
*/
type Window struct {
	Child Widget
}

func (w *Window) getChild() Widget {
	return w.Child
}

type WindowState struct {
	Size Size
}

func (w *Window) CreateState() State {
	return &WindowState{}
}

func (ws *WindowState) Build(context StatefulContext) (Widget, error) {
	return &SizedBox{
		Size:  ws.Size,
		Child: context.GetWidget().(*Window).Child,
	}, nil
}
