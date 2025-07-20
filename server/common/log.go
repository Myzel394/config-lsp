package common

import (
	"context"
	"log/slog"

	sentryslog "github.com/getsentry/sentry-go/slog"
)

var Log *slog.Logger

func SetupLogging() {
	if ServerOptions.AllowedUsageReports == UsageReportAll {
		ctx := context.Background()
		handler := sentryslog.Option{
			EventLevel: []slog.Level{slog.LevelError},
			LogLevel:   []slog.Level{slog.LevelWarn, slog.LevelInfo},
		}.NewSentryHandler(ctx)

		Log = slog.New(handler)
	} else if ServerOptions.AllowedUsageReports == UsageReportErrorOnly {
		ctx := context.Background()
		handler := sentryslog.Option{
			EventLevel: []slog.Level{slog.LevelError},
			LogLevel:   []slog.Level{},
		}.NewSentryHandler(ctx)

		Log = slog.New(handler)
	} else {
		Log = slog.Default()
	}

	Log = Log.With("release", Version)
}
