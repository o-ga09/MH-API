package profiler

import (
	"log"
	"runtime"

	"mh-api/pkg/config"

	"github.com/grafana/pyroscope-go"
)

// StartPyroscope はPyroscopeプロファイラを起動し、停止用の関数を返す。
// PYROSCOPE_SERVER_ADDRESS が未設定の場合は何もしない。
// 呼び出し元は返却されたstop関数を defer で呼び出すこと。
func StartPyroscope(cfg *config.Config, appName string) func() {
	if cfg.PyroscopeServerAddress == "" {
		log.Println("Pyroscope server address is not configured, skipping initialization")
		return func() {}
	}

	// MutexProfileおよびBlockProfileはGoランタイムのサンプリングが
	// デフォルトで無効(0)のため、明示的に有効化する必要がある。
	runtime.SetMutexProfileFraction(5)
	runtime.SetBlockProfileRate(5)

	p, err := pyroscope.Start(pyroscope.Config{
		ApplicationName: appName,
		ServerAddress:   cfg.PyroscopeServerAddress,
		AuthToken:       cfg.PyroscopeAPIKey,
		Logger:          pyroscope.StandardLogger,
		ProfileTypes: []pyroscope.ProfileType{
			pyroscope.ProfileCPU,
			pyroscope.ProfileAllocObjects,
			pyroscope.ProfileAllocSpace,
			pyroscope.ProfileInuseObjects,
			pyroscope.ProfileInuseSpace,
			pyroscope.ProfileGoroutines,
			pyroscope.ProfileMutexCount,
			pyroscope.ProfileMutexDuration,
			pyroscope.ProfileBlockCount,
			pyroscope.ProfileBlockDuration,
		},
		Tags: map[string]string{
			"env": cfg.Env,
		},
	})

	if err != nil {
		log.Printf("Failed to start Pyroscope profiler: %v", err)
		return func() {}
	}

	log.Printf("Pyroscope profiler started for %s", appName)
	return func() {
		if err := p.Stop(); err != nil {
			log.Printf("Failed to stop Pyroscope profiler: %v", err)
		}
	}
}
