package core

/*

Atom ...
*/
type Atom string

/*

Eq ...
*/
func (a Atom) Eq(b Value) bool {
	switch v := b.(type) {
	case Atom:
		return a == v
	default:
		return false
	}
}

/*

Any ...
*/
const Any = Atom("_")
