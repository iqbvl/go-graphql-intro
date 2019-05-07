package usecase

import (
	"time"
	"context"
	"iqbvl/go-graphql-intro/business-services/admin-user"
	"iqbvl/go-graphql-intro/models"
)

type AdminUserUsecase struct {
	AdminUserRepository adminuser.AdminUserRepository
	contextTimeout      time.Duration
}

func NewAdminUserUsecase(
	adminUserRepo adminuser.AdminUserRepository) adminuser.AdminUserUsecase {
	return &AdminUserUsecase{
		AdminUserRepository: adminUserRepo,
		contextTimeout:      time.Second * time.Duration(models.Timeout()),
	}
}

func (a *AdminUserUsecase) Fetch(c context.Context, order string, count int32) ( []*models.AdminUser, error) {

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	listAdminUser, err := a.AdminUserRepository.Fetch(ctx, order, count)
	if err != nil {
		return nil,err
	}
	return listAdminUser, nil
}
