package bill

import "fmt"

func MyBills() {
	firstBill := newBill("Mario's bill")

	fmt.Println(firstBill.format())
}

type bill struct {
	name  string
	items map[string]float64
	tips  float64
}

func newBill(name string) bill {
	myBill := bill{
		name: name,
		items: map[string]float64{
			"Meat pie": 2.99,
			"Cake":     3.99,
		},
		tips: 0,
	}

	return myBill
}

func (b bill) format() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0

	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v)
		total += v
	}

	fs += fmt.Sprintf("%-25v ....$%0.2f", "total: ", total)

	return fs
}
