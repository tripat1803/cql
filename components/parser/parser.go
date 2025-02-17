package parser

import (
	"tripat3k2/cql/components/lexer"
	"tripat3k2/cql/config"
)

func handle_keyword(statement []lexer.Token, start *int, beginning_node *Node) {
	token := statement[*start]
	switch token.Category {
	case config.CREATE:
		Handle_Create_Command(statement, start, beginning_node)
	case config.SHOW:
		Handle_Show_Command(statement, start, beginning_node)
	case config.INSERT:
		Handle_Insert_Command(statement, start, beginning_node)
	case config.DELETE:
		Handle_Delete_Command(statement, start, beginning_node)
	case config.TRUNCATE:
		Handle_Truncate_Command(statement, start, beginning_node)
	}
}

func Parser(statements [][]lexer.Token) []ASTNode {
	ast_statements := make([]ASTNode, 0)
	for _, statement := range statements {
		token_length := len(statement)
		beginning_node := Node{
			Token: lexer.Get_Token(config.BEGINNING, "beginning", ""),
		}
		for start := 0; start < token_length; {
			token := statement[start]
			switch token.Token {
			case config.KEYWORD:
				handle_keyword(statement, &start, &beginning_node)
			}
			start++
		}
		ast_statements = append(ast_statements, &beginning_node)
	}
	return ast_statements
}
