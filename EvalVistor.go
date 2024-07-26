package go_antlr4_cl

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	parser "go-antlr4-cl/gen"
	"strconv"
)

type EvalVistor struct {
	vars map[string]int
}

func NewEvalVisitor() *EvalVistor {
	return &EvalVistor{vars: make(map[string]int)}
}

func (v *EvalVistor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *EvalVistor) VisitChildren(node antlr.RuleNode) interface{} {
	switch node := node.(type) {
	case *parser.PrintExprContext:
		return v.VisitPrintExpr(node)
	case *parser.AssignContext:
		return v.VisitAssign(node)
	case *parser.BlankContext:
		return v.VisitBlank(node)
	case *parser.ParensContext:
		return v.VisitParens(node)
	case *parser.MulDivContext:
		return v.VisitMulDiv(node)
	case *parser.AddSubContext:
		return v.VisitAddSub(node)
	case *parser.IdContext:
		return v.VisitId(node)
	case *parser.IntContext:
		return v.VisitInt(node)
	case *parser.ClearContext:
		return v.VisitClear(node)
	default:
		return nil
	}
}

func (v *EvalVistor) VisitTerminal(node antlr.TerminalNode) interface{} {
	return nil
}

func (v *EvalVistor) VisitErrorNode(node antlr.ErrorNode) interface{} {
	return nil
}

func (v *EvalVistor) VisitProg(ctx *parser.ProgContext) interface{} {
	for _, statCtx := range ctx.AllStat() {
		v.Visit(statCtx)
	}
	return nil
}

func (v *EvalVistor) VisitClear(ctx *parser.ClearContext) interface{} {
	text := ctx.CLEAR().GetText()
	if "clear" == text {
		v.vars = make(map[string]int)
	}
	fmt.Println(v.vars)
	return nil
}

func (v *EvalVistor) VisitPrintExpr(ctx *parser.PrintExprContext) interface{} {
	result := v.Visit(ctx.Expr())
	fmt.Println(result)
	return 0
}

func (v *EvalVistor) VisitAssign(ctx *parser.AssignContext) interface{} {
	id := ctx.ID().GetText()
	value := v.Visit(ctx.Expr()).(int)
	v.vars[id] = value
	return value
}

func (v *EvalVistor) VisitBlank(ctx *parser.BlankContext) interface{} {
	return nil
}

func (v *EvalVistor) VisitParens(ctx *parser.ParensContext) interface{} {
	return v.Visit(ctx.Expr())
}

func (v *EvalVistor) VisitMulDiv(ctx *parser.MulDivContext) interface{} {
	left := v.Visit(ctx.Expr(0)).(int)
	right := v.Visit(ctx.Expr(1)).(int)
	op := ctx.GetOp().GetTokenType()

	switch op {
	case parser.LabeledExprLexerMUL:
		fmt.Println(left, "*", right)
		return left * right
	case parser.LabeledExprLexerDIV:
		fmt.Println(left, "/", right)
		return left / right
	default:
		panic(fmt.Sprintf("Unexpected operator: %d", op))
	}
}

func (v *EvalVistor) VisitAddSub(ctx *parser.AddSubContext) interface{} {
	left := v.Visit(ctx.Expr(0)).(int)
	right := v.Visit(ctx.Expr(1)).(int)
	op := ctx.GetOp().GetTokenType()

	switch op {
	case parser.LabeledExprLexerADD:
		fmt.Println(left, "+", right)
		return left + right
	case parser.LabeledExprLexerSUB:
		fmt.Println(left, "-", right)
		return left - right
	default:
		panic(fmt.Sprintf("Unexpected operator: %d", op))
	}
}

func (v *EvalVistor) VisitId(ctx *parser.IdContext) interface{} {
	id := ctx.ID().GetText()
	if val, ok := v.vars[id]; ok {
		return val
	}
	return 0
}

func (v *EvalVistor) VisitInt(ctx *parser.IntContext) interface{} {
	val := ctx.INT().GetText()
	d, _ := strconv.Atoi(val)
	fmt.Println("VisitInt,", d)
	return d
}
