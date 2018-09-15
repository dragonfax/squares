package main

import (
	"github.com/dragonfax/squares/examples/wordlist/wordpairs"
	"github.com/dragonfax/squares/squares"
	"github.com/dragonfax/squares/squares/listview"
)

var _ squares.StatelessWidget = &MyApp{}
var _ squares.StatefulWidget = &RandomWords{}
var _ squares.State = &RandomWordsState{}

func main() {
	err := squares.RunApp(&MyApp{})
	if err != nil {
		panic(err)
	}
}

type MyApp struct {
}

func (ma *MyApp) Build(context squares.StatelessContext) (squares.Widget, error) {

	return &squares.Center{Child: &RandomWords{}}, nil
}

func isOdd(i int) bool {
	return i%2 == 1
}

type RandomWords struct {
}

func (*RandomWords) CreateState() squares.State {
	return &RandomWordsState{suggestions: make([]wordpairs.WordPair, 0, 10)}
}

type RandomWordsState struct {
	suggestions []wordpairs.WordPair
}

func (rws *RandomWordsState) Build(context squares.StatefulContext) (squares.Widget, error) {
	return &squares.Composite{Child: &listview.Builder{
		Padding: squares.EdgeInsetsAll(16),
		ItemBuilder: func(i int) squares.Widget {
			if isOdd(i) {
				return &squares.Divider{}
			}

			r := i / 2
			if r >= len(rws.suggestions) {
				for x := 0; x < 10; x++ {
					rws.suggestions = append(rws.suggestions, wordpairs.GenerateWordPair())
				}
			}

			wp := rws.suggestions[r]
			return BuildRow(wp)
		},
	}}, nil
}

func BuildRow(wp wordpairs.WordPair) squares.Widget {
	return &listview.ListTile{
		Title: &squares.Text{
			Text: wp.AsPascalCase(),
		},
	}
}
