package flutter

type Divider struct {
	sizeData
	parentData
}

func (ce *Divider) layout(c constraints) error {

	ce.size = Size{width: c.maxWidth, height: 5}

	return nil
}

const CHARACTER_WIDTH = 10
const CHARACTER_HEIGHT = 10

type Text struct {
	Text string
	sizeData
	parentData
}

func (t *Text) layout(c constraints) error {
	cWidth := len(t.Text) * CHARACTER_WIDTH
	cHeight := 1 * CHARACTER_HEIGHT

	t.size = Size{width: MaxUint16(uint16(cWidth), c.minWidth), height: MaxUint16(uint16(cHeight), c.minHeight)}

	return nil
}
