package task

import (
	"testing"

	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

func TestBanking(t *testing.T){
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "Banking")
}

var _ = ginkgo.Describe("Banking", func() {
	ginkgo.Describe("Creating Account", func() {
		ginkgo.It("should have match all the account details", func() {
			account1, err := NewAccount(44, "savings", 5000.54)
			gomega.Expect(account1.AccountID).To(gomega.Equal(44))
			gomega.Expect(account1.AccountType).To(gomega.Equal("savings"))
			gomega.Expect(account1.Balance).To(gomega.Equal(5000.54))
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
		})
		ginkgo.It("should have a balance set to zero", func() {
			account1, err := NewAccount(44, "CURRENT", 0.0)
			gomega.Expect(account1.Balance).To(gomega.Equal(0.0))
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
		})
		ginkgo.It("should fail to create the account and should return an error message indicating that the account type is required ", func() {
			//given accountype is empty
			acc , err := NewAccount(44, "", 4500.45)
			gomega.Expect(err).To(gomega.HaveOccurred())
			gomega.Expect(acc).To(gomega.BeNil())
		})
		ginkgo.It("should fail to create the account and should display errormessage indicating account IDmust be positive", func() {
			//giving a negative accountid
			account1, err := NewAccount(-67890, "SAVINGS", 789000)
			gomega.Expect(account1).To(gomega.BeNil())
			gomega.Expect(err).To(gomega.HaveOccurred())
		})
	})
	ginkgo.Describe("Deposit", func() {
		ginkgo.It("should deposit successfully when we give an positive amount", func() {
			//giving a positive amount
			account1, err:=NewAccount(123,"savings",1000)
			dep1 := account1.Deposit(10)
			gomega.Expect(dep1).To(gomega.Equal("success"))
			gomega.Expect(account1.Balance).To(gomega.Equal(1010.0))
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
		})
		ginkgo.It("should deposit successfully when we give an positive decimal amount", func() {
			//giving a positive decimal amount
			account1, err:=NewAccount(123,"savings",1000.5)
			dep1 := account1.Deposit(10.5)
			gomega.Expect(dep1).To(gomega.Equal("success"))
			gomega.Expect(account1.Balance).To(gomega.Equal(1011.0))
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

		})
		ginkgo.It("should fail to deposit when we give an negative amount", func() {
			//giving a negative amount
			account1, err:=NewAccount(123,"savings",1000.10)
			dep1 := account1.Deposit(-1098.5)
			gomega.Expect(dep1).To(gomega.Equal("failure"))
			gomega.Expect(account1.Balance).To(gomega.Equal(1000.10))
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
		})
		ginkgo.It("should not deposit when we give an zero amount", func() {
			//giving a zero amount
			account1,err:=NewAccount(123,"savings",1000.10)
			dep1 := account1.Deposit(0.0)
			gomega.Expect(dep1).To(gomega.Equal("failure"))
			gomega.Expect(account1.Balance).To(gomega.Equal(1000.10))
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
		})
	})
	ginkgo.Describe("Withdrawal", func() {
		ginkgo.It("should withdraw successfully, when withdrawal amount is less than balance amount", func() {
			//withdrawing an amount which is less than balance amount
			wDraw , err := NewAccount(123, "savings", 4000)
			wDraw1 := wDraw.Withdrawal(200)
			gomega.Expect(wDraw1).To(gomega.Equal("success"))
			gomega.Expect(wDraw.Balance).To(gomega.Equal(3800.0))
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
		})
		ginkgo.It("should withdraw successfully when withdrawal amount is equal to balance amount", func() {
			//withdrawing an amount which is equal to the balance amount
			wDraw , err:= NewAccount(123, "savings", 2000)
			wDraw1 := wDraw.Withdrawal(2000)
			gomega.Expect(wDraw1).To(gomega.Equal("success"))
			gomega.Expect(wDraw.Balance).To(gomega.Equal(0.0))
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
		})
		ginkgo.It("should not withdrawamount the amount, when withdrawal amount is greater than balance amount ", func() {
			//withdrawing the amount greater than the balance
			wDraw , err:= NewAccount(123, "savings", 2000)
			wDraw1 := wDraw.Withdrawal(20000)
			gomega.Expect(wDraw1).To(gomega.Equal("failure"))
			gomega.Expect(wDraw.Balance).To(gomega.Equal(2000.0))
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

		})
		ginkgo.It("should not withdraw amount, when we are trying to withdraw an negative amount ", func() {
			//withdrawing the negative amount from the balance
			wDraw , err := NewAccount(123, "savings", 2000)
			wDraw1 := wDraw.Withdrawal(-200)
			gomega.Expect(wDraw1).To(gomega.Equal("failure"))
			gomega.Expect(wDraw.Balance).To(gomega.Equal(2000.0))
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
		})
	})
	ginkgo.Describe("Transfer", func() {
		ginkgo.It("should transfer successfully a valid amount to another account", func() {
			//transfering a valid amount to another valid account
			transfer , err := NewAccount(123, "savings", 2000)
			account1 := &Account{123,"1230",123.5 }
			transfer1 := transfer.Transfer(20, account1 )
			gomega.Expect(transfer1).To(gomega.Equal("success"))
			gomega.Expect(transfer.Balance).To(gomega.Equal(1980.0))
			gomega.Expect(account1.Balance).To(gomega.Equal(143.5))
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

		})
		ginkgo.It("should not transfer amount, when transfering amount is greater than balance amount ", func() {
			//transfering a amount to another account greater than balance amount
			transfer, err := NewAccount(123, "savings", 2000)
			account1 := &Account{123,"1230",123.5 }
			transfer1 := transfer.Transfer(2000000, account1)
			gomega.Expect(transfer1).To(gomega.Equal("failure"))
			gomega.Expect(transfer.Balance).To(gomega.Equal(2000.0))
			gomega.Expect(account1.Balance).To(gomega.Equal(123.5))
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
		})
		ginkgo.It("should not transfer, when transfering an amount to invalid account ", func() {
			//transfering a valid amount to invalid account
			transfer , err:= NewAccount(123, "savings", 2000)
			account1 := &Account{-123,"1230",0 }
			transfer1 := transfer.Transfer(20, account1)
			gomega.Expect(transfer1).To(gomega.Equal("failure"))
			gomega.Expect(transfer.Balance).To(gomega.Equal(2000.0))
			gomega.Expect(account1.Balance).To(gomega.Equal(0.0))
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
		})
		ginkgo.It("should not transfer, when transfering an negative amount to valid account ", func() {
			//transfering a neagtive amount to valid account
			transfer , err:= NewAccount(123, "savings", 2000)
			account1 := &Account{123,"1230",123.5 }
			transfer1 := transfer.Transfer(-20, account1)
			gomega.Expect(transfer1).To(gomega.Equal("failure"))
			gomega.Expect(transfer.Balance).To(gomega.Equal(2000.0))
			gomega.Expect(account1.Balance).To(gomega.Equal(123.5))
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
		})
	})
})
