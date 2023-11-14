package main

import "fmt"

type AbstractBanker interface {
	DoBusi()
}

type SaveBanker struct {
}

func (sb *SaveBanker) DoBusi() {
	fmt.Println("进行了 存款业务...")
}

// ++++++
type TransferBanker struct {
}

func (sb *TransferBanker) DoBusi() {
	fmt.Println("进行了 转账业务...")
}

func BankerBusiness(banker AbstractBanker) {
	banker.DoBusi()
}

func main() {
	BankerBusiness(&SaveBanker{})
	BankerBusiness(&TransferBanker{})
}
