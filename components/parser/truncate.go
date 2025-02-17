package parser

import (
	"log"
	"tripat3k2/cql/components/lexer"
	"tripat3k2/cql/config"
)

func Handle_Truncate_Command(statement []lexer.Token, start *int, beginning_node *Node) {
	statement_length := len(statement)

	if statement_length != 3 {
		log.Fatalf("Invalid truncate statement.")
	}

	truncate_node := new(TruncateNode)
	truncate_node.Token = config.KEYWORD
	truncate_node.Command = config.TRUNCATE
	*start++

	if *start >= statement_length {
		log.Fatalf("Invalid delete statement.")
	}

	if statement[*start].Token != config.DATATYPE && statement[*start].Token != config.STRING {
		log.Fatalf("Invalid truncate statement.")
	}
	truncate_node.File = statement[*start].Value
	*start++

	if !(*start < statement_length && statement[*start].Token == config.SEPERATOR && statement[*start].Category == config.SEMICOLON) {
		log.Fatalf("Invalid truncate statement. Missing Semicolon at the end.")
	}

	beginning_node.Next = truncate_node
}
