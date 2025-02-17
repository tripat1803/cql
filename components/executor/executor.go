package executor

import (
	"fmt"
	"log"
	"tripat3k2/cql/components/executor/actions"
	"tripat3k2/cql/components/parser"
	"tripat3k2/cql/config"
)

func traverse_ast(ast_root parser.ASTNode) {
	switch ast_root.GetToken() {
	case config.KEYWORD:
		switch ast_root.GetCommand() {
		case config.SHOW:
			_, ok := ast_root.(*parser.ShowNode)
			if !ok {
				log.Fatalln("Invalid show statement")
			}
			fmt.Println(actions.ShowCSV())
		case config.CREATE:
			create_node := ast_root.(*parser.CreateNode)
			actions.CreateCSV(create_node.Columns, create_node.File)
		case config.INSERT:
			insert_node := ast_root.(*parser.InsertNode)
			actions.InsertCSV(insert_node.File, insert_node.Values, insert_node.Columns)
		case config.DELETE:
			delete_node := ast_root.(*parser.DeleteNode)
			actions.DeleteCSV(delete_node.File)
		case config.TRUNCATE:
			truncate_node := ast_root.(*parser.TruncateNode)
			actions.TruncateCSV(truncate_node.File)
		}
	}
	if ast_root.GetNext() == nil {
		return
	}
	traverse_ast(ast_root.GetNext())
}

func Executor(ast_statements []parser.ASTNode) {
	actions.ValidateOutputDir()
	actions.ValidateConfig()

	for _, ast := range ast_statements {
		traverse_ast(ast)
	}
}
