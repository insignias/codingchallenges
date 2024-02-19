package main

import (
	"fmt"
)

type Parser struct {
	l *Lexer
}

func NewParser(input string) *Parser {
	l := Newlexer(input)
	return &Parser{l}
}

func (p *Parser) Parse() (interface{}, error){
	var output interface{}
	var err error

	tok := p.l.NextToken()

	output, err = p.ParseToken(tok)
	
	tok = p.l.NextToken()

	if tok.Type != EOF {
		err = fmt.Errorf("expected end of file but got %s", tok.Value)
	}

	return output, err

}

func (p *Parser) ParseToken(tok Token) (interface{}, error) {
	var value interface{}
	var err error

	switch tok.Type {
	case BeginObject:
		value, err = p.ParseObject(make(map[string]interface{}))
	case EOF:
		err = fmt.Errorf("unexpected end of file")
	}

	return value, err
}

func (p *Parser) ParseObject(obj map[string]interface{}) (interface{}, error){
	var err error
	tok := p.l.NextToken()

	if tok.Type == EndObject {
		return obj, err
	}
	return obj, err
}