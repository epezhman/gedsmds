package profiling

import (
	"github.com/IBM/gedsmds/internal/logger"
	"github.com/pkg/profile"
	"net/http"
	//_ "net/http/pprof"
)

type profiler interface {
	Stop()
}

var (
	cpuProfiler          profiler
	memoryProfiler       profiler
	IsBandwidthProfiling bool
)

// For Profiling in the experiment's config file, set either profiling: enable_server or profiling: cpu or profiling: memory or profiling: bandwidth

func StartCPUProfiling() {
	logger.WarningLogger.Println("Running CPU Profiling!")
	cpuProfiler = profile.Start(profile.CPUProfile, profile.ProfilePath(logger.LogsPath))
}

func StopCPUProfiling() {
	if cpuProfiler != nil {
		cpuProfiler.Stop()
	}
}

func StartMemoryProfiling() {
	logger.WarningLogger.Println("Running Memory Profiling!")
	memoryProfiler = profile.Start(profile.MemProfile, profile.ProfilePath(logger.LogsPath))
}

func StopMemoryProfiling() {
	if memoryProfiler != nil {
		memoryProfiler.Stop()
	}
}

func StartBandWithProfiling() {
	logger.WarningLogger.Println("Running Bandwidth Profiling!")
	IsBandwidthProfiling = true
}

func StartProfilingServer(profilingPort string) {
	logger.WarningLogger.Println("MDS Profiling is Running!")
	go func(port string) {
		if err := http.ListenAndServe(port, nil); err != nil {
			logger.ErrorLogger.Println(err)
		}
	}(profilingPort)
}
