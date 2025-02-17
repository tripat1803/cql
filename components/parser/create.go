package parser

import (
	"log"
	"tripat3k2/cql/components/lexer"
	"tripat3k2/cql/config"
)

func Handle_Create_Command(statement []lexer.Token, start *int, beginning_node *Node) {
	statement_length := len(statement)
	create_node := new(CreateNode)
	create_node.Token = config.KEYWORD
	create_node.Command = config.CREATE
	*start++
	columns := make([]string, 0)
	for *start < statement_length && statement[*start].Category != config.AS {
		if statement[*start].Token == config.IDENTIFIER {
			columns = append(columns, statement[*start].Value)
		}
		*start++
	}

	if *start >= statement_length {
		log.Fatalf("Invalid create statement.")
	}

	if statement[*start].Category != config.AS {
		log.Fatalln("Expected AS here in create statement.")
	}
	*start++

	if *start >= statement_length {
		log.Fatalf("Invalid create statement.")
	}

	if statement[*start].Token != config.DATATYPE && statement[*start].Category != config.STRING {
		log.Fatalln("Expected file name here in create statement.")
	}
	file := statement[*start].Value

	create_node.Columns = columns
	create_node.File = file
	beginning_node.Next = create_node
}
