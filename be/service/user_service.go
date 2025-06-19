package service

import (
    "context"
    "time"
	"unicode/utf8"

    "golang.org/x/crypto/bcrypt"

    "github.com/noahsignt/blackout/be/model"
    "github.com/noahsignt/blackout/be/repository"
	beErrors "github.com/noahsignt/blackout/be/errors"
    "go.mongodb.org/mongo-driver/v2/bson"
)

type UserService struct {
    repo *repository.UserRepo
}

func NewUserService(repo *repository.UserRepo) *UserService {
    return &UserService{repo: repo}
}

const (
	minPasswordLength = 6
	maxPasswordLength = 15
)

func validatePasswordLength(password string) error {
    length := utf8.RuneCountInString(password)

    switch {
		case length < minPasswordLength:
			return beErrors.PasswordNotLongEnough
		case length > maxPasswordLength:
			return beErrors.PasswordTooLong
		default:
			return nil
    }
}

func (s *UserService) SignUp(ctx context.Context, username, password string) (*model.User, error) {
	err := validatePasswordLength(password)
	if err != nil {
		return nil, err
	}

    _, err = s.repo.FindByUsername(ctx, username)
    if err == nil {
        return nil, beErrors.ErrDuplicateUsernameOnSignup
    }

    hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }

    user := &model.User{
        Username:     username,
        PasswordHash: string(hash),
        CreatedAt:    time.Now(),
    }

    return s.repo.CreateUser(ctx, user)
}

func (s *UserService) ChangePassword(ctx context.Context, userID bson.ObjectID, newPassword string) error {
	err := validatePasswordLength(newPassword)
	if err != nil {
		return err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
    if err != nil {
        return err
    }

	err = s.repo.UpdatePassword(ctx, userID, string(hash))
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) UpdateProfileImage(ctx context.Context, userID bson.ObjectID, imageURL string) error {
    return s.repo.UpdateImage(ctx, userID, imageURL)
}
