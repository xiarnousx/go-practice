package testexample

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	s := Greet("James")
	if s != "Welcome dear James" {
		t.Error("expected", "Welcome dear James", "got", s)
	}
}

func ExampleGreet() {
	fmt.Println(Greet("James"))
	// Output:
	// Welcome dear James
}

func BenchmarkGreet(b *testing.B) {
	for _ = range b.N {
		Greet("James")
	}
}
