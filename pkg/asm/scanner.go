// Copyright 2020 Meador Inge.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package asm

type Scanner struct {
	src      []byte // The source code.
	ch       rune   // The current character.
	chOffset int    // The current character offset.
	offset   int    // The next character offset.
	savedTok *Token // A saved token from an earlier scan.
}

func (s *Scanner) next() {
	if s.offset < len(s.src) {
		s.chOffset = s.offset
		s.ch = rune(s.src[s.offset])
		s.offset += 1
	} else {
		s.chOffset = s.offset
		s.offset = len(s.src)
		s.ch = -1
	}
}

func (s *Scanner) skipWhitespace() {
	for s.ch == ' ' || s.ch == '\t' || s.ch == '\n' || s.ch == '\r' {
		s.next()
	}
}

func (s *Scanner) isLetter(ch rune) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z'
}

func (s *Scanner) isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9'
}

func (s *Scanner) scanComment() string {
	start := s.chOffset - 1

	for s.ch != '\n' {
		s.next()
	}

	return string(s.src[start:s.offset])
}

func (s *Scanner) scanName() *Token {
	start := s.chOffset
	for s.isLetter(s.ch) || s.isDigit(s.ch) {
		s.next()
	}

	str := s.src[start:s.chOffset]
	kind := NAME
	literal := string(str)
	if len(str) > 1 {
		kind = LookupName(literal)
	}
	return NewToken(kind, literal)
}

func (s *Scanner) scanNumber() *Token {
	start := s.chOffset
	for s.isDigit(s.ch) {
		s.next()
	}
	return NewToken(NUMBER, string(s.src[start:s.chOffset]))
}

func (s *Scanner) scanStringConst() *Token {
	start := s.offset - 1
	s.next()
	for s.ch != '"' {
		s.next()
	}
	s.next()
	return NewToken(STRINGCONST, string(s.src[start:s.chOffset]))
}

func (s *Scanner) scanOperator(ch rune) *Token {
	kind := ILLEGAL
	lit := string(ch)
	s.next()

	switch ch {
	case ',':
		kind = COMMA
	case '#':
		kind, lit = COMMENT, s.scanComment()
	case ':':
		kind = COLON
	case -1:
		kind, lit = EOF, ""
	}

	return NewToken(kind, lit)
}

func (s *Scanner) Init(src []byte) {
	s.src = src
	s.ch = ' '
	s.offset = 0
	s.savedTok = nil
	s.next()
}

func (s *Scanner) Next() (tok *Token) {
next:
	if s.savedTok != nil {
		tok = s.savedTok
		s.savedTok = nil
	} else {
		s.skipWhitespace()

		switch ch := s.ch; {
		case s.isLetter(ch):
			tok = s.scanName()
		case s.isDigit(ch):
			tok = s.scanNumber()
		case ch == '"':
			tok = s.scanStringConst()
		case ch == '\n':
			goto next
		default:
			tok = s.scanOperator(ch)
		}
	}

	return
}
