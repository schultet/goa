package heuristic

import (
	"github.com/schultet/goa/pkg/state"
	"github.com/schultet/goa/pkg/task"
)

// blindHeuristic returns 1 for all non-goal states and 0 for goal-states
type BlindHeuristic struct {
	goal task.ConditionSet
}

func ConditionSatisfied(c task.Condition, s state.State) bool {
	return s[c.Variable] == int(c.Value)
}

func (b *BlindHeuristic) Evaluate(s state.State) int {
	for _, c := range b.goal {
		if !ConditionSatisfied(c, s) {
			return 1
		}
	}
	return 0
}

func NewBlindHeuristic(t *task.Task) StateEvaluator {
	return &BlindHeuristic{t.Goal}
}

// registers the blind heuristic
func init() {
	Register(&EvaluatorInfo{"blind", "the blind heuristic", NewBlindHeuristic})
}
