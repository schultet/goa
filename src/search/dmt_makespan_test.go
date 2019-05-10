package search

import (
	"testing"

	"github.com/schultet/goa/src/task"

	"github.com/stretchr/testify/assert"
)

func TestPathToRoot(t *testing.T) {
	//           root {expensiveAction, cheapAction}
	//           /  \
	//  leftChild    rightChild
	// {leftAction}  {rightAction, secondRightAction}

	expensiveAction := &task.Action{Cost: 25}
	cheapAction := &task.Action{Cost: 10}

	root := &DmtMakespanNode{
		parent:  nil,
		actions: []*task.Action{expensiveAction, cheapAction},
		costs:   []int{expensiveAction.Cost, cheapAction.Cost},
	}

	actions, costs := root.pathToRoot()
	assert.Equal(t, []*task.Action{expensiveAction, cheapAction}, actions)
	assert.Equal(t, []int{expensiveAction.Cost, cheapAction.Cost}, costs)

	leftAction := &task.Action{Cost: 100}
	leftChild := &DmtMakespanNode{
		parent:  root,
		actions: []*task.Action{leftAction},
		costs:   []int{leftAction.Cost},
	}

	actions, costs = leftChild.pathToRoot()

	assert.Equal(t, []*task.Action{
		expensiveAction,
		cheapAction,
		leftAction,
	}, actions)

	assert.Equal(t, []int{
		expensiveAction.Cost,
		cheapAction.Cost,
		leftAction.Cost,
	}, costs)

	rightAction := &task.Action{Cost: 200}
	secondRightAction := &task.Action{Cost: 300}
	rightChild := &DmtMakespanNode{
		parent:  root,
		actions: []*task.Action{rightAction, secondRightAction},
		costs:   []int{rightAction.Cost, secondRightAction.Cost},
		h:       1,
		visits:  1,
	}

	actions, costs = rightChild.pathToRoot()
	assert.Equal(t, []*task.Action{
		expensiveAction,
		cheapAction,
		rightAction,
		secondRightAction,
	}, actions)
	assert.Equal(t, []int{
		expensiveAction.Cost,
		cheapAction.Cost,
		rightAction.Cost,
		secondRightAction.Cost,
	}, costs)
}

func TestMakespanApprox(t *testing.T) {
	//                          root {root1: 17}
	//                          /  \
	// {left1: 9, left2: 11} left  right {right1: 2}
	//
	// right1 is dependent on root1.
	// left1 and left2 are independent of root1.
	// left1 and left2 have the same owner.
	// All other actions have different owners.

	root1 := task.Action{
		Owner:   0,
		Cost:    17,
		Effects: []task.Effect{task.Effect{Variable: 1, Value: 10}},
	}
	left1 := task.Action{
		Owner: 1,
		Cost:  9,
	}
	left2 := task.Action{
		Owner: 1,
		Cost:  11,
	}
	right1 := task.Action{
		Owner:         2,
		Cost:          2,
		Preconditions: []task.Condition{task.Condition{Variable: 1, Value: 10}},
	}

	root := DmtMakespanNode{
		actions: []*task.Action{&root1},
		costs:   []int{root1.Cost},
	}
	left := DmtMakespanNode{
		parent:  &root,
		actions: []*task.Action{&left1, &left2},
		costs:   []int{left1.Cost, left2.Cost},
	}
	right := DmtMakespanNode{
		parent:  &root,
		actions: []*task.Action{&right1},
		costs:   []int{right1.Cost},
	}

	// Execute the single action in root.
	assert.Equal(t, root1.Cost, makespanApprox(root.pathToRoot()))

	// Execute root1 in parallel to the sequential execution of left1 and left2.
	leftMakespan := left1.Cost + left2.Cost // c(left1) + c(left2) > c(root1)
	assert.Equal(t, leftMakespan, makespanApprox(left.pathToRoot()))

	// Execute root1, then execute right1, as right1 depends on root1.
	rightMakespan := root1.Cost + right1.Cost
	assert.Equal(t, rightMakespan, makespanApprox(right.pathToRoot()))
}

