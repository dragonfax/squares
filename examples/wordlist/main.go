package main

import (
	"github.com/dragonfax/gltr/examples/wordlist/wordpairs"
	"github.com/dragonfax/gltr/gltr/listview"
)

var _ gltr.StatelessWidget = &MyApp{}
var _ gltr.StatefulWidget = &RandomWords{}
var _ gltr.State = &RandomWordsState{}

func main() {
	err := gltr.RunApp(&MyApp{})
	if err != nil {
		panic(err)
	}
}

type MyApp struct {
}

func (ma *MyApp) Build(context gltr.BuildContext) (gltr.Widget, error) {

	return &gltr.Center{Child: &RandomWords{}}, nil
}

func isOdd(i int) bool {
	return i%2 == 1
}

type RandomWords struct {
}

func (*RandomWords) CreateState() gltr.State {
	return &RandomWordsState{suggestions: make([]wordpairs.WordPair, 0, 10)}
}

type RandomWordsState struct {
	suggestions []wordpairs.WordPair
}

func (rws *RandomWordsState) Build(context gltr.BuildContext) (gltr.Widget, error) {
	return &gltr.Composite{Child: &listview.Builder{
		Padding: gltr.EdgeInsetsAll(16),
		ItemBuilder: func(i int) gltr.Widget {
			if isOdd(i) {
				return &gltr.Divider{}
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

func BuildRow(wp wordpairs.WordPair) gltr.Widget {
	return &listview.ListTile{
		Title: &gltr.Text{
			Text: wp.AsPascalCase(),
		},
	}
}
