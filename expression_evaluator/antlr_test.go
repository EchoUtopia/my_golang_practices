package main

import (
	"demo.com/examples/antlr/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"testing"
)

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parse()
	}
}

func BenchmarkWithCache(b *testing.B) {
	is := antlr.NewInputStream("(3+2)*2 >= 10")

	// Create the Lexer
	lexer := parser.NewExpressionLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewExpressionParser(stream)
	start := p.Start()
	for i := 0; i < b.N; i++ {
		parseWithCache(start)
	}
}
