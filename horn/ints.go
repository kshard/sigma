package horn

/*

ints ...
*/
type ints struct {
	r   Atom
	seq []int
	ofs int
}

/*


Stream ...
*/
func IStream(r Atom, seq []int) Stream {
	return &ints{r: r, seq: seq, ofs: 0}
}

func (seq *ints) Read(heap *Heap) error {
	if len(seq.seq) == seq.ofs {
		return EOS
	}

	v := seq.seq[seq.ofs]
	heap.Join(seq.r, v)

	seq.ofs++
	return nil
}

func (seq *ints) Skip() {
	seq.ofs = 0
}

/*


Σ ...
*/
// type σ struct {
// 	seq []int
// }

// /*

// Σ ...
// */
// func IStreamΣ(seq []int) Σ {
// 	return σ{seq: seq}
// }

// func (σ σ) Stream(q Heap) Stream {
// 	sio := IStream(σ.seq)

// 	return sio
// }
