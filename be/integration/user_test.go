package integration

import (
	"context"
	"testing"

	"github.com/noahsignt/blackout/be/errors"
	"github.com/noahsignt/blackout/be/repository"
	"github.com/noahsignt/blackout/be/service"

	"github.com/stretchr/testify/require"
)

func createUserService() (context.Context, *repository.UserRepo, *service.UserService) {
	ctx := context.Background()

	db := client.Database("testdb")
	userRepo := repository.NewUserRepo(db)

	userService := service.NewUserService(userRepo, "automated-testing-secret")

	return ctx, userRepo, userService
}

func TestSignUp(t *testing.T) {
	SetupTest(t)
	ctx, userRepo, userService := createUserService()

	username := "automated_testing_testsignup"
	_, err := userService.SignUp(ctx, username, "password")
	require.NoError(t, err)

	user, err := userRepo.FindByUsername(ctx, username)
	require.NoError(t, err)
	require.NotNil(t, user)
	require.Equal(t, username, user.Username)
}

func TestBadPasswords(t *testing.T) {
	SetupTest(t)
	ctx, _, userService := createUserService()

	username := "automated_testing_testbadpasswords"
	_, err := userService.SignUp(ctx, username, "1")
	require.EqualError(t, err, errors.ErrPasswordNotLongEnough.Error())

	_, err = userService.SignUp(ctx, username, "123456789101112131415")
	require.EqualError(t, err, errors.ErrPasswordTooLong.Error())
}

func TestDuplicateUsernames(t *testing.T) {
	SetupTest(t)
	ctx, _, userService := createUserService()

	username := "automated_testing_testduplicateusernames"
	_, err := userService.SignUp(ctx, username, "password")
	require.NoError(t, err)

	_, err = userService.SignUp(ctx, username, "password")
	require.EqualError(t, err, errors.ErrDuplicateUsernameOnSignup.Error())
}

func TestLogIn(t *testing.T) {
    SetupTest(t)
    ctx, _, userService := createUserService()

    username := "automated_testing_testlogin"
    password := "securepassword"

    _, err := userService.SignUp(ctx, username, password)
    require.NoError(t, err)

    // Login with user
    token, err := userService.LogIn(ctx, username, password)
    require.NoError(t, err)
    require.NotEmpty(t, token)

    // Login with bad password
    _, err = userService.LogIn(ctx, username, "wrongpassword")
    require.EqualError(t, err, errors.ErrPasswordsDontMatch.Error())

    // Login with non-existent user
    _, err = userService.LogIn(ctx, "nonexistent_user_abcdef", "any_password")
    require.Error(t, err)
}