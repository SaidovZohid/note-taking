package postgres_test

import (
	"testing"

	"github.com/SaidovZohid/note-taking/storage/repo"
	"github.com/bxcodec/faker/v4"
	"github.com/stretchr/testify/require"
)

func createUser(t *testing.T) *repo.User {
	User, err := dbManager.User().Create(&repo.User{
		FirstName: faker.FirstName(),
		LastName:  faker.LastName(),
		Email:     faker.Email(),
		Password:  faker.Password(),
		Username:  faker.Username(),
	})
	require.NoError(t, err)
	return User
}

func deleteUser(t *testing.T, user_id int64) {
	err := dbManager.User().Delete(user_id)
	require.NoError(t, err)
}

func TestCreateUser(t *testing.T) {
	user := createUser(t)
	require.NotEmpty(t, user)
	deleteUser(t, user.ID)
}

func TestGetUser(t *testing.T) {
	user := createUser(t)
	require.NotEmpty(t, user)
	user2, err := dbManager.User().Get(user.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	deleteUser(t, user.ID)
}

func TestUpdateUser(t *testing.T) {
	user := createUser(t)
	user2, err := dbManager.User().Update(&repo.User{
		ID:        user.ID,
		FirstName: faker.FirstName(),
		LastName:  faker.LastName(),
		Email:     faker.Email(),
		Password:  faker.Password(),
	})
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.NotEmpty(t, user2)
	deleteUser(t, user.ID)
}

func TestDeleteUser(t *testing.T) {
	user := createUser(t)
	deleteUser(t, user.ID)
}

func TestGetAllUsers(t *testing.T) {
	user := createUser(t)
	users, err := dbManager.User().GetAll(&repo.GetAllUsersParams{
		Limit:  10,
		Page:   1,
		SortBy: "ASC",
	})
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(users.Users), 1)
	require.NotEmpty(t, user)
	deleteUser(t, user.ID)
}
