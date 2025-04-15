package facade

import "fmt"

// What
//	- is a structural design pattern that provides a simplified interface to a library,
//	- a framework, or any other complex set of classes.

// Why
// - to provide a simple interface to a complex system
// - to hide the complexity of the system from the client
// - to decouple the client from the system
// - to make the system easier to use and understand

// How
// - by creating a facade class that contains the complex system and provides a simple interface to it

// Entity Logic Layer
type Product struct {
	Name  string
	Price float32
}

type Inventory struct {
	products []Product
}

type Account struct {
	Name    string
	balance float32
}

type AccountStorage struct {
	accounts []Account
}

func (iv Inventory) lookup(name string) (*Product, error) {
	for _, product := range iv.products {
		if product.Name == name {
			return &product, nil
		}
	}
	return nil, fmt.Errorf("product not found")
}

func (acc *Account) Deposit(amount float32) {
	acc.balance += amount
}

func (acc *Account) Withdraw(amount float32) {
	acc.balance -= amount
}

func (acc *Account) GetBalance() float32 {
	return acc.balance
}

func (as AccountStorage) Lookup(name string) (*Account, error) {
	for _, account := range as.accounts {
		if account.Name == name {
			return &account, nil
		}
	}
	return nil, fmt.Errorf("account not found")
}

// thay vì để client tạo từng obj và call từng method thì tạo 1 facade class
// để gom hết lại thành 1 service và client chỉ cần gọi 1 method là xong

// Business Logic Layer
type FacadeSevice struct {
	inventory      Inventory
	accountStorage AccountStorage
}

func (s *FacadeSevice) BuyProduct(accountName string, productName string, amount float32) error {
	// lookup account and product
	account, err := s.accountStorage.Lookup(accountName)
	if err != nil {
		return err
	}
	product, err := s.inventory.lookup(productName)
	if err != nil {
		return err
	}

	// check balance and withdraw money
	if account.GetBalance() < product.Price*amount {
		return fmt.Errorf("not enough money")
	}
	account.Withdraw(product.Price * amount)

	fmt.Printf("Bought %f %s for %s\n", amount, product.Name, account.Name)
	return nil
}

func (s *FacadeSevice) Deposit(accountName string, amount float32) error {
	account, err := s.accountStorage.Lookup(accountName)
	if err != nil {
		return err
	}
	account.Deposit(amount)
	fmt.Printf("Deposited %f to %s\n", amount, account.Name)
	return nil
}

func (s *FacadeSevice) FetchBalance(accountName string, amount float32) error {
	account, err := s.accountStorage.Lookup(accountName)
	if err != nil {
		return err
	}
	fmt.Printf("Balance of %s is %f\n", account.Name, account.GetBalance())
	return nil
}

func NewFacadeService() FacadeSevice {
	return FacadeSevice{
		inventory: Inventory{
			products: []Product{
				{"apple", 2.5},
				{"banana", 1.5},
				{"orange", 3.0},
			},
		},
		accountStorage: AccountStorage{
			accounts: []Account{
				{"John", 1000},
				{"Jane", 500},
			},
		},
	}
}

func Caller() {
	facade := NewFacadeService()
	err := facade.BuyProduct("John", "apple", 2)
	if err != nil {
		fmt.Println(err)
	}
	err = facade.Deposit("Jane", 100)
	if err != nil {
		fmt.Println(err)
	}
	err = facade.FetchBalance("John", 0)
	if err != nil {
		fmt.Println(err)
	}
}
