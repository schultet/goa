package search

//import (
//	"gkigit.informatik.uni-freiburg.de/rgroensfeld/makespan"
//	"gkigit.informatik.uni-freiburg.de/tschulte/dimap/src/task"
//)
//
//func convertAction(from task.Action) makespan.Action {
//	return makespan.Action{
//		Owner:        from.Owner,
//		Costs:        float64(from.Cost),
//		Precondition: convertConditions(from.Preconditions),
//		Effect:       convertEffects(from.Effects),
//	}
//}
//
//func convertConditions(from []task.Condition) makespan.Assignments {
//	var assignments []task.VariableValuePair
//	for _, condition := range from {
//		assignments = append(assignments, task.VariableValuePair(condition))
//	}
//	return convertAssignments(assignments)
//}
//
//func convertAssignments(from []task.VariableValuePair) makespan.Assignments {
//	var assignments makespan.Assignments
//	for _, condition := range from {
//		assignments.Assign(condition.Variable, condition.Value)
//	}
//	return assignments
//}
//
//func convertEffects(from []task.Effect) makespan.Assignments {
//	var assignments []task.VariableValuePair
//	for _, condition := range from {
//		assignments = append(assignments, task.VariableValuePair(condition))
//	}
//	return convertAssignments(assignments)
//}
