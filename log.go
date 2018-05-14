package quickstart

import (
	"go.uber.org/zap"
	"sync"
)

var rootLogger *zap.Logger
var rootLoggerMutex sync.Mutex

// GetRootLogger retrieves the root logger, potentially constructing the logger if required
func GetRootLogger() *zap.Logger {
	rootLoggerMutex.Lock()
	defer rootLoggerMutex.Unlock()
	if rootLogger == nil {
		var err error
		rootLogger, err = zap.NewDevelopment()
		if err != nil {
			panic(err)
		}
	}
	return rootLogger
}
