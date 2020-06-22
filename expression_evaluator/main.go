package main

import (
	"demo.com/examples/antlr/parser"
	"flag"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"strconv"
)

type expressionListener struct {
	*parser.BaseExpressionListener

	stack []interface{}
	parameters map[string]interface{}
}

func (l *expressionListener) push(i interface{}) {
	l.stack = append(l.stack, i)
}

func (l *expressionListener) pop() interface{} {
	if len(l.stack) < 1 {
		panic("stack is empty unable to pop")
	}

	// Get the last value from the stack.
	result := l.stack[len(l.stack)-1]

	// Remove the last element from the stack.
	l.stack = l.stack[:len(l.stack)-1]

	return result
}

func (l *expressionListener) ExitMulDiv(c *parser.MulDivContext) {
	right, left := l.pop().(int), l.pop().(int)

	switch c.GetOp().GetTokenType() {
	case parser.ExpressionParserMUL:
		l.push(left * right)
	case parser.ExpressionParserDIV:
		l.push(left / right)
	default:
		panic(fmt.Sprintf("unexpected op: %s", c.GetOp().GetText()))
	}
}

func (l *expressionListener) ExitAddSub(c *parser.AddSubContext) {
	right, left := l.pop().(int), l.pop().(int)

	switch c.GetOp().GetTokenType() {
	case parser.ExpressionParserADD:
		l.push(left + right)
	case parser.ExpressionParserSUB:
		l.push(left - right)
	default:
		panic(fmt.Sprintf("unexpected op: %s", c.GetOp().GetText()))
	}
}

func (l *expressionListener) ExitNumber(c *parser.NumberContext) {
	i, err := strconv.Atoi(c.GetText())
	if err != nil {
		panic(err.Error())
	}
	l.push(i)
}


func (l *expressionListener) ExitMinus(c *parser.MinusContext) {
	n := l.pop().(int)
	l.push(-n)
}

func (l *expressionListener) ExitMathExpressionCompare(c *parser.MathExpressionCompareContext) {
	right, left := l.pop().(int), l.pop().(int)
	switch c.GetOp().GetTokenType(){
	case parser.ExpressionLexerGT:
		l.push(left > right)
	case parser.ExpressionLexerEQ:
		l.push(left == right)
	case parser.ExpressionLexerGTE:
		l.push(left >= right)
	case parser.ExpressionLexerLT:
		l.push(left < right)
	case parser.ExpressionLexerLTE:
		l.push(left <= right)
	case parser.ExpressionLexerNEQ:
		l.push(left != right)
	}
}


func (l *expressionListener) ExitAnd(c *parser.AndContext) {
	right, left := l.pop().(bool), l.pop().(bool)
	l.push(left && right)
}


func (l *expressionListener) ExitOr(c *parser.OrContext) {
	right, left := l.pop().(bool), l.pop().(bool)
	l.push(left || right)
}


func (l *expressionListener) ExitFalse(c *parser.FalseContext) {
	l.pop()
	l.push(false)
}


func (l *expressionListener) ExitTrue(c *parser.TrueContext) {
	l.pop()
	l.push(true)
}


func (l *expressionListener) ExitIdentifier(c *parser.IdentifierContext) {
	key := c.GetText()
	l.push(l.parameters[key].(bool))
}


func (l *expressionListener) ExitMathIdentifier(c *parser.MathIdentifierContext) {
	key := c.GetText()
	l.push(l.parameters[key].(int))
}

func (l *expressionListener) ExitEq(c *parser.EqContext) {
	right, left := l.pop().(bool), l.pop().(bool)
	switch c.GetOp().GetTokenType() {
	case parser.ExpressionLexerNEQ:
		l.push(right != left)
	case parser.ExpressionParserEQ:
		l.push(right == left)
	}
}


func (l *expressionListener) ExitNot(c *parser.NotContext) {
	l.push(!l.pop().(bool))
}


var number = flag.Int(`n`, 3, ``)
var boolean = flag.Bool(`b`, false, ``)

func main() {
	flag.Parse()
	parse()
}

func parse(){
	is := antlr.NewInputStream("3*-a*2 >= -10 && !b2")

	// Create the Lexer
	lexer := parser.NewExpressionLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewExpressionParser(stream)

	// Finally parse the expression
	var listener = expressionListener{
		parameters: map[string]interface{}{
			`a`: *number,
			`b2`: *boolean,
		},
	}
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Start())
	//fmt.Println(listener.pop())
}

func parseWithCache(ctx parser.IStartContext){
	var listener expressionListener
	antlr.ParseTreeWalkerDefault.Walk(&listener, ctx)
}
