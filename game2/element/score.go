package element

type Score struct {
	Left  int
	Right int
}

func (s *Score) IncLeft() {
	s.Left++
}

func (s *Score) IncRight() {
	s.Right++
}
