package main

import (
	"fmt"
	"time"
)

// 고루틴은 main 이 종료되면 끝난다. 따라서 main 에 결과를 전달해주어야 한다.
// 고루틴간의 커뮤니케이션도 가능하다.
func main() {
	// go getCount("gyunam")
	// go getCount("dongyu")
	// time.Sleep(time.Second * 5)

	c := make(chan string) // make channel
	people := [4]string{"nico", "gyunam", "dongyu", "techan"}
	for _, person := range people {
		go isPerson(person, c)
	}

	// resultOne := <-c // waiting message (blocking operation)
	// resultTwo := <-c
	// fmt.Println(resultOne)
	// fmt.Println(resultTwo)

	for i := 0; i < len(people); i++ {
		fmt.Println(<-c)
	}

}

func isPerson(person string, c chan string) {
	fmt.Println(person)
	time.Sleep(time.Second * 5)
	c <- person + " is Person" // send message to channel
}

// func getCount(person string) {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(person, "is count", i)
// 		time.Sleep(time.Second)
// 	}
// }
