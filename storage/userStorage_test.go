package storage

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestFindUserByName(t *testing.T) {
	//trying to find user that don't exist
	unexistingUsername := "someusername"
	_, err := userST.FindUserByUsername(unexistingUsername)
	assert.NotNil(t, err)
	existingUsername := "user1"
	user, err := userST.FindUserByUsername(existingUsername)
	assert.Nil(t, err)
	assert.Equal(t, "user1", user.Username)
	assert.Equal(t, "3614bbf5-01ad-4a86-a9cb-cc0fbebda6fb", user.Id.String())
}

func TestEnoughCoins(t *testing.T) {
	unexistingUserId := uuid.New()
	res := userST.CheckEnoughCoins(100, unexistingUserId)
	assert.False(t, res)
	existingUserId, _ := uuid.Parse("3614bbf5-01ad-4a86-a9cb-cc0fbebda6fb")
	res = userST.CheckEnoughCoins(10000, existingUserId)
	assert.False(t, res)
	res = userST.CheckEnoughCoins(10, existingUserId)
	assert.True(t, res)
}

func TestGetUserById(t *testing.T) {
	unexistingUserId := uuid.New()
	_, err := userST.GetUserById(unexistingUserId)
	assert.NotNil(t, err)
	existingUserId, _ := uuid.Parse("3614bbf5-01ad-4a86-a9cb-cc0fbebda6fb")
	user, err := userST.GetUserById(existingUserId)
	assert.Nil(t, err)
	assert.Equal(t, existingUserId.String(), user.Id.String())
}

func TestSaveUser(t *testing.T) {
	newUser := User{
		Username:       "newusername",
		HashedPassword: "somehash",
	}
	createdId, err := userST.SaveUser(newUser)
	assert.Nil(t, err)
	foundUser, _ := userST.FindUserByUsername(newUser.Username)
	assert.Equal(t, createdId.String(), foundUser.Id.String())
	assert.Equal(t, newUser.Username, foundUser.Username)
	assert.Equal(t, newUser.HashedPassword, foundUser.HashedPassword)
	conn.Exec("DELETE FROM users WHERE id = $1", createdId.String())
}
