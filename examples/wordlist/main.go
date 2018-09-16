package main

import (
	"github.com/dragonfax/squares/examples/wordlist/wordpairs"
	. "github.com/dragonfax/squares/squares"
	"github.com/dragonfax/squares/squares/listview"
)

var _ StatelessWidget = &MyApp{}
var _ StatefulWidget = &RandomWords{}
var _ State = &RandomWordsState{}

func main() {
	err := RunApp(&MyApp{})
	if err != nil {
		panic(err)
	}
}

type MyApp struct {
}

func (ma *MyApp) Build(context StatelessContext) (Widget, error) {

	return &Center{Child: &RandomWords{}}, nil
}

func isOdd(i int) bool {
	return i%2 == 1
}

type RandomWords struct {
}

func (*RandomWords) CreateState() State {
	return &RandomWordsState{suggestions: make([]wordpairs.WordPair, 0, 10)}
}

type RandomWordsState struct {
	suggestions []wordpairs.WordPair
}

func (rws *RandomWordsState) Build(context StatefulContext) (Widget, error) {
	return &Composite{Child: &listview.Builder{
		Padding: EdgeInsetsAll(16),
		ItemBuilder: func(i int) Widget {
			if isOdd(i) {
				return &Divider{}
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

func BuildRow(wp wordpairs.WordPair) Widget {
	return &listview.ListTile{
		Title: &Text{
			Text: wp.AsPascalCase(),
		},
	}
}
