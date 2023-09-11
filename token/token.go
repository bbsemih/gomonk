package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT"
	INT   = "INT"
	FLOAT = "FLOAT"

	ASSIGN   = "="
	EQ       = "=="
	NOT_EQ   = "!="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERIKS = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"

	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	FUNCTION = "FUNCTION"
	TAN      = "TAN"
	IF       = "IF"
	ELSE     = "ELSE"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	RETURN   = "RETURN"
)

var keywords = map[string]TokenType{
	"fu":    FUNCTION,
	"tan":   TAN,
	"if":    IF,
	"else":  ELSE,
	"true":  TRUE,
	"false": FALSE,
	"re":    RETURN,
}

func LookUpIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
