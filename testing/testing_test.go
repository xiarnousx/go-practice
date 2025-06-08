package testexample

import (
	"fmt"
	"testing"
)

// Table test example
func TestAbc(t *testing.T) {
	type test struct {
		data   string
		answer string
	}

	tests := []test{
		test{data: "a", answer: "a"},
		test{data: "b", answer: "b"},
		test{data: "c", answer: "c"},
	}

	for _, v := range tests {
		a := Abc(v.data)
		if a != v.answer {
			t.Error("Expected", v.answer, "Got", a)
		}
	}
}

func TestSum(t *testing.T) {
	s := Sum(1, 2)
	if s != 3 {
		t.Error("expected", 3, "got", s)
	}
}

func ExampleSum() {
	fmt.Println(Sum(1, 2))
	//Output:
	// 3
}
