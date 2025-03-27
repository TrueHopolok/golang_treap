package main

import (
	"fmt"
	"main/treap"
	"math/rand/v2"
	"time"
)

const tests_amount = 100_000_000

func main() {
	var timestamp time.Time
	values := make([]int, tests_amount)
	indexes := make([]int, tests_amount)
	for i := 0; i < tests_amount; i++ {
		indexes = append(indexes, rand.IntN(i+1))
		values = append(values, rand.IntN(100))
	}

	//* SLICE TESTING

	s := make([]int, 0)
	timestamp = time.Now()

	for i := 0; i < tests_amount; i++ {
		index := indexes[i]
		value := values[i]
		if len(s)-1 >= index {
			s = append(s, value)
		} else if index <= 0 {
			s = append([]int{value}, s...)
		} else {
			s = append(s[:index], append([]int{value}, s[index:]...)...)
		}
	}

	fmt.Println(time.Since(timestamp).Seconds())

	//* TREAP TESTING

	t := treap.Create()
	timestamp = time.Now()

	for i := 0; i < tests_amount; i++ {
		t.Insert(indexes[i], values[i])
	}

	fmt.Println(time.Since(timestamp).Seconds())

	//* COMPARE RESULTS
	e := t.Export()
	for i := 0; i < tests_amount; i++ {
		if e[i] != s[i] {
			fmt.Println("Bad")
		}
	}
	fmt.Println("Good")
}
