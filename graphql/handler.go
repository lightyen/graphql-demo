package graphql

import (
	"app/graphql/generated"
	"app/graphql/model"
	"app/internal/auth"
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type GinContextKey struct{}
type RoleKey struct{}

func GinContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), GinContextKey{}, c)
		c.Request = c.Request.WithContext(ctx)
	}
}

type CustomClaims struct {
	Role model.RoleEnumType
	jwt.StandardClaims
}

func NewHandler() gin.HandlerFunc {
	resolver := &Resolver{}
	auth := func(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
		c := ctx.Value(GinContextKey{}).(*gin.Context)
		tokenString, err := c.Cookie("app_token")
		if err != nil {
			return nil, ErrAuthentication
		}
		claims := &CustomClaims{}
		if err := auth.VerifyJwt(tokenString, claims); err != nil {
			return nil, ErrAuthentication
		}
		return next(context.WithValue(ctx, RoleKey{}, claims.Role))
	}
	hasRole := func(ctx context.Context, obj interface{}, next graphql.Resolver, role model.RoleEnumType) (res interface{}, err error) {
		userRole := ctx.Value(RoleKey{}).(model.RoleEnumType)
		if role == model.RoleEnumTypeAdministrator && userRole != model.RoleEnumTypeAdministrator {
			return nil, ErrForbidden
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
		return errors.New("Internal server error")
	})
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), GinContextKey{}, c)
		c.Request = c.Request.WithContext(ctx)
		srv.ServeHTTP(c.Writer, c.Request)
	}
}
