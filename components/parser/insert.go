package parser

import (
	"log"
	"tripat3k2/cql/components/lexer"
	"tripat3k2/cql/config"
)

func Handle_Insert_Command(statement []lexer.Token, start *int, beginning_node *Node) {
	statement_length := len(statement)
	insert_node := new(InsertNode)
	insert_node.Token = config.KEYWORD
	insert_node.Command = config.INSERT
	*start++

	if *start >= statement_length {
		log.Fatalf("Invalid insert statement.")
	}

	if !(statement[*start].Token == config.KEYWORD && statement[*start].Category == config.INTO) {
		log.Fatalln("Invalid insert query")
	}
	*start++

	if *start >= statement_length {
		log.Fatalf("Invalid insert statement.")
	}

	if !(statement[*start].Token == config.DATATYPE && statement[*start].Category == config.STRING) {
		log.Fatalln("Invalid insert query. Missing file name")
	}
	insert_node.File = statement[*start].Value
	*start++

	if *start >= statement_length {
		log.Fatalf("Invalid insert statement.")
	}

	columns := make([]string, 0)
	if statement[*start].Token == config.SEPERATOR && statement[*start].Category == config.OPEN_PARENTHESIS {
		*start++
		if *start >= statement_length {
			log.Fatalf("Invalid insert statement.")
		}

		first_identifier := -1
		for *start < statement_length && statement[*start].Category != config.CLOSE_PARENTHESIS {
			if statement[*start].Token == config.IDENTIFIER {
				if first_identifier == -1 {
					first_identifier = *start
				}
				columns = append(columns, statement[*start].Value)
			}
			if (*start-first_identifier)%2 != 0 {
				if statement[*start].Token != config.COMMA {
					log.Fatalln("Invalid insert statement")
				}
			} else {
				if statement[*start].Token != config.IDENTIFIER {
					log.Fatalln("Invalid insert statement")
				}
			}
			*start++
		}
		*start++

		if *start >= statement_length {
			log.Fatalf("Invalid insert statement.")
		}
	}

	if len(columns) == 0 {
		insert_node.Columns = nil
	} else {
		insert_node.Columns = columns
	}

	if statement[*start].Token == config.KEYWORD && statement[*start].Category == config.VALUES {
		*start++
		if *start >= statement_length {
			log.Fatalf("Invalid insert statement.")
		}

		values := make([][]string, 0)
		if statement[*start].Token == config.SEPERATOR && statement[*start].Category == config.OPEN_PARENTHESIS {
			*start++
			if *start >= statement_length {
				log.Fatalf("Invalid insert statement.")
			}
			row := make([]string, 0)
			first_datatype := -1
			for *start < statement_length && statement[*start].Category != config.SEMICOLON {
				if statement[*start].Token == config.SEPERATOR && statement[*start].Category == config.CLOSE_PARENTHESIS {
					values = append(values, row)
					row = make([]string, 0)
					first_datatype = -1
					*start++
					if statement[*start].Category == config.SEMICOLON {
						break
					}
					if *start < statement_length {
						if statement[*start].Token == config.COMMA {
							*start++
							if *start < statement_length {
								if statement[*start].Category == config.OPEN_PARENTHESIS {
									*start++
									if *start >= statement_length {
										log.Fatalf("Invalid insert statement.")
									}
								} else {
									log.Fatalf("Invalid insert statement.")
								}
							} else {
								log.Fatalf("Invalid insert statement.")
							}
						}
					}
					if statement[*start].Category == config.CLOSE_PARENTHESIS {
						continue
					}
				}

				if statement[*start].Token == config.DATATYPE {
					if first_datatype == -1 {
						first_datatype = *start
					}
					row = append(row, statement[*start].Value)
				}
				if (*start-first_datatype)%2 != 0 {
					if statement[*start].Token != config.COMMA {
						log.Fatalln("Invalid insert statement")
					}
				} else {
					if statement[*start].Token != config.DATATYPE {
						log.Fatalln("Invalid insert statement")
					}
				}
				*start++
			}

			if *start >= statement_length {
				log.Fatalf("Invalid insert statement.")
			}
		}

		if !(statement[*start].Token == config.SEPERATOR && statement[*start].Category == config.SEMICOLON) {
			log.Fatalf("Invalid insert statement. Missing Semicolon at the end.")
		}

		insert_node.Values = values
	} else {
		log.Fatalln("Invalid insert statement")
	}

	beginning_node.Next = insert_node
}
