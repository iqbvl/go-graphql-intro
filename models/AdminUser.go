package models

import (
	"database/sql"

	_ "github.com/denisenkom/go-mssqldb"
)

type AdminUser struct {
	ID          string
	AdminName   sql.NullString
	Username    string
	Password    string
	Status      int32
	CreatedDate string
	CreatedBy   string
	UpdatedDate sql.NullString
	UpdatedBy   sql.NullString
}
