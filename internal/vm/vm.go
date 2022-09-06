package vm

func Eval(s Stream, h *Heap) error {
	for {
		s.Head(h)
		if err := s.Tail(h); err != nil {
			return err
		}
	}
}

func Debug(s Stream, h *Heap) error {
	for {
		s.Head(h)
		h.Dump()
		if err := s.Tail(h); err != nil {
			return err
		}
	}
}
