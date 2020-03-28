// Package sexpr implements symbolic expressions
package gominikanren

import (
	sexpr "github.com/GoLangsam/sexpr/internal"
	"github.com/GoLangsam/sexpr/internal/ast"
)

type Expression = ast.SExpr

type Variable = ast.Variable

//type Symbol = ast.Symbol
type Atom = ast.Atom

func Parse(s string) (*Expression, error) {
	return sexpr.Parse(s)
}

// in Test-Programmen will/man oft speziell erzeugen:

func NewSymbol(s string) *Expression {
	return ast.NewSymbol(s)
}

func NewVariable(s string) *Expression {
	return ast.NewVariable(s)
}

func NewList(ss ...*Expression) *Expression {
	return ast.NewList(ss...)
}

/*

   func Cons(car *SExpr, cdr *SExpr) *SExpr
   func NewFloat(f float64) *Expression
   func NewInt(i int64) *Expression
   func NewString(s string) *Expression
*/
