package pddl

import "errors"

type Object struct {
	name    string
	kind    string
	private bool
}

func NewTypedObject(name, kind string) *Object {
	return &Object{name, kind, false}
}

func DummyObject(name string) *Object { return &Object{name: name} }

func (o Object) Name() string { return o.name }
func (o Object) Kind() string { return o.kind }

// isPrivate returns a boolean representing whether the object is private or not.
func (o *Object) IsPrivate() bool { return o.private }

// Pddl returns a PDDL representation of a TypedObject as string.
func (o Object) Pddl() string {
	if o.kind != "" {
		return o.name + " - " + o.kind
	} else {
		return o.name
	}
}

// String returns a string representation of a TypedObject.
func (t Object) String() string {
	typeStr := "nil"
	if t.kind != "" {
		typeStr = t.kind
	}
	return "Object{" + t.name + ", " + typeStr + "}"
}

// Array of typed objects, supports filter and comparison methods
type ObjectList []Object

// ObjectList implement sort.Interface
func (ol ObjectList) Len() int           { return len(ol) }
func (ol ObjectList) Swap(i, j int)      { ol[i], ol[j] = ol[j], ol[i] }
func (ol ObjectList) Less(i, j int) bool { return ol[i].Name() < ol[j].Name() }

func (ol ObjectList) String() string {
	result := ""
	for i, o := range ol {
		result += o.Pddl()
		if i < len(ol)-1 {
			result += " "
		}
	}
	return result
}

// Insert object o at position i into the object list
func (ol *ObjectList) Insert(o Object, i int) {
	*ol = append(*ol, o)
	copy((*ol)[i+1:], (*ol)[i:])
	(*ol)[i] = o
}

// Inserts a new element into the ObjectList, such that the list remains sorted
func (ol *ObjectList) OrderedInsert(o Object) {
	for i := 0; i < len(*ol); i++ {
		if o.Name() < (*ol)[i].Name() {
			ol.Insert(o, i)
			return
		} else if o.Name() == (*ol)[i].Name() {
			return
		}
	}
	*ol = append(*ol, o)
}

// ObjectLists are ordered equal iff they contain the same objects in the same order
func (ol ObjectList) OrderedEquals(other ObjectList) bool {
	if len(ol) != len(other) {
		return false
	}
	for i, e := range ol {
		if e != other[i] {
			return false
		}
	}
	return true
}

//TODO: testing
// returns all objects of type t (not including subtypes)
func (ol ObjectList) ofType(t *Type) ObjectList {
	res := make(ObjectList, 0)
	for i := range ol {
		if ol[i].kind == t.name {
			res = append(res, ol[i])
		}
	}
	return res
}

//TODO: testing
// returns all objects of type t (including subtypes)
// TODO; objects could be stored sorted by type/subtype... -> saves computation
func (ol ObjectList) ofTypeAndSubtypes(t *Type) ObjectList {
	res := make(ObjectList, 0)
	types := []*Type{t}
	for len(types) > 0 {
		t := types[0]
		types = types[1:]
		types = append(types, t.children...)
		res = append(res, ol.ofType(t)...)
	}
	return res
}

//TODO: testing
// from an ObjectList returns the object with name 'name' or nil if no such
// object can be found
func (ol ObjectList) Search(name string) (*Object, error) {
	for i := range ol {
		if ol[i].name == name {
			return &ol[i], nil
		}
	}
	return nil, errors.New("Object not in ObjectList!")
}

// returns true if ObjectList contains object with name 'name'
func (ol ObjectList) contains(name string) bool {
	_, err := ol.Search(name)
	return err == nil
}

// TODO: testing + rename to sth. like UnorderedEquals to differentiate
// two ObjectLists are equal if they contain objects with the same name and have
// the same length
func (ol ObjectList) equals(other ObjectList) bool {
	if len(ol) != len(other) {
		return false
	}
	elems := make(map[string]int)
	for _, o := range ol {
		if _, ok := elems[o.name]; ok {
			elems[o.name] += 1
		} else {
			elems[o.name] = 1
		}
	}
	for _, o := range other {
		if _, ok := elems[o.name]; ok {
			elems[o.name] -= 1
		} else {
			return false
		}
	}
	for _, v := range elems {
		if v != 0 {
			return false
		}
	}
	return true
}

// filter returns a list of objects that match the filter criterion f
func (ol ObjectList) filter(f func(*Object) bool) ObjectList {
	matches := make(ObjectList, 0)
	for i := range ol {
		if f(&ol[i]) {
			matches = append(matches, ol[i])
		}
	}
	return matches
}

func (ol *ObjectList) Privates() ObjectList {
	return ol.filter((*Object).IsPrivate)
}

func (ol *ObjectList) Publics() ObjectList {
	return ol.filter(func(o *Object) bool { return !o.IsPrivate() })
}
