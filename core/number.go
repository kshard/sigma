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

/*

Float64 ...
*/
type Float64 float64

/*

Eq ...
*/
func (a Float64) Eq(b Value) bool {
	switch v := b.(type) {
	case Float64:
		return a == v
	default:
		return false
	}
}
