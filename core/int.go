package core

/*

Int ...
*/
type Int int

/*

Eq ...
*/
func (a Int) Eq(b Value) bool {
	switch v := b.(type) {
	case Int:
		return a == v
	default:
		return false
	}
}
