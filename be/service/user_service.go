package service

import (
    "context"
    "errors"
    "time"

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

func (s *UserService) SignUp(ctx context.Context, username, password string) (*model.User, error) {
    _, err := s.repo.FindByUsername(ctx, username)
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

// You'll need to add FindByID method to repo for this method to work
func (s *UserService) ChangePassword(ctx context.Context, userID bson.ObjectID, oldPassword, newPassword string) error {
    // TODO: add FindByID to repo and use here instead of FindByUsername("")
    return errors.New("ChangePassword requires repo.FindByID - implement first")
}

func (s *UserService) UpdateProfileImage(ctx context.Context, userID bson.ObjectID, imageURL string) error {
    return s.repo.UpdateImage(ctx, userID, imageURL)
}
