package main

import (
	"fmt"
	"log"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/Novitsh/dummy-training-app/config"
	"github.com/Novitsh/dummy-training-app/handlers"
	"github.com/Novitsh/dummy-training-app/middleware"
)

func createRenderer() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles("home.html", "templates/layout.html", "templates/home.html")
	r.AddFromFiles("about.html", "templates/layout.html", "templates/about.html")
	return r
}

func main() {
	cfg := config.Load()

	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())

	r.HTMLRender = createRenderer()
	r.Static("/static", "./static")

	r.GET("/", handlers.Home)
	r.GET("/about", handlers.About)

	addr := fmt.Sprintf(":%s", cfg.Port)
	if err := r.Run(addr); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
