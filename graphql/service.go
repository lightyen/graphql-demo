package graphql

import (
	"app/graphql/generated"
	"app/jwt"
	"app/model"
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

var (
	GinContextKey struct{}
	RoleKey       struct{}
)

func GinContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), GinContextKey, c)
		c.Request = c.Request.WithContext(ctx)
	}
}

func GetGinContext(ctx context.Context) (*gin.Context, bool) {
	c, ok := ctx.Value(GinContextKey).(*gin.Context)
	return c, ok
}

type CustomClaims struct {
	Role generated.RoleEnumType
	jwtgo.StandardClaims
}

func Stack(skip int) []byte {
	FuncForPC := func(pc uintptr) string {
		fn := runtime.FuncForPC(pc)
		if fn == nil {
			return "???"
		}
		name := fn.Name()
		return name
	}
	buf := new(bytes.Buffer)
	for i := skip; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		fmt.Fprintf(buf, "%s:%d\n\t%s\n", file, line, FuncForPC(pc))
	}
	return buf.Bytes()
}

func newHandler() gin.HandlerFunc {
	resolver := &Resolver{}

	go func(r *Resolver) {
		tick := time.NewTicker(time.Second)
		for now := range tick.C {
			r.BroadcastTime(&model.Time{Time: now})
		}
	}(resolver)

	auth := func(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
		c, ok := GetGinContext(ctx)
		if !ok {
			return nil, ErrAuthentication
		}
		tokenString, err := c.Cookie("app_token")
		if err != nil {
			return nil, ErrAuthentication
		}
		claims := &CustomClaims{}
		if err := jwt.Verify(tokenString, claims); err != nil {
			return nil, ErrAuthentication
		}
		return next(context.WithValue(ctx, RoleKey, claims.Role))
	}
	hasRole := func(ctx context.Context, obj interface{}, next graphql.Resolver, role generated.RoleEnumType) (res interface{}, err error) {
		userRole := ctx.Value(RoleKey).(generated.RoleEnumType)
		if role == generated.RoleEnumTypeAdministrator && userRole != generated.RoleEnumTypeAdministrator {
			return nil, ErrAuthorization
		}
		return next(ctx)
	}
	c := generated.Config{
		Resolvers: resolver,
		Directives: generated.DirectiveRoot{
			Auth:    auth,
			HasRole: hasRole,
		},
	}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(c))
	srv.SetRecoverFunc(func(ctx context.Context, err interface{}) error {
		message := fmt.Sprint(err)
		begin := "\x1b[31m"
		end := "\x1b[0m"
		fmt.Fprintln(os.Stderr, begin+message)
		fmt.Fprintln(os.Stderr, string(Stack(5))+end)
		return errors.New("internal system error")
	})
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), GinContextKey, c)
		c.Request = c.Request.WithContext(ctx)
		srv.ServeHTTP(c.Writer, c.Request)
	}
}

func Service() http.Handler {
	e := gin.Default()
	gql := newHandler()
	e.POST("/graphql", gql)
	web := static.Serve("/graphql", static.LocalFile("web", false))
	e.GET("/graphql", func(c *gin.Context) {
		accept := c.Request.Header.Get("Accept")
		if strings.Contains(accept, "text/html") {
			web(c)
			return
		}
		gql(c)
	})
	e.GET("/", func(c *gin.Context) { c.Redirect(http.StatusFound, "/graphql") })
	e.NoRoute(web)
	return e
}
