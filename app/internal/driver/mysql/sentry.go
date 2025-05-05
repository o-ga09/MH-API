package mysql

import (
	"context"
	"fmt"
	"log"
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
		log.Printf("[GORM INFO] %s\n", fmt.Sprintf(msg, data...))
	}
}

func (l *SentryLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.logLevel >= logger.Warn {
		log.Printf("[GORM WARN] %s\n", fmt.Sprintf(msg, data...))
	}
}

func (l *SentryLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.logLevel >= logger.Error {
		log.Printf("[GORM ERROR] %s\n", fmt.Sprintf(msg, data...))
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
	if err != nil {
		span.Description = err.Error()
	} else if l.slowThreshold != 0 && elapsed > l.slowThreshold {
		span.Description = "slow query"
	}
	span.Finish()

	if l.logLevel >= logger.Info {
		log.Printf("[GORM TRACE] [%.3fms] [rows:%v] %s\n", float64(elapsed.Nanoseconds())/1e6, rows, sql)
	}
}
