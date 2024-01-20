package binder

import "context"

type Binder interface {
	BindExpr(tree.Expr, int32, bool) (*plan.Expr, error)
	BindColRef(*tree.UnresolvedName, int32, bool) (*plan.Expr, error)

	BindFunc(string, *tree.FuncExpr, int32, bool) (*plan.Expr, error)
	BindAggFunc(string, *tree.FuncExpr, int32, bool) (*plan.Expr, error)

	GetContext() context.Context
}
