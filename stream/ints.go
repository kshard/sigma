package stream

// TODO: this is test

//
//
// type Int int

// func (v Int) String() *string {
// 	x := fmt.Sprintf("%d", v)
// 	return &x
// }

// func (v Int) Int() *int {
// 	return (*int)(&v)
// }

// type IntValues []Value

// func (seq IntValues) FMap(f func(Value)) {
// 	for _, v := range seq {
// 		f(v)
// 	}
// }

// func (seq IntValues) Merge(values Values) Values {
// 	switch v := values.(type) {
// 	case IntValues:
// 		return append(seq, v...)
// 	default:
// 		panic("xxx")
// 	}
// }

// func (seq IntValues) Value(Value) Value {
// 	return nil
// }

// func (seq IntValues) Len() int {
// 	return len(seq)
// }

// //
// //
// type ints struct {
// 	seq [][]int
// }

// //
// func Ints(seq [][]int) Stream {
// 	return ints{seq: seq}
// }

// func (seq ints) Head() Values {
// 	values := make(IntValues, len(seq.seq[0]))
// 	for k, v := range seq.seq[0] {
// 		values[k] = Int(v)
// 	}
// 	return values
// }

// func (seq ints) Tail() Stream {
// 	if len(seq.seq) == 1 {
// 		return nil
// 	}

// 	return ints{seq: seq.seq[1:]}
// }
