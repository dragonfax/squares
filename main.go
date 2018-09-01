package main

func main() {
	runApp(&MyApp{})
}

type MyApp struct {
}

func (ma *MyApp) Build(context *BuildContext) (Widget, error) {

	wp := wordPair.random()

	return &MaterialApp{
		title: "Welcome to Flutter",
		home: &Scaffold{
			appBar: &AppBar{
				title: &Text{"Welcome to Flutter"},
			},
			body: &Center{
				child: &Text{wp.AsPascalCase()},
			},
		},
	}, nil
}
