package helpers

import (
	"fmt"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// DiffReporter is a simple custom reporter that only records differences
// detected during comparison.
type DiffReporter struct {
	path  cmp.Path
	diffs []string
}

func (r *DiffReporter) PushStep(ps cmp.PathStep) {
	r.path = append(r.path, ps)
}

func (r *DiffReporter) Report(rs cmp.Result) {
	if !rs.Equal() {
		vx, vy := r.path.Last().Values()

		format := "\n" + Ko("%#v") + ":\n\t%s: %+v\n\t%s: %+v"
		r.diffs = append(r.diffs, fmt.Sprintf(format, r.path, KO, vx, OK, vy))
	}
}

func (r *DiffReporter) PopStep() {
	r.path = r.path[:len(r.path)-1]
}

func (r *DiffReporter) String() string {
	return strings.Join(r.diffs, "\n")
}

// NotNil fails the test if the subject is nil.
func NotNil(t *testing.T, data any) {
	t.Helper()

	if data == nil {
		t.Errorf(Ko("\ngot %#v"), nil)
	}
}

// NoError fails the test on error.
func NoError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Errorf(Ko("\nunexpected error:\n\tERR: %#v"), err)
	}
}

// Equal fails the test if not equal.
func Equal(t *testing.T, actual, expected any) {
	t.Helper()

	r := DiffReporter{}
	if !cmp.Equal(actual, expected, cmp.Reporter(&r)) {
		t.Error(r.String())
	}
}

// NotEqual fails the test if equal.
func NotEqual(t *testing.T, actual, expected any) {
	t.Helper()

	r := DiffReporter{}
	if cmp.Equal(actual, expected, cmp.Reporter(&r)) {
		t.Error(r.String())
	}
}
