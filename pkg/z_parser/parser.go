package parser

import (
	"github.com/pingcap/tidb/parser"
	"github.com/pingcap/tidb/parser/ast"
)

type Parser struct {
	p *parser.Parser
}

func NewParser() *Parser {
	return &Parser{p: parser.New()}
}

func (p *Parser) Parse(sql string) (stmt []ast.StmtNode, err error) {
	stmt, _, err = p.p.Parse(sql, "", "")
	return stmt, err
}
