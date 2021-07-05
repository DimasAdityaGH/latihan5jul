package main

import "fmt"

func main () {
	//hello world
	fmt.Println("hello, world!")

	//string
	var nama string = "this a string"
	fmt.Println(nama)

	//variabel
	var data = "variabel dimulai dengan kata var dan diikuti nama variabelnya"
	fmt.Println(data)

	//constant

	const data2 = "constant dimulai dengan kata const dan diikuti nama constantnya"
	fmt.Println(data2)

	//type data number

	var num int = 12012120 //int adalah int32 dan 64
	fmt.Println(num)

	//conversion

	var nilai32 int32 = 112222
	var nilai8 int8 = int8(nilai32)
	fmt.Println(nilai8)

	var produk = "botol"
	var p = produk[2]
	var pProduk = string(p)
	fmt.Println(pProduk)

	//type declaration

	type str string
	type numb int8

	var namae str = "dimas"
	var code numb = 7
	fmt.Println(namae)
	fmt.Println(code)

	//boolean

	var question = "apakah kamu pelajar?"
	var answer = true
	fmt.Println(question, answer)

	//operasi matematika
	var f = 10
	f++
	fmt.Println(f)

	f += 1
	fmt.Println(f)

	var hasil = 10 + 100
	fmt.Println(hasil)

	//perbandingan

	var dimas = "laki laki"
	var venny = "perempuan"
	fmt.Println(dimas != venny)
}