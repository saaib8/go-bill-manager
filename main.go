package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func takeInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	if err != nil {
		return "", err
	}
	input = strings.TrimSpace(input)
	return input, nil
}

func promptOption(b Bill) {
	r := bufio.NewReader(os.Stdin)
	opt, _ := takeInput("Choose an option:\n1. Add Item\n2. Update Tip\n3. View Bill\n4. Save Bill\n5. Exit\n", r)
	switch {
	case opt == "1":
		itemName, _ := takeInput("Enter item name: ", r)
		itemPriceStr, _ := takeInput("Enter item price: ", r)
		itemPrice, err := strconv.ParseFloat(itemPriceStr, 64)
		if err != nil {
			fmt.Println("Invalid price, please enter a valid number.")
			promptOption(b)
			return
		}
		b.addItem(itemName, itemPrice)
		promptOption(b)
	case opt == "2":
		tipStr, _ := takeInput("Enter tip amount: ", r)
		tip, err := strconv.ParseFloat(tipStr, 64)
		if err != nil {
			fmt.Println("Invalid tip amount, please enter a valid number.")
			promptOption(b)
			return
		}
		b.updateTip(tip)
		promptOption(b)
	case opt == "3":
		go func() {
			fmt.Println(b.format())
		}()
		promptOption(b)

	case opt == "4":
		b.save()
		fmt.Println("Bill saved successfully.")
		promptOption(b)
	case opt == "5":
		fmt.Println("Exiting...")
		os.Exit(0)
	default:
		fmt.Println("Invalid option, please try again.")
		promptOption(b)

	}
}
func main() {
	r := bufio.NewReader(os.Stdin)
	name, _ := takeInput("Enter your name: ", r)
	bill := newBill(name)
	fmt.Println("Welcome to the Bill Manager!")
	fmt.Println("You can add items, update the tip, and view your bill.")
	promptOption(bill)
	// fmt.Println(bill.name)
	// fmt.Println(bill.format())
	// bill.addItem("Burger", 10.00)
	// bill.addItem("Fries", 3.50)
	// bill.addItem("Soda", 1.50)
	// bill.updateTip(2.00)
	// fmt.Println(bill.format())
}
