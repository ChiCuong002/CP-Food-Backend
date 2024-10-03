package services

import (
	"context"
	"database/sql"
	"fmt"
	"food-recipes-backend/internal/queries"
	"food-recipes-backend/internal/repo"
	"food-recipes-backend/internal/vo"
	"food-recipes-backend/pkg/auth"
	"food-recipes-backend/pkg/hash"
)

type IUserService interface {
	Register(ctx context.Context, name string, email string, password string) (*vo.UserRegisterResponse, error)
	Login(ctx context.Context, email string, password string) (*vo.UserLoginResponse, error)
}

type userService struct {
	userRepo repo.IUserRepository
	keyRepo  repo.IKeyRepository
	//...
}

func NewUserService(
	userRepo repo.IUserRepository,
	keyRepo repo.IKeyRepository,
) IUserService {
	return &userService{
		userRepo: userRepo,
		keyRepo:  keyRepo,
	}
}

func (us *userService) Register(ctx context.Context, name, email, password string) (*vo.UserRegisterResponse, error) {
	fmt.Printf(`service: %s, %s, %s\n`, name, email, password)
	foundUser, _ := us.userRepo.GetUserByEmail(ctx, email)
	if foundUser != "" {
		return nil, fmt.Errorf("user already exists")
	}
	// hash password
	hashedPassword, err := hash.HashPassword(password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password")
	}
	// insert user
	user, err := us.userRepo.CreateUser(ctx, name, email, hashedPassword)
	if err != nil {
		return nil, fmt.Errorf("failed to create user")
	}
	accessToken, refreshToken, err := createTokens(int(user.ID))
	if err != nil {
		return nil, fmt.Errorf("failed to create tokens")
	}
	// insert token
	_, err = us.keyRepo.UpsertKey(ctx, queries.UpsertRefreshTokenParams{
		UserID:       user.ID,
		RefreshToken: sql.NullString{String: refreshToken, Valid: true},
	})
	if err != nil {
		fmt.Println("failed to insert token: ", err.Error())
		return nil, fmt.Errorf("failed to insert token")
	}
	result := &vo.UserRegisterResponse{
		Name:  user.Name,
		Email: user.Email,
		Tokens: vo.TokensResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}
	return result, nil
}

func (us *userService) Login(ctx context.Context, email, password string) (*vo.UserLoginResponse, error) {
	foundUser := us.userRepo.GetUserObjByEmail(ctx, email)
	if foundUser == nil {
		return nil, fmt.Errorf("email or password is incorrect")
	}
	err := hash.VerifyPassword(foundUser.Password, password)
	if err != nil {
		return nil, fmt.Errorf("email or password is incorrect")
	}
	accessToken, refreshToken, err := createTokens(int(foundUser.ID))
	if err != nil {
		return nil, fmt.Errorf("failed to create tokens")
	}
	// update refresh token
	err = us.upsertRefreshToken(ctx, int(foundUser.ID), refreshToken)
	if err != nil {
		return nil, fmt.Errorf("failed to insert token")
	}
	result := &vo.UserLoginResponse{
		ID:   int(foundUser.ID),
		Name: foundUser.Name,
		Tokens: vo.TokensResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}
	return result, nil
}

// generate tokens
func createTokens(payload int) (string, string, error) {
	accessToken, refreshToken, err := auth.CreateTokenPair(int(payload))
	if err != nil {
		return "", "", fmt.Errorf("failed to create tokens")
	}
	return accessToken, refreshToken, nil
}

// upsert refresh token
func (us *userService) upsertRefreshToken(ctx context.Context, userID int, refreshToken string) error {
	_, err := us.keyRepo.UpsertKey(ctx, queries.UpsertRefreshTokenParams{
		UserID:       int32(userID),
		RefreshToken: sql.NullString{String: refreshToken, Valid: true},
	})
	if err != nil {
		return fmt.Errorf("failed to insert token")
	}
	return nil
}
