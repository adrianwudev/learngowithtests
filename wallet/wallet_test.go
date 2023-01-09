package wallet

import (
	"fmt"
	"testing"
)

type BitCoin int

type Wallet struct {
	balance BitCoin
}

func (w *Wallet) Deposit(amount BitCoin) {
	w.balance += amount
}

func (w *Wallet) Balance() BitCoin {
	return w.balance
}

type Stringer interface {
	String() string
}

func (b BitCoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(BitCoin(10))
	got := wallet.Balance()

	want := BitCoin(10)

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
