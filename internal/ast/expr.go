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

// AsVariable reports whether the expression expresses nothing but a variable,
// and if so, that variable is returned (or else: nil).
func (s *SExpr) AsVariable() (v *Variable, isVariable bool) {
	isVariable = s.IsVariable()
	if isVariable {
		v = s.Atom.Var
	}
	return
}
