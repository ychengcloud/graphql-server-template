package middleware

import (
	"context"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"{{ .Extra.pkgpath }}/internal/models"
)

// A private key for context that only this package can access. This is important
// to prevent collisions between different context uses
var UserCtxKey = "user"

// Options is  configuration of database
type Options struct {
	AuthType string `yaml:"authType"`
	Skip     bool   `yaml:"skip"`
}

// A stand-in for our database backed user object
type User struct {
	Name string
}

type Auth interface {
	Authenticate(c *gin.Context) (*models.CurrentUser, error)
}

// Middleware decodes the share session cookie and packs the session into context
func AuthMiddleware(auth Auth, opts *models.ServiceOptions) gin.HandlerFunc {
	return func(c *gin.Context) {

		user, err := auth.Authenticate(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}
		c.Set(UserCtxKey, user)
		c.Next()
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *User {
	raw, _ := ctx.Value(UserCtxKey).(*User)
	return raw
}
