package main

import "fmt"

func main() {
	/*

		for j:=1; j<=9; j++ {
			for i := 1; i<=9; i++ {
				fmt.Printf("%d * %d = %d\t", j,i, j*i)
			}
			fmt.Print("\n")
		}
	*/

	for j := 1; j <= 9; j++ {
		for i := 1; i <= j; i++ {
			fmt.Printf("%d * %d = %d\t", j, i, j*i)
		}
		fmt.Print("\n")
	}
}
