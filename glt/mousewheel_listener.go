package glt

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
	return &MouseWheelListenerState{widget: mwl}
}

type MouseWheelListenerState struct {
	widget *MouseWheelListener
}

func (mwls *MouseWheelListenerState) Build() (Widget, error) {
	mouseWheelCallback = mwls.widget.Callback
	return mwls.widget.Child, nil
}
