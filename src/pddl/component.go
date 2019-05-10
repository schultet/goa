package pddl

import "strings"

type Component interface {
	Init(pddl string) Component // init a pddl component from string
	Pddl() string               // pddl representation
	NameIdentifier() string     // name to unambigiously identify component
}

type PrivacyHolder interface {
	IsPrivate() bool
}

type TypeHolder interface {
	Type() string
}

type ComponentList []Component

// ComponentList implements sort.Interface
func (l ComponentList) Len() int           { return len(l) }
func (l ComponentList) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l ComponentList) Less(i, j int) bool { return l[i].NameIdentifier() < l[j].NameIdentifier() }

func (l ComponentList) String() string {
	names := make([]string, len(l))
	for i, c := range l {
		names[i] = c.NameIdentifier()
	}
	return strings.Join(names, "\n")
}
