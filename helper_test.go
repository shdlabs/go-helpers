package helpers

import (
	"fmt"
	"testing"
)

func TestHelperEqualFail(t *testing.T) {
	Equal(t, []string{"foo", "bar", "lan"}, []string{"foo", "bar", "lan"})
}

func ExampleEqual() {
	t := &testing.T{}
	Equal(t, []string{"foo", "bar", "ln"}, []string{"foo", "bar", "lan"})
	fmt.Printf("%v", t.Failed())
	// Output:
	// true
}
