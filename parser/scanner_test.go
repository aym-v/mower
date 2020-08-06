package parser

import (
	"bufio"
	"strings"
	"testing"

	"github.com/valaymerick/mower/test"
)

func TestToken(t *testing.T) {
	in := `
	5 5
	1 2 N
	LFLFRFRFBBF
	3 4 S
	BBBFZ
	`

	out := []struct {
		expTyp TokenType
		expLit string
	}{
		{Lawn, "5 5"},
		{Mower, "1 2 N"},
		{Instructions, "LFLFRFRFBBF"},
		{Mower, "3 4 S"},
		{Instructions, "BBBF"},
		{EOF, ""},
	}

	r := bufio.NewReader(strings.NewReader(in))

	l := NewScanner(r)

	for _, c := range out {
		tok := l.next()

		test.AssertEqual(t, tok.typ, c.expTyp)
		test.AssertEqual(t, tok.lit, c.expLit)
	}
}
