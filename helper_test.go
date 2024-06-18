package helpers

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestHelpers(t *testing.T) {
	t.Run("Equal", func(t *testing.T) {
		Equal(t, []string{"foo", "bar", "lan"}, []string{"foo", "bar", "lan"})
	})
	t.Run("NotNil", func(t *testing.T) {
		a := struct{ l string }{l: "foo"}
		NotNil(t, a)
	})
	t.Run("NoError", func(t *testing.T) {
		NoError(t, nil)
	})
	t.Run("NotEqual", func(t *testing.T) {
		NotEqual(t, "foo", "bar")
		NotEqual(t, []string{"food", "br", "lan"}, []string{"foo", "bar", "lan"})
	})
	t.Run("DurationLog", func(t *testing.T) {
		start := time.Now()
		defer DurationLog(start, "foo")
		time.Sleep(time.Millisecond * 500)
	})
}

func ExampleEqual() {
	t := &testing.T{}
	Equal(t, []string{"foo", "bar", "ln"}, []string{"foo", "bar", "lan"})
	fmt.Printf("%v", t.Failed())
	// Output:
	// true
}

func ExampleNotNil() {
	t := &testing.T{}
	NotNil(t, nil)
	fmt.Printf("%v", t.Failed())
	// Output:
	// true
}

func ExampleNoError() {
	t := &testing.T{}
	NoError(t, errors.New("unexpected error"))
	fmt.Printf("%v", t.Failed())
	// Output:
	// true
}

func ExampleNotEqual() {
	t := &testing.T{}
	NotEqual(t, 2, 2)
	fmt.Printf("%v", t.Failed())
	// Output:
	// true
}

func ExampleAh() {
	fmt.Printf("%s", Ah("foo"))
	// Output:
	// [34;3m 🤨 foo [0m
}

func ExampleOk() {
	fmt.Printf("%s", Ok("foo"))
	// Output:
	// [32m ✅ foo [0m
}
