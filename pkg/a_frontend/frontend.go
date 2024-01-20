package frontend

import (
	"fmt"
	"github.com/polarsignals/frostdb/query"
	binder "tiny_binder/pkg/b_binder"
	parser "tiny_binder/pkg/z_parser"
)

type Frontend struct {
	parser *parser.Parser
}

func NewFrontend() *Frontend {
	return &Frontend{
		parser: parser.NewParser(),
	}
}

func (p *Frontend) SqlToLogicalPlan(builder query.Builder, sql string) (query.Builder, error) {
	asts, err := p.parser.Parse(sql)
	if err != nil {
		return nil, err
	}

	if len(asts) != 1 {
		return nil, fmt.Errorf("cannot handle multiple asts, found %d", len(asts))
	}

	v := binder.NewASTVisitor(builder)
	asts[0].Accept(v)
	if v.Err != nil {
		return nil, v.Err
	}

	return v.Builder, nil
}
