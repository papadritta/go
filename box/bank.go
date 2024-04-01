package main

import "fmt"
import "os"
import "strconv"

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() float64 {
	data, _ := os.ReadFile(accountBalanceFile)
	balanceText := string(data)
	balance, _ := strconv.ParseFloat(balanceText, 64)
	return balance
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {
	var accountBalance float64 = getBalanceFromFile()

	fmt.Println("Welcome to Go Bank!")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check ballance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		// wantsCheckBalance := choice ==1
		if choice == 1 {
			fmt.Println("Your balance is", accountBalance)
		} else if choice == 2 {
			fmt.Print("How much do you want to deposit?: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				// return
				continue
			}

			accountBalance += depositAmount //accountBalance = accountBalance + depositAmount
			fmt.Println("Balance updated! New amount: ", accountBalance)
			writeBalanceToFile(accountBalance)
		} else if choice == 3 {
			fmt.Print("How much do you want to withdrawal?: ")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				// return
				continue
			}
			if withdrawalAmount > accountBalance {
				fmt.Println("Invalid amount. You can't withdraw more than you have.")
				// return
				continue
			}

			accountBalance -= withdrawalAmount
			fmt.Println("Balance updated! New amount: ", accountBalance)
			writeBalanceToFile(accountBalance)
		} else {
			fmt.Println("Goodbye!")
			// return
			break
		}

		fmt.Println("Your choice: ", choice)
	}
	fmt.Println("Thanks for choosing our bank")
}