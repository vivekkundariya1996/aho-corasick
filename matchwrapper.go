package ahocorasick

// MatchWrapper contains all matched pattern in the input.
type MatchWrapper struct {
	total int
	match []*Match
}

func newMatchWrapper(matches []*Match) *MatchWrapper {
	return &MatchWrapper{total: len(matches), match: matches}
}

// GetMatch return  get the match at index
func (m *MatchWrapper) GetMatch(index int) []byte {
	return m.match[index].match
}

// TotalLength return total matched count
func (m *MatchWrapper) TotalLength() int {
	return m.total
}
