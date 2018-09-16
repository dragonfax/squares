package squares

/* Some biolerplate code duplicated for various structs.
 * A results of certain intential missing features in golang.
 * This code could get auto-generated in the future, if I'm so inclined
 */

func (ce Center) getChild() Widget {
	return ce.Child
}

func (p Padding) getChild() Widget {
	return p.Child
}

func (c Column) getChildren() []Widget {
	return c.Children
}

func (cw Composite) getChild() Widget {
	return cw.Child
}

func (sb SizedBox) getChild() Widget {
	return sb.Child
}
