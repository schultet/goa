package pddl

import (
	"testing"
)

func TestObjectListInsert(t *testing.T) {
	ol := ObjectList{}
	objects := []Object{{"o3", "object", false}, {"o2", "kx", false}, {"o1", "joda", false}}
	ol.Insert(objects[0], 0)
	ol.Insert(objects[1], 0)
	ol.Insert(objects[2], 2)
	if ol[0] != objects[1] || ol[1] != objects[0] || ol[2] != objects[2] {
		t.Errorf("Wrong insertion in ObjectList!")
	}
}

func TestObjectListOrderedInsert(t *testing.T) {
	o2 := ObjectList{}
	objects := []Object{{"o3", "object", false}, {"o2", "kx", false},
		{"o1", "joda", false}, {"o3", "", false}}
	for _, o := range objects {
		o2.OrderedInsert(o)
	}
	if !o2.equals(ObjectList{objects[2], objects[1], objects[0]}) {
		t.Errorf("Wrong OrderedInsert into ObjectList!")
	}
}
