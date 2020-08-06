package parser

import (
	"bufio"
	"errors"
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/google/go-cmp/cmp"
	"github.com/valaymerick/mower/lawn"
	"github.com/valaymerick/mower/mower"
)

// Config holds a mow configuration state
type Config struct {
	Lawn   lawn.Lawn
	Mowers []*mower.Mower
}

// Parser holds the state of the parser.
type Parser struct {
	s    *Scanner
	r    io.RuneReader // input reader
	buf  Token
	conf Config
}

// NewParser returns a new instance of Parser.
func NewParser(r bufio.Reader) *Parser {
	return &Parser{s: NewScanner(&r)}
}

func (p *Parser) scan() *Token {
	return p.s.next()
}

// Parse iterates over the scanner's tokens and returns a mow config object.
func (p *Parser) Parse() *Config {
	conf := &Config{}

	for {
		tok := p.scan()
		switch tok.typ {
		case Lawn:
			// Allow only one lawn configuration line (header)
			if !cmp.Equal(conf.Lawn, Config{}.Lawn) {
				p.error("only one lawn config is expected")
			}
			conf.Lawn = p.parseLawn(tok.lit)
		case Mower:
			mower := p.parseMower(tok.lit)
			nextTok := p.scan()
			if nextTok.typ != Instructions {
				p.error("expected an instruction set")
			}

			mower.Instruct(p.parseInstructions(nextTok.lit))

			conf.Mowers = append(conf.Mowers, &mower)
		case EOF:
			return conf
		}
	}
}

// parseLawn parses a lawn literal.
func (p *Parser) parseLawn(lit string) lawn.Lawn {
	s := strings.Split(lit, " ")
	x, err := strconv.Atoi(s[0])
	if err != nil {
		p.litError(lit)
	}

	y, err := strconv.Atoi(s[1])
	if err != nil {
		p.litError(lit)
	}

	return lawn.New(x, y)
}

// parseMower parses a mower literal.
func (p *Parser) parseMower(lit string) mower.Mower {
	s := strings.Split(lit, " ")
	x, err := strconv.Atoi(s[0])
	if err != nil {
		p.litError(lit)
	}

	y, err := strconv.Atoi(s[1])
	if err != nil {
		p.litError(lit)
	}

	o, err := getCardinal(s[2])
	if err != nil {
		p.litError(lit)
	}

	return mower.New(x, y, o)
}

// parseInstructions parses a literal containing a set of instructions.
func (p *Parser) parseInstructions(lit string) []mower.Move {
	var m []mower.Move

	for _, char := range lit {
		switch char {
		case 'F':
			m = append(m, mower.Forwards)
		case 'B':
			m = append(m, mower.Backwards)
		case 'L':
			m = append(m, mower.Left)
		case 'R':
			m = append(m, mower.Right)
		}
	}

	return m
}

// error raises a parser error.
func (p *Parser) error(msg string) {
	log.Fatalf("parser error: %s", msg)
}

// litError raises an error when a literal is illegal.
func (p *Parser) litError(lit string) {
	p.error("illegal literal " + lit)
}

func getCardinal(ch string) (int, error) {
	switch ch {
	case "N":
		return mower.North, nil
	case "E":
		return mower.East, nil
	case "S":
		return mower.South, nil
	case "W":
		return mower.West, nil
	default:
		return 0, errors.New("invalid cardinal")
	}
}
