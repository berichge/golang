package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kuangchanglang/graceful"
)

type Engine struct {
	server *graceful.Server
	*gin.Engine
}

func main() {
	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func Default() *Engine {
	// r := New()

}

func New() *Engine {
	r := &Engine{
		Engine: gin.New(),
	}
	return r
}

func (engine *Engine) Use(middleware ...HandlerFunc) IRoutes {

}
