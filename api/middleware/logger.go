package middleware

import (
	"context"
	"errors"
	"fmt"
	"os"

	"cloud.google.com/go/logging"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"

	"golang.org/x/exp/slog"
)

// cloud logging の Log level 定義
var (
	Severitydefault = slog.Level(logging.Default)
	SeverityInfo = slog.Level(logging.Info)
	SeverityWarn = slog.Level(logging.Warning)
	SeverityError = slog.Level(logging.Error)
	SeverityNotice = slog.Level(logging.Notice)
)

// traceId , spanId 追加
type traceHandler struct {
	slog.Handler
	projectID string
}

// traceHandler 実装
func (h *traceHandler) Enabled(ctx context.Context, l slog.Level) bool {
	return h.Handler.Enabled(ctx,l)
}

func (h *traceHandler) Handle(ctx context.Context, r slog.Record) error {
	if sc := trace.SpanContextFromContext(ctx); sc.IsValid() {
		trace := fmt.Sprintf("projects/%s/traces/%s",h.projectID,sc.TraceID().String())
		r.AddAttrs(slog.String("logging.googleapis.com/trace",trace),
				slog.String("logging.googleapis.com/spanId",sc.SpanID().String()))
	}

	return h.Handler.Handle(ctx,r)
}

func (h *traceHandler) WithAttr(attrs []slog.Attr) slog.Handler {
	return &traceHandler{h.Handler.WithAttrs(attrs),h.projectID}
}

func (h *traceHandler) WithGroup(g string) slog.Handler {
	return h.Handler.WithGroup(g)
}

// logger 生成関数
func New() {
	replacer := func(groups []string, a slog.Attr) slog.Attr {
		if a.Key == slog.MessageKey {
			a.Key = "message"
		}

		if a.Key == slog.LevelKey {
			a.Key = "severity"
			a.Value = slog.StringValue(logging.Severity(a.Value.Any().(slog.Level)).String())
		}

		if a.Key == slog.SourceKey {
			a.Key = "logging.googleapis.com/sourceLocation"
		}

		return a
	}

	tracer := otel.Tracer("example.com/example-service")
	// 新しいSpanを開始
	ctx, _ := tracer.Start(context.Background(), "example-operation")

	projectID := "32423553"
	se := NewStackError(errors.New("Error"))

	h := traceHandler{slog.NewJSONHandler(os.Stdout,&slog.HandlerOptions{ReplaceAttr: replacer}),projectID}
	newh := h.WithAttr([]slog.Attr{
		slog.Group("logging.googleapis.com/labels",
					slog.String("app","REST API"),
					slog.String("env","prod"),
		),
	})
	slog.SetDefault(slog.New(newh))

	slog.Log(ctx,SeverityWarn,"Hello")
	slog.Log(ctx,SeverityWarn,"Hello")
	slog.Log(ctx,SeverityError,"Hello")
	slog.Log(ctx,SeverityNotice,"Hello")
	slog.Log(ctx,SeverityError,"something wrong !","stack_trace",se.Error()+"\n\n"+string(se.stack))
}