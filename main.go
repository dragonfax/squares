package main

import (
	"github.com/dragonfax/flutter-go-example/flutter"
	"github.com/dragonfax/flutter-go-example/listview"
	"github.com/dragonfax/flutter-go-example/wordpairs"
)

func main() {
	flutter.RunApp(&MyApp{})
}

type MyApp struct {
}

func (ma *MyApp) Build(context *flutter.BuildContext) (flutter.Widget, error) {

	return &flutter.MaterialApp{
		Title: "Welcome to Flutter",
		Home: &flutter.Scaffold{
			AppBar: &flutter.AppBar{
				Title: &flutter.Text{"Welcome to Flutter"},
			},
			Body: &flutter.Center{
				Child: &RandomWords{},
			},
		},
	}, nil
}

type RandomWords struct {
}

func (r *RandomWords) CreateState() flutter.State {
	return &RandomWordsState{suggestions: make([]wordpairs.WordPair, 0)}
}

type RandomWordsState struct {
	suggestions []wordpairs.WordPair
}

func isOdd(i int) bool {
	return i%2 == 1
}

func (rws *RandomWordsState) Build(bc *flutter.BuildContext) (flutter.Widget, error) {
	return listview.Builder{
		Padding: flutter.EdgeInsets{All: 16.0},
		ItemBuilder: func(context flutter.BuildContext, i int) flutter.Widget {
			if isOdd(i) {
				return &flutter.Divider{}
			}

			r := i / 2
			if r >= len(rws.suggestions) {
				rws.suggestions = append(rws.suggestions, wordpairs.RandomNum(10)...)
			}

			return rws.BuildRow(rws.suggestions[r])
		},
	}, nil
}

func (rws *RandomWordsState) BuildRow(wp wordpairs.WordPair) flutter.Widget {
	return &listview.ListTile{
		Title: &flutter.Text{
			Text: wp.AsPascalCase(),
		},
	}
}
