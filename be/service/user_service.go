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
	"github.com/golang-jwt/jwt/v5"
)

type UserService struct {
    repo *repository.UserRepo
	jwtSecret string
}

func NewUserService(repo *repository.UserRepo, jwtSecret string) *UserService {
    return &UserService{repo: repo, jwtSecret: jwtSecret}
}

const (
	minPasswordLength = 6
	maxPasswordLength = 15
)

func validatePasswordLength(password string) error {
    length := utf8.RuneCountInString(password)

    switch {
		case length < minPasswordLength:
			return beErrors.ErrPasswordNotLongEnough
		case length > maxPasswordLength:
			return beErrors.ErrPasswordTooLong
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

func (s *UserService) LogIn(ctx context.Context, username, password string) (string, error) {
	user, err := s.repo.FindByUsername(ctx, username)
	if(err != nil) {	
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if(err != nil) {
		return "", beErrors.ErrPasswordsDontMatch
	}

	keyBytes := []byte(s.jwtSecret)
	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{ // claims are pieces of information related to user encoded in the token
			"sub": user.ID.Hex(),
			"username": username,
			"exp":      time.Now().Add(24 * time.Hour).Unix(), // lasts 24 hours
		})

	signed, err := t.SignedString(keyBytes)
	if(err != nil) {
		return "", err
	}

	return signed, nil
}