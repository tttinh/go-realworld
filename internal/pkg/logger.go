package pkg

import (
	"io"
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"
	slogformatter "github.com/samber/slog-formatter"
)

func NewLogger(mode string) *slog.Logger {
	if mode == "test" {
		return noopLogger()
	}

	if mode == "release" {
		return releaseLogger()
	}

	return debugLogger()
}

func noopLogger() *slog.Logger {
	return slog.New(
		slog.NewTextHandler(io.Discard, nil),
	)
}

func debugLogger() *slog.Logger {
	return slog.New(
		tint.NewHandler(os.Stderr, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.Kitchen,
			ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
				if err, ok := a.Value.Any().(error); ok {
					aErr := tint.Err(err)
					aErr.Key = a.Key
					return aErr
				}
				return a
			},
		}),
	)
}

func releaseLogger() *slog.Logger {
	errFmt := slogformatter.ErrorFormatter("error")
	opts := &slog.HandlerOptions{
		Level: new(slog.LevelVar),
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			// Modify time attribute
			if a.Key == slog.TimeKey {
				t := a.Value.Time()
				return slog.Attr{
					Key:   slog.TimeKey,
					Value: slog.Int64Value(t.UnixMilli()),
				}
			}
			return a
		},
	}
	return slog.New(
		slogformatter.NewFormatterHandler(errFmt)(
			slog.NewJSONHandler(os.Stderr, opts),
		),
	)
}
