package integer

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(4, 2)
	expected := 6

	if sum != expected {
		t.Errorf("got: %q want: %q", sum, expected)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
