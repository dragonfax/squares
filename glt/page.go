package glt

type Scaffold struct {
	Key  *GlobalKey
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
