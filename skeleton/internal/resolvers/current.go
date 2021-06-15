package resolvers

import (
	"context"
	"fmt"

	"{{ .Extra.pkgpath }}/internal/models"
)

func (r *queryResolver) CurrentUser(ctx context.Context) (*models.CurrentUser, error) {
	fmt.Println("CurrentUser")
	return r.CurrentService.ToSession(ctx)
}

