package example

// Current only work on OSX, the sed command is different on darwin

//go:generate java -jar ../../lib/antlr4-4.5.4-SNAPSHOT.jar ExprLR.g4 -o parser -visitor -Dlanguage=Go
//go:generate sh -c "(cd parser; sed -i '' -e 's!github.com/antlr/antlr4/runtime/Go/antlr!../../../!' *.go )"

import (
	"bytes"
	"fmt"
	"testing"

	"./parser"

	antlr "../../"
)

type treeShapeListener struct {
	*parser.BaseExprLRListener
	buf bytes.Buffer
}

func (t *treeShapeListener) VisitTerminal(node antlr.TerminalNode) {
	if node.GetSymbol().GetTokenType() == antlr.TokenEOF {
		return
	}
	t.buf.WriteString(node.GetText())
}

func TestMain(t *testing.T) {
	input := antlr.NewInputStream("1+2^3")
	lexer := parser.NewExprLRLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewExprLRParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Formula()
	fmt.Println("--")
	l := &treeShapeListener{}
	antlr.ParseTreeWalkerDefault.Walk(l, tree)
	fmt.Println(l.buf.String())
}
