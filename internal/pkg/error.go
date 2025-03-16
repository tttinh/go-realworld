package pkg

import (
	"path/filepath"
	"runtime"
)

type Error struct {
	err    error
	frames []StackFrame
}

type StackFrame struct {
	Function string `json:"func"`
	File     string `json:"file"`
	Line     int    `json:"line"`
}

func (e *Error) Error() string {
	return e.err.Error()
}

func (e *Error) Unwrap() error {
	return e.err
}

func (e *Error) Trace() []StackFrame {
	return e.frames
}

func stack() []StackFrame {
	var frames []StackFrame
	pc := make([]uintptr, 50) // Adjust max depth as needed
	n := runtime.Callers(5, pc)
	callStack := pc[:n]
	for _, pc := range callStack {
		fn := runtime.FuncForPC(pc)
		if fn == nil {
			continue
		}
		file, line := fn.FileLine(pc)
		frames = append(frames, StackFrame{
			File: filepath.Join(
				filepath.Base(filepath.Dir(file)),
				filepath.Base(file),
			),
			Function: filepath.Base(fn.Name()),
			Line:     line,
		})
	}

	return frames
}

func NewError(err error) *Error {
	return &Error{
		err:    err,
		frames: stack(),
	}

}
