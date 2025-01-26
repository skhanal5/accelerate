package example

import "testing"

func TestAddTwoNums(t *testing.T) {
	inputA, inputB := 1, 2
	want := 3
	result := AddTwoNums(inputA, inputB)

	if want != result {
		t.Errorf("got %q, wanted %q", result, want)
	}

}
