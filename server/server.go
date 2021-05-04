package server

import (
	"app/graphql"
	"net/http"
	"strings"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func NewRouter() http.Handler {
	e := gin.Default()

	// handle graphql
	h := graphql.NewHandler()
	e.POST("/graphql", h)
	web := static.Serve("/graphql", static.LocalFile("web", false))
	e.GET("/graphql", func(c *gin.Context) {
		accept := c.Request.Header.Get("Accept")
		if strings.Contains(accept, "text/html") {
			web(c)
			return
		}
		h(c)
	})
	e.GET("/", func(c *gin.Context) { c.Redirect(http.StatusFound, "/graphql") })
	e.NoRoute(web)

	return e
}
