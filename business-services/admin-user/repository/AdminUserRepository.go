package repository

import (
	"context"
	"database/sql"
	"iqbvl/go-graphql-intro/business-services/admin-user"
	models "iqbvl/go-graphql-intro/models"
)

type sqlAdminUserRepository struct {
	Conn *sql.DB
}

func NewSqlAdminUserRepository(Conn *sql.DB) adminuser.AdminUserRepository {
	conn := Conn
	return &sqlAdminUserRepository{conn}
}

func (sj *sqlAdminUserRepository) Fetch(ctx context.Context, order string, count int32) ([]*models.AdminUser, error) {

	query := `
	SELECT  *  FROM dbo.AdminUser `
	rows, err := sj.Conn.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := make([]*models.AdminUser, 0)
	for rows.Next() {
		j := new(models.AdminUser)
		err := rows.Scan(
			&j.ID,
			&j.AdminName,
			&j.Username,
			&j.Password,
			&j.Status,
			&j.CreatedDate,
			&j.CreatedBy,
			&j.UpdatedDate,
			&j.UpdatedBy,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, j)
	}
	return result, nil
}
