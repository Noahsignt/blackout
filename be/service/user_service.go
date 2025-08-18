package service

import (
    "context"
    "time"
	"unicode/utf8"
	"net/http"
	"strings"
	"fmt"

    "golang.org/x/crypto/bcrypt"

    "github.com/noahsignt/blackout/be/model"
    "github.com/noahsignt/blackout/be/repository"
	beErrors "github.com/noahsignt/blackout/be/errors"
    "go.mongodb.org/mongo-driver/v2/bson"
	"github.com/golang-jwt/jwt/v5"
)

type contextKey string
const UserClaimsKey contextKey = "userClaims"

type UserService struct {
    repo *repository.UserRepo
	JwtSecret string
}

func NewUserService(repo *repository.UserRepo, jwtSecret string) *UserService {
    return &UserService{repo: repo, JwtSecret: jwtSecret}
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
		return "", beErrors.ErrUserNotFound
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if(err != nil) {
		return "", beErrors.ErrPasswordsDontMatch
	}

	keyBytes := []byte(s.JwtSecret)
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

func (s *UserService) ParseAndValidateToken(tokenString string) (jwt.MapClaims, error) {
	// parse header, payload and signature and ensure signature matches token signed with this secret
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return []byte(s.JwtSecret), nil
    })

    if err != nil {
        return nil, err
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok || !token.Valid {
        return nil, fmt.Errorf("invalid token")
    }

    return claims, nil
}

func (s *UserService) AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            http.Error(w, "missing Authorization header", http.StatusUnauthorized)
            return
        }

		// token looks like Bearer: {TOKEN} -> split into 2 and take the second
        parts := strings.SplitN(authHeader, " ", 2)
        if len(parts) != 2 || parts[0] != "Bearer" {
            http.Error(w, "invalid Authorization header format", http.StatusUnauthorized)
            return
        }

		// parse the token
        claims, err := s.ParseAndValidateToken(parts[1])
        if err != nil {
            http.Error(w, "invalid token", http.StatusUnauthorized)
            return
        }

        // embed claims in context
        ctx := context.WithValue(r.Context(), UserClaimsKey, claims)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}
