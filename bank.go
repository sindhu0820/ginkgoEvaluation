package task

import "fmt"

type Account struct {
	AccountID   int
	AccountType string
	Balance     float64
}

func NewAccount(accountId int, accountType string, balance float64) (acc *Account, err error) {
	if accountId>0 && accountType != "" && balance >=0 {
	return &Account{
	AccountID: accountId,
	AccountType: accountType,
	Balance: balance,
	}, nil
	}  else {
	return nil, fmt.Errorf("invalid Details")
	}
}

func (d *Account) Deposit(amt float64) string {
	if amt > 0 {
		d.Balance += amt
		return "success"
	} else {
		return "failure"
	}

}

func (d *Account) Withdrawal(amt float64) interface{} {
	if d.Balance >= amt && amt > 0 {
		d.Balance -= amt
		return "success"
	} else {
		return "failure"
	}

}
func (d *Account) Transfer(amt float64, targetAccount * Account) string {
	if d.Balance > amt && amt > 0 && targetAccount != nil && targetAccount.AccountID>0 {
		d.Balance -= amt
		targetAccount.Balance += amt
		return "success"
	} else {
		return "failure"
	}
}
