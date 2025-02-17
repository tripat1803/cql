package parser

import (
	"log"
	"tripat3k2/cql/components/lexer"
	"tripat3k2/cql/config"
)

func Handle_Delete_Command(statement []lexer.Token, start *int, beginning_node *Node) {
	statement_length := len(statement)

	if statement_length != 3 {
		log.Fatalf("Invalid delete statement.")
	}

	delete_node := new(DeleteNode)
	delete_node.Token = config.KEYWORD
	delete_node.Command = config.DELETE
	*start++

	if *start >= statement_length {
		log.Fatalf("Invalid delete statement.")
	}

	if statement[*start].Token != config.DATATYPE && statement[*start].Token != config.STRING {
		log.Fatalf("Invalid delete statement.")
	}
	delete_node.File = statement[*start].Value
	*start++

	if !(*start < statement_length && statement[*start].Token == config.SEPERATOR && statement[*start].Category == config.SEMICOLON) {
		log.Fatalf("Invalid delete statement. Missing Semicolon at the end.")
	}

	beginning_node.Next = delete_node
}
