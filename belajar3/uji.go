package main

import(
"fmt"
"strconv"
)

func main() {
	nama := "6070"
	nama2, _ := strconv.Atoi(nama)
	umur := 6000
	total := nama2 + umur
	fmt.Println("nama : ", nama, "--> umur : ", umur)
	fmt.Println("total : ", total)
	}