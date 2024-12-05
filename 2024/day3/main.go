package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)

const (
	EOF = iota
	ILLEGAL
	MUL
	COLON
	NUM
	BRACKET_L
	BRACKET_R
	DO
	DONT
)

var sequences = map[string]Token{
	"mul": MUL,
	// I cheated
	"do()":    DO,
	"don't()": DONT,
}

// :TODO could be removed
var tokens = []string{
	EOF:       "EOF",
	ILLEGAL:   "ILLEGAL",
	MUL:       "*",
	COLON:     ",",
	BRACKET_L: "(",
	BRACKET_R: ")",
	NUM:       "NUM",
	DO:        "DO",
	DONT:      "DON'T",
}

type Token int

func (t Token) String() string {
	return tokens[t]
}

// :TODO could be removed
type Position struct {
	line   int
	column int
}

type Lexer struct {
	pos    Position
	reader *bufio.Reader
}

func NewLexer(reader io.Reader) *Lexer {
	return &Lexer{
		pos:    Position{line: 1, column: 0},
		reader: bufio.NewReader(reader),
	}
}

func (l *Lexer) Lex() (Position, Token, string) {
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return l.pos, EOF, ""
			}

			panic(err)
		}

		l.pos.column++

		switch r {
		case '\n':
			l.resetPosition()
		case ',':
			return l.pos, COLON, ","
		case '*':
			return l.pos, MUL, "*"
		case '(':
			return l.pos, BRACKET_L, "("
		case ')':
			return l.pos, BRACKET_R, ")"
		default:
			if unicode.IsDigit(r) {
				startPos := l.pos
				l.backup()
				lit := l.lexInt()
				return startPos, NUM, lit
			}
			if unicode.IsLetter(r) {
				l.backup()
				if peek, err := l.reader.Peek(7); err == nil {
					for seq, tok := range sequences {
						if strings.Index(string(peek), seq) == 0 {
							startPos := l.pos
							for i := 0; i < len(seq); i++ {
								_, _, _ = l.reader.ReadRune()
								l.pos.column++
							}

							return startPos, tok, tok.String()
						}
					}
				}
				_, _, _ = l.reader.ReadRune()
			}
			return l.pos, ILLEGAL, string(r)
		}
	}
}

func (l *Lexer) lexInt() string {
	var lit string
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return lit
			}
		}

		l.pos.column++
		if unicode.IsDigit(r) {
			lit = lit + string(r)
		} else {
			l.backup()
			return lit
		}
	}
}

func (l *Lexer) backup() {
	if err := l.reader.UnreadRune(); err != nil {
		panic(err)
	}

	l.pos.column--
}

func (l *Lexer) resetPosition() {
	l.pos.line++
	l.pos.column = 0
}

type TokenData struct {
	Tok Token
	Val string
}

var fomulars = map[string][]Token{
	"mul": {MUL, BRACKET_L, NUM, COLON, NUM, BRACKET_R},
	// "do":   {MUL, BRACKET_L, NUM, COLON, NUM, BRACKET_R},
	// "dont": {MUL, BRACKET_L, NUM, COLON, NUM, BRACKET_R},
}

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()
	lexer := NewLexer(f)
	tokens := []TokenData{}

	for {
		_, tok, lit := lexer.Lex()
		if tok == EOF {
			break
		}

		tokens = append(tokens, TokenData{Tok: tok, Val: lit})
	}

	total := 0

	mul := fomulars["mul"]
	do := true
	for i := 0; i <= len(tokens)-len(mul); i++ {
		match := true
		res := 1

		if tokens[i].Tok == DO {
			do = true
		} else if tokens[i].Tok == DONT {
			do = false
		}

		for j := 0; j < len(mul); j++ {
			if tokens[i+j].Tok != mul[j] {
				match = false
				break
			}

			if mul[j] == NUM {
				num, _ := strconv.Atoi(tokens[i+j].Val)
				res *= num
			}
		}

		if match && do {
			total += res
			i += len(mul) - 1
		}
	}

	fmt.Println(total)
}
