package agent

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"google.golang.org/genai"

	"google.golang.org/adk/agent"
	"google.golang.org/adk/agent/llmagent"
	"google.golang.org/adk/cmd/launcher"
	"google.golang.org/adk/model/gemini"
	"google.golang.org/adk/server/adkrest"
	"google.golang.org/adk/session"

	"mh-api/internal/agent/middleware"
	"mh-api/internal/database/mysql"
	"mh-api/internal/service/items"
	"mh-api/internal/service/monsters"
	"mh-api/internal/service/skills"
	"mh-api/internal/service/weapons"
)

// Server represents the ADK agent server
type Server struct {
	handler http.Handler
	port    string
}

// Config holds configuration for the agent server
type Config struct {
	GeminiAPIKey string
	GeminiModel  string
	Port         string
}

// NewServer creates a new agent server
func NewServer(cfg *Config) (*Server, error) {
	ctx := context.Background()

	// Initialize services
	monsterRepo := mysql.NewMonsterRepository()
	monsterQS := mysql.NewmonsterQueryService()
	monsterService := monsters.NewMonsterService(monsterRepo, monsterQS)

	weaponQS := mysql.NewWeaponQueryService()
	weaponService := weapons.NewWeaponService(weaponQS)

	itemRepo := mysql.NewItemQueryService()
	itemService := items.NewService(monsterQS, itemRepo)

	skillRepo := mysql.NewSkillQueryService()
	skillService := skills.NewService(skillRepo)

	// Create tools
	monHunTools := NewMonHunTools(monsterService, weaponService, itemService, skillService)
	tools, err := monHunTools.GetTools()
	if err != nil {
		return nil, fmt.Errorf("failed to create tools: %w", err)
	}

	// Initialize Gemini model
	model, err := gemini.NewModel(ctx, cfg.GeminiModel, &genai.ClientConfig{
		APIKey: cfg.GeminiAPIKey,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Gemini: %w", err)
	}

	// Create agent with tools
	a, err := llmagent.New(llmagent.Config{
		Name:        "monhun_ai_agent",
		Model:       model,
		Description: "モンスターハンターの攻略情報に特化したAIアシスタント",
		Instruction: `あなたはモンスターハンターの攻略情報に特化したAIアシスタントです。
ユーザーの質問に対して、利用可能なツールを使用してモンスター、武器、アイテム、スキルなどの情報を検索し、
わかりやすく日本語で回答してください。

情報を提供する際は以下を心がけてください：
- 正確な情報を提供する
- 必要に応じて複数のツールを組み合わせて使用する
- ユーザーにとって有用な追加情報も提供する
- 専門用語は適切に説明する
- 情報が見つからない場合は、その旨を明確に伝える`,
		Tools: tools,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create agent: %w", err)
	}

	// Configure the ADK REST API
	config := &launcher.Config{
		AgentLoader:    agent.NewSingleLoader(a),
		SessionService: session.InMemoryService(),
	}

	// Create the REST API handler
	apiHandler := adkrest.NewHandler(config)

	// Create a mux for routing
	mux := http.NewServeMux()

	// Register the API handler at the /v1/agent/ path and inject DB session via middleware
	mux.Handle("/v1/agent/", http.StripPrefix("/v1/agent", middleware.DBSessionMiddleware(apiHandler)))

	// Add a health check endpoint
	mux.HandleFunc("/v1/agent/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})

	return &Server{
		handler: mux,
		port:    cfg.Port,
	}, nil
}

// Start starts the HTTP server
func (s *Server) Start() error {
	log.Printf("Starting MonHun AI Agent server on port %s", s.port)
	log.Printf("API available at http://localhost:%s/v1/agent/", s.port)
	log.Printf("Health check at http://localhost:%s/v1/agent/health", s.port)
	return http.ListenAndServe(":"+s.port, s.handler)
}
