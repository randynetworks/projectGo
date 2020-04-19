package main

import "fmt"

func main() {
	var opt, title string
	var result float64
	fmt.Println("==================")
	fmt.Println("Calculate Triangle")
	fmt.Println("==================")
	fmt.Println("What's calculate?")
	fmt.Println("(ae) for Areas")
	fmt.Println("(ao) for Around")
	fmt.Print("your option? : ")
	fmt.Scanf("%s", &opt)
	fmt.Println("==================")

	if opt == "ae" || opt == "AE" {
		title = "Areas"
		fmt.Println("=====  Areas =====")
		var pedestal, height float64
		fmt.Print("Number of Pedestal: ")
		fmt.Scanf("%f", &pedestal)
		fmt.Print("Number of height: ")
		fmt.Scanf("%f", &height)
		result = areas(pedestal, height)

	} else if opt == "ao" || opt == "AO" {
		title = "Around"
		fmt.Println("===== Around =====")

		var a, b, c float64
		fmt.Print("Number of a: ")
		fmt.Scanf("%f", &a)
		fmt.Print("Number of b: ")
		fmt.Scanf("%f", &b)
		fmt.Print("Number of c: ")
		fmt.Scanf("%f", &c)
		result = around(a, b, c)
	}

	println("your Result of", title, "is =", result)
}

func areas(pedestal float64, height float64) float64 {
	return 0.5 * pedestal * height
}

func around(a float64, b float64, c float64) float64 {
	return a + b + c
}
