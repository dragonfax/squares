package main

import (
	"github.com/dragonfax/glitter/examples/wordlist/wordpairs"
	"github.com/dragonfax/glitter/glt"
	"github.com/dragonfax/glitter/glt/listview"
)

func main() {
	err := glt.RunApp(&MyApp{})
	if err != nil {
		panic(err)
	}
}

type MyApp struct {
}

func (ma *MyApp) Build() (glt.Widget, error) {

	return &glt.Center{Child: &RandomWords{}}, nil
}

func isOdd(i int) bool {
	return i%2 == 1
}

type RandomWords struct {
}

func (*RandomWords) CreateState() glt.State {
	return &RandomWordsState{make([]wordpairs.WordPair, 0, 10)}
}

type RandomWordsState struct {
	suggestions []wordpairs.WordPair
}

func (rws *RandomWordsState) Build() glt.Widget {
	return listview.Builder{
		Padding: glt.EdgeInsets{All: 16.0},
		ItemBuilder: func(i int) glt.Widget {
			if isOdd(i) {
				return &glt.Divider{}
			}

			r := i / 2
			if r >= len(rws.suggestions) {
				for i := 0; i < 10; i++ {
					rws.suggestions = append(rws.suggestions, wordpairs.GenerateWordPair())
				}
			}

			return BuildRow(rws.suggestions[r])
		},
	}
}

func BuildRow(wp wordpairs.WordPair) glt.Widget {
	return &listview.ListTile{
		Title: &glt.Text{
			Text: wp.AsPascalCase(),
		},
	}
}
