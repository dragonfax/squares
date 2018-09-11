package main

import (
	"github.com/dragonfax/glitter/examples/wordlist/wordpairs"
	"github.com/dragonfax/glitter/glt"
	"github.com/dragonfax/glitter/glt/listview"
)

var _ glt.StatelessWidget = &MyApp{}
var _ glt.StatefulWidget = &RandomWords{}
var _ glt.State = &RandomWordsState{}

func main() {
	err := glt.RunApp(&MyApp{})
	if err != nil {
		panic(err)
	}
}

type MyApp struct {
}

func (ma *MyApp) Build(context glt.BuildContext) (glt.Widget, error) {

	return &glt.Center{Child: &RandomWords{}}, nil
}

func isOdd(i int) bool {
	return i%2 == 1
}

type RandomWords struct {
}

func (*RandomWords) CreateState() glt.State {
	return &RandomWordsState{suggestions: make([]wordpairs.WordPair, 0, 10)}
}

type RandomWordsState struct {
	suggestions []wordpairs.WordPair
}

func (rws *RandomWordsState) Build(context glt.BuildContext) (glt.Widget, error) {
	return &listview.Builder{
		Padding: glt.EdgeInsetsAll(16),
		ItemBuilder: func(i int) glt.Widget {
			if isOdd(i) {
				return &glt.Divider{}
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
	}, nil
}

func BuildRow(wp wordpairs.WordPair) glt.Widget {
	return &listview.ListTile{
		Title: &glt.Text{
			Text: wp.AsPascalCase(),
		},
	}
}
