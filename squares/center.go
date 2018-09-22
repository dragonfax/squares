package squares

var _ StatelessWidget = &Center{}

type Center struct {
	Child Widget
}

func (c Center) Build(context StatelessContext) Widget {
	return Align{Alignment: AlignmentCenter, Child: c.Child}
}
