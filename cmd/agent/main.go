package main

import (
	"log"

	"mh-api/internal/agent"
	"mh-api/pkg/config"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	if cfg.GeminiAPIKey == "" {
		log.Fatal("GEMINI_API_KEY environment variable is required")
	}

	agentCfg := &agent.Config{
		GeminiAPIKey: cfg.GeminiAPIKey,
		GeminiModel:  cfg.GeminiModel,
		Port:         cfg.Port,
	}

	server, err := agent.NewServer(agentCfg)
	if err != nil {
		log.Fatalf("Failed to create agent server: %v", err)
	}

	if err := server.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
