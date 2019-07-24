package ahocorasick

// MatchWrapper contains all matched pattern in the input.
type MatchWrapper struct {
	total int
	curr  int
	match []*Match
}

func newMatchWrapper(matches []*Match) *MatchWrapper {
	return &MatchWrapper{curr: 0, total: len(matches), match: matches}
}

// NextString return next matched String
func (m *MatchWrapper) NextString() []byte {
	cur := m.curr
	m.curr = m.curr + 1
	return m.match[cur].match
}

// TotalLength return total matched count
func (m *MatchWrapper) TotalLength() int {
	return m.total
}
