package main

import (
	"fmt"
	"os"
)

type Bill struct {
	name string
	item map[string]float64
	tip  float64
}

func newBill(name string) Bill {
	b := Bill{
		name: name,
		item: map[string]float64{},
		tip:  0,
	}
	return b
}



func (b Bill) format() string {
	fs := fmt.Sprintf("Bill: %v\n", b.name)
	var total float64 = 0

	// list items
	for k, v := range b.item {
		fs += fmt.Sprintf("%-25v ...$%v\n", k+":", v)
		total += v
	}
	total += b.tip
	fs += fmt.Sprintf("%-25v ...$%0.2f\n", "tip:", b.tip)

	// add total
	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total)

	return fs
}

func (b Bill) addItem(name string, price float64) {
	b.item[name] = price
}

func (b *Bill) updateTip(tip float64) {

	b.tip = tip
}

func (b*Bill) save() {
	data:=[]byte(b.format())
	fileName := b.name + ".txt"
	err := os.WriteFile(fileName, data, 0644)
	if err != nil {	
		panic(err)
	
}
	fmt.Printf("Bill saved to %v\n", fileName)
}
