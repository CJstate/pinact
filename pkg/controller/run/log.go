package run

import (
	"fmt"
	"io"

	"github.com/fatih/color"
)

type colorFunc func(a ...any) string

type Logger struct {
	stderr io.Writer
	red    colorFunc
	green  colorFunc
}

// NewLogger creates a new Logger with colored output.
// It initializes color functions for red and green text output
// and configures the stderr writer for log output.
//
// Parameters:
//   - stderr: writer for error output
//
// Returns a pointer to the configured Logger.
func NewLogger(stderr io.Writer) *Logger {
	return &Logger{
		red:    color.New(color.FgRed).SprintFunc(),
		green:  color.New(color.FgGreen).SprintFunc(),
		stderr: stderr,
	}
}

const (
	levelError = "error"
	levelInfo  = "info"
)

const (
	logINFO = "INFO"
)

const (
	msgActionNotPinned = "action isn't pinned"
	msgActionUpdated   = "action is updated"
)

// Output writes formatted log messages with color coding and line information.
// It displays the log level, message, file location, and optionally shows
// before/after line changes with color-coded diff format.
//
// The level prefix is always rendered, even when message is empty; callers
// that emit a diff without a message still surface whether the output is
// informational or an error.
//
// Parameters:
//   - level: log level ("error" for red, others for default)
//   - message: log message to display (empty to show the level only)
//   - line: line information including file path and line number
//   - newLine: new line content for diff display (empty for no diff)
func (l *Logger) Output(level, message string, line *Line, newLine string) {
	s := logINFO
	if level == levelError {
		s = l.red("ERROR")
	}
	header := s
	if message != "" {
		header += " " + message
	}
	if newLine == "" {
		fmt.Fprintf(l.stderr, `%s
%s:%d
%s
`, header, line.File, line.Number, line.Line)
		return
	}
	fmt.Fprintf(l.stderr, `%s
%s:%d
%s
%s
`, header, line.File, line.Number, l.red("- "+line.Line), l.green("+ "+newLine))
}
