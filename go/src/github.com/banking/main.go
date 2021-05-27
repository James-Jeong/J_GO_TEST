package main

import (
	"fmt"

	banking "github.com/banking/base"
)

func main() {
	account1 := banking.NewAccount("abc") // struct pointer

	banking.SetAccountType(account1, false)
	fmt.Println(banking.GetAccountType(account1))

	fmt.Println(account1.Balance())
	account1.Deposit(1000)
	fmt.Println(account1.Balance())

	withDrawResult1 := account1.Withdraw(2000)
	if withDrawResult1 != nil {
		fmt.Println(withDrawResult1)
		//log.Fatalln(withDrawResult1)
	}
	fmt.Println(account1.Balance())

	withDrawResult2 := account1.Withdraw(500)
	if withDrawResult2 != nil {
		fmt.Println(withDrawResult2)
		//log.Fatalln(withDrawResult2)
	}
	fmt.Println(account1.Balance())

	fmt.Println(account1.Balance(), account1.Owner())
	account1.ChangeOwner("xyz")
	fmt.Println(account1.Balance(), account1.Owner())

	///////////////////////////////////////////////////////
	dictionary := banking.Dictionary{"first": "First word"}
	dictionary["hello"] = "hello"
	fmt.Println(dictionary["first"])

	definition, err := dictionary.Search("second")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}

	err = dictionary.Add("first", "First word")
	if err != nil {
		fmt.Println(err)
	}

	err = dictionary.Add("second", "Second word")
	if err != nil {
		fmt.Println(err)
	}

	err = dictionary.Add("second", "Second word")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(dictionary)

	err = dictionary.Update("first", "123")
	if err != nil {
		fmt.Println(err)
	}

	err = dictionary.Update("third", "123")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(dictionary)

	err = dictionary.Delete("hello")
	if err != nil {
		fmt.Println(err)
	}

	err = dictionary.Delete("hello")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(dictionary)
	///////////////////////////////////////////////////////

}
