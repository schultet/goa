package pddl

import (
	"strconv"
)

type Function struct {
	name string
	args ObjectList
	kind string
}

func (f *Function) Init(tl TokenList) *Function {
	f.name = tl[1]
	vargs := extractParameters(tl[2:])
	f.args = make(ObjectList, len(vargs))
	for i, v := range vargs {
		f.args[i] = *NewTypedObject(v.name, v.kind)
	}
	f.kind = tl[len(tl)-1]
	return f
}

func (f Function) String() string {
	result := "(" + f.name
	if len(f.args) > 0 {
		result += " "
	}
	return result + f.args.String() + ") - " + f.kind
}

func (f Function) Pddl() string { return f.String() }

type NumericFluent struct {
	name   string
	args   ObjectList
	number int
}

func (n *NumericFluent) Init(tl TokenList) *NumericFluent {
	tl = tl[2:] // skip first two elements: ( =
	if tl[0] == "(" {
		n.name = tl[1]
		vargs := extractParameters(tl[2:])
		n.args = make(ObjectList, len(vargs))
		for i, v := range vargs {
			n.args[i] = *NewTypedObject(v.name, v.kind)
		}
		n.number, _ = strconv.Atoi(tl[len(tl)-2])
	} else {
		n.name = tl[0]
		n.args = make(ObjectList, 0)
		n.number, _ = strconv.Atoi(tl[len(tl)-2])
	}
	return n
}

func (n NumericFluent) String() string {
	if len(n.args) > 0 {
		return "(= (" + n.name + " " + n.args.String() + ") " + strconv.Itoa(n.number) + ")"
	} else {
		return "(= " + n.name + " " + strconv.Itoa(n.number) + ")"
	}
}

func (n NumericFluent) Pddl() string { return n.String() }
