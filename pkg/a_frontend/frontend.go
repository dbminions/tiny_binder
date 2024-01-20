package frontend

import (
	"github.com/polarsignals/frostdb/query/logicalplan"
	binder "tiny_binder/pkg/b_binder"
	security "tiny_binder/pkg/c_security"
	parser "tiny_binder/pkg/z_parser"
)

type Frontend struct {
	parser *parser.Parser
	acl    security.AclManager
}

func NewFrontend() *Frontend {
	return &Frontend{
		parser: parser.NewParser(),
		acl:    security.NewMockAclManager(),
	}
}

func (fe *Frontend) SqlToLogicalPlan(sql string) (*logicalplan.LogicalPlan, error) {
	asts, err := fe.parser.Parse(sql)
	if err != nil {
		return nil, err
	}

	lpBuilder := logicalplan.Builder{}
	v := binder.NewASTVisitor(lpBuilder)
	asts[0].Accept(v)
	if v.Err != nil {
		return nil, v.Err
	}

	lp, err := v.Builder.Build()
	if err != nil {
		return nil, err
	}

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
