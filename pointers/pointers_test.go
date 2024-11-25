package pointers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		// given
		wallet := Wallet{}

		// when
		wallet.Deposit(10)

		// then
		assert.Equal(t, Bitcoin(10), wallet.Balance())
	})

	t.Run("Withdraw", func(t *testing.T) {
		// given
		wallet := Wallet{balance: 20}

		// when
		err := wallet.Withdraw(Bitcoin(10))

		// then
		assert.NoError(t, err)
		assert.Equal(t, Bitcoin(10), wallet.Balance())
	})

	t.Run("Withdraw with insufficient funds", func(t *testing.T) {
		// given
		initialBalance := Bitcoin(20)
		wallet := Wallet{balance: initialBalance}

		// when
		err := wallet.Withdraw(Bitcoin(100))

		// then
		assert.EqualError(t, err, "insufficient funds")
		assert.Equal(t, initialBalance, wallet.Balance())
	})
}

func TestBitcoin_String(t *testing.T) {
	// given
	btc := Bitcoin(10)

	// when
	result := btc.String()

	// then
	assert.Equal(t, "10 BTC", result)
}
