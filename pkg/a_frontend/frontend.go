package frontend

import (
	"github.com/polarsignals/frostdb/query/logicalplan"
	security "tiny_binder/pkg/b_security"
	binder "tiny_binder/pkg/c_binder"
	rewriter "tiny_binder/pkg/x_rewriter"
	parser "tiny_binder/pkg/z_parser"
)

type Frontend struct {
	parser   *parser.Parser
	acl      security.AclManager
	rewriter *rewriter.Rewriter
}

func NewFrontend() *Frontend {
	return &Frontend{
		parser:   parser.NewParser(),
		acl:      security.NewMockAclManager(),
		rewriter: rewriter.NewRewriter(),
	}
}

func (fe *Frontend) SqlToLogicalPlan(sql string) (*logicalplan.LogicalPlan, error) {
	// 1. Parse SQL to AST
	ast, err := fe.parser.Parse(sql)
	if err != nil {
		return nil, err
	}

	// 2. Rewrite AST
	ast, err = fe.rewriter.Rewrite(ast)
	if err != nil {
		return nil, err
	}

	// 3. Convert AST to LogicalPlan
	lpBuilder := logicalplan.Builder{}
	v := binder.NewASTVisitor(lpBuilder)
	ast.Accept(v)
	if v.Err != nil {
		return nil, v.Err
	}
	lp, err := v.Builder.Build()
	if err != nil {
		return nil, err
	}

	// 4. Check Access Control
	err = fe.acl.CheckQueryIntegrity(lp, fe.GetCurrentUserDef())
	if err != nil {
		return nil, err
	}

	return lp, err
}

func (fe *Frontend) GetCurrentUserDef() *security.UserDef {
	return &security.UserDef{
		UserName: "test",
		Roles: []security.RoleDef{
			{
				Id:       1,
				RoleName: "test",
				Privileges: []security.Privilege{
					security.Select,
				},
			},
		},
	}
}
