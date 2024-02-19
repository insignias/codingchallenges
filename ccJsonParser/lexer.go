package main

type TokenType string

type Token struct {
	Type TokenType
	Value string
}

var (
	Illegal TokenType = "ILLEGAL"
	EOF TokenType = "EOF"
	// Structure
	BeginObject TokenType = "{"
	EndObject TokenType = "}"
)

type Lexer struct {
	input string
	lastPosition int
	currentPosition int
	ch byte
}

func Newlexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar(){
	if l.currentPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.currentPosition]
	}
	l.lastPosition = l.currentPosition
	l.currentPosition += 1
}

func (l *Lexer) NextToken() Token{
	var tok Token

	switch l.ch {
	case '{':
		tok = Token{BeginObject, string(l.ch)}
	case '}':
		tok = Token{EndObject, string(l.ch)}
	case 0:
		tok = Token{EOF, ""}
	default:
		tok = Token{Illegal, string(l.ch)}
	}
	l.readChar()
	return tok
}