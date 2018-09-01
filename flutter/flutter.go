package flutter

type Widget interface {
}

type EdgeInsets struct {
	All int
}

type Divider struct {
}

type StatefulWidget interface {
	CreateState() State
}

type State interface {
	Build(*BuildContext) (Widget, error)
}

type StatelessWidget interface {
	Build(*BuildContext) (Widget, error)
}

func RunApp(w Widget) {
}

type AppBar struct {
	Title Widget
}

func (w AppBar) Build(bc *BuildContext) (Widget, error) {
	return nil, nil
}

type Center struct {
	Child Widget
}

func (w Center) Build(bc *BuildContext) (Widget, error) {
	return nil, nil
}

type Text struct {
	Text string
}

func (w Text) Build(bc *BuildContext) (Widget, error) {
	return nil, nil
}

type BuildContext struct {
}

type MaterialApp struct {
	Title string
	Home  Widget
}

func (w MaterialApp) Build(bc *BuildContext) (Widget, error) {
	return nil, nil
}

type Scaffold struct {
	AppBar *AppBar
	Body   Widget
}

func (w Scaffold) Build(bc *BuildContext) (Widget, error) {
	return nil, nil
}
