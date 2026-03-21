package main

import (
	"fmt"
	"reflect"
)

func main() {

	TaskSliceArray()

	currencies := []Currency{
		USD{},
		Bitcoin{},
	}

	for _, currency := range currencies {
		typeName := reflect.TypeOf(currency).Name()
		fmt.Printf("%s: Rate = %.2f, IsCrypto = %t\n",
			typeName,
			currency.USDRate(),
			currency.IsCrypto(),
		)
	}
}
