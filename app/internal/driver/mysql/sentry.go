package mysql

import (
	"context"
	"fmt"
	"log/slog"

	"mh-api/app/pkg"

	"time"

	"github.com/getsentry/sentry-go"
	"gorm.io/gorm/logger"
)

type SentryLogger struct {
	slowThreshold time.Duration
	logLevel      logger.LogLevel
}

func NewSentryLogger() *SentryLogger {
	return &SentryLogger{
		slowThreshold: 200 * time.Millisecond, // スロークエリの閾値
		logLevel:      logger.Info,            // ログレベル
	}
}

func (l *SentryLogger) LogMode(level logger.LogLevel) logger.Interface {
	newlogger := *l
	newlogger.logLevel = level
	return &newlogger
}

func (l *SentryLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.logLevel >= logger.Info {
		slog.Log(ctx, pkg.SeverityInfo, fmt.Sprintf(msg, data...))
	}
}

func (l *SentryLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.logLevel >= logger.Warn {
		slog.Log(ctx, pkg.SeverityWarn, fmt.Sprintf(msg, data...))
	}
}

func (l *SentryLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.logLevel >= logger.Error {
		slog.Log(ctx, pkg.SeverityError, fmt.Sprintf(msg, data...))
	}
}

func (l *SentryLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.logLevel <= 0 {
		return
	}

	elapsed := time.Since(begin)
	sql, rows := fc()

	span := sentry.StartSpan(ctx, "gorm.query")
	span.SetData("sql", sql)
	span.SetData("rows", rows)
	span.SetData("elapsed", elapsed.String())
	defer span.Finish()
	switch {
	case err != nil && l.logLevel >= logger.Error:
		if rows == -1 {
			slog.Log(ctx, pkg.SeverityError, fmt.Sprintf("[%.3fms] [rows:%s]", float64(elapsed.Nanoseconds())/1e6, "-"), "sql", sql, "error", err.Error())
		} else {
			slog.Log(ctx, pkg.SeverityError, fmt.Sprintf("[%.3fms] [rows:%d]", float64(elapsed.Nanoseconds())/1e6, rows), "sql", sql, "error", err.Error())
		}
	case elapsed > l.slowThreshold && l.slowThreshold != 0 && l.logLevel >= logger.Warn:
		if rows == -1 {
			slog.Log(ctx, pkg.SeverityWarn, fmt.Sprintf("[%.3fms] [rows:%s]", float64(elapsed.Nanoseconds())/1e6, "-"), "sql", sql)
		} else {
			slog.Log(ctx, pkg.SeverityWarn, fmt.Sprintf("[%.3fms] [rows:%d]", float64(elapsed.Nanoseconds())/1e6, rows), "sql", sql)
		}
	case l.logLevel == logger.Info:
		if rows == -1 {
			slog.Log(ctx, pkg.SeverityInfo, fmt.Sprintf("[%.3fms] [rows:%s]", float64(elapsed.Nanoseconds())/1e6, "-"), "sql", sql)
		} else {
			slog.Log(ctx, pkg.SeverityInfo, fmt.Sprintf("[%.3fms] [rows:%d]", float64(elapsed.Nanoseconds())/1e6, rows), "sql", sql)
		}
	}
}
