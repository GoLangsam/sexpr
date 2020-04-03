// Package sexpr implements symbolic expressions
package sexpr

import (
	"github.com/GoLangsam/sexpr/internal"
	"github.com/GoLangsam/sexpr/internal/ast"
)

type Expression = ast.SExpr

type Variable = ast.Variable

type Atom = ast.Atom
type Pair = ast.Pair

var (
	Parse = sexpr.Parse

	Cons = ast.Cons

	NewString   = ast.NewString
	NewSymbol   = ast.NewSymbol
	NewInt      = ast.NewInt
	NewFloat    = ast.NewFloat
	NewVariable = ast.NewVariable

	NewList = ast.NewList
)
