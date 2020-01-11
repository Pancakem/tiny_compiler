package main

import (
	//"os"
	"bufio"
	"unicode"
)

const (
	EOP = -1
	_   = iota
	ID
	NUM
	OP1
	OP2
	LBR
	RBR
	SEM
	EQ
)

const MAX_LEN = 100

type Token struct {
	tokenType int
	attr      int
}

var (
	source   *bufio.Scanner
	lookDone = false
	lookTok  Token
)

func nextTok() Token {
	var ch rune
	var tok Token

	// If lookaheaded
	if lookDone {
		lookDone = false
		return lookTok
	}

eat:
	if source.Scan() {
		ch = rune(source.Text()[0])
		switch ch {
		case ' ':
		case '\n':
			goto eat
		case '+':
			tok.tokenType = OP1
			tok.attr = ADD_TYPE
		case '-':
			tok.tokenType = OP1
			tok.attr = SUB_TYPE
		case '*':
			tok.tokenType = OP2
			tok.attr = MUL_TYPE
		case '/':
			tok.tokenType = OP2
			tok.attr = DIV_TYPE
		case '(':
			tok.tokenType = LBR
		case ')':
			tok.tokenType = RBR
		case '=':
			tok.tokenType = EQ
		case ';':
			tok.tokenType = SEM
		default:
			// ID
			if unicode.IsLetter(ch) {
				idName := make([]rune, 0)
				len := 0
				idName = append(idName, ch)
				len += 1

				for {
					ch = rune(source.Text()[0])
					idName = append(idName, ch)
					len += 1
					if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
						break
					}
				}

				// idName[len-1] = \0
				idName = append(idName, '0')

				tok.tokenType = ID
				tok.attr = addSym(string(idName))
				// NUM
			} else if unicode.IsDigit(ch) {
				val := ch - '0'

				for source.Scan() {
					if unicode.IsDigit(rune(source.Text()[0])) {
						val = val*10 + (ch - '0')
					} else {
						break
					}
				}

				tok.tokenType = NUM
				tok.attr = int(val)
				// ERROR
			} else {
				fatalError("Lexer: unexpected symbol")
			}
		}
	} else {
		// EOF
		tok.tokenType = EOP
	}

	return tok

}

func lookahead() Token {
	lookTok = nextTok()
	lookDone = true

	return lookTok
}
