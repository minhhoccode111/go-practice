package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

/*
I want to support DSL like this:

```txt
[rd_name1] == true && [rd_name1] == [rd_another_name1] || ([d_num2] != 5 && [t_name3] == "hello")
```

Grammar (EBNF-ish)

```txt
expression  = or_expr ;

or_expr     = and_expr { "||" and_expr } ;
and_expr    = equality_expr { "&&" equality_expr } ;
equality_expr = relational_expr { ("==" | "!=") relational_expr } ;
relational_expr = primary { (">" | "<" | ">=" | "<=") primary } ;

primary     = identifier | literal | "(" expression ")" ;

identifier  = "[" name "]" ;
literal     = string | number | boolean ;
```
*/

func main() {
	fmt.Print("input: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	lexer := NewLexer(input)
	parser := NewParser(lexer)
	ast := parser.Parse()

	PrintAST(ast, 0)
	PrintJSONOfAST(ast)
}

func PrintAST(expr Expr, indent int) {
	prefix := strings.Repeat("  ", indent)

	switch e := expr.(type) {
	case *BinaryExpr:
		fmt.Printf("%sBinaryExpr (%s)\n", prefix, e.Op)
		PrintAST(e.Left, indent+1)
		PrintAST(e.Right, indent+1)

	case *Identifier:
		fmt.Printf("%sIdentifier: [%s]\n", prefix, e.Name)

	case *Literal:
		fmt.Printf("%sLiteral: %v (%T)\n", prefix, e.Value, e.Value)
	}
}

func PrintJSONOfAST(expr Expr) {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "  ")
	if err := enc.Encode(expr); err != nil {
		panic(err)
	}
	fmt.Println(buf.String())
}

type Expr any

type BinaryExpr struct {
	Kind  string    `json:"kind"` // can skip this
	Left  Expr      `json:"left"`
	Op    TokenType `json:"op"`
	Right Expr      `json:"right"`
}

type Identifier struct {
	Kind string `json:"kind"` // can skip this
	Name string `json:"name"`
}

type Literal struct {
	Kind  string `json:"kind"` // can skip this
	Value any    `json:"value"`
}

type TokenType string

const (
	IDENT  TokenType = "IDENT"
	STRING TokenType = "STRING"
	NUMBER TokenType = "NUMBER"
	BOOL   TokenType = "BOOL"

	AND TokenType = "&&"
	OR  TokenType = "||"

	EQ  TokenType = "=="
	NEQ TokenType = "!="

	GT  TokenType = ">"
	LT  TokenType = "<"
	GTE TokenType = ">="
	LTE TokenType = "<="

	LPAREN TokenType = "("
	RPAREN TokenType = ")"

	EOF TokenType = "EOF"
)

type Token struct {
	Type  TokenType
	Value string
}

type Lexer struct {
	input    []rune
	position int
}

func NewLexer(input string) *Lexer {
	return &Lexer{input: []rune(input)}
}

