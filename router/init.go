package router

import (
	"biu-x.org/TikTok/module/config"
	"biu-x.org/TikTok/module/log"
	v1 "biu-x.org/TikTok/router/api/v1"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

func Init() {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	var mode string
	if config.Get("server.mode") != "prod" {
		mode = gin.DebugMode
	} else {
		mode = gin.ReleaseMode
	}
	log.Logger.Debugf("gin mode: %v", mode)
	gin.SetMode(mode)

	e := v1.NewAPI()
	log.Logger.Debugf("server port: %v", config.GetString("server.port"))
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", config.GetString("server.port")),
		Handler: e,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Logger.Fatalf("listen: %s", err)
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Logger.Info("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Logger.Fatal("Server forced to shutdown: ", err)
	}
}
