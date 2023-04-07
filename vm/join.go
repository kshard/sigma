/*

  Sigma Virtual Machine
  Copyright (C) 2016  Dmitry Kolesnikov

  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU Affero General Public License as published
  by the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.

  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU Affero General Public License for more details.

  You should have received a copy of the GNU Affero General Public License
  along with this program.  If not, see <https://www.gnu.org/licenses/>.

*/

package vm

//
// The file defines join (⨝) operator over streams
//

// Join is a fundamental stream combinator, it builds a new stream by
// evaluating a tail stream for each element of head stream.
//
//	for err := head.Init(&h); err == nil; err = head.Read(&h) {
//	  for err := tail.Init(&h); err == nil; err = tail.Read(&h) {
//	    ...
//	  }
//	}
func Join(lhs, rhs Stream) Stream {
	return &join{lhs: lhs, rhs: rhs}
}

// stream left join operator
type join struct {
	lhs Stream
	rhs Stream
}

func (join *join) Init(heap *Heap) error {
	if err := join.lhs.Init(heap); err != nil {
		return err
	}

	if err := join.rhs.Init(heap); err != nil {
		return err
	}

	return nil
}

func (join *join) Read(heap *Heap) error {
	if err := join.rhs.Read(heap); err == nil {
		return nil
	}

	if err := join.lhs.Read(heap); err != nil {
		return err
	}

	if err := join.rhs.Init(heap); err != nil {
		return err
	}

	return nil
}
