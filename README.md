# sexpr
No, it's not about Public Relations for a popular activity.

It's a bluntly stolen package - originally found in "github.com/awalterschulze/gominikanren".

Package sexpr implements symbolic expressions

## API

```go
type V = *ast.Variable
type X = *ast.SExpr

func Parse(s string) (X, error)
```

Note: Expression constructors such as  

    func NewSymbol(s string) X
    func NewVariable(s string) X
    func NewList(ss ...X) X

are provided as a convenience
for test implementations only. 


---
## bnf

```bnf
SExpr           : Atom
                | Pair
                ;

Pair            : "(" ")"                           
                | "(" SExpr ")"                     
                | "(" SExpr space ContinueList ")"  
                | "(" SExpr space "." space SExpr ")"  
                ;

ContinueList    : SExpr
                | SExpr space ContinueList
                ;

Atom            : symbol
                | int_lit
                | float_lit
                | string_lit
                | variable
                ;
```

## go

```go
type SExpr struct {
	*Pair
	*Atom
}

type Pair struct {
	Car, Cdr *SExpr
}

type Atom struct {
	Str    *string
	Symbol *string
	Float  *float64
	Int    *int64
	Var    *Variable
}

```
