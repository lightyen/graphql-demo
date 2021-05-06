package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"app/graphql/generated"
	"app/jwt"
	"context"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func (r *mutationResolver) Login(ctx context.Context, input generated.UserLoginInput) (*string, error) {
	if input.Username != "guest" && input.Password == nil {
		return nil, ErrAuthentication
	}

	if *input.Password != "helloworld" {
		return nil, ErrAuthentication
	}
	c := ctx.Value(GinContextKey{}).(*gin.Context)

	role := generated.RoleEnumTypeNormal
	if input.Username == "lightyen" {
		role = generated.RoleEnumTypeAdministrator
	}
	now := time.Now()
	tokenString, err := jwt.Sign(&CustomClaims{
		Role: role,
		StandardClaims: jwtgo.StandardClaims{
			IssuedAt:  now.Unix(),
			ExpiresAt: now.Add(600 * time.Second).Unix(),
			Issuer:    "lightyen",
			Subject:   "my_project",
		},
	})
	if err != nil {
		return nil, err
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "app_token",
		Value:    tokenString,
		MaxAge:   3600,
		Domain:   "",
		Path:     "/",
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})
	tk := "Welcome!"
	return &tk, nil
}

func (r *mutationResolver) Operations(ctx context.Context) (*generated.Operations, error) {
	return &generated.Operations{}, nil
}

func (r *mutationResolver) SingleUpload(ctx context.Context, file graphql.Upload) (string, error) {
	return file.Filename, nil
}

func (r *operationsResolver) Show(ctx context.Context, obj *generated.Operations, input int) (*int, error) {
	return &input, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Operations returns generated.OperationsResolver implementation.
func (r *Resolver) Operations() generated.OperationsResolver { return &operationsResolver{r} }

type mutationResolver struct{ *Resolver }
type operationsResolver struct{ *Resolver }
