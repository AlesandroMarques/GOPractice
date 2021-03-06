package main

import ("fmt"
"time"
)

func count(id int){
	for i:=0;i<10 ;i++  {
		fmt.Println(id , ":", i)

		time.Sleep(time.Millisecond * 1000)
	}

}

func main()  {
	// A go routine is a function that runs in parallel with other functions
	// We define one by using go followed by the function name

	for i := 0; i < 10; i++ {
		go count(i)
	}

	// Wait for the timer to make sure the go routine has time to
	// finish otherwise the program would end before that happens
	time.Sleep(time.Millisecond * 11000)
}