package main

import "fmt"

func noEmpty(data []string) (out []string)  {
	out = data[:0]
	for _, str := range data {
		if str != "" {
			out = append(out, str)
		}
	}
	return

}


// 在原串上进行修改。
func noEmpty2(data []string) []string {
	i := 0
	for _, str := range data {
		if str != "" {
			data[i] = str
			i++
		}
	}
	return data[:i] // 返回data[:i] ,而不是返回i+1, 因为此时的i就是下标位置。
}
func main() {
	s1 := []string{"read","", "black","","","pink","blue"}
	//outData := noEmpty(s1)
	outData := noEmpty2(s1)
	fmt.Println(outData)
}
