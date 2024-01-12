package logger

import (
	"io"
	"log/slog"
	"testing"
)

func BenchmarkAtomicLogger(b *testing.B) {
	textHandler := slog.NewTextHandler(io.Discard, &slog.HandlerOptions{
		AddSource:   true,
		Level:       slog.LevelInfo,
		ReplaceAttr: ReplaceSourceAttr,
	})
	SetHandler(textHandler)

	for i := 0; i < b.N; i++ {
		Debug("msg")
	}
}

func BenchmarkNonAtomicLogger(t *testing.B) {
	textHandler := slog.NewTextHandler(io.Discard, &slog.HandlerOptions{
		AddSource:   true,
		Level:       slog.LevelInfo,
		ReplaceAttr: ReplaceSourceAttr,
	})

	_logger := slog.New(textHandler)

	for i := 0; i < t.N; i++ {
		_logger.Debug("msg")
	}
}

func BenchmarkAtomicLoggerParallel(b *testing.B) {
	textHandler := slog.NewTextHandler(io.Discard, &slog.HandlerOptions{
		AddSource:   true,
		Level:       slog.LevelInfo,
		ReplaceAttr: ReplaceSourceAttr,
	})
	SetHandler(textHandler)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Debug("msg")
		}
	})
}

func BenchmarkNonAtomicLoggerParallel(b *testing.B) {
	textHandler := slog.NewTextHandler(io.Discard, &slog.HandlerOptions{
		AddSource:   true,
		Level:       slog.LevelInfo,
		ReplaceAttr: ReplaceSourceAttr,
	})

	_logger := slog.New(textHandler)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_logger.Debug("msg")
		}
	})
}
