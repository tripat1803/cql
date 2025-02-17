package lexer

import (
	"slices"
	"strings"
	"unicode"

	"tripat3k2/cql/config"
)

func get_string(fileContent *string, index *int) (token Token) {
	length := len(*fileContent)
	*index++
	value := make([]byte, 0)
	for *index < length {
		character := (*fileContent)[*index]
		if string(character) == `"` {
			break
		}
		value = append(value, character)
		*index++
	}
	token.Token = config.DATATYPE
	token.Category = config.STRING
	token.Value = string(value)
	return
}

func Get_Token(tokenType string, value string, categoryType string) (token Token) {
	token.Token = tokenType
	token.Category = categoryType
	if categoryType == "" {
		token.Category = "none"
	}
	token.Value = value
	return
}

func handle_comments(fileContent *string, index *int) {
	*index++
	length := len(*fileContent)
	if *index < length && string((*fileContent)[*index]) == `/` {
		for *index < length && (*fileContent)[*index] != 10 && (*fileContent)[*index] != 13 {
			*index++
		}
	}
}

func get_keyword_or_identifier(fileContent *string, index *int) (token Token) {
	length := len(*fileContent)
	literal := make([]byte, 0)
	for *index < length && (*fileContent)[*index] != 32 && !slices.Contains(IdentifierBreak, string((*fileContent)[*index])) {
		literal = append(literal, (*fileContent)[*index])
		*index++
	}
	if slices.Contains(IdentifierBreak, string((*fileContent)[*index])) {
		*index--
	}
	value := string(literal)
	selected_token := config.KEYWORD
	selected_category := ""
	switch strings.TrimSpace(strings.ToLower(value)) {
	case config.LOAD:
		selected_category = config.LOAD
	case config.SELECT:
		selected_category = config.SELECT
	case config.INSERT:
		selected_category = config.INSERT
	case config.SET:
		selected_category = config.SET
	case config.DELETE:
		selected_category = config.DELETE
	case config.WHERE:
		selected_category = config.WHERE
	case config.CREATE:
		selected_category = config.CREATE
	case config.AS:
		selected_category = config.AS
	case config.ID:
		selected_category = config.ID
	case config.SHOW:
		selected_category = config.SHOW
	case config.INTO:
		selected_category = config.INTO
	case config.VALUES:
		selected_category = config.VALUES
	case config.TRUNCATE:
		selected_category = config.TRUNCATE
	default:
		selected_token = config.IDENTIFIER
	}
	token = Get_Token(selected_token, value, selected_category)
	return
}

/*
  - @param file's data
  - @return Statement
  - @variable []Statement -> [Token]
  - @variable []Token -> [{
    -- type: <TOKEN TYPE>
    -- literal: <value>
  - }]
*/
func Lexer(fileContent *string) [][]Token {
	length := len(*fileContent)
	tokens := make([]Token, 0)
	statements := make([][]Token, 0)
	for index := 0; index < length; {
		character := (*fileContent)[index]
		switch true {
		case character == 10, character == 32, character == 13:
		case string(character) == `/`:
			handle_comments(fileContent, &index)
		case string(character) == `"`:
			value := get_string(fileContent, &index)
			tokens = append(tokens, value)
		case string(character) == `,`:
			value := Get_Token(config.COMMA, string(character), "")
			tokens = append(tokens, value)
		case string(character) == `*`:
			value := Get_Token(config.STAR, string(character), "")
			tokens = append(tokens, value)
		case string(character) == `(`:
			value := Get_Token(config.SEPERATOR, string(character), config.OPEN_PARENTHESIS)
			tokens = append(tokens, value)
		case string(character) == `)`:
			value := Get_Token(config.SEPERATOR, string(character), config.CLOSE_PARENTHESIS)
			tokens = append(tokens, value)
		case string(character) == `;`:
			value := Get_Token(config.SEPERATOR, string(character), config.SEMICOLON)
			tokens = append(tokens, value)
			statements = append(statements, tokens)
			tokens = make([]Token, 0)
		case unicode.IsLetter(rune(character)), character == 95:
			value := get_keyword_or_identifier(fileContent, &index)
			tokens = append(tokens, value)
		}
		index++
	}
	if len(tokens) > 0 {
		statements = append(statements, tokens)
	}
	return statements
}
