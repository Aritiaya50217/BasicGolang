package main

import "fmt"

func main() {
	// แบบที่ 1
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}

	// แบบที่ 2
	for x := 1; x <= 10; x++ {
		fmt.Println(x)
	}
}
