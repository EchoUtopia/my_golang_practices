// Code generated from Expression.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Expression

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ExpressionListener is a complete listener for a parse tree produced by ExpressionParser.
type ExpressionListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterBracket is called when entering the Bracket production.
	EnterBracket(c *BracketContext)

	// EnterNot is called when entering the Not production.
	EnterNot(c *NotContext)

	// EnterIdentifier is called when entering the Identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterOr is called when entering the Or production.
	EnterOr(c *OrContext)

	// EnterAnd is called when entering the And production.
	EnterAnd(c *AndContext)

	// EnterTrue is called when entering the True production.
	EnterTrue(c *TrueContext)

	// EnterMathExpressionCompare is called when entering the MathExpressionCompare production.
	EnterMathExpressionCompare(c *MathExpressionCompareContext)

	// EnterFalse is called when entering the False production.
	EnterFalse(c *FalseContext)

	// EnterEq is called when entering the Eq production.
	EnterEq(c *EqContext)

	// EnterMathBracket is called when entering the MathBracket production.
	EnterMathBracket(c *MathBracketContext)

	// EnterNumber is called when entering the Number production.
	EnterNumber(c *NumberContext)

	// EnterMulDiv is called when entering the MulDiv production.
	EnterMulDiv(c *MulDivContext)

	// EnterAddSub is called when entering the AddSub production.
	EnterAddSub(c *AddSubContext)

	// EnterMathIdentifier is called when entering the MathIdentifier production.
	EnterMathIdentifier(c *MathIdentifierContext)

	// EnterMinus is called when entering the Minus production.
	EnterMinus(c *MinusContext)

	// EnterUnaryOperator is called when entering the unaryOperator production.
	EnterUnaryOperator(c *UnaryOperatorContext)

	// EnterComparisonOperator is called when entering the comparisonOperator production.
	EnterComparisonOperator(c *ComparisonOperatorContext)

	// EnterLogicalOperator is called when entering the logicalOperator production.
	EnterLogicalOperator(c *LogicalOperatorContext)

	// EnterMathOperator is called when entering the mathOperator production.
	EnterMathOperator(c *MathOperatorContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitBracket is called when exiting the Bracket production.
	ExitBracket(c *BracketContext)

	// ExitNot is called when exiting the Not production.
	ExitNot(c *NotContext)

	// ExitIdentifier is called when exiting the Identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitOr is called when exiting the Or production.
	ExitOr(c *OrContext)

	// ExitAnd is called when exiting the And production.
	ExitAnd(c *AndContext)

	// ExitTrue is called when exiting the True production.
	ExitTrue(c *TrueContext)

	// ExitMathExpressionCompare is called when exiting the MathExpressionCompare production.
	ExitMathExpressionCompare(c *MathExpressionCompareContext)

	// ExitFalse is called when exiting the False production.
	ExitFalse(c *FalseContext)

	// ExitEq is called when exiting the Eq production.
	ExitEq(c *EqContext)

	// ExitMathBracket is called when exiting the MathBracket production.
	ExitMathBracket(c *MathBracketContext)

	// ExitNumber is called when exiting the Number production.
	ExitNumber(c *NumberContext)

	// ExitMulDiv is called when exiting the MulDiv production.
	ExitMulDiv(c *MulDivContext)

	// ExitAddSub is called when exiting the AddSub production.
	ExitAddSub(c *AddSubContext)

	// ExitMathIdentifier is called when exiting the MathIdentifier production.
	ExitMathIdentifier(c *MathIdentifierContext)

	// ExitMinus is called when exiting the Minus production.
	ExitMinus(c *MinusContext)

	// ExitUnaryOperator is called when exiting the unaryOperator production.
	ExitUnaryOperator(c *UnaryOperatorContext)

	// ExitComparisonOperator is called when exiting the comparisonOperator production.
	ExitComparisonOperator(c *ComparisonOperatorContext)

	// ExitLogicalOperator is called when exiting the logicalOperator production.
	ExitLogicalOperator(c *LogicalOperatorContext)

	// ExitMathOperator is called when exiting the mathOperator production.
	ExitMathOperator(c *MathOperatorContext)
}
