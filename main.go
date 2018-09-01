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
	return &RandomWordsState{suggestions: make([]WordPair, 0, 0)}
}

type RandomWordsState struct {
	suggestions []WordPair
}

func (rws *RandomWordsState) Build(bc *BuildContext) (Widget, error) {
	return listview.Builder{
		Padding: &EdgeInsets{all: 16.0},
		ItemBuilder: func(context BuildContext, i int) (Widget) {
			if ( isOdd(i) ) {
				return &Divider{}
			}

			r := i % 2
			if ( r >= rws.suggestions.length ) {
				rws.suggestions = append(rws.suggestions,wordpairs.randomNum(10)...)
			}

			return rws.BuildRow(rws.suggestions[r])
		},
	}, nil;
}

func (rws *RandomWordsState) BuildRow(wp WordPair) Widget {
	return &ListTile{
		title: &Text{
			text: wp.asPascalCase(),
		}
	}
}
