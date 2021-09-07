package core

/*


Seq ...
*/
type Seq []Value

/*

Merge ...
*/
func (seq Seq) Merge(Values) Values {
	return seq
}

/*

Len ...
*/
func (seq Seq) Len() int {
	return len(seq)
}

/*

Get ...
*/
func (seq Seq) Get(key Value) Value {
	switch v := key.(type) {
	case Int:
		return seq[v]
	}

	return Any
}

/*

Copy ...
*/
func (seq Seq) Copy(keys []Value, values Values) Values {
	for i, key := range keys {
		seq[i] = values.Get(key)
	}
	return seq
}
