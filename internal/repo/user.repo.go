package repo

import (
	"context"
	"database/sql"
	"fmt"
	"food-recipes-backend/internal/queries"
)

type IUserRepository interface {
	GetUserByEmail(ctx context.Context, email string) (string, error)
	CreateUser(ctx context.Context, name string, email string, password string) (queries.User,error)
	GetUserObjByEmail(ctx context.Context, email string) *queries.User
	GetUserTokenById(ctx context.Context, id int) (queries.GetUserTokenByIdRow, error)
}

type userRepository struct {
	queries *queries.Queries
}

func NewUserRepository(db *sql.DB) IUserRepository {
	return &userRepository{
		queries: queries.New(db),
	}
}
// GetUserByEmail implements IUserRepository
func (us *userRepository) GetUserByEmail(ctx context.Context, email string) (string, error) {
	foundUser, err := us.queries.GetUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}
	return foundUser, nil
}
func (us *userRepository) CreateUser(ctx context.Context, name, email, password string) (queries.User,error) {
	fmt.Printf(`repo: %s, %s, %s`, name, email, password)
	user ,err := us.queries.CreateUser(ctx, queries.CreateUserParams{
		Name:     name,
		Email:    email,
		Password: password,
	})
	if err != nil {
		return queries.User{}, err
	}
	return user, nil
}
func (us *userRepository) GetUserObjByEmail(ctx context.Context, email string) *queries.User {
	foundUser, err := us.queries.GetUserObjectByEmail(ctx, email)
	if err != nil {
		return &queries.User{}
	}
	return &foundUser
}
func (us *userRepository) GetUserTokenById(ctx context.Context, id int) (queries.GetUserTokenByIdRow, error) {
	token, err := us.queries.GetUserTokenById(ctx, int32(id))
	if err != nil {
		return queries.GetUserTokenByIdRow{}, err
	}
	return token, nil
}