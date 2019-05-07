-- GoGraphQLIntroDB
IF DB_ID (N'GoGraphQLIntroDB') IS NOT NULL
DROP DATABASE GoGraphQLIntroDB;
GO
CREATE DATABASE GoGraphQLIntroDB;
GO
USE GoGraphQLIntroDB

DROP TABLE IF EXISTS AdminUser; 

-- Admin User 
CREATE TABLE AdminUser (
    ID uniqueidentifier primary key not null,
    AdminName nvarchar(300) not null,
	Username nvarchar(50) not null,
	Password nvarchar(100) not null,
	Status smallint not null,
    CreatedDate datetime not null,
	CreatedBy nvarchar(50) not null,
	UpdatedDate datetime null,
	UpdatedBy nvarchar(50) null 
);