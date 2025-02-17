package parser

import "tripat3k2/cql/components/lexer"

type ASTNode interface {
	GetToken() string
	GetCommand() string
	GetNext() ASTNode
}

// Default Node
type Node struct {
	Token lexer.Token
	Next  ASTNode
}

func (node *Node) GetToken() string {
	return node.Token.Token
}

func (node *Node) GetCommand() string {
	return node.Token.Category
}

func (node *Node) GetNext() ASTNode {
	return node.Next
}

// Create Node
type CreateNode struct {
	Token   string
	Command string
	Columns []string
	File    string
	Next    ASTNode
}

func (node *CreateNode) GetToken() string {
	return node.Token
}

func (node *CreateNode) GetCommand() string {
	return node.Command
}

func (node *CreateNode) GetNext() ASTNode {
	return node.Next
}

// Show Node
type ShowNode struct {
	Token   string
	Command string
	Next    ASTNode
}

func (node *ShowNode) GetToken() string {
	return node.Token
}

func (node *ShowNode) GetCommand() string {
	return node.Command
}

func (node *ShowNode) GetNext() ASTNode {
	return node.Next
}

// Insert Node
type InsertNode struct {
	Token   string
	Command string
	File    string
	Columns []string
	Values  [][]string
	Next    ASTNode
}

func (node *InsertNode) GetToken() string {
	return node.Token
}

func (node *InsertNode) GetCommand() string {
	return node.Command
}

func (node *InsertNode) GetNext() ASTNode {
	return node.Next
}

// Delete Node
type DeleteNode struct {
	Token   string
	Command string
	File    string
	Next    ASTNode
}

func (node *DeleteNode) GetToken() string {
	return node.Token
}

func (node *DeleteNode) GetCommand() string {
	return node.Command
}

func (node *DeleteNode) GetNext() ASTNode {
	return node.Next
}

// Truncate Node
type TruncateNode struct {
	Token   string
	Command string
	File    string
	Next    ASTNode
}

func (node *TruncateNode) GetToken() string {
	return node.Token
}

func (node *TruncateNode) GetCommand() string {
	return node.Command
}

func (node *TruncateNode) GetNext() ASTNode {
	return node.Next
}
