package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// IDENT Identifiers add, foobar, x, y, ...
	IDENT = "IDENT"
	// INT literals 12345
	INT = "INT"
	// STRING "wow!"
	STRING = "STRING"

	// Operators

	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"
	EQ       = "=="
	NOT_EQ   = "!="

	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"
	LPAREN    = "("
	RPAREN    = ")"
	LSQUIGGLE = "{"
	RSQUIGGLE = "}"
	LBRACKET  = "["
	RBRACKET  = "]"

	// Keywords

	FUNCTION = "FN"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var TokenTypes = map[byte]TokenType{
	'=': ASSIGN,
	';': SEMICOLON,
	':': COLON,
	'(': LPAREN,
	')': RPAREN,
	',': COMMA,
	'+': PLUS,
	'{': LSQUIGGLE,
	'}': RSQUIGGLE,
	'[': LBRACKET,
	']': RBRACKET,
	'-': MINUS,
	'!': BANG,
	'*': ASTERISK,
	'/': SLASH,
	'<': LT,
	'>': GT,
	0:   EOF,
}

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if toke, ok := keywords[ident]; ok {
		return toke
	}
	return IDENT
}
