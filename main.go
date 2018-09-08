package main

import (
	"github.com/dragonfax/flutter-go-example/flutter"
	"github.com/dragonfax/flutter-go-example/listview"
	"github.com/dragonfax/flutter-go-example/wordpairs"
)

func main() {
	err := flutter.RunApp(&MyApp{})
	if err != nil {
		panic(err)
	}
}

type MyApp struct {
}

func (ma *MyApp) Build(context *flutter.BuildContext) (flutter.Widget, error) {

	return &flutter.Center{Child: NewRandomWords()}, nil
}

func isOdd(i int) bool {
	return i%2 == 1
}

var suggestions = make([]wordpairs.WordPair, 0)

func NewRandomWords() flutter.Widget {
	return listview.Builder{
		Padding: flutter.EdgeInsets{All: 16.0},
		ItemBuilder: func(context flutter.BuildContext, i int) flutter.Widget {
			if isOdd(i) {
				return &flutter.Divider{}
			}

			r := i / 2
			if r >= len(suggestions) {
				suggestions = append(suggestions, wordpairs.RandomNum(10)...)
			}

			return BuildRow(suggestions[r])
		},
	}
}

func BuildRow(wp wordpairs.WordPair) flutter.Widget {
	return &listview.ListTile{
		Title: &flutter.Text{
			Text: wp.AsPascalCase(),
		},
	}
}
