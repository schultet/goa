package heuristic

import "github.com/schultet/goa/pkg/state"
import "github.com/schultet/goa/pkg/task"

// goalCount returns the number of fluents in which the given state differs from
// the goal state
type goalcountHeuristic struct {
	goal task.ConditionSet
}

func GCHeuristic(t *task.Task) StateEvaluator {
	return &goalcountHeuristic{t.Goal}
}

func (gc *goalcountHeuristic) Evaluate(s state.State) int {
	n := 0
	for _, c := range gc.goal {
		if s[c.Variable] != int(c.Value) {
			n += 1
		}
	}
	return n
}

// register heuristic
func init() {
	Register(&EvaluatorInfo{"gc", "goal count heuristic", GCHeuristic})
}
