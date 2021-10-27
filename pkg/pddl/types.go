package pddl

import "fmt"

type Type struct {
	name     string
	parent   *Type
	children []*Type
}

type TypeMap map[string]*Type

// Returns a string representation of a PDDL-Type
func (t Type) String() string {
	children_names := make([]string, 0, len(t.children))
	for _, c := range t.children {
		children_names = append(children_names, c.name)
	}
	parentName := "nil"
	if t.parent != nil {
		parentName = t.parent.name
	}
	return fmt.Sprintf("Type{name: %v, parent: %v, children: %q}", t.name,
		parentName, children_names)
}
