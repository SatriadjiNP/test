package main

import (
	"fmt"

	"codeid.goquizpackage/orders"
	"codeid.goquizpackage/store"
	"codeid.goquizpackage/users"
)

// Ini Masukkan ke cart.go
func ValidasiStock(qty int, Product *store.Product, User *users.Users) string {
	var exception string
	if qty > Product.Stock {
		exception = "Jumlah Produk yg anda inginkan tidak tersedia di katalog kami, silahkan mengubah jumlah qty yg anda order atau memilih produk lain"
	} else if Product.Stock == 0 {
		exception = "Stock Sudah Habis"
	} else {
		exception = "Produk Anda sudah dimasukkan keranjang"
		Product.Stock -= qty
		orders.AddToCart(Product, qty, User)

	}
	return exception
}

func main() {
	// create data category
	fmt.Println("---------------------------Category---------------------------")
	category1 := store.NewCategory("Card")
	category2 := store.NewCategory("Eletronic")
	category3 := store.NewCategory("Gaming")
	fmt.Println(category1, "\n", category2, "\n", category3)
	fmt.Println()

	fmt.Println("---------------------------Product---------------------------")
	// create data product
	product1 := store.NewProduct("Yugioh", *category1, 70, 120)
	product2 := store.NewProduct("TV", *category2, 5, 700)
	product3 := store.NewProduct("PS4", *category3, 15, 1500)
	product4 := store.NewProduct("Pokemon", *category1, 100, 50)
	product5 := store.NewProduct("Kipas Angin", *category2, 10, 120)
	product6 := store.NewProduct("PS5", *category3, 5, 1325)
	product7 := store.NewProduct("Uno", *category1, 100, 20)
	product8 := store.NewProduct("AC", *category2, 5, 2250)
	product9 := store.NewProduct("Nintendo Switch", *category3, 10, 850)
	fmt.Println(product1, "\n", product2, "\n", product3, "\n", product4, "\n", product5, "\n", product6, "\n", product7, "\n", product8, "\n", product9)
	fmt.Println()

	fmt.Println("---------------------------User---------------------------")
	// create data user
	user1 := users.NewUsers("Yugi")
	user2 := users.NewUsers("Kaiba")
	fmt.Println(user1, "\n", user2)
	fmt.Println()

	fmt.Println("---------------------------Item Product---------------------------")
	// create data itemproduct
	itemproduct1 := orders.NewItemProduct(*product1, 10, 1200, *user1)
	fmt.Println(itemproduct1)
	fmt.Println()

	//AddToCart
	qty := 10
	pilProduct := product1
	userLogin := user1
	item := ValidasiStock(qty, pilProduct, userLogin)
	fmt.Println(item)
	fmt.Println(orders.AddToCart(pilProduct, qty, userLogin))

}
