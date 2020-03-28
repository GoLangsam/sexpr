// Code generated by gocc; DO NOT EDIT.

package parser

type (
	actionTable [numStates]actionRow
	actionRow   struct {
		canRecover bool
		actions    [numSymbols]action
	}
)

var actionTab = actionTable{
	actionRow{ // S0
		canRecover: false,
		actions: [numSymbols]action{
			nil,      /* INVALID */
			nil,      /* $ */
			shift(4), /* ( */
			nil,      /* ) */
			nil,      /* space */
			nil,      /* . */
			shift(5), /* symbol */
			shift(6), /* int_lit */
			shift(7), /* float_lit */
			shift(8), /* string_lit */
			shift(9), /* variable */
		},
	},
	actionRow{ // S1
		canRecover: false,
		actions: [numSymbols]action{
			nil,          /* INVALID */
			accept(true), /* $ */
			nil,          /* ( */
			nil,          /* ) */
			nil,          /* space */
			nil,          /* . */
			nil,          /* symbol */
			nil,          /* int_lit */
			nil,          /* float_lit */
			nil,          /* string_lit */
			nil,          /* variable */
		},
	},
	actionRow{ // S2
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(1), /* $, reduce: SExpr */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* space */
			nil,       /* . */
			nil,       /* symbol */
			nil,       /* int_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* variable */
		},
	},
	actionRow{ // S3
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(2), /* $, reduce: SExpr */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* space */
			nil,       /* . */
			nil,       /* symbol */
			nil,       /* int_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* variable */
		},
	},
	actionRow{ // S4
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			shift(13), /* ( */
			shift(14), /* ) */
			nil,       /* space */
			nil,       /* . */
			shift(15), /* symbol */
			shift(16), /* int_lit */
			shift(17), /* float_lit */
			shift(18), /* string_lit */
			shift(19), /* variable */
		},
	},
	actionRow{ // S5
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(9), /* $, reduce: Atom */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* space */
			nil,       /* . */
			nil,       /* symbol */
			nil,       /* int_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* variable */
		},
	},
	actionRow{ // S6
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(10), /* $, reduce: Atom */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* space */
			nil,        /* . */
			nil,        /* symbol */
			nil,        /* int_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* variable */
		},
	},
	actionRow{ // S7
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(11), /* $, reduce: Atom */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* space */
			nil,        /* . */
			nil,        /* symbol */
			nil,        /* int_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* variable */
		},
	},
	actionRow{ // S8
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(12), /* $, reduce: Atom */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* space */
			nil,        /* . */
			nil,        /* symbol */
			nil,        /* int_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* variable */
		},
	},
	actionRow{ // S9
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(13), /* $, reduce: Atom */
			nil,        /* ( */
			nil,        /* ) */
			nil,        /* space */
			nil,        /* . */
			nil,        /* symbol */
			nil,        /* int_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* variable */
		},
	},
	actionRow{ // S10
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* ( */
			shift(20), /* ) */
			shift(21), /* space */
			nil,       /* . */
			nil,       /* symbol */
			nil,       /* int_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* variable */
		},
	},
	actionRow{ // S11
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* ( */
			reduce(1), /* ), reduce: SExpr */
			reduce(1), /* space, reduce: SExpr */
			nil,       /* . */
			nil,       /* symbol */
			nil,       /* int_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* variable */
		},
	},
	actionRow{ // S12
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* ( */
			reduce(2), /* ), reduce: SExpr */
			reduce(2), /* space, reduce: SExpr */
			nil,       /* . */
			nil,       /* symbol */
			nil,       /* int_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* variable */
		},
	},
	actionRow{ // S13
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			shift(13), /* ( */
			shift(23), /* ) */
			nil,       /* space */
			nil,       /* . */
			shift(15), /* symbol */
			shift(16), /* int_lit */
			shift(17), /* float_lit */
			shift(18), /* string_lit */
			shift(19), /* variable */
		},
	},
	actionRow{ // S14
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(3), /* $, reduce: Pair */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* space */
			nil,       /* . */
			nil,       /* symbol */
			nil,       /* int_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* variable */
		},
	},
	actionRow{ // S15
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* ( */
			reduce(9), /* ), reduce: Atom */
			reduce(9), /* space, reduce: Atom */
			nil,       /* . */
			nil,       /* symbol */
			nil,       /* int_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* variable */
		},
	},
	actionRow{ // S16
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* ( */
			reduce(10), /* ), reduce: Atom */
			reduce(10), /* space, reduce: Atom */
			nil,        /* . */
			nil,        /* symbol */
			nil,        /* int_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* variable */
		},
	},
	actionRow{ // S17
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* ( */
			reduce(11), /* ), reduce: Atom */
			reduce(11), /* space, reduce: Atom */
			nil,        /* . */
			nil,        /* symbol */
			nil,        /* int_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* variable */
		},
	},
	actionRow{ // S18
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* ( */
			reduce(12), /* ), reduce: Atom */
			reduce(12), /* space, reduce: Atom */
			nil,        /* . */
			nil,        /* symbol */
			nil,        /* int_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* variable */
		},
	},
	actionRow{ // S19
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* ( */
			reduce(13), /* ), reduce: Atom */
			reduce(13), /* space, reduce: Atom */
			nil,        /* . */
			nil,        /* symbol */
			nil,        /* int_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* variable */
		},
	},
	actionRow{ // S20
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(4), /* $, reduce: Pair */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* space */
			nil,       /* . */
			nil,       /* symbol */
			nil,       /* int_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* variable */
		},
	},
	actionRow{ // S21
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			shift(13), /* ( */
			nil,       /* ) */
			nil,       /* space */
			shift(26), /* . */
			shift(15), /* symbol */
			shift(16), /* int_lit */
			shift(17), /* float_lit */
			shift(18), /* string_lit */
			shift(19), /* variable */
		},
	},
	actionRow{ // S22
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* ( */
			shift(27), /* ) */
			shift(28), /* space */
			nil,       /* . */
			nil,       /* symbol */
			nil,       /* int_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* variable */
		},
	},
	actionRow{ // S23
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* ( */
			reduce(3), /* ), reduce: Pair */
			reduce(3), /* space, reduce: Pair */
			nil,       /* . */
			nil,       /* symbol */
			nil,       /* int_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* variable */
		},
	},
	actionRow{ // S24
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* ( */
			reduce(7), /* ), reduce: ContinueList */
			shift(29), /* space */
			nil,       /* . */
			nil,       /* symbol */
			nil,       /* int_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* variable */
		},
	},
	actionRow{ // S25
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* ( */
			shift(30), /* ) */
			nil,       /* space */
			nil,       /* . */
			nil,       /* symbol */
			nil,       /* int_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* variable */
		},
	},
	actionRow{ // S26
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* ( */
			nil,       /* ) */
			shift(31), /* space */
			nil,       /* . */
			nil,       /* symbol */
			nil,       /* int_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* variable */
		},
	},
	actionRow{ // S27
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* ( */
			reduce(4), /* ), reduce: Pair */
			reduce(4), /* space, reduce: Pair */
			nil,       /* . */
			nil,       /* symbol */
			nil,       /* int_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* variable */
		},
	},
	actionRow{ // S28
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			shift(13), /* ( */
			nil,       /* ) */
			nil,       /* space */
			shift(33), /* . */
			shift(15), /* symbol */
			shift(16), /* int_lit */
			shift(17), /* float_lit */
			shift(18), /* string_lit */
			shift(19), /* variable */
		},
	},
	actionRow{ // S29
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			shift(13), /* ( */
			nil,       /* ) */
			nil,       /* space */
			nil,       /* . */
			shift(15), /* symbol */
			shift(16), /* int_lit */
			shift(17), /* float_lit */
			shift(18), /* string_lit */
			shift(19), /* variable */
		},
	},
	actionRow{ // S30
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(5), /* $, reduce: Pair */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* space */
			nil,       /* . */
			nil,       /* symbol */
			nil,       /* int_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* variable */
		},
	},
	actionRow{ // S31
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			shift(38), /* ( */
			nil,       /* ) */
			nil,       /* space */
			nil,       /* . */
			shift(39), /* symbol */
			shift(40), /* int_lit */
			shift(41), /* float_lit */
			shift(42), /* string_lit */
			shift(43), /* variable */
		},
	},
	actionRow{ // S32
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* ( */
			shift(44), /* ) */
			nil,       /* space */
			nil,       /* . */
			nil,       /* symbol */
			nil,       /* int_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* variable */
		},
	},
	actionRow{ // S33
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* ( */
			nil,       /* ) */
			shift(45), /* space */
			nil,       /* . */
			nil,       /* symbol */
			nil,       /* int_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* variable */
		},
	},
	actionRow{ // S34
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* ( */
			reduce(8), /* ), reduce: ContinueList */
			nil,       /* space */
			nil,       /* . */
			nil,       /* symbol */
			nil,       /* int_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* variable */
		},
	},
	actionRow{ // S35
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* ( */
			shift(46), /* ) */
			nil,       /* space */
			nil,       /* . */
			nil,       /* symbol */
			nil,       /* int_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* variable */
		},
	},
	actionRow{ // S36
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* ( */
			reduce(1), /* ), reduce: SExpr */
			nil,       /* space */
			nil,       /* . */
			nil,       /* symbol */
			nil,       /* int_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* variable */
		},
	},
	actionRow{ // S37
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* ( */
			reduce(2), /* ), reduce: SExpr */
			nil,       /* space */
			nil,       /* . */
			nil,       /* symbol */
			nil,       /* int_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* variable */
		},
	},
	actionRow{ // S38
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			shift(13), /* ( */
			shift(48), /* ) */
			nil,       /* space */
			nil,       /* . */
			shift(15), /* symbol */
			shift(16), /* int_lit */
			shift(17), /* float_lit */
			shift(18), /* string_lit */
			shift(19), /* variable */
		},
	},
	actionRow{ // S39
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* ( */
			reduce(9), /* ), reduce: Atom */
			nil,       /* space */
			nil,       /* . */
			nil,       /* symbol */
			nil,       /* int_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* variable */
		},
	},
	actionRow{ // S40
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* ( */
			reduce(10), /* ), reduce: Atom */
			nil,        /* space */
			nil,        /* . */
			nil,        /* symbol */
			nil,        /* int_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* variable */
		},
	},
	actionRow{ // S41
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* ( */
			reduce(11), /* ), reduce: Atom */
			nil,        /* space */
			nil,        /* . */
			nil,        /* symbol */
			nil,        /* int_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* variable */
		},
	},
	actionRow{ // S42
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* ( */
			reduce(12), /* ), reduce: Atom */
			nil,        /* space */
			nil,        /* . */
			nil,        /* symbol */
			nil,        /* int_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* variable */
		},
	},
	actionRow{ // S43
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* ( */
			reduce(13), /* ), reduce: Atom */
			nil,        /* space */
			nil,        /* . */
			nil,        /* symbol */
			nil,        /* int_lit */
			nil,        /* float_lit */
			nil,        /* string_lit */
			nil,        /* variable */
		},
	},
	actionRow{ // S44
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* ( */
			reduce(5), /* ), reduce: Pair */
			reduce(5), /* space, reduce: Pair */
			nil,       /* . */
			nil,       /* symbol */
			nil,       /* int_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* variable */
		},
	},
	actionRow{ // S45
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			shift(38), /* ( */
			nil,       /* ) */
			nil,       /* space */
			nil,       /* . */
			shift(39), /* symbol */
			shift(40), /* int_lit */
			shift(41), /* float_lit */
			shift(42), /* string_lit */
			shift(43), /* variable */
		},
	},
	actionRow{ // S46
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(6), /* $, reduce: Pair */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* space */
			nil,       /* . */
			nil,       /* symbol */
			nil,       /* int_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* variable */
		},
	},
	actionRow{ // S47
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* ( */
			shift(50), /* ) */
			shift(51), /* space */
			nil,       /* . */
			nil,       /* symbol */
			nil,       /* int_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* variable */
		},
	},
	actionRow{ // S48
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* ( */
			reduce(3), /* ), reduce: Pair */
			nil,       /* space */
			nil,       /* . */
			nil,       /* symbol */
			nil,       /* int_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* variable */
		},
	},
	actionRow{ // S49
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* ( */
			shift(52), /* ) */
			nil,       /* space */
			nil,       /* . */
			nil,       /* symbol */
			nil,       /* int_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* variable */
		},
	},
	actionRow{ // S50
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* ( */
			reduce(4), /* ), reduce: Pair */
			nil,       /* space */
			nil,       /* . */
			nil,       /* symbol */
			nil,       /* int_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* variable */
		},
	},
	actionRow{ // S51
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			shift(13), /* ( */
			nil,       /* ) */
			nil,       /* space */
			shift(54), /* . */
			shift(15), /* symbol */
			shift(16), /* int_lit */
			shift(17), /* float_lit */
			shift(18), /* string_lit */
			shift(19), /* variable */
		},
	},
	actionRow{ // S52
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* ( */
			reduce(6), /* ), reduce: Pair */
			reduce(6), /* space, reduce: Pair */
			nil,       /* . */
			nil,       /* symbol */
			nil,       /* int_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* variable */
		},
	},
	actionRow{ // S53
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* ( */
			shift(55), /* ) */
			nil,       /* space */
			nil,       /* . */
			nil,       /* symbol */
			nil,       /* int_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* variable */
		},
	},
	actionRow{ // S54
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* ( */
			nil,       /* ) */
			shift(56), /* space */
			nil,       /* . */
			nil,       /* symbol */
			nil,       /* int_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* variable */
		},
	},
	actionRow{ // S55
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* ( */
			reduce(5), /* ), reduce: Pair */
			nil,       /* space */
			nil,       /* . */
			nil,       /* symbol */
			nil,       /* int_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* variable */
		},
	},
	actionRow{ // S56
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			shift(38), /* ( */
			nil,       /* ) */
			nil,       /* space */
			nil,       /* . */
			shift(39), /* symbol */
			shift(40), /* int_lit */
			shift(41), /* float_lit */
			shift(42), /* string_lit */
			shift(43), /* variable */
		},
	},
	actionRow{ // S57
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* ( */
			shift(58), /* ) */
			nil,       /* space */
			nil,       /* . */
			nil,       /* symbol */
			nil,       /* int_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* variable */
		},
	},
	actionRow{ // S58
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* ( */
			reduce(6), /* ), reduce: Pair */
			nil,       /* space */
			nil,       /* . */
			nil,       /* symbol */
			nil,       /* int_lit */
			nil,       /* float_lit */
			nil,       /* string_lit */
			nil,       /* variable */
		},
	},
}
