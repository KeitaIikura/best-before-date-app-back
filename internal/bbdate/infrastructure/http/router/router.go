package router

import (
	controller "bbdate/internal/bbdate/adapter/controller"
	"bbdate/internal/bbdate/adapter/handler/helper"
	"bbdate/internal/bbdate/adapter/handler/helper/user_helper"
	"bbdate/internal/bbdate/infrastructure/http/session"
	"bbdate/pkg/config"
	"bbdate/pkg/logging"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

var (
	ginRouter *gin.Engine
	routerG   *gin.RouterGroup
)

func init() {
	ginRouter = gin.New()
	ginRouter.Use(gin.Recovery(), gin.Logger())
}

func SetRouters(con *dig.Container) {
	commonSettings()
	setRouter(con)
}

func commonSettings() {
	// session
	store, err := redis.NewStore(10, "tcp", config.EnvStore.RedisAddress, config.EnvStore.RedisPass, []byte("secret"))
	if err != nil {
		logging.Fatal("system", fmt.Sprintf("cannot connect to redis =>%v", err))
	}

	sessionNames := []string{session.TmxSessionKey}
	ginRouter.Use(sessions.SessionsMany(sessionNames, store))

}

// ルーティング
func setRouter(con *dig.Container) {
	logging.Info("system", "setRouter start")
	corsRouter := NewCorsRouterImpl(ginRouter, config.EnvStore.CORSAcceptOrigin)
	corsRouter.Use(helper.GenerateXRequestID())
	corsRouter.Use(user_helper.AuthRequired())
	fmt.Println()

	err := con.Invoke(func(
		authController controller.IAuthController,
	) {
		corsRouter.POST("/auth/login", authController.Login)
		corsRouter.POST("/auth/check", authController.Check)
		corsRouter.POST("/auth/logout", authController.Logout)
		corsRouter.POST("/auth", authController.UpdatePassword)
	})
	if err != nil {
		logging.Fatal("system", fmt.Sprintf("cannot Invoke =>%v", err))
	}
	fmt.Println("=====con=====")
}

func Run() {
	appServer := &http.Server{
		Handler: ginRouter,
		Addr:    fmt.Sprintf("0.0.0.0:%d", 8000),
	}
	go func() {
		logging.Debug("system", fmt.Sprintf("Listening and serving HTTP on %s", appServer.Addr))

		if err := appServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logging.Fatal("system", fmt.Sprintf("Unable to start the application http server. %v", err))
		}
	}()

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logging.Debug("system", "Shutting down server...")

	// 60秒待ってから強制的に終了する
	to := time.Duration(60)
	ctx, cancel := context.WithTimeout(context.Background(), to*time.Second)
	defer cancel()
	if err := appServer.Shutdown(ctx); err != nil {
		logging.Fatal("system", fmt.Sprintf("Server forced to shutdown: %v", err))
	}

}
