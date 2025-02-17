package lexer

type Token struct {
	Token    string
	Category string
	Value    string
}

var IdentifierBreak = []string{
	";",
	",",
	"(",
	")",
}
