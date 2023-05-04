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

import "github.com/kshard/xsd"

type eq struct {
	x      []Addr
	y      []Addr
	stream Stream
}

func Eq(x, y []Addr, stream Stream) Stream {
	return &eq{x: x, y: y, stream: stream}
}

func (eq *eq) Init(heap *Heap) error {
	if err := eq.stream.Init(heap); err != nil {
		return err
	}

	for i, addr := range eq.x {
		vx := (*heap)[addr]
		vy := (*heap)[eq.y[i]]
		if xsd.Compare(vx, vy) != 0 {
			return eq.Read(heap)
		}
	}

	return nil
}

func (eq *eq) Read(heap *Heap) error {
	if err := eq.stream.Read(heap); err != nil {
		return err
	}

	for i, addr := range eq.x {
		vx := (*heap)[addr]
		vy := (*heap)[eq.y[i]]
		if xsd.Compare(vx, vy) != 0 {
			return eq.Read(heap)
		}
	}

	return nil
}
