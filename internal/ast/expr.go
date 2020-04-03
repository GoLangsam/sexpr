package ast

// Expr returns a singleton symbolic expression composed only of
// the variable v.
func (v *Variable) Expr() *SExpr {
	return &SExpr{
		Atom: &Atom{
			Var: v,
		},
	}
}

// Expr returns a singleton symbolic expression composed only of
// the atom a.
func (a *Atom) Expr() *SExpr {
	return &SExpr{
		Atom: a,
	}
}

// Expr returns a singleton symbolic expression composed only of
// the pair p.
func (p *Pair) Expr() *SExpr {
	return &SExpr{
		Pair: p,
	}
}
