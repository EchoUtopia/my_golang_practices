// Code generated from Expression.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Expression

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseExpressionListener is a complete listener for a parse tree produced by ExpressionParser.
type BaseExpressionListener struct{}

var _ ExpressionListener = &BaseExpressionListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseExpressionListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseExpressionListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseExpressionListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseExpressionListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseExpressionListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseExpressionListener) ExitStart(ctx *StartContext) {}

// EnterBracket is called when production Bracket is entered.
func (s *BaseExpressionListener) EnterBracket(ctx *BracketContext) {}

// ExitBracket is called when production Bracket is exited.
func (s *BaseExpressionListener) ExitBracket(ctx *BracketContext) {}

// EnterNot is called when production Not is entered.
func (s *BaseExpressionListener) EnterNot(ctx *NotContext) {}

// ExitNot is called when production Not is exited.
func (s *BaseExpressionListener) ExitNot(ctx *NotContext) {}

// EnterIdentifier is called when production Identifier is entered.
func (s *BaseExpressionListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production Identifier is exited.
func (s *BaseExpressionListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterOr is called when production Or is entered.
func (s *BaseExpressionListener) EnterOr(ctx *OrContext) {}

// ExitOr is called when production Or is exited.
func (s *BaseExpressionListener) ExitOr(ctx *OrContext) {}

// EnterAnd is called when production And is entered.
func (s *BaseExpressionListener) EnterAnd(ctx *AndContext) {}

// ExitAnd is called when production And is exited.
func (s *BaseExpressionListener) ExitAnd(ctx *AndContext) {}

// EnterTrue is called when production True is entered.
func (s *BaseExpressionListener) EnterTrue(ctx *TrueContext) {}

// ExitTrue is called when production True is exited.
func (s *BaseExpressionListener) ExitTrue(ctx *TrueContext) {}

// EnterMathExpressionCompare is called when production MathExpressionCompare is entered.
func (s *BaseExpressionListener) EnterMathExpressionCompare(ctx *MathExpressionCompareContext) {}

// ExitMathExpressionCompare is called when production MathExpressionCompare is exited.
func (s *BaseExpressionListener) ExitMathExpressionCompare(ctx *MathExpressionCompareContext) {}

// EnterFalse is called when production False is entered.
func (s *BaseExpressionListener) EnterFalse(ctx *FalseContext) {}

// ExitFalse is called when production False is exited.
func (s *BaseExpressionListener) ExitFalse(ctx *FalseContext) {}

// EnterEq is called when production Eq is entered.
func (s *BaseExpressionListener) EnterEq(ctx *EqContext) {}

// ExitEq is called when production Eq is exited.
func (s *BaseExpressionListener) ExitEq(ctx *EqContext) {}

// EnterMathBracket is called when production MathBracket is entered.
func (s *BaseExpressionListener) EnterMathBracket(ctx *MathBracketContext) {}

// ExitMathBracket is called when production MathBracket is exited.
func (s *BaseExpressionListener) ExitMathBracket(ctx *MathBracketContext) {}

// EnterNumber is called when production Number is entered.
func (s *BaseExpressionListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production Number is exited.
func (s *BaseExpressionListener) ExitNumber(ctx *NumberContext) {}

// EnterMulDiv is called when production MulDiv is entered.
func (s *BaseExpressionListener) EnterMulDiv(ctx *MulDivContext) {}

// ExitMulDiv is called when production MulDiv is exited.
func (s *BaseExpressionListener) ExitMulDiv(ctx *MulDivContext) {}

// EnterAddSub is called when production AddSub is entered.
func (s *BaseExpressionListener) EnterAddSub(ctx *AddSubContext) {}

// ExitAddSub is called when production AddSub is exited.
func (s *BaseExpressionListener) ExitAddSub(ctx *AddSubContext) {}

// EnterMathIdentifier is called when production MathIdentifier is entered.
func (s *BaseExpressionListener) EnterMathIdentifier(ctx *MathIdentifierContext) {}

// ExitMathIdentifier is called when production MathIdentifier is exited.
func (s *BaseExpressionListener) ExitMathIdentifier(ctx *MathIdentifierContext) {}

// EnterMinus is called when production Minus is entered.
func (s *BaseExpressionListener) EnterMinus(ctx *MinusContext) {}

// ExitMinus is called when production Minus is exited.
func (s *BaseExpressionListener) ExitMinus(ctx *MinusContext) {}

// EnterUnaryOperator is called when production unaryOperator is entered.
func (s *BaseExpressionListener) EnterUnaryOperator(ctx *UnaryOperatorContext) {}

// ExitUnaryOperator is called when production unaryOperator is exited.
func (s *BaseExpressionListener) ExitUnaryOperator(ctx *UnaryOperatorContext) {}

// EnterComparisonOperator is called when production comparisonOperator is entered.
func (s *BaseExpressionListener) EnterComparisonOperator(ctx *ComparisonOperatorContext) {}

// ExitComparisonOperator is called when production comparisonOperator is exited.
func (s *BaseExpressionListener) ExitComparisonOperator(ctx *ComparisonOperatorContext) {}

// EnterLogicalOperator is called when production logicalOperator is entered.
func (s *BaseExpressionListener) EnterLogicalOperator(ctx *LogicalOperatorContext) {}

// ExitLogicalOperator is called when production logicalOperator is exited.
func (s *BaseExpressionListener) ExitLogicalOperator(ctx *LogicalOperatorContext) {}

// EnterMathOperator is called when production mathOperator is entered.
func (s *BaseExpressionListener) EnterMathOperator(ctx *MathOperatorContext) {}

// ExitMathOperator is called when production mathOperator is exited.
func (s *BaseExpressionListener) ExitMathOperator(ctx *MathOperatorContext) {}
