package squares

var _ StatelessWidget = DecoratedBox{}
var _ HasChild = DecoratedBox{}
var _ StatelessWidget = SliverAppBar{}
var _ HasChildren = SliverAppBar{}
var _ StatelessWidget = FlexibleSpaceBar{}
var _ HasChildren = FlexibleSpaceBar{}

type BoxDecoration struct {
	Border   Border
	Gradient LinearGradient
}

type DecoratedBox struct {
	Decoration BoxDecoration
	Child      Widget
}

type LinearGradient struct {
	Begin  Alignment
	End    Alignment
	Colors []Color
}

type Alignment struct {
	X, Y float32
}

type Border struct {
	Bottom BorderSide
}

type BorderSide struct {
}
