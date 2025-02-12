package storage

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestSendCoin(t *testing.T) {
	user1Id, _ := uuid.Parse("3614bbf5-01ad-4a86-a9cb-cc0fbebda6fb")
	user2Id, _ := uuid.Parse("2641b07b-ef83-4eeb-9734-71e78248cd5f")
	username1 := "user1"
	username2 := "user2"
	user1Before, _ := userST.FindUserByUsername(username1)
	user2Before, _ := userST.FindUserByUsername(username2)
	coinsToSend := 100
	err := coinST.SendCoin(coinsToSend, user1Id, user2Id)
	//Ошибка только если неправильные данные, может произойти только в случае не тех типов
	assert.Nil(t, err)
	user1After, _ := userST.FindUserByUsername(username1)
	user2After, _ := userST.FindUserByUsername(username2)
	assert.Equal(t, user1Before.Coins-coinsToSend, user1After.Coins)
	assert.Equal(t, user2Before.Coins+coinsToSend, user2After.Coins)
}
