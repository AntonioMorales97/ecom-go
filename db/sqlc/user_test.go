package db

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/AntonioMorales97/ecom-go/pkg/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {

	arg := CreateUserParams{
		Username:       util.RandomString(5),
		HashedPassword: util.RandomString(5),
		FullName:       fmt.Sprintf("%s %s", util.RandomString(5), util.RandomString(5)),
		Email:          fmt.Sprintf("%s_test@test.test", util.RandomString(5)),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)
	require.True(t, time.Time(user.PasswordChangedAt).IsZero())
	require.NotEmpty(t, user.CreatedAt)

	return user

}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}
