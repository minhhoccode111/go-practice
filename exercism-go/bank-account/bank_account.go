package account

import "sync"

// Define the Account type here.
type Account struct {
	balance  int64
	isClosed bool
	mu       sync.Mutex
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}
	return &Account{balance: amount}
}

func (a *Account) Balance() (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.isClosed {
		return 0, false
	}
	return a.balance, true
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if tmp := a.balance + amount; tmp < 0 || a.isClosed {
		return 0, false
	}
	a.balance += amount
	return a.balance, true
}

func (a *Account) Close() (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.isClosed {
		return 0, false
	}
	tmp := a.balance
	a.balance = 0
	a.isClosed = true
	return tmp, true
}
