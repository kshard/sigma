package stream_test

// func TestX(t *testing.T) {
// 	a := stream.Ints([][]int{{1}, {2}, {3}})

// 	stream.ForEach(
// 		stream.Map(a, func(v stream.Values) stream.Values {
// 			fmt.Println(v)
// 			return nil
// 		}),
// 	)
// }

// func TestY(t *testing.T) {
// 	a := stream.Ints([][]int{{1}, {2}, {3}})
// 	b := stream.Ints([][]int{{1}, {2}, {3}})

// 	seq := stream.FMap(a, func(x stream.Values) stream.Stream {
// 		return stream.Map(b,
// 			func(y stream.Values) stream.Values {
// 				return y.Merge(x)
// 			})
// 	})

// 	stream.ForEach(
// 		stream.Map(seq, func(v stream.Values) stream.Values {
// 			fmt.Println(v)
// 			return nil
// 		}),
// 	)
// }

// func TestZ(t *testing.T) {
// 	a := stream.Ints([][]int{{1}, {2}, {3}})
// 	b := stream.Ints([][]int{{1}, {2}, {3}})

// 	seq := stream.Join([]stream.Stream{a, b})

// 	stream.ForEach(
// 		stream.Map(seq, func(v stream.Values) stream.Values {
// 			fmt.Println(v)
// 			return nil
// 		}),
// 	)
// }

// func TestFilter(t *testing.T) {
// 	fmt.Println("===========")
// 	a := stream.Ints([][]int{{1}, {2}, {3}})

// 	seq := stream.Filter(a, func(v stream.Values) bool {
// 		switch v := v.(type) {
// 		case stream.IntValues:
// 			return *v[0].Int()%2 == 0
// 		default:
// 			panic("zzz")
// 		}
// 	})

// 	stream.ForEach(
// 		stream.Map(seq, func(v stream.Values) stream.Values {
// 			fmt.Println(v)
// 			return nil
// 		}),
// 	)
// }
