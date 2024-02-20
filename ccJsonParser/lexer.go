package main

import (

	"unicode"
)

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
	BeginArray TokenType = "["
	EndArray TokenType = "]"
	ValueSeparator TokenType = ":"
	ObjectSeparator TokenType = ","
	//Literal
	String TokenType = "string"
	Number TokenType = "number"
	True TokenType = "true"
	False TokenType = "false"
	Null TokenType = "null"
	//Operator
	Minus TokenType = "-"
	Dot TokenType = "."
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
	case '[':
		tok = Token{BeginArray, string(l.ch)}
	case ']':
		tok = Token{EndArray, string(l.ch)}
	case ':':
		tok = Token{ValueSeparator, string(l.ch)}
	case ',':
		tok = Token{ObjectSeparator, string(l.ch)}
	case '"':
		tok = l.readString('"')
	case '\'':
		tok = l.readString('\'')
	case 0:
		tok = Token{EOF, ""}
	default:
		if (l.isDigit(l.ch) || l.ch == '-' || l.ch == '.') {
			tok = l.readNumber()
		} else if (l.isLiteral(l.ch)) {
			tok = l.readLiteral()
		} else {
			tok = l.readRegex()
		}
	}
	l.readChar()
	return tok
}


func (l *Lexer) readRegex() Token {
	position := l.lastPosition
	for {
		if l.peekChar() == ' ' || l.peekChar() == '\\' || l.peekChar() == ',' || l.peekChar() == ':' {
			break
		}
		l.readChar()
	}
	return Token{Illegal, l.input[position:l.currentPosition]}
}

func (l *Lexer) isDigit(input byte) bool {
	return unicode.IsNumber(rune(input))
}

func (l *Lexer) readNumber() Token {
	position := l.lastPosition
	for (l.isDigit(l.peekChar()) || l.peekChar() == '.')  {
		l.readChar()
	}
	return Token{Number, l.input[position:l.currentPosition]}
}

func (l *Lexer) isLiteral(ch byte) bool {
	return ch == 't' || ch == 'f' || ch == 'n'
}

func (l *Lexer) readLiteral() Token {
	var tok Token
	position := l.lastPosition

	switch l.ch {
	case 't':
		for i, c := range True[1:] {
			if rune(l.peekChar()) != c {
				return Token{Illegal, l.input[position:position+i]}
			}
			l.readChar()
		}
		tok = Token{True, "true"}
	case 'f':
		for i, c := range False[1:] {
			if rune(l.peekChar()) != c {
				return Token{Illegal, l.input[position:position+i]}
			}
			l.readChar()
		}
		tok = Token{False, "false"}
	case 'n':
		for i, c := range Null[1:] {
			if rune(l.peekChar()) != c {
				return Token{Illegal, l.input[position:position+i]}
			}
			l.readChar()
		}
		tok = Token{Null, "null"}
	}

	return tok
}

func (l *Lexer) readString(val byte) Token {
	var tok Token
	position := l.lastPosition + 1

	for {
		l.readChar()
		if l.ch == val {
			if val == '"' {
				tok = Token{String, l.input[position:l.lastPosition]}
			} else if val == '\'' {
				tok = Token{Illegal, l.input[position:l.lastPosition]}
			}
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