func (l *Lexer) NextToken() Token {
	for l.position < len(l.input) && l.input[l.position] == ' ' {
		l.position++
	}

	if l.position >= len(l.input) {
		return Token{Type: EOF}
	}

	ch := l.input[l.position]

	if ch == '[' {
		start := l.position + 1
		for l.position < len(l.input) && l.input[l.position] != ']' {
			l.position++
		}
		name := string(l.input[start:l.position])
		l.position++
		return Token{Type: IDENT, Value: name}
	}

	if ch == '"' {
		l.position++
		start := l.position
		for l.position < len(l.input) && l.input[l.position] != '"' {
			l.position++
		}
		val := string(l.input[start:l.position])
		l.position++
		return Token{Type: STRING, Value: val}
	}

	switch ch {
	case '&':
		l.position += 2
		return Token{Type: AND}
	case '|':
		l.position += 2
		return Token{Type: OR}
	case '=':
		l.position += 2
		return Token{Type: EQ}
	case '!':
		l.position += 2
		return Token{Type: NEQ}
	case '>':
		if l.peek() == '=' {
			l.position += 2
			return Token{Type: GTE}
		}
		l.position++
		return Token{Type: GT}
	case '<':
		if l.peek() == '=' {
			l.position += 2
			return Token{Type: LTE}
		}
		l.position++
		return Token{Type: LT}
	case '(':
		l.position++
		return Token{Type: LPAREN}
	case ')':
		l.position++
		return Token{Type: RPAREN}
	}

	// Boolean: true / false
	if ch == 't' || ch == 'f' {
		start := l.position
		for l.position < len(l.input) && isLetter(l.input[l.position]) {
			l.position++
		}
		word := string(l.input[start:l.position])
		if word == "true" || word == "false" {
			return Token{Type: BOOL, Value: word}
		}
		panic("unknown identifier: " + word)
	}

	// Number: digits, optional decimal
	if isDigit(ch) {
		start := l.position
		for l.position < len(l.input) && isDigit(l.input[l.position]) {
			l.position++
		}
		if l.position < len(l.input) && l.input[l.position] == '.' {
			l.position++
			for l.position < len(l.input) && isDigit(l.input[l.position]) {
				l.position++
			}
		}
		return Token{Type: NUMBER, Value: string(l.input[start:l.position])}
	}

	panic("unknown token")
}

func (l *Lexer) peek() rune {
	if l.position+1 >= len(l.input) {
		return 0
	}
	return l.input[l.position+1]
}

func isDigit(ch rune) bool { return ch >= '0' && ch <= '9' }

func isLetter(
	ch rune,
) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_'
}

type Parser struct {
	lexer   *Lexer
	current Token
}

func NewParser(l *Lexer) *Parser {
	p := &Parser{lexer: l}
	p.next()
	return p
}

func (p *Parser) next() {
	p.current = p.lexer.NextToken()
}

func (p *Parser) Parse() Expr {
	return p.parseOr()
}

func (p *Parser) parseOr() Expr {
	left := p.parseAnd()

	for p.current.Type == OR {
		op := p.current.Type
		p.next()
		right := p.parseAnd()
		left = &BinaryExpr{Left: left, Op: op, Right: right, Kind: "binary"}
	}

	return left
}

func (p *Parser) parseAnd() Expr {
	left := p.parseEquality()

	for p.current.Type == AND {
		op := p.current.Type
		p.next()
		right := p.parseEquality()
		left = &BinaryExpr{Left: left, Op: op, Right: right, Kind: "binary"}
	}

	return left
}

func (p *Parser) parseEquality() Expr {
	left := p.parseRelational()

	for p.current.Type == EQ || p.current.Type == NEQ {
		op := p.current.Type
		p.next()
		right := p.parseRelational()
		left = &BinaryExpr{Left: left, Op: op, Right: right, Kind: "binary"}
	}

	return left
}

func (p *Parser) parseRelational() Expr {
	left := p.parsePrimary()

	for p.current.Type == GT || p.current.Type == LT ||
		p.current.Type == GTE || p.current.Type == LTE {

		op := p.current.Type
		p.next()
		right := p.parsePrimary()
		left = &BinaryExpr{Left: left, Op: op, Right: right, Kind: "binary"}
	}

	return left
}

func (p *Parser) parsePrimary() Expr {
	switch p.current.Type {

	case IDENT:
		val := p.current.Value
		p.next()
		return &Identifier{Name: val, Kind: "identifier"}

	case STRING:
		val := p.current.Value
		p.next()
		return &Literal{Value: val, Kind: "literal"}

	case LPAREN:
		p.next()
		expr := p.parseOr()
		if p.current.Type != RPAREN {
			panic("missing )")
		}
		p.next()
		return expr

	case BOOL:
		val := p.current.Value == "true" // convert to actual bool
		p.next()
		return &Literal{Value: val, Kind: "literal"}

	case NUMBER:
		raw := p.current.Value
		p.next()
		// store as float64, same as JSON/most dynamic languages
		var n float64
		fmt.Sscanf(raw, "%f", &n)
		return &Literal{Value: n, Kind: "literal"}
	}

	panic("unexpected token")
}
