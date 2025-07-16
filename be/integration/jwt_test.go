package integration

import (
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
    "time"

    "github.com/go-chi/chi/v5"
    "github.com/golang-jwt/jwt/v5"
    "github.com/stretchr/testify/require"
	"github.com/noahsignt/blackout/be/service"
)

func generateTestJWT(secret, userID, username string, expiry time.Duration) (string, error) {
    claims := jwt.MapClaims{
        "sub":      userID,
        "username": username,
        "exp":      time.Now().Add(expiry).Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(secret))
}

func TestAuthMiddleware(t *testing.T) {
    secret := "test-secret"
    userService := service.UserService{JwtSecret: secret}

    r := chi.NewRouter()
    r.Use(userService.AuthMiddleware)

    // dummy handler
    r.Get("/protected", func(w http.ResponseWriter, r *http.Request) {
        claims := r.Context().Value(service.UserClaimsKey).(jwt.MapClaims)
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(claims["username"].(string)))
    })

    // attempt request on protected route
    req := httptest.NewRequest("GET", "/protected", nil)
    resp := httptest.NewRecorder()
    r.ServeHTTP(resp, req)
    require.Equal(t, http.StatusUnauthorized, resp.Code)

    // attempt request with bad JWT token
    req = httptest.NewRequest("GET", "/protected", nil)
    req.Header.Set("Authorization", "Bearer invalid.token.string")
    resp = httptest.NewRecorder()
    r.ServeHTTP(resp, req)
    require.Equal(t, http.StatusUnauthorized, resp.Code)

    // attempt request with valid token
    token, err := generateTestJWT(secret, "user123", "testuser", time.Minute)
    require.NoError(t, err)

    req = httptest.NewRequest("GET", "/protected", nil)
    req.Header.Set("Authorization", "Bearer "+token)
    resp = httptest.NewRecorder()
    r.ServeHTTP(resp, req)
    require.Equal(t, http.StatusOK, resp.Code)
    require.Equal(t, "testuser", strings.TrimSpace(resp.Body.String()))
}

