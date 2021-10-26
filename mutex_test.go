package learn_go_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutexGo(t *testing.T) {
	x := 0
	var mutex sync.Mutex

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter = ", x)
}

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock()
	account.Balance = account.Balance + amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
	account.RWMutex.RLock()
	balance := account.Balance
	account.RWMutex.RUnlock()
	return balance
}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Total Balance", account.GetBalance())
}

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (usr *UserBalance) Lock() {
	usr.Mutex.Lock()
}

func (usr *UserBalance) Unlock() {
	usr.Mutex.Unlock()
}

func (usr *UserBalance) Change(amount int) {
	usr.Balance = usr.Balance + amount
}

func Transfer(usr1 *UserBalance, usr2 *UserBalance, amount int) {
	usr1.Lock()
	fmt.Println("Lock usr1", usr1.Name)
	usr1.Change(-amount)

	time.Sleep(1 * time.Second)

	usr2.Lock()
	fmt.Println("Lock usr2", usr2.Name)
	usr2.Change(amount)

	time.Sleep(1 * time.Second)

	usr1.Unlock()
	usr2.Unlock()
}

func TestDeadLock(t *testing.T) {
	user1 := UserBalance{
		Name:    "Raisa Supriatna",
		Balance: 2000,
	}
	user2 := UserBalance{
		Name:    "Gozenx",
		Balance: 2000,
	}

	go Transfer(&user1, &user2, 500)
	go Transfer(&user2, &user1, 500)

	time.Sleep(10 * time.Second)
	fmt.Println("User ", user1.Name, ", Balance ", user1.Balance)
	fmt.Println("User ", user2.Name, ", Balance ", user2.Balance)

}
