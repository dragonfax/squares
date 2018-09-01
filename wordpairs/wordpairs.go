package wordpairs

type WordPair struct {
}

func (wp WordPair) AsPascalCase() string {
	return "HelloWorld"
}

func RandomNum(i int) []WordPair {
	wpl := make([]WordPair, i)
	for x := 0; x < i; x++ {
		wpl[x] = WordPair{}
	}
	return wpl
}
