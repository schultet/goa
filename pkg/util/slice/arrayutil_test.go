package slice

import "testing"

func TestCopyIntArray(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := CopyIntArray(a)
	if !EqualIntArrays(a, b) {
		t.Errorf("ALARM")
	}
}

func TestReverseIntArray(t *testing.T) {
	a := []int{1, 2, 3, 4}
	exp := []int{4, 3, 2, 1}
	ReverseIntArray(a)
	if !EqualIntArrays(a, exp) {
		t.Errorf("ALARM")
	}
}

func TestInsertIntArray(t *testing.T) {
	a := []int{0, 4, 5}
	InsertIntArray(&a, 35, 2)
	InsertIntArray(&a, 35, 0)
	exp := []int{35, 0, 4, 35, 5}
	if !EqualIntArrays(a, exp) {
		t.Errorf("ALARM!\nexp:%v\nwas:%v\n", exp, a)
	}
}

func TestAppendIntArray(t *testing.T) {
	v1 := make([]int, 10)
	v2 := []int{1, 2, 3, 4, 5}
	result := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 2, 3, 4, 5}
	v2copy := []int{1, 2, 3, 4, 5}
	AppendIntArray(&v1, &v2)
	if !EqualIntArrays(v1, result) {
		t.Errorf("ALARM")
	}
	if !EqualIntArrays(v2, v2copy) {
		t.Errorf("ALARM")
	}
}

func TestContainSameInts(t *testing.T) {
	type testcase struct {
		a1  []int
		a2  []int
		res bool
	}

	tests := []testcase{
		{[]int{0, 1, 2}, []int{0, 2, 1}, true},
		{[]int{}, []int{}, true},
		{make([]int, 5), []int{0, 0, 0, 0, 0}, true},
		{[]int{0, 0, 1}, []int{0, 0}, false},
		{[]int{0, 1}, []int{0, 0}, false},
		{[]int{1, 1}, []int{1}, false},
		{[]int{1, 1, 2}, []int{2, 2, 1}, false},
	}

	for _, tc := range tests {
		if ContainSameInts(tc.a1, tc.a2) != tc.res {
			t.Errorf("arrays 'set'-equality check failed!\nExp:%v\nWas:%v\n",
				tc.res, ContainSameInts(tc.a1, tc.a2))
		}
	}

}

func TestEqualIntArrays(t *testing.T) {
	type testcase struct {
		a1  []int
		a2  []int
		res bool
	}

	tests := []testcase{
		{[]int{0, 1, 2}, []int{0, 1, 2}, true},
		{[]int{}, []int{}, true},
		{make([]int, 5), []int{0, 0, 0, 0, 0}, true},
		{[]int{0, 0, 1}, []int{0, 0}, false},
	}

	for _, tc := range tests {
		if EqualIntArrays(tc.a1, tc.a2) != tc.res {
			t.Errorf("arrays equality check failed!\nExp:%v\nWas:%v\n",
				tc.res, EqualIntArrays(tc.a1, tc.a2))
		}
	}
}
