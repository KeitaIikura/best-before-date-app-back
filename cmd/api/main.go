package main

import (
	"bbdate/injector"
	"bbdate/internal/bbdate/infrastructure/http/router"
	"bbdate/pkg/logging"
	"fmt"
	"time"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			logging.Fatal("system", fmt.Sprintf("panic occurred in start process: %v", err))
		}
	}()
	time.Local = time.UTC

	// injection
	container := injector.RunAPI()
	// ルーティング設定
	router.SetRouters(container)
	router.Run()
}
