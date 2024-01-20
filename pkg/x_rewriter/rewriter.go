package rewriter

import (
	"github.com/blastrain/vitess-sqlparser/tidbparser/ast"
)

type Rewriter struct {
}

func NewRewriter() *Rewriter {
	return &Rewriter{}
}

func (r *Rewriter) Rewrite(stmt ast.StmtNode) (ast.StmtNode, error) {
	return stmt, nil
}
