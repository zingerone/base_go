package user

import "github.com/zingerone/base_go/internal/repository/db"

type UserRepository struct {
	Db db.DBInterface
}

// tableName this is for getting table name
func (s *UserRepository) tableName() string {
	return UserTableName
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}
