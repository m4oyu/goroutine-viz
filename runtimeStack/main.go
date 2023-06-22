package main

import (
	"fmt"
	"runtime"
	"time"
)

func stackExample() {
	stackSlice := make([]byte, 2048)
	s := runtime.Stack(stackSlice, true)
	fmt.Printf("\n======stack1======\n%s", stackSlice[0:s])

	go func() {

		go func() {
			s = runtime.Stack(stackSlice, true)
			fmt.Printf("\n======stack2======\n%s", stackSlice[0:s])

		}()

		go func() {
			s = runtime.Stack(stackSlice, true)
			fmt.Printf("\n======stack3======\n%s", stackSlice[0:s])

		}()

		s = runtime.Stack(stackSlice, true)
		fmt.Printf("\n======stack4======\n%s", stackSlice[0:s])

	}()

	time.Sleep(time.Second * 1)

	s = runtime.Stack(stackSlice, true)
	fmt.Printf("\n======stack5======\n%s", stackSlice[0:s])
}

// func First() {
// 	Second()
// }

// func Second() {
// 	Third()
// }

// func Third() {
// 	for c := 0; c < 5; c++ {
// 		fmt.Println(runtime.Caller(c))
// 	}
// }

func main() {
	fmt.Println("######### STACK ################")
	stackExample()
	// fmt.Println("\n\n######### CALLER ################")
	// First()
}
