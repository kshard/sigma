/*

  Copyright 2016 Dmitry Kolesnikov, All Rights Reserved

  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.

*/

package vm

//
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
		vx := (*heap)[addr].(*any)
		vy := (*heap)[eq.y[i]].(*any)
		if *vx != *vy {
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
		vx := (*heap)[addr].(*any)
		vy := (*heap)[eq.y[i]].(*any)
		if *vx != *vy {
			return eq.Read(heap)
		}
	}

	return nil
}
