package bill

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func MyBills() {
	myFirstBills := createBill()
	promptOptions(myFirstBills)
}

type bill struct {
	name  string
	items map[string]float64
	tips  float64
}

func newBill(name string) bill {
	myBill := bill{
		name:  name,
		items: map[string]float64{},
		tips:  0,
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

	fs += fmt.Sprintf("%-25v ....$%v\n", "tip: ", b.tips)

	fs += fmt.Sprintf("%-25v ....$%0.2f", "total: ", total+b.tips)

	return fs
}

func (b *bill) updateTip(tip float64) {
	b.tips = tip
}

func (b *bill) addItems(name string, price float64) {
	b.items[name] = price
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getUserInput("Create a new bill name: ", reader)

	b := newBill(name)

	fmt.Println("Created the bill - ", b.name)

	return b
}

func getUserInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	options, _ := getUserInput("Choose options (a -add item, s - save bill, t - add tip): ", reader)

	switch options {
	case "a":
		name, _ := getUserInput("Item name: ", reader)
		price, _ := getUserInput("Item price: ", reader)

		p, err := strconv.ParseFloat(price, 64)

		if err != nil {
			fmt.Println("The price must be a number")
			promptOptions(b)
		}
		b.addItems(name, p)
		fmt.Println("Item added - :", name, price)
		promptOptions(b)
	case "t":
		tip, _ := getUserInput("Enter tip amount ($): ", reader)
		t, err := strconv.ParseFloat(tip, 64)

		if err != nil {
			fmt.Println("The tip must be a number")
			promptOptions(b)
		}
		b.updateTip(t)
		fmt.Println("Tip added -", tip)
		promptOptions(b)
	case "s":
		b.saveBill()
		fmt.Println("You saved the file -", b.name)
	default:
		fmt.Println("that was not a valid options...")
		promptOptions(b)
	}
}

func (b *bill) saveBill() {
	data := []byte(b.format())

	err := os.WriteFile("bill/"+b.name+".txt", data, 0644)

	if err != nil {
		panic(err)
	}

	fmt.Println("bills was save to file")
}
