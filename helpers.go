package helpers

import (
	"log/slog"
	"strings"
	"time"
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

// DurationLog measure the duration of a function
//
// Usage:
//
//	defer helpers.DurationLog(time.Now())
func DurationLog(start time.Time, name string) {
	slog.Info("DURATION", "func", name, "duration", time.Since(start))
}

func concat(parts ...string) string {
	return strings.Join(parts, " ")
}
