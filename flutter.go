package main

type Widget interface {
	Build(*BuildContext) (Widget, error)
}

func runApp(w Widget) {
}

type AppBar struct {
	title Widget
}

func (w AppBar) Build(bc *BuildContext) (Widget, error) {
	return nil, nil
}

type Center struct {
	child Widget
}

func (w Center) Build(bc *BuildContext) (Widget, error) {
	return nil, nil
}

type Text struct {
	text string
}

func (w Text) Build(bc *BuildContext) (Widget, error) {
	return nil, nil
}

type BuildContext struct {
}

type MaterialApp struct {
	title string
	home  Widget
}

func (w MaterialApp) Build(bc *BuildContext) (Widget, error) {
	return nil, nil
}

type Scaffold struct {
	appBar *AppBar
	body   Widget
}

func (w Scaffold) Build(bc *BuildContext) (Widget, error) {
	return nil, nil
}
