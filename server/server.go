package server

import (
	"app/graphql"
	"net/http"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func NewRouter() http.Handler {
	e := gin.Default()

	// handle graphql
	h := graphql.NewHandler()
	e.POST("/graphql", h)
	e.GET("/subscriptions", h)
	e.POST("/subscriptions", h)

	// handle graphql subscription(websocket)
	// e.GET("/subscriptions", func(c *gin.Context) {
	// 	if c.Request.Header.Get("Sec-WebSocket-Protocol") != "graphql-ws" {
	// 		c.Status(http.StatusNotFound)
	// 		return
	// 	}
	// })

	// handle graphql playground
	e.GET("/", func(c *gin.Context) { c.Redirect(http.StatusFound, "/graphql") })
	e.NoRoute(static.Serve("/graphql", static.LocalFile("web", false)))

	return e
}
