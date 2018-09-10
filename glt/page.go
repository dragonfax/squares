package glt

type Scaffold struct {
	Key  *GlobalKey
	Body Widget
}

type SnackBar struct {
	Content Widget
}

type PopupMenuButton struct {
	OnSelected  func(interface{})
	ItemBuilder BuildFunc
}

type PopupMenuItem struct {
	Value interface{}
	Child Widget
}

type MaterialApp struct {
}
