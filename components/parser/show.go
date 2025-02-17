package parser

import (
	"log"
	"tripat3k2/cql/components/lexer"
	"tripat3k2/cql/config"
)

func Handle_Show_Command(statement []lexer.Token, start *int, beginning_node *Node) {
	statement_length := len(statement)

	if statement_length != 3 {
		log.Fatalf("Invalid show statement.")
	}

	show_node := new(ShowNode)
	show_node.Token = config.KEYWORD
	show_node.Command = config.SHOW
	*start++

	if *start >= statement_length {
		log.Fatalf("Invalid show statement.")
	}

	if statement[*start].Token != config.STAR {
		log.Fatalf("Invalid show statement.")
	}
	*start++

	if !(*start < statement_length && statement[*start].Token == config.SEPERATOR && statement[*start].Category == config.SEMICOLON) {
		log.Fatalf("Invalid show statement. Missing Semicolon at the end.")
	}

	beginning_node.Next = show_node
}
