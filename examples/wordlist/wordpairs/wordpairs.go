package wordpairs

import (
	"math/rand"
	"strings"
	"time"
)

var words = []string{
	"first",
	"second",
	"third",
	"fourth",
	"hello",
	"world",
	"person",
	"year",
	"way",
	"day",
	"thing",
	"man",
	"world",
	"life",
	"hand",
	"part",
	"child",
	"eye",
	"woman",
	"place",
	"work",
	"week",
	"case",
	"point",
	"government",
	"company",
	"number",
	"group",
	"problem",
	"fact",
}

type WordPair struct {
	First  string
	Second string
}

func (wp WordPair) AsPascalCase() string {
	return PascalCase(wp.First) + PascalCase(wp.Second)
}

func GenerateWordPair() WordPair {
	return WordPair{First: RandomWord(), Second: RandomWord()}
}

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

func RandomWord() string {
	return words[random.Intn(len(words))]
}

func PascalCase(s string) string {
	return strings.ToUpper(string(s[0])) + s[1:]
}
