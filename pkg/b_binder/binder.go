package binder

import (
	"fmt"
	"github.com/pingcap/tidb/parser/ast"
	"github.com/pingcap/tidb/parser/opcode"
	"github.com/polarsignals/frostdb/query"
	"github.com/polarsignals/frostdb/query/logicalplan"
	"strings"
)

type AstVisitor struct {
	Builder   query.Builder
	Err       error
	exprStack []logicalplan.Expr
}

var _ ast.Visitor = &AstVisitor{}

func NewASTVisitor(builder query.Builder) *AstVisitor {
	return &AstVisitor{
		Builder: builder,
	}
}

func (v *AstVisitor) Enter(n ast.Node) (nRes ast.Node, skipChildren bool) {
	switch expr := n.(type) {
	case *ast.SelectStmt:
		// The SelectStmt is handled in during pre-visit given that it has many
		// clauses we need to handle independently (e.g. a group by with a
		// filter).
		if expr.Where != nil {
			expr.Where.Accept(v)
			lastExpr, newExprs := pop(v.exprStack)
			v.exprStack = newExprs
			v.Builder = v.Builder.Filter(lastExpr)
		}
		expr.Fields.Accept(v)
		switch {
		case expr.GroupBy != nil:
			expr.GroupBy.Accept(v)
			var agg []logicalplan.Expr
			var groups []logicalplan.Expr

			for _, expr := range v.exprStack {
				switch expr.(type) {
				case *logicalplan.AliasExpr, *logicalplan.AggregationFunction:
					agg = append(agg, expr)
				default:
					groups = append(groups, expr)
				}
			}
			v.Builder = v.Builder.Aggregate(agg, groups)
		case expr.Distinct:
			v.Builder = v.Builder.Distinct(v.exprStack...)
		default:
			v.Builder = v.Builder.Project(v.exprStack...)
		}
		return n, true
	}
	return n, false
}

func (v *AstVisitor) Leave(n ast.Node) (nRes ast.Node, ok bool) {
	if err := v.leaveImpl(n); err != nil {
		v.Err = err
		return n, false
	}
	return n, true
}

func (v *AstVisitor) leaveImpl(n ast.Node) error {
	switch expr := n.(type) {
	case *ast.SelectStmt:
		// Handled in Enter.
		return nil
	case *ast.AggregateFuncExpr:
		// At this point, the child node is the column name, so it has just been
		// added to exprs.
		lastExpr := len(v.exprStack) - 1
		switch strings.ToLower(expr.F) {
		case "sum":
			v.exprStack[lastExpr] = logicalplan.Sum(v.exprStack[lastExpr])
		default:
			return fmt.Errorf("unhandled aggregate function %s", expr.F)
		}
	case *ast.BinaryOperationExpr:
		// Note that we're resolving exprs as a stack, so the last two
		// expressions are the leaf expressions.
		rightExpr, newExprs := pop(v.exprStack)
		leftExpr, newExprs := pop(newExprs)
		v.exprStack = newExprs

		var frostDBOp logicalplan.Op
		switch expr.Op {
		case opcode.LT:
			frostDBOp = logicalplan.OpLt
		}
		v.exprStack = append(v.exprStack, &logicalplan.BinaryExpr{
			Left:  logicalplan.Col(leftExpr.Name()),
			Op:    frostDBOp,
			Right: rightExpr,
		})
	case *ast.ColumnName:
		colName := columnNameToString(expr)
		var col logicalplan.Expr
		//TODO catalog
		col = logicalplan.Col(colName)
		v.exprStack = append(v.exprStack, col)
	case *ast.SelectField:
		if as := expr.AsName.String(); as != "" {
			lastExpr := len(v.exprStack) - 1
			v.exprStack[lastExpr] = v.exprStack[lastExpr].(*logicalplan.AggregationFunction).Alias(as) // TODO should probably just be an alias expr and not from an aggregate function
		}

	case *ast.FieldList, *ast.ColumnNameExpr, *ast.GroupByClause, *ast.ByItem, *ast.RowExpr,
		*ast.ParenthesesExpr:
		// Deliberate pass-through nodes.
	default:
		return fmt.Errorf("unhandled ast node %T", expr)
	}
	return nil
}

func columnNameToString(c *ast.ColumnName) string {
	colName := ""
	if c.Table.String() != "" {
		colName = c.Table.String() + "."
	}
	colName += c.Name.String()
	return colName
}

func pop[T any](s []T) (T, []T) {
	lastIdx := len(s) - 1
	return s[lastIdx], s[:lastIdx]
}
