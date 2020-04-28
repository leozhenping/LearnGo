package main

import "fmt"

func TypeJudge(items...interface{})  {
	for _, v := range items {
		switch v.(type) {
		case bool:
			fmt.Printf("item %v is bool type.", v)
		case float64:
			fmt.Printf("item %v is float64 type.", v)
		case int:
			fmt.Printf("item %v is int type.", v)
		case nil:
			fmt.Printf("item %v is nil type.", v)
		default:
			fmt.Print("item type unknown.", v)


		}
	}

}

func main() {
	var num int = 100
	TypeJudge(num)
}
