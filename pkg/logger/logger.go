package logger

import (
	"log"
	"sync"

	"go.uber.org/zap"
)

var (
	// onceInit guarantee only once logger initialization
	onceInit sync.Once
)

// Init initializes logger
func Init(development bool) {
	onceInit.Do(func() {
		var logger *zap.Logger
		var err error
		if development {
			logger, err = zap.NewDevelopment()
		} else {
			logger, err = zap.NewProduction()
		}
		if err != nil {
			log.Fatal("failed to initialize zap logger:", err)
		}
		zap.RedirectStdLog(logger)
		zap.ReplaceGlobals(logger)
	})
}
