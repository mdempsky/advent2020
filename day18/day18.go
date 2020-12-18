package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"strings"

	advent "github.com/mdempsky/advent2020"
)

func main() {
	part1, part2 := 0, 0
	for _, line := range advent.InputLines() {
		part1 += solve(line, token.SUB)
		part2 += solve(line, token.EQL)
	}
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func solve(s string, mul token.Token) int {
	expr, err := parser.ParseExpr(strings.ReplaceAll(s, "*", mul.String()))
	if err != nil {
		log.Fatal(err)
	}

	var eval func(ast.Expr) int
	eval = func(expr ast.Expr) int {
		switch expr := expr.(type) {
		case *ast.BasicLit:
			return int(advent.Atoi(expr.Value))
		case *ast.ParenExpr:
			return eval(expr.X)
		case *ast.BinaryExpr:
			x, y := eval(expr.X), eval(expr.Y)
			switch expr.Op {
			case token.ADD:
				return x + y
			case mul:
				return x * y
			}
		}
		panic(fmt.Sprint(expr))
	}

	return eval(expr)
}
