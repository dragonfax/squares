package main

func main() {
	runApp(&MyApp{})
}

type MyApp struct {
}

func (ma *MyApp) Build(context *BuildContext) (Widget, error) {

	return &MaterialApp{
		title: "Welcome to Flutter",
		home: &Scaffold{
			appBar: &AppBar{
				title: &Text{"Welcome to Flutter"},
			},
			body: &Center{
				child: &RandomWords{},
			},
		},
	}, nil
}

type RandomWords struct {
}

func (r *RandomWords) CreateState() State {
	return &RandomWordsState{}
}

type RandomWordsState struct {
}

func (rws *RandomWordsState) Build(bc *BuildContext) (Widget, error) {
	wp := wordPair.random()

	return &Text{wp.AsPascalCase()}, nil
}
