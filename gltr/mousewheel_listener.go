package gltr

var _ StatefulWidget = &MouseWheelListener{}
var _ State = &MouseWheelListenerState{}
var _ HasChild = &MouseWheelListener{}

var mouseWheelCallback MouseWheelFunc

type MouseWheelDirection uint8

const (
	MOUSEWHEEL_UP MouseWheelDirection = iota
	MOUSEWHEEL_DOWN
)

type MouseWheelFunc func(MouseWheelDirection)

type MouseWheelListener struct {
	Child    Widget
	Callback MouseWheelFunc
}

func (mwl *MouseWheelListener) getChild() Widget {
	return mwl.Child
}

func (mwl *MouseWheelListener) CreateState() State {
	return &MouseWheelListenerState{}
}

type MouseWheelListenerState struct {
}

func (mwls *MouseWheelListenerState) Build(context BuildContext) (Widget, error) {
	widget := context.GetWidget().(*MouseWheelListener)
	mouseWheelCallback = widget.Callback
	return widget.Child, nil
}
