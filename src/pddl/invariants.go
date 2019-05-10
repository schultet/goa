package pddl

// TODO: the invariants-module is currently not implemented. In the long run this
// module provides the functionality to find or identify mutex-groups/invariants.
//

// representing invariants and invariant candidates
type (
	VariableSet  []Variable
	AtomSet      []Predicate
	InvariantSet []invariant

	invariant struct {
		parameters VariableSet // set of fol variables (V)
		atoms      AtomSet     // set of atoms (Phi)
	}
)

// returns the set of initial invariant candidates
func initialCandidates(dom Domain) InvariantSet {
	// TODO: implement
	return InvariantSet{}
}

// proves whether an invariant candidate (V, Phi) is an invariant
func proveInvariant(V VariableSet, Phi AtomSet) bool {
	// TODO: implement
	return true
}

func operatorTooHeavy(o Action, V VariableSet, Phi AtomSet) bool {
	// TODO: implement
	return true
}

func unbalancedAddEffect(o Action, e FolFormula, V VariableSet, Phi AtomSet) bool {
	// TODO: implement
	return true // the add effect is unbalanced
}

// refines an unbalanced invariant candidate (V, Phi)
func refineCandidate(V VariableSet, Phi AtomSet) {
	// TODO: implement
}

// returns the set of invariants, basically performs breadth-first search
func computeInvariants() InvariantSet {
	// TODO: implement
	return InvariantSet{}
}
