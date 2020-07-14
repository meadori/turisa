// Copyright 2020 Meador Inge.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package asm

type TokenKind int

type Token struct {
	Kind TokenKind
	Lit  string
}

const (
	ILLEGAL TokenKind = iota

	// End of file
	EOF

	// Comment
	COMMENT

	// Literals
	NAME
	NUMBER
	STRINGCONST

	// Operators
	COLON
	COMMA

	// Instructions
	reserved_begin
	CWRITE
	HALT
	reserved_end
)

var restoks = [...]string{
	ILLEGAL:     "ILLEGAL",
	EOF:         "EOF",
	COMMENT:     "COMMENT",
	NAME:        "NAME",
	NUMBER:      "NUMBER",
	STRINGCONST: "STRINGCONST",
	CWRITE:      "cwrite",
	HALT:        "halt",
	COLON:       ":",
	COMMA:       ",",
}

// Map from reserved system words to token kind.
var reswords map[string]TokenKind

func init() {
	reswords = make(map[string]TokenKind)
	for i := reserved_begin + 1; i < reserved_end; i++ {
		reswords[restoks[i]] = i
	}
}

// Lookup the given string and determine if it is a name or
// a reserved system word.
func LookupName(str string) TokenKind {
	if tok, is_reserved := reswords[str]; is_reserved {
		return tok
	}
	return NAME
}

// Create a new token.
func NewToken(kind TokenKind, lit string) *Token {
	t := new(Token)
	t.Kind = kind
	t.Lit = lit
	return t
}

// Return the string representation of the token.
func (tok Token) String() string {
	return tok.Lit
}

// Return the string representation of the token kind.
func (kind TokenKind) String() string {
	return restoks[kind]
}
