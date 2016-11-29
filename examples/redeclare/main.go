package main

//go:generate java -jar $GOPATH/src/github.com/millergarym/antlr4-go/lib/antlr4-4.6-SNAPSHOT.jar -o parser -visitor -Dlanguage=Go redeclare.g4
//go:generate sh -c "(cd parser; sed -i '' -e 's!github.com/antlr/antlr4/runtime/Go/antlr!github.com/millergarym/antlr4-go!' *.go )"

import (
	"fmt"

	"github.com/millergarym/antlr4-go"
	"github.com/millergarym/antlr4-go/examples/redeclare/parser"
)

type treeShapeListener struct {
	*parser.BaseredeclareListener
}

func (t *treeShapeListener) VisitTerminal(node antlr.TerminalNode) {
	if node.GetSymbol().GetTokenType() == antlr.TokenEOF {
		return
	}
}

func (t *treeShapeListener) EnterS(node *parser.SContext) {
	fmt.Printf("A:%v B:%v X:%v Y:%v\n", node.GetA(), node.GetB(), node.GetX(), node.GetY())
}

func main() {
	lexparsewalk("aa")
	lexparsewalk("bb")
}

func lexparsewalk(str string) {
	input := antlr.NewInputStream(str)
	lexer := parser.NewredeclareLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewredeclareParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.S()
	fmt.Println("--")
	l := &treeShapeListener{}
	antlr.ParseTreeWalkerDefault.Walk(l, tree)
}
