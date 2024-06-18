package helpers

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/charmbracelet/log"
	"github.com/google/go-cmp/cmp"
)

const (
	red     = "\033[31;1;1m"
	green   = "\033[32m"
	yellow  = "\033[33m;3;1m"
	blue    = "\033[34;3m"
	magenta = "\033[35m"
	cyan    = "\033[36m"
	gray    = "\033[37;1m"
	white   = "\033[97;1m"
	end     = "\033[0m"
)

const (
	OK = "✅" // no comments needed
	KO = "❌" // not OK
	OH = "😯" // WARNING
	AH = "🤨" // DEBUG or TODO
)

func concat(parts ...string) string {
	return strings.Join(parts, " ")
}

// Ko helper function for terminal output decoration in red with ❌.
func Ko(line string) string {
	return concat(red, KO, line, end)
}

// Ok helper function for terminal output decoration in green with ✅.
func Ok(line string) string {
	return concat(green, OK, line, end)
}

// Ah helper function for terminal output decoration in yellow with 😯
// Debugging and exploartion purpose.
func Ah(line string) string {
	return concat(blue, AH, line, end)
}

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
		t.Errorf(Ko("was not expecting %#v"), nil)
	}
}

// NoError fails the test on error.
func NoError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Errorf(Ko("was not expecting error...\n\tGOT: %#v"), err)
	}
}

// Equal fails the test if not equal (DeepEqual).
func Equal(t *testing.T, actual, expected any) {
	t.Helper()

	r := DiffReporter{}
	if !cmp.Equal(actual, expected, cmp.Reporter(&r)) {
		t.Error(r.String())
	}
}

// NotEqual fails the test if equal (DeepEqual).
func NotEqual(t *testing.T, actual, expected any) {
	t.Helper()

	r := DiffReporter{}
	if cmp.Equal(actual, expected, cmp.Reporter(&r)) {
		t.Error(r.String())
	}
}

// DurationLog measure the duration of a function
//
// Usage:
//
//	defer helpers.DurationLog(time.Now())
func DurationLog(start time.Time, name string) {
	log.Info("DURATION", "func", name, "duration", time.Since(start))
}
