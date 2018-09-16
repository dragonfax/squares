package squares

var _ StatelessWidget = DecoratedBox{}
var _ HasChild = DecoratedBox{}
var _ StatelessWidget = CustomScrollView{}
var _ HasChildren = CustomScrollView{}
var _ StatelessWidget = SliverAppBar{}
var _ HasChildren = SliverAppBar{}
var _ StatelessWidget = FlexibleSpaceBar{}
var _ HasChildren = FlexibleSpaceBar{}

// TODO var _ StatefulWidget = &SliverList{}
var _ HasChildren = &SliverList{}

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

type CustomScrollView struct {
	Slivers []Widget
}

type SliverAppBar struct {
	ExpandedHeight float64
	Pinned         bool
	Floating       bool
	Snap           bool
	Actions        []Widget
	FlexibleSpace  FlexibleSpaceBar // Widget
}

type FlexibleSpaceBar struct {
	Title      Widget
	Background Widget
}

type SliverList struct {
	Delegate SliverChildListDelegate
}

type SliverChildListDelegate struct {
	Children []Widget
}
