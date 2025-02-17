package main

import (
	"log"
	"os"
	"tripat3k2/cql/components/executor"
	"tripat3k2/cql/components/lexer"
	"tripat3k2/cql/components/parser"
)

func main() {
	args := os.Args
	if len(args) < 2 || len(args) > 2 {
		log.Fatalln("Invalid command.")
	}

	file, err := os.ReadFile(args[1])
	if err != nil {
		log.Fatalf("File not found %v\n", err)
	}

	fileContent := string(file)
	statements := lexer.Lexer(&fileContent)
	ast := parser.Parser(statements)
	executor.Executor(ast)
}
