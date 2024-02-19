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
	ValueSeparator TokenType = ":"
	ObjectSeparator TokenType = ","
	//Literal
	String TokenType = "string"
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

func (l *Lexer) peekChar() byte {
	if l.currentPosition >= len(l.input) {
		return 0
	} 
	return l.input[l.currentPosition]
}

func (l *Lexer) NextToken() Token{
	var tok Token

	l.skipWhiteSpaces()

	switch l.ch {
	case '{':
		tok = Token{BeginObject, string(l.ch)}
	case '}':
		tok = Token{EndObject, string(l.ch)}
	case ':':
		tok = Token{ValueSeparator, string(l.ch)}
	case ',':
		tok = Token{ObjectSeparator, string(l.ch)}
	case '"':
		tok = l.readString()
	case 0:
		tok = Token{EOF, ""}
	default:
		tok = Token{Illegal, string(l.ch)}
	}
	l.readChar()
	return tok
}

func (l *Lexer) readString() Token {
	var tok Token
	position := l.lastPosition + 1

	for {
		l.readChar()
		if l.ch == '"' {
			tok = Token{String, l.input[position:l.lastPosition]}
			break
		} else if (l.ch == 0) {
			tok = Token{Illegal, l.input[position:l.lastPosition]}
			break
		}
	}
	return tok
}

func (l *Lexer) skipWhiteSpaces() {
	for {
		if l.ch == ' ' {
			l.readChar()
			continue
		} else if (l.ch == '\\' && (l.peekChar() == 't' || l.peekChar() == 'n' || l.peekChar() == 'r')) {
			l.readChar()
			l.readChar()
			continue
		}
		break
	}

}