package metrics

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
)

const agentMeterName = "mh-agent"

// AgentMetrics holds OTEL counters for Gemini token usage.
type AgentMetrics struct {
	promptTokens    metric.Int64Counter
	candidateTokens metric.Int64Counter
	thoughtsTokens  metric.Int64Counter
	requests        metric.Int64Counter
}

// NewAgentMetrics initialises the Gemini token counters.
func NewAgentMetrics() (*AgentMetrics, error) {
	meter := otel.Meter(agentMeterName)

	promptTokens, err := meter.Int64Counter(
		"gemini_prompt_tokens_total",
		metric.WithDescription("Total number of Gemini input (prompt) tokens"),
	)
	if err != nil {
		return nil, err
	}

	candidateTokens, err := meter.Int64Counter(
		"gemini_candidate_tokens_total",
		metric.WithDescription("Total number of Gemini output (candidate) tokens"),
	)
	if err != nil {
		return nil, err
	}

	thoughtsTokens, err := meter.Int64Counter(
		"gemini_thoughts_tokens_total",
		metric.WithDescription("Total number of Gemini thinking tokens"),
	)
	if err != nil {
		return nil, err
	}

	requests, err := meter.Int64Counter(
		"gemini_requests_total",
		metric.WithDescription("Total number of Gemini API requests"),
	)
	if err != nil {
		return nil, err
	}

	return &AgentMetrics{
		promptTokens:    promptTokens,
		candidateTokens: candidateTokens,
		thoughtsTokens:  thoughtsTokens,
		requests:        requests,
	}, nil
}

// RecordTokens records token counts from a Gemini response's UsageMetadata.
func (m *AgentMetrics) RecordTokens(ctx context.Context, model string, prompt, candidate, thoughts int32) {
	attr := metric.WithAttributes()
	_ = model // reserved for future label if needed

	m.promptTokens.Add(ctx, int64(prompt), attr)
	m.candidateTokens.Add(ctx, int64(candidate), attr)
	m.thoughtsTokens.Add(ctx, int64(thoughts), attr)
	m.requests.Add(ctx, 1, attr)
}