func TestMakespanApprox2(t *testing.T) {
	actions := []*task.Action{
		{ID: 1, Owner: 3},
		{ID: 2, Owner: 3},
		{ID: 3, Owner: 2},
		{ID: 4, Owner: 0},
		{ID: 5, Owner: 0},
		{ID: 6, Owner: 0},
		{ID: 7, Owner: 0},
		{ID: 8, Owner: 0},
		{ID: 9, Owner: 0},
		{ID: 10, Owner: 0},
	}
	costs := []int{
		5, 1, 3, 1, 1, 1, 1, 1, 1, 1,
	}

	exp := 7
	has := makespanApprox(actions, costs)
	if exp != has {
		t.Fatalf("Exp:%v Has:%v\n", exp, has)
	}

}

func TestNonconcurrent(t *testing.T) {
	a := &task.Action{ID: 0, Owner: 0,
		Effects: []task.Effect{{0, 0}},
	}
	b := &task.Action{ID: 1, Owner: 1,
		Preconditions: []task.Condition{{0, 0}, {1, 1}},
	}
	c := &task.Action{ID: 2, Owner: 2,
		Effects: []task.Effect{{0, 1}},
	}
	d := &task.Action{ID: 3, Owner: 3}
	e := &task.Action{ID: 4, Owner: 3}
	f := &task.Action{ID: 4, Owner: 4}

	type testcase struct {
		a, b       *task.Action
		concurrent bool
	}

	// TODO: add more test cases
	tests := []testcase{
		{a, b, false}, // prod - cons
		{a, c, false}, // prod - threat (conflicting effects)
		{b, c, false}, // cons - threat
		{d, e, false}, // same owner
		{d, f, true},  // diff owner, same action otherwise
		{a, d, true},
		{d, a, true},
	}

	for _, test := range tests {
		if nonconcurrent(test.a, test.b) == test.concurrent {
			t.Errorf("For\n%+v\n%+v\nExpected: %v, Got: %v",
				*test.a,
				*test.b,
				false,
				true,
			)
		}
	}
}

func TestGetPublicCosts(t *testing.T) {
	actions := []*task.Action{
		{
			Owner: 0,
			Cost:  3,
		},
		{
			Owner: 1,
			Cost:  17,
		},
		{
			Owner: 0,
			Cost:  0,
		},
	}
	assert.Equal(t, []int{3, 17, 0}, getPublicCosts(actions, task.NormalCost))
	assert.Equal(t, []int{1, 1, 1}, getPublicCosts(actions, task.UnitCost))
	assert.Equal(t, []int{3, 17, 1}, getPublicCosts(actions, task.NormalMinOne))
}

//func TestThreat(t *testing.T) {
//	// Variable		Effect		Safe precondition	Thretened precondition
//	// 1			2			-					3
//	// 2			3			3					3
//	// 3			-			4					-
//
//	effects := []task.Effect{
//		task.Effect{Variable: 1, Value: 2},
//		task.Effect{Variable: 2, Value: 3},
//	}
//	safePreconditions := []task.Condition{
//		task.Condition{Variable: 2, Value: 3},
//		task.Condition{Variable: 3, Value: 4},
//	}
//	threatenedPreconditions := []task.Condition{
//		task.Condition{Variable: 1, Value: 3},
//		task.Condition{Variable: 2, Value: 3},
//	}
//
//	assert.False(t, threat(effects, safePreconditions))
//	assert.True(t, threat(effects, threatenedPreconditions))
//}
//
//func TestConflict(t *testing.T) {
//	// Variable	Effect	Conflicting effect	Safe effect
//	// 1		2		3					-
//	// 2		3		-					3
//	// 3		-		-					4
//
//	effects := []task.Effect{
//		task.Effect{Variable: 1, Value: 2},
//		task.Effect{Variable: 2, Value: 3},
//	}
//
//	conflictingEffects := []task.Effect{
//		task.Effect{Variable: 1, Value: 3},
//	}
//
//	safeEffects := []task.Effect{
//		task.Effect{Variable: 2, Value: 3},
//		task.Effect{Variable: 3, Value: 4},
//	}
//
//	assert.False(t, conflict(effects, safeEffects))
//	assert.False(t, conflict(safeEffects, effects))
//
//	assert.True(t, conflict(conflictingEffects, effects))
//	assert.True(t, conflict(effects, conflictingEffects))
//}
