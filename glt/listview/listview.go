package listview

import "github.com/dragonfax/glitter/glt"

var _ glt.StatefulWidget = &Builder{}
var _ glt.State = &BuilderState{}
var _ glt.StatelessWidget = &ListTile{}

type ItemBuilderFunc func(int) glt.Widget

type Builder struct {
	Padding     glt.EdgeInsets
	ItemBuilder ItemBuilderFunc
}

func (b *Builder) CreateState() glt.State {
	return &BuilderState{}
}

type BuilderState struct {
	firstItem int
}

func (bs *BuilderState) Build(context glt.BuildContext) (glt.Widget, error) {
	widget := context.GetWidget().(*Builder)
	children := make([]glt.Widget, 10)
	println("rebuilding at ", bs.firstItem)
	for i := 0; i < 10; i++ {
		child := widget.ItemBuilder(bs.firstItem + i)
		children[i] = &glt.Padding{Padding: widget.Padding, Child: child}
	}

	return &glt.MouseWheelListener{
		Callback: func(d glt.MouseWheelDirection) {
			if d == glt.MOUSEWHEEL_UP {
				println("scrolling up")
				bs.firstItem += 1
				println("first item was set to ", bs.firstItem)
			} else if d == glt.MOUSEWHEEL_DOWN && bs.firstItem > 0 {
				println("scrolling down")
				bs.firstItem -= 1
				println("first item was set to ", bs.firstItem)
			}
		},
		Child: &glt.Column{Children: children},
	}, nil
}

type ListTile struct {
	Title glt.Widget
}

func (w ListTile) Build(context glt.BuildContext) (glt.Widget, error) {
	return w.Title, nil
}
