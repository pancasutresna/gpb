package main

import "fmt"

func (cli *CLI) createWallet() {
	wallets, _ := NewWallets()
	address := wallets.CreateWallet()
	wallets.SaveToFile()

	fmt.Println("Your new address : %s\n", address)
}
