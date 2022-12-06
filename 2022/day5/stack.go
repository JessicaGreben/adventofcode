package day5

type Stack struct {
	data []string
}

func (s *Stack) Len() int {
	return len(s.data)
}

func (s *Stack) Peek() string {
	if len(s.data) > 0 {
		return s.data[len(s.data)-1]
	}
	return ""
}

func (s *Stack) Pop() string {
	if len(s.data) > 0 {
		x := s.data[len(s.data)-1]
		s.data = s.data[:len(s.data)-1]
		return x
	}
	return ""
}

func (s *Stack) Push(val string) {
	s.data = append(s.data, val)
}
