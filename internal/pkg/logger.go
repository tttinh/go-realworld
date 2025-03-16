package pkg

import (
	"io"
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"
)

func NewLogger(_ string) *slog.Logger {
	// if mode == "test" {
	// 	return noopLogger()
	// }

	// if mode == "release" {
	// 	return releaseLogger()
	// }

	return debugLogger()
	// return releaseLogger()
}

func noopLogger() *slog.Logger {
	return slog.New(
		slog.NewTextHandler(io.Discard, nil),
	)
}

func debugLogger() *slog.Logger {
	o := &tint.Options{
		Level:      slog.LevelDebug,
		TimeFormat: time.Kitchen,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			// Trace error
			if err, ok := a.Value.Any().(*Error); ok {
				a.Value = formatErrorAttr(err)
			}

			return a
		},
	}
	h := tint.NewHandler(os.Stderr, o)
	return slog.New(h)
}

func releaseLogger() *slog.Logger {
	o := &slog.HandlerOptions{
		Level: slog.LevelInfo,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			// Trace error
			if err, ok := a.Value.Any().(*Error); ok {
				a.Value = formatErrorAttr(err)
			}

			return a
		},
	}
	h := slog.NewJSONHandler(os.Stderr, o)
	return slog.New(h)
}

func formatErrorAttr(err *Error) slog.Value {
	return slog.GroupValue(
		slog.String("msg", err.Error()),
		slog.Any("trace", err.frames),
	)
}
