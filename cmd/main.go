package main

import (
	"PowerShop/buy"
	"PowerShop/service"
)

func main(){

	go service.Use()
	buy.Do()
}

