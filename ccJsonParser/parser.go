package main

import (
	"fmt"
)

type Parser struct {
	l *Lexer
}

func NewParser(input string) *Parser {
	l := Newlexer(input)
	return &Parser{l: l}
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
	case String:
		value = tok.Value
	case Number:
		value = tok.Value
	case True:
		value = tok.Value
	case False:
		value = tok.Value
	case Null:
		value = nil
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

	for{
		if tok.Type != String {
			return obj, fmt.Errorf("expected key but got %s", tok.Value)
		} 

		key := tok.Value

		tok = p.l.NextToken()

		if tok.Type != ValueSeparator {
			return obj, fmt.Errorf("expected ':' but got %s", tok.Value)
		}

		tok = p.l.NextToken()

		value, err := p.ParseToken(tok)
		if err != nil {
			return value, err
		}

		obj[key] = value

		tok = p.l.NextToken()

		if tok.Type != ObjectSeparator {
			break
		}

		tok = p.l.NextToken()
	}

	if tok.Type != EndObject {
		return obj, fmt.Errorf("expected } but got %s", tok.Value)
	}
	
	return obj, err
}