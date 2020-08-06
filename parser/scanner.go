package parser

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"unicode"
)

// Scanner holds the state of the scanner.
type Scanner struct {
	r         bufio.Reader // input reader
	peekRunes []rune       // peek runes queue
	buf       bytes.Buffer // input buffer to hold literals
}

// NewScanner creates a new Scanner.
func NewScanner(r *bufio.Reader) *Scanner {
	return &Scanner{
		r: *r,
	}
}

// nextRune reads the next rune from the input.
func (s *Scanner) nextRune() rune {
	r, _, err := s.r.ReadRune()
	if err != nil {
		if err != io.EOF {
			fmt.Fprintln(os.Stderr)
		}
		r = -1 // EOF rune
	}
	return r
}

// read consumes the peekRunes queue then calls nextRune.
func (s *Scanner) read() rune {
	if len(s.peekRunes) > 0 {
		r := s.peekRunes[0]
		s.peekRunes = s.peekRunes[1:]
		return r
	}
	return s.nextRune()
}

// skipSpace ignores spaces and returns the next rune
func (s *Scanner) skipSpace() rune {
	r := s.read()
	if unicode.IsSpace(r) {
		return s.skipSpace()
	}
	return r
}

// peek returns but does not consume the next n rune in the input.
func (s *Scanner) peek(n int) rune {
	if len(s.peekRunes) >= n {
		return s.peekRunes[n-1]
	}

	var p rune

	for i := 1; i <= n; i++ {
		p = s.nextRune()
		s.peekRunes = append(s.peekRunes, p)
	}

	return p
}

// next returns the next token.
func (s *Scanner) next() *Token {
	for {
		r := s.read()
		switch {
		case unicode.IsSpace(r):
		case isInstruction(r):
			return s.instructions(r)

		case unicode.IsDigit(r):
			if isCardinal(s.peek(4)) {
				return s.mower(r)
			}
			return s.lawn(r)
		case r == -1:
			return &Token{EOF, ""}
		default:
			return s.error(r)
		}
	}
}

// accum appends the current rune to the buffer until
// the valid function returns false
func (s *Scanner) accum(r rune, valid func(rune) bool) {
	s.buf.Reset()
	for {
		s.buf.WriteRune(r)
		r = s.read()
		if r == -1 {
			return
		}
		if !valid(r) {
			return
		}
	}
}

// instructions scans an instruction set
func (s *Scanner) instructions(r rune) *Token {
	s.accum(r, isInstruction)
	return &Token{Instructions, s.buf.String()}
}

// mower scans a mower configuration
func (s *Scanner) mower(r rune) *Token {
	c := string(r)

	for i := 1; i <= 4; i++ {
		c += string(s.read())
	}

	return &Token{Mower, c}
}

// lawn scans a lawn configuration
func (s *Scanner) lawn(r rune) *Token {
	c := string(r)        // x
	c += string(s.read()) // space
	c += string(s.read()) // y

	return &Token{Lawn, c}
}

func (s *Scanner) error(r rune) *Token {
	log.Printf("syntax error: illegal character %d", r)
	return &Token{SyntaxError, "illegal character"}
}

func isInstruction(r rune) bool {
	return r == 'L' || r == 'R' || r == 'F' || r == 'B'
}

func isCardinalOrDigit(r rune) bool {
	return isCardinal(r) || unicode.IsDigit(r)
}

func isCardinal(r rune) bool {
	return r == 'N' || r == 'E' || r == 'S' || r == 'W'
}
