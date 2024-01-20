package parser

import (
	"github.com/blastrain/vitess-sqlparser/tidbparser/ast"
	"github.com/blastrain/vitess-sqlparser/tidbparser/parser"
)

type Parser struct {
	src *parser.Parser
}

func NewParser() *Parser {
	return &Parser{src: parser.New()}
}

func (p *Parser) Parse(sql string) (stmt []ast.StmtNode, err error) {
	stmt, err = p.src.Parse(sql, "", "")
	return stmt, err
}
