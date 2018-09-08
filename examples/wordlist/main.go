package main

import (
	"github.com/dragonfax/glitter/listview"
	"github.com/dragonfax/glitter/wordpairs"
)

func main() {
	err := glt.RunApp(&MyApp{})
	if err != nil {
		panic(err)
	}
}

type MyApp struct {
}

func (ma *MyApp) Build(context *glt.BuildContext) (glt.Widget, error) {

	return &glt.Center{Child: NewRandomWords()}, nil
}

func isOdd(i int) bool {
	return i%2 == 1
}

var suggestions = make([]wordpairs.WordPair, 0)

func NewRandomWords() glt.Widget {
	return listview.Builder{
		Padding: glt.EdgeInsets{All: 16.0},
		ItemBuilder: func(context *glt.BuildContext, i int) glt.Widget {
			if isOdd(i) {
				return &glt.Divider{}
			}

			r := i / 2
			if r >= len(suggestions) {
				for i := 0; i < 10; i++ {
					suggestions = append(suggestions, wordpairs.GenerateWordPair())
				}
			}

			return BuildRow(suggestions[r])
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
