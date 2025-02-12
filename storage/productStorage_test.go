package storage

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindProductByName(t *testing.T) {
	//test case 1 no such product
	_, err := productST.FindProductByName("unexisting")
	assert.NotNil(t, err)
	assert.True(t, errors.Is(sql.ErrNoRows, err))
	//test case there is such product
	product, err := productST.FindProductByName("wallet")
	assert.Nil(t, err)
	assert.Equal(t, "wallet", product.Name)
	assert.Equal(t, 50, product.Price)
}

func TestBuyProduct(t *testing.T) {
	//check balance gone
	//check boughts done
	//find user
	userBefore, _ := userST.FindUserByUsername("user3")
	product, _ := productST.FindProductByName("wallet")
	var transactionsBefore int
	conn.QueryRow("SELECT COUNT(*) FROM boughts WHERE user_id = $1 AND product_id = $2", userBefore.Id.String(), product.Id).Scan(&transactionsBefore)
	err := productST.Buy(userBefore.Id, product)
	var transactionsAfter int
	conn.QueryRow("SELECT COUNT(*) FROM boughts WHERE user_id = $1 AND product_id = $2", userBefore.Id.String(), product.Id).Scan(&transactionsAfter)
	userAfter, _ := userST.FindUserByUsername("user3")
	assert.Nil(t, err)
	assert.Equal(t, userBefore.Coins-product.Price, userAfter.Coins)
	assert.Equal(t, transactionsBefore+1, transactionsAfter)
}
