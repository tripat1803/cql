package config

// seperators
const (
	SEMICOLON         = "semicolon"
	OPEN_PARENTHESIS  = "open_parenthesis"
	CLOSE_PARENTHESIS = "close_parenthesis"
)

// operators
const (
	// arthimetic operators
	PLUS     = "plus"
	MINUS    = "minus"
	MULTIPLY = "multiply"
	DIVIDE   = "divide"

	// logical operators
	AND = "and"
	OR  = "or"

	// assignment operators
	EQUAL = "EQUAL"

	// comparision operators
	EQUALTO = "equalto"
)

// datatypes
const (
	STRING = "string"
	INT    = "int"
)

// keywords
const (
	LOAD     = "load"
	SELECT   = "select"
	INSERT   = "insert"
	SET      = "set"
	DELETE   = "delete"
	TRUNCATE = "truncate"
	WHERE    = "where"
	CREATE   = "create"
	AS       = "as"
	ID       = "id"
	SHOW     = "show"
	INTO     = "into"
	VALUES   = "values"
)

// token types
const (
	KEYWORD    = "keyword"
	SEPERATOR  = "seperator"
	IDENTIFIER = "identifier"
	STAR       = "STAR"
	COMMA      = "comma"
	OPERATOR   = "operator"
	DATATYPE   = "datatype"
	BEGINNING  = "beginning"
)
