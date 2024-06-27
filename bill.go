package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(n string) bill {
	b := bill{
		name: n,
		items: map[string]float64{
			"pie":   3.14,
			"donut": 1.81,
		},
		tip: 0,
	}
	return b
}

func (b bill) format() string {
	fs := "=====Bill breakdown===== \n"
	var total float64 = 0
	for k, v := range b.items {
		total += v
		fs += fmt.Sprintf("%-25v  ....$%v\n", k+":", v)
	}

	fs += fmt.Sprintf("%-25v  ....$%v\n", "Tip:", b.tip)
	fs += fmt.Sprintf("%-25v  ....$%v\n", "Total:", total)
	return fs
}

func (b *bill) updateTip(tip float64) {
	// we don't need asterisk on struct b, go willl automatically dereference
	b.tip = tip
}
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}
