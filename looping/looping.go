package main

import "fmt"

func main() {
	loop(100000)
}

func loop(duration int) {
	for i := 0; i < duration; i++ {
		fmt.Println(fmt.Sprintf("Loop number %d", i+1))
	}
}
