package main

import "fmt"

func main() {
	var num int = 10
	Update(&num)
	fmt.Println(num)
}

func Update(p *int)  {
	*p = 60

}
