package scanner

type Scanner struct {
  source string
}

func New(source string) *Scanner {
  return &Scanner{source: source}
}

func (s *Scanner) ScanTokens() {
  //
}
