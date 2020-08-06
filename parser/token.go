package parser

// TokenType represents a token type
type TokenType int

// Token represents a token
type Token struct {
	typ TokenType
	lit string
}

// Definition of tokens
const (
	Lawn TokenType = iota
	Mower
	Instructions
	SyntaxError
	EOF
)
