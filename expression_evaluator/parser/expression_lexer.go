// Code generated from Expression.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 24, 122,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6,
	3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10,
	3, 10, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3,
	15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 6, 19, 90,
	10, 19, 13, 19, 14, 19, 91, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3,
	21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 22, 6, 22, 106, 10, 22, 13, 22, 14,
	22, 107, 3, 22, 3, 22, 3, 23, 6, 23, 113, 10, 23, 13, 23, 14, 23, 114,
	3, 23, 7, 23, 118, 10, 23, 12, 23, 14, 23, 121, 11, 23, 2, 2, 24, 3, 3,
	5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13,
	25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22,
	43, 23, 45, 24, 3, 2, 6, 3, 2, 50, 59, 5, 2, 11, 12, 15, 15, 34, 34, 4,
	2, 67, 92, 99, 124, 5, 2, 50, 59, 67, 92, 99, 124, 2, 125, 2, 3, 3, 2,
	2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2,
	2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3,
	2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27,
	3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2,
	35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2,
	2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 3, 47, 3, 2, 2, 2, 5, 49, 3, 2, 2,
	2, 7, 52, 3, 2, 2, 2, 9, 55, 3, 2, 2, 2, 11, 57, 3, 2, 2, 2, 13, 60, 3,
	2, 2, 2, 15, 62, 3, 2, 2, 2, 17, 65, 3, 2, 2, 2, 19, 68, 3, 2, 2, 2, 21,
	71, 3, 2, 2, 2, 23, 74, 3, 2, 2, 2, 25, 76, 3, 2, 2, 2, 27, 78, 3, 2, 2,
	2, 29, 80, 3, 2, 2, 2, 31, 82, 3, 2, 2, 2, 33, 84, 3, 2, 2, 2, 35, 86,
	3, 2, 2, 2, 37, 89, 3, 2, 2, 2, 39, 93, 3, 2, 2, 2, 41, 98, 3, 2, 2, 2,
	43, 105, 3, 2, 2, 2, 45, 112, 3, 2, 2, 2, 47, 48, 7, 39, 2, 2, 48, 4, 3,
	2, 2, 2, 49, 50, 7, 126, 2, 2, 50, 51, 7, 126, 2, 2, 51, 6, 3, 2, 2, 2,
	52, 53, 7, 40, 2, 2, 53, 54, 7, 40, 2, 2, 54, 8, 3, 2, 2, 2, 55, 56, 7,
	64, 2, 2, 56, 10, 3, 2, 2, 2, 57, 58, 7, 64, 2, 2, 58, 59, 7, 63, 2, 2,
	59, 12, 3, 2, 2, 2, 60, 61, 7, 62, 2, 2, 61, 14, 3, 2, 2, 2, 62, 63, 7,
	62, 2, 2, 63, 64, 7, 63, 2, 2, 64, 16, 3, 2, 2, 2, 65, 66, 7, 63, 2, 2,
	66, 67, 7, 63, 2, 2, 67, 18, 3, 2, 2, 2, 68, 69, 7, 35, 2, 2, 69, 70, 7,
	63, 2, 2, 70, 20, 3, 2, 2, 2, 71, 72, 7, 107, 2, 2, 72, 73, 7, 112, 2,
	2, 73, 22, 3, 2, 2, 2, 74, 75, 7, 35, 2, 2, 75, 24, 3, 2, 2, 2, 76, 77,
	7, 44, 2, 2, 77, 26, 3, 2, 2, 2, 78, 79, 7, 49, 2, 2, 79, 28, 3, 2, 2,
	2, 80, 81, 7, 45, 2, 2, 81, 30, 3, 2, 2, 2, 82, 83, 7, 47, 2, 2, 83, 32,
	3, 2, 2, 2, 84, 85, 7, 42, 2, 2, 85, 34, 3, 2, 2, 2, 86, 87, 7, 43, 2,
	2, 87, 36, 3, 2, 2, 2, 88, 90, 9, 2, 2, 2, 89, 88, 3, 2, 2, 2, 90, 91,
	3, 2, 2, 2, 91, 89, 3, 2, 2, 2, 91, 92, 3, 2, 2, 2, 92, 38, 3, 2, 2, 2,
	93, 94, 7, 118, 2, 2, 94, 95, 7, 116, 2, 2, 95, 96, 7, 119, 2, 2, 96, 97,
	7, 103, 2, 2, 97, 40, 3, 2, 2, 2, 98, 99, 7, 104, 2, 2, 99, 100, 7, 99,
	2, 2, 100, 101, 7, 110, 2, 2, 101, 102, 7, 117, 2, 2, 102, 103, 7, 103,
	2, 2, 103, 42, 3, 2, 2, 2, 104, 106, 9, 3, 2, 2, 105, 104, 3, 2, 2, 2,
	106, 107, 3, 2, 2, 2, 107, 105, 3, 2, 2, 2, 107, 108, 3, 2, 2, 2, 108,
	109, 3, 2, 2, 2, 109, 110, 8, 22, 2, 2, 110, 44, 3, 2, 2, 2, 111, 113,
	9, 4, 2, 2, 112, 111, 3, 2, 2, 2, 113, 114, 3, 2, 2, 2, 114, 112, 3, 2,
	2, 2, 114, 115, 3, 2, 2, 2, 115, 119, 3, 2, 2, 2, 116, 118, 9, 5, 2, 2,
	117, 116, 3, 2, 2, 2, 118, 121, 3, 2, 2, 2, 119, 117, 3, 2, 2, 2, 119,
	120, 3, 2, 2, 2, 120, 46, 3, 2, 2, 2, 121, 119, 3, 2, 2, 2, 7, 2, 91, 107,
	114, 119, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'%'", "'||'", "'&&'", "'>'", "'>='", "'<'", "'<='", "'=='", "'!='",
	"'in'", "'!'", "'*'", "'/'", "'+'", "'-'", "'('", "')'", "", "'true'",
	"'false'",
}

