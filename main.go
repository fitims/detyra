package main

import (
	"crypto_task/store"
	"fmt"
	"log"
)

func main() {
	// populate store with data ----------------------------------------------------
	fmt.Println("** Populate store with users **")
	fmt.Println("-------------------------------")
	for _, u := range getUserData() {
		usr, err := store.RegisterUser(u.Email, u.Name)
		if err != nil {
			log.Println("Error registering the user. Error: ", err)
			panic(err)
		}
		fmt.Printf("Register user with name: %s and email: %s. The wallet balance for this user is :%f\n", usr.Name, usr.Email, usr.Wallet.CheckBalance())
	}

	// Deposit money  --------------------------------------------------------------
	fmt.Println("\n\n** Deposit money into wallet **")
	fmt.Println("-------------------------------")
	chelsy, err := store.GetUser("Chelsy@microsoft.com")
	if err != nil {
		log.Println("Error getting the user. Error: ", err)
		panic(err)
	}

	amountToDeposit := 255.50

	fmt.Printf("The current wallet balance for  user %s is :%f\n", chelsy.Name, chelsy.Wallet.CheckBalance())
	fmt.Printf("Depositing  the amount of %f\n", amountToDeposit)
	err = chelsy.DepositMoneyIntoWallet(amountToDeposit)
	if err != nil {
		log.Println("Error depositing money into the wallet. Error: ", err)
		panic(err)
	}
	fmt.Printf("The updated wallet balance for  user %s is :%f\n", chelsy.Name, chelsy.Wallet.CheckBalance())

	// Withdraw money  --------------------------------------------------------------
	fmt.Println("\n\n** Withdraw money from wallet **")
	fmt.Println("--------------------------------")

	amountToWithdraw := 50.0

	fmt.Printf("The current wallet balance for  user %s is :%f\n", chelsy.Name, chelsy.Wallet.CheckBalance())
	fmt.Printf("Withdrawing the amount of %f\n", amountToWithdraw)
	err = chelsy.WithdrawMoneyFromWallet(amountToWithdraw)
	if err != nil {
		log.Println("Error withdrawing money from the wallet. Error: ", err)
		panic(err)
	}
	fmt.Printf("The updated wallet balance for  user %s is :%f\n", chelsy.Name, chelsy.Wallet.CheckBalance())

	// Send money to another user  --------------------------------------------------------------
	fmt.Println("\n\n** Send money to another user **")
	fmt.Println("--------------------------------")

	joey, err := store.GetUser("Joey@amazon.com")
	if err != nil {
		log.Println("Error getting the user. Error: ", err)
		panic(err)
	}

	amountToSend := 52.50

	fmt.Printf("The current wallet balance for  user %s is :%f\n", chelsy.Name, chelsy.Wallet.CheckBalance())
	fmt.Printf("The current wallet balance for  user %s is :%f\n", joey.Name, joey.Wallet.CheckBalance())
	fmt.Printf("Sending the amount of %f from %s to %s\n", amountToSend, chelsy.Name, joey.Name)
	err = chelsy.SendMoneyTo(joey, amountToSend)
	if err != nil {
		log.Println("Error sending money to another user. Error: ", err)
		panic(err)
	}
	fmt.Printf("The updated wallet balance after sending money for user %s is :%f\n", chelsy.Name, chelsy.Wallet.CheckBalance())
	fmt.Printf("The updated wallet balance after sending money for user %s is :%f\n", joey.Name, joey.Wallet.CheckBalance())
}

type userData struct {
	Email string
	Name  string
}

func getUserData() []userData {
	return []userData{
		{Email: "Martyna@gmail.com", Name: "Martyna Mcdougall"},
		{Email: "Priscilla@gmail.com", Name: "Priscilla Benson"},
		{Email: "Aiza@hotmail.com", Name: "Aiza Mccullough"},
		{Email: "Jorja@hotmail.com", Name: "Jorja Montoya"},
		{Email: "Chelsy@microsoft.com", Name: "Chelsy Maldonado"},
		{Email: "Sumaiyah@microsoft.com", Name: "Sumaiyah Burnett"},
		{Email: "Miya@crypto.com", Name: "Miya Hoover"},
		{Email: "Cruz@crypto.com", Name: "Cruz Mcdaniel"},
		{Email: "Montague@crypto.com", Name: "Montague Stanton"},
		{Email: "Joey@amazon.com", Name: "Joey Nixon"},
		{Email: "Miriam@amazon.com", Name: "Miriam Kelly"},
		{Email: "Amos@amazon.com", Name: "Amos Wilcox"},
		{Email: "Antonina@amazon.com", Name: "Antonina Wheeler"},
		{Email: "Nathan@enterprise.com", Name: "Nathan Dyer"},
		{Email: "Alanah@enterprise.com", Name: "Alanah Khan"},
		{Email: "Tommy@enterprise.com", Name: "Tommy-Lee Irving"},
		{Email: "Clark@starwars.com", Name: "Clark Mathews"},
		{Email: "Kavan@starwars.com", Name: "Kavan Harrison"},
		{Email: "Summer@starwars.com", Name: "Summer-Rose Ward"},
		{Email: "Keith@starwars.com", Name: "Keith Bean"},
	}
}
