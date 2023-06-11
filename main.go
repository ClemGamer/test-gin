package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/ClemGamer/test-gin/database"
	"github.com/ClemGamer/test-gin/http/controllers"
	"github.com/gin-gonic/gin"
)

func main() {

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	LoadAppConfig()

	database.Connect(AppConfig.DatabaseConfig.ConnectionString)
	database.Migrate()

	// 在停止server時關閉sql連線
	defer database.Close()

	router := gin.Default()

	router.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "hello",
		})
	})

	apiGroup := router.Group("/api")
	{
		apiGroup.GET("/users", controllers.User{}.All)
	}

	srv := &http.Server{
		Addr:    ":" + AppConfig.Port,
		Handler: router,
	}

	// 用goroutine初始化server，讓server不會卡住接下來要監聽shutdown事件的流程
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待interrupt signal
	<-ctx.Done()

	stop()
	log.Println()
	log.Println("shutting down gracefully press Ctrl+C again to force")

	// 告訴server還有5秒可以處理現有未處理完的request
	// cancel應該是在5秒後沒處理完的話，就全部取消
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
