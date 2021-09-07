package core

/*

Value ...
*/
type Value interface {
	Eq(Value) bool
}

/*

Values ...
*/
type Values interface {
	Len() int
	Get(Value) Value
	Merge(values Values) Values
}
