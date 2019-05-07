package adminuser

import (
	"context"

	"iqbvl/go-graphql-intro/models"
)

type AdminUserRepository interface {
	Fetch(ctx context.Context, order string, count int32) ([]*models.AdminUser, error)
}
