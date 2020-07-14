// Copyright 2020 Meador Inge.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package asm

import (
	"testing"
)

// Helper test functions.

func assertTokensEqual(t *testing.T, tok, etok *Token) {
	if tok.Kind != etok.Kind {
		t.Errorf("bad token: got '%s', expected '%s'", tok, etok)
	}
	if tok.Lit != etok.Lit {
		t.Errorf("bad token: got '%s', expected '%s'", tok, etok)
	}
}

func assertTokensEqualSource(t *testing.T, toks []*Token, str string) {
	var s Scanner
	s.Init([]byte(str))
	for _, etok := range toks {
		tok := s.Next()
		assertTokensEqual(t, tok, etok)
	}

}

var test_single_token = [...]*Token{
	NewToken(EOF, ""),
	NewToken(COMMENT, "# this is a comment!\n"),

	// Names.
	NewToken(NAME, "foo"),
	NewToken(NAME, "start"),
	NewToken(NAME, "a"),
	NewToken(NAME, "Z"),

	// Numbers.
	NewToken(NUMBER, "1"),
	NewToken(NUMBER, "98765"),

	// String constants.
	NewToken(STRINGCONST, "\"0\""),
	NewToken(STRINGCONST, "\"1\""),
	NewToken(STRINGCONST, "\">\""),

	// Operators.
	NewToken(COMMA, ","),
	NewToken(COLON, ":"),

	// Instructions.
	NewToken(CWRITE, "cwrite"),
	NewToken(HALT, "halt"),
}

func TestSingleToken(t *testing.T) {
	var s Scanner

	for _, etok := range test_single_token {
		s.Init([]byte(etok.Lit))
		tok := s.Next()
		assertTokensEqual(t, tok, etok)
	}
}

var test_program_str = `
start:
cwrite "0", "0", left, start
halt
`

var test_program_tokens = []*Token{
	NewToken(NAME, "start"),
	NewToken(COLON, ":"),
	NewToken(CWRITE, "cwrite"),
	NewToken(STRINGCONST, "\"0\""),
	NewToken(COMMA, ","),
	NewToken(STRINGCONST, "\"0\""),
	NewToken(COMMA, ","),
	NewToken(NAME, "left"),
	NewToken(COMMA, ","),
	NewToken(NAME, "start"),
	NewToken(HALT, "halt"),
}

func TestMultiple(t *testing.T) {
	assertTokensEqualSource(t, test_program_tokens, test_program_str)
}
