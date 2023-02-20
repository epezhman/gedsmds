package mds

import (
	"context"
	"github.com/IBM/gedsMDS/internal/logger"
	"github.com/IBM/gedsMDS/internal/mdsprocessor"
	"github.com/IBM/gedsMDS/internal/profiling"
	protos "github.com/IBM/gedsMDS/protos/goprotos"
	"io"
	"os"
)

type Service struct {
	mdsProcessor *mdsprocessor.Processor
}

func NewService() *Service {
	return &Service{
		mdsProcessor: mdsprocessor.InitProcessor(),
	}
}

func (t *Service) SubscribeBucket(_ context.Context, _ *protos.BucketEventSubscription) (*protos.Empty, error) {
	return &protos.Empty{}, nil
}

func (t *Service) StopAndGetProfilingResult(pr *protos.Profiling, respStream protos.MDSService_StopAndGetProfilingResultServer) error {
	reportPath := logger.LogsPath
	if pr.ProfilingType == protos.Profiling_CPU {
		profiling.StopCPUProfiling()
		reportPath += "cpu.pprof"
	}
	if pr.ProfilingType == protos.Profiling_MEMORY {
		profiling.StopMemoryProfiling()
		reportPath += "mem.pprof"
	}

	profilingReport, err := os.Open(reportPath)
	if err != nil {
		logger.ErrorLogger.Println(err)
		return err
	}
	defer func(report *os.File) {
		if err = report.Close(); err != nil {
			logger.ErrorLogger.Println(err)
		}
	}(profilingReport)
	buffer := make([]byte, 64*1024)
	for {
		bytesRead, readErr := profilingReport.Read(buffer)
		if readErr != nil {
			if readErr != io.EOF {
				logger.ErrorLogger.Println(readErr)
			}
			break
		}
		response := &protos.ProfilingResult{
			Content: buffer[:bytesRead],
		}
		readErr = respStream.Send(response)
		if readErr != nil {
			logger.ErrorLogger.Println("Error while sending chunk:", readErr)
			return readErr
		}
	}
	return nil
}
