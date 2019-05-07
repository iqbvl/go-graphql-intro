package resolvers

import (
	"context"
	adminuser "iqbvl/go-graphql-intro/business-services/admin-user"
	usecase "iqbvl/go-graphql-intro/business-services/admin-user/usecase"
	"iqbvl/go-graphql-intro/models"
)

type AdminUserResolver struct {
	m *models.AdminUser
}

type AdminUserQuery struct {
	id string
	m  models.AdminUser
}

type AdminUserInput struct {
	ID          *string
	AdminName   *string
	Username    string
	Password    string
	Status      int
	CreatedDate string
	CreatedBy   string
	UpdatedDate *string
	UpdatedBy   *string
}

func (r *Resolver) ListAdminUser(ctx context.Context, args struct {
	Order *string
	Count *int32
}) (*[]*AdminUserResolver, error) {
	adminUserRepo := ctx.Value("AdminUserRepository").(adminuser.AdminUserRepository)
	bu := usecase.NewAdminUserUsecase(adminUserRepo)
	adminUser, err := bu.Fetch(ctx, *args.Order, *args.Count)
	if err != nil {
		return nil, err
	}
	var resolvers = make([]*AdminUserResolver, 0, len(adminUser))
	for _, admUsr := range adminUser {
		bq := AdminUserQuery{id: admUsr.ID, m: *admUsr}
		resolver := AdminUser(ctx, bq)
		resolvers = append(resolvers, resolver)
	}
	return &resolvers, nil
}

func AdminUser(ctx context.Context, args AdminUserQuery) *AdminUserResolver {
	var b AdminUserResolver
	b.m = &args.m
	return &b
}


func (m *AdminUserResolver) ID() *string {
	return &m.m.ID
}

func (m *AdminUserResolver) AdminName() *string {
	return &m.m.AdminName.String
}

func (m *AdminUserResolver) Username() *string {
	return &m.m.Username
}

func (m *AdminUserResolver) Password() *string {
	return &m.m.Password
}

func (m *AdminUserResolver) Status() *int32 {
	return &m.m.Status
}

func (m *AdminUserResolver) CreatedDate() *string {
	return &m.m.CreatedDate
}

func (m *AdminUserResolver) CreatedBy() *string {
	return &m.m.CreatedBy
}

func (m *AdminUserResolver) UpdatedDate() *string {
	return &m.m.UpdatedDate.String
}

func (m *AdminUserResolver) UpdatedBy() *string {
	return &m.m.UpdatedBy.String
}