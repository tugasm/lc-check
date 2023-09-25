package main

import "fmt"

// go cli :
// Release 2
//Pharmacy Menu
//
//Cough Syrup - $4.99
//Pain Killers - $8.99
//Antibiotics - $15.99
//Antacids - $3.99

func main() {
	//---------------------------------------------
	pharmacy := map[string]float64{
		"Cough Syrup":  4.99,
		"Pain Killers": 8.99,
		"Antibiotics":  15.99,
		"Antacids":     3.99,
	}
	cart := map[string]int{}
	var item int
	var qty int
	totalAmount := 0.0
	discount := 0.0
	//---------------------------------------------
	for {
		fmt.Println("Welcome to the Pharmacy")
		fmt.Println("Silahkan pilih obat yang ingin dibeli :")
		fmt.Println("1. Cough Syrup - $4.99")
		fmt.Println("2. Pain Killers - $8.99")
		fmt.Println("3. Antibiotics - $15.99")
		fmt.Println("4. Antacids - $3.99")
		fmt.Println("5. Selesai Pemesanan")

		fmt.Println("Masukkan pilihan obat :")
		fmt.Scanln(&item)

		if item == 5 {
			var addMore string
			fmt.Println("Apakah anda ingin menambahkan obat lainnya ? (y/n)")
			fmt.Scanln(&addMore)
			if addMore == "n" {
				break
			} else {
				continue
			}
		}

		fmt.Println("Masukkan jumlah obat :")
		fmt.Scanln(&qty)

		switch item {
		case 1:
			cart["Cough Syrup"] += qty
		case 2:
			cart["Pain Killers"] += qty
		case 3:
			cart["Antibiotics"] += qty
		case 4:
			cart["Antacids"] += qty

		}

	}

	for item, qty := range cart {
		fmt.Printf("Item : %s (%d) \n", item, qty)
	}
	for item, qty := range cart {
		price, _ := pharmacy[item]
		totalAmount += price * float64(qty)
	}
	fmt.Printf("Subtotal : $%.2f\n", totalAmount)
	if totalAmount > 10 {
		discount = totalAmount * 0.10
		totalAmount -= discount
	}

	fmt.Printf("Tax: $%.2f\n", totalAmount*0.05) // tax after discount
	fmt.Printf("Discount : $%.2f\n", discount)
	fmt.Printf("Total : $%.2f\n", totalAmount)

}
