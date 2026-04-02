package profiler

import (
	"log"
	"mh-api/pkg/config"

	"github.com/grafana/pyroscope-go"
)

func StartPyroscope(cfg *config.Config, appName string) {
	if cfg.PyroscopeServerAddress == "" {
		log.Println("Pyroscope server address is not configured, skipping initialization")
		return
	}

	_, err := pyroscope.Start(pyroscope.Config{
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
	} else {
		log.Printf("Pyroscope profiler started for %s", appName)
	}
}
