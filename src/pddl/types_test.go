package pddl

import (
	"fmt"
	"strings"
	"testing"

	"github.com/schultet/goa/src/util/slice"
)

func init() {
	fmt.Println("testing types...")
}

// _____________________________________________________________________________
func TestParseTypes(t *testing.T) {
	typetokens := strings.Fields(`( :types location agent - object taxi passenger - agent ) `)
	types, _ := ParseTypes(typetokens)

	typeNames := []string{}
	for _, t := range types {
		typeNames = append(typeNames, t.name)
	}
	expectedTypes := []string{"object", "location", "agent", "taxi", "passenger"}
	if !slice.ContainSameStrings(typeNames, expectedTypes) {
		t.Errorf("types dont match!\nexp:%v\nwas:%v\n", expectedTypes, typeNames)
	}
}
