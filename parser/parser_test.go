package parser

import (
	"bufio"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/valaymerick/mower/lawn"
	"github.com/valaymerick/mower/mower"
)

func TestParse(t *testing.T) {
	in := `
	5 5
	1 2 N
	LFLFLFLFF 
	3 3 E 
	FFRFFRFRRF
	`

	mowerA := mower.New(1, 2, mower.North)
	mowerA.Instruct([]mower.Move{
		mower.Left,
		mower.Forwards,
		mower.Left,
		mower.Forwards,
		mower.Left,
		mower.Forwards,
		mower.Left,
		mower.Forwards,
		mower.Forwards,
	})

	mowerB := mower.New(3, 3, mower.East)
	mowerB.Instruct([]mower.Move{
		mower.Forwards,
		mower.Forwards,
		mower.Right,
		mower.Forwards,
		mower.Forwards,
		mower.Right,
		mower.Forwards,
		mower.Right,
		mower.Right,
		mower.Forwards,
	})

	want := Config{
		Lawn: lawn.New(5, 5),
		Mowers: []*mower.Mower{
			&mowerA,
			&mowerB,
		},
	}

	p := NewParser(*bufio.NewReader(strings.NewReader(in)))
	got := p.Parse()

	if !cmp.Equal(*got, want) {
		t.Errorf("got %+v; want %+v", got, want)
	}
}

func TestParseLawn(t *testing.T) {
	for i, test := range []struct {
		lit  string
		want lawn.Lawn
	}{
		{"5 5", lawn.New(5, 5)},
		{"0 0", lawn.New(0, 0)},
		{"0 0 0", lawn.New(0, 0)},
	} {
		p := NewParser(bufio.Reader{})

		got := p.parseLawn(test.lit)

		if !cmp.Equal(got, test.want) {
			t.Errorf("test #%d: got %+v; want %+v", i, got, test.want)
		}
	}
}

func TestParseMower(t *testing.T) {
	for i, test := range []struct {
		lit  string
		want mower.Mower
	}{
		{"5 5 N", mower.New(5, 5, mower.North)},
		{"1 0 E", mower.New(1, 0, mower.East)},
		{"0 0 W", mower.New(0, 0, mower.West)},
	} {
		p := NewParser(bufio.Reader{})

		got := p.parseMower(test.lit)

		if !cmp.Equal(got, test.want) {
			t.Errorf("test #%d: got %+v; want %+v", i, got, test.want)
		}
	}
}

func TestParseInstructions(t *testing.T) {
	for i, test := range []struct {
		lit  string
		want []mower.Move
	}{
		{"LRFB", []mower.Move{mower.Left, mower.Right, mower.Forwards, mower.Backwards}},
	} {
		p := NewParser(bufio.Reader{})

		got := p.parseInstructions(test.lit)

		if !cmp.Equal(got, test.want) {
			t.Errorf("test #%d: got %v; want %v", i, got, test.want)
		}
	}
}
