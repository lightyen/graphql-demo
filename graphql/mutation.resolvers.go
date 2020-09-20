package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"app/graphql/generated"
	"app/graphql/model"
	"app/internal/auth"
	"context"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func (r *mutationResolver) Login(ctx context.Context, input model.UserLoginInput) (*string, error) {
	if input.Username != "guest" && input.Password == nil {
		return nil, ErrAuthentication
	}

	if *input.Password != "helloworld" {
		return nil, ErrAuthentication
	}
	c := ctx.Value(GinContextKey{}).(*gin.Context)

	role := model.RoleEnumTypeNormal
	if input.Username == "lightyen" {
		role = model.RoleEnumTypeAdministrator
	}
	now := time.Now()
	tokenString, err := auth.SignJwt(&CustomClaims{
		Role: role,
		StandardClaims: jwt.StandardClaims{
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

func (r *mutationResolver) Operations(ctx context.Context) (*model.Operations, error) {
	return &model.Operations{}, nil
}

func (r *operationsResolver) Show(ctx context.Context, obj *model.Operations, input uint) (*uint, error) {
	return &input, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Operations returns generated.OperationsResolver implementation.
func (r *Resolver) Operations() generated.OperationsResolver { return &operationsResolver{r} }

type mutationResolver struct{ *Resolver }
type operationsResolver struct{ *Resolver }
