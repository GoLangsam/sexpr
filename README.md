# sexpr
No, it's not about Public Relations for a popular activity.

It's a bluntly stolen package - originally found in "github.com/awalterschulze/gominikanren".

Package sexpr implements symbolic expressions

func Parse(s string) (*Expression, error)

type Atom = ast.Atom
type Variable = ast.Variable

type Expression = ast.SExpr

    func NewSymbol(s string) *Expression
    func NewVariable(s string) *Expression
    func NewList(ss ...*Expression) *Expression
