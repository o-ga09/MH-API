package config

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	Env                      string `env:"ENV" envDefault:"dev"`
	Port                     string `env:"PORT" envDefault:"80"`
	Database_url             string `env:"DATABASE_URL" envDefult:""`
	ProjectID                string `env:"PROJECTID" envDefault:""`
	OtelExporterOtlpEndpoint string `env:"OTEL_EXPORTER_OTLP_ENDPOINT" envDefault:"localhost:4317"`
	// OtelInsecure: trueにすると TLS なしで OTLP エンドポイントへ接続する (ローカル・内部NW向け)
	// 本番環境では false (デフォルト) にして TLS 通信を使うこと
	OtelInsecure             bool   `env:"OTEL_EXPORTER_OTLP_INSECURE" envDefault:"true"`
	GeminiAPIKey             string `env:"GEMINI_API_KEY" envDefault:""`
	GeminiModel              string `env:"GEMINI_MODEL" envDefault:"gemini-2.0-flash-exp"`
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
