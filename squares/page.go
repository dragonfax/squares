package squares

var _ StatelessWidget = &Scaffold{}
var _ HasChild = &Scaffold{}
var _ StatelessWidget = &SnackBar{}
var _ HasChild = &SnackBar{}
var _ StatelessWidget = &PopupMenuButton{}
var _ StatelessWidget = &PopupMenuItem{}
var _ HasChild = &PopupMenuItem{}
var _ StatelessWidget = &MaterialApp{}
var _ HasChild = &MaterialApp{}

type Scaffold struct {
	Body Widget
}

func (s *Scaffold) ShowSnackBar(snackBar *SnackBar) {

}

type SnackBar struct {
	Content Widget
}

type PopupMenuItemBuilderFunc func(context BuildContext) ([]*PopupMenuItem, error)

type PopupMenuButton struct {
	OnSelected  func(interface{})
	ItemBuilder PopupMenuItemBuilderFunc
}

type PopupMenuItem struct {
	Value interface{}
	Child Widget
}

type MaterialApp struct {
	Title string
	Color Color
	Child Widget
}
