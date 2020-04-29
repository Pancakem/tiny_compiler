package main

import (
	"bufio"
	"unicode"
)

const (
	EOP = -1
	ID  = 1
	NUM = 2
	OP1 = 3
	OP2 = 4
	LBR = 5
	RBR = 6
	SEM = 7
	EQ  = 8
)

const MAX_LEN = 100

type Token struct {
	tokenType int
	attr      int
}

var (
	source   *bufio.Reader
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

	ch, _, err := source.ReadRune()
	if err == nil {
		switch ch {
		case ' ':
			fallthrough
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
				idName = append(idName, ch)

				for unicode.IsLetter(ch) || unicode.IsDigit(ch) {
					ch, _, err = source.ReadRune()
					if err != nil {
						fatalError("lexer: unexpected end of input")
					}
					idName = append(idName, ch)
				}

				tok.tokenType = ID
				tok.attr = addSym(string(idName))
				// NUM
			} else if unicode.IsDigit(ch) {
				val := int(ch - '0')
				for {
					ch, _, err := source.ReadRune()
					if err != nil {
						break
					}
					if unicode.IsDigit(ch) {
						val = val*10 + int(ch-'0')
					} else {
						break
					}
				}

				tok.tokenType = NUM
				tok.attr = val
				// ERROR
			} else {
				fatalError("lexer: unexpected symbol")
			}

			source.UnreadRune()
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
