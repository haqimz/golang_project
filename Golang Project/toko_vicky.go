package main

import "fmt"

func main() {
	myFunction()
}

func myFunction() {
	var buku, pensil, penghapus int
	fmt.Println("Masukkan harga jual buku, pensil, dan penghapus: ")
	fmt.Scan(&buku, &pensil, &penghapus)

	var unt_buku, unt_pensil, unt_penghapus int
	unt_buku = buku + ((buku * 5) / 100)
	unt_pensil = pensil + ((pensil * 5) / 100)
	unt_penghapus = penghapus + ((penghapus * 5) / 100)
	fmt.Println("Jadi total harga jual untuk ketiga barang adalah:", "\nbuku: ", unt_buku, "\npensil: ", unt_pensil, "\npenghapus: ", unt_penghapus)
}
