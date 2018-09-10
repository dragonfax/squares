package glt

// var _ ElementWidget = &Container{}
// var _ HasChild = &Container{}

type Container struct {
	Child      Widget
	Padding    EdgeInsets
	Width      uint16
	Decoration BoxDecoration
}

type BoxDecoration struct {
	Border   Border
	Gradient LinearGradient
}

type DecoratedBox struct {
	Decoration BoxDecoration
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

type SafeArea struct {
	Top    bool
	Bottom bool
	Child  Widget
}

type Row struct {
	CrossAxisAlignment CrossAxisAlignment
	MainAxisAlignment  MainAxisAlignment
	Children           []Widget
}

type CrossAxisAlignment uint8

const (
	CrossAxisAlignmentStart CrossAxisAlignment = iota
)

type MainAxisAlignment uint8

const (
	MainAxisAlignmentSpaceBetween MainAxisAlignment = iota
)

type Expanded struct {
	Child Widget
}

type SizedBox struct {
	Width uint16
	Child Widget
}

type MergeSemantics struct {
	Child Widget
}

type CustomScrollView struct {
	Slivers []Widget
}

type SliverAppBar struct {
	ExpandedHeight uint16
	Pinned         bool
	Floating       bool
	Snap           bool
	Actions        []Widget
	FlexibleSpace  *FlexibleSpaceBar // Widget
}

type FlexibleSpaceBar struct {
	Title      Widget
	Background Widget
}

type Stack struct {
	Fit      StackFit
	Children []Widget
}

type StackFit uint8

const (
	StackFitExpand StackFit = iota
)

type SliverList struct {
	Delegate *SliverChildListDelegate
}

type SliverChildListDelegate struct {
	Children []Widget
}

type SystemUiOverlayStyle uint8

const (
	SystemUiOverlayStyleDark SystemUiOverlayStyle = iota
)

type AnnotatedRegion struct {
	Value SystemUiOverlayStyle
	Child Widget
}
