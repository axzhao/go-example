package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/swaggo/gin-swagger/example/basic/docs"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func main() {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

func main2() {
	// gin settting
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(gin.DebugMode)

	// router
	r.GET("/auth", handleAuth)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(handleJWT())
	{
		apiv1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}

	r.Run()
}

func main3() {
	// gin settting
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(gin.DebugMode)

	// router
	r.GET("/auth", handleAuth)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(handleJWT())
	{
		apiv1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}

	// endless
	server := endless.NewServer(":8080", r)
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}
	if err := server.ListenAndServe(); err != nil {
		log.Printf("Server err: %v", err)
	}
}

func main4() {
	// gin settting
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(gin.DebugMode)

	// router
	r.GET("/auth", handleAuth)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(handleJWT())
	{
		apiv1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}

	// shutdown
	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		log.Printf("Actual pid is %d", syscall.Getpid())
		if err := s.ListenAndServe(); err != nil {
			log.Printf("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")
}

// http://0.0.0.0:8080/auth
// http://0.0.0.0:8080/api/v1/ping?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOjEsInVzZXJOYW1lIjoiYWRtaW4iLCJleHAiOjE1OTQ2MzcxODIsImlzcyI6ImRlbW8ifQ.nGpkXJ-l3zuGu9er7OaImLL19aV-yXzBEm0W2p7QRx4

// kill -1 pid
// http://0.0.0.0:8080/api/v1/ping?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOjEsInVzZXJOYW1lIjoiYWRtaW4iLCJleHAiOjE1OTQ2MzcxODIsImlzcyI6ImRlbW8ifQ.nGpkXJ-l3zuGu9er7OaImLL19aV-yXzBEm0W2p7QRx4