var lexerSymbolicNames = []string{
	"", "", "OR", "AND", "GT", "GTE", "LT", "LTE", "EQ", "NEQ", "IN", "NOT",
	"MUL", "DIV", "ADD", "SUB", "LPAREN", "RPAREN", "NUMBER", "TRUE", "FALSE",
	"WS", "IDENTIFIER",
}

var lexerRuleNames = []string{
	"T__0", "OR", "AND", "GT", "GTE", "LT", "LTE", "EQ", "NEQ", "IN", "NOT",
	"MUL", "DIV", "ADD", "SUB", "LPAREN", "RPAREN", "NUMBER", "TRUE", "FALSE",
	"WS", "IDENTIFIER",
}

type ExpressionLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewExpressionLexer(input antlr.CharStream) *ExpressionLexer {

	l := new(ExpressionLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Expression.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ExpressionLexer tokens.
const (
	ExpressionLexerT__0       = 1
	ExpressionLexerOR         = 2
	ExpressionLexerAND        = 3
	ExpressionLexerGT         = 4
	ExpressionLexerGTE        = 5
	ExpressionLexerLT         = 6
	ExpressionLexerLTE        = 7
	ExpressionLexerEQ         = 8
	ExpressionLexerNEQ        = 9
	ExpressionLexerIN         = 10
	ExpressionLexerNOT        = 11
	ExpressionLexerMUL        = 12
	ExpressionLexerDIV        = 13
	ExpressionLexerADD        = 14
	ExpressionLexerSUB        = 15
	ExpressionLexerLPAREN     = 16
	ExpressionLexerRPAREN     = 17
	ExpressionLexerNUMBER     = 18
	ExpressionLexerTRUE       = 19
	ExpressionLexerFALSE      = 20
	ExpressionLexerWS         = 21
	ExpressionLexerIDENTIFIER = 22
)
