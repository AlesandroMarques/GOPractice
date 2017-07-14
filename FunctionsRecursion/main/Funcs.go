package main

import "fmt"
func main(){
	/*listNums := []float64{1,2,3,4,5}
	fmt.Println("Sum: ", addThemUp(listNums))

	num1, num2 := next2Values(5)
	fmt.Println(num1,num2)

	fmt.Println(subtractThem(1,2,3,4,5))
	num3:=3
	// nested func
	doubleNum:= func() int {
		num3 =num3*2
		return num3


	}

	fmt.Println(doubleNum())
	fmt.Println(doubleNum())

	fmt.Println("Factorial of 3 is " , factorial(3))

	// PrintOne called after main finishes
	defer printOne()
	printTwo()
	fmt.Println(safeDiv(3,0))
	fmt.Println(safeDiv(3,2))*/
 	demoPanic()


}






// demo how to call panic and handle with recover
func demoPanic(){

	defer func(){
	// if i didnt print the message will show nothing
		fmt.Println(recover())

	}()

	panic("PANIC")
}


func safeDiv(num1 , num2 int)int {
	defer func(){
		fmt.Println(recover())
	}()// catches error ie) divide by zero then continues
	solution:= num1 / num2
	return solution



}


// to demo defer
func printOne(){fmt.Println(1)}
func printTwo(){fmt.Println(2)}


func factorial(num int) int{

	if num ==0{
		return 1
	}
	return num* factorial(num-1)

}



func subtractThem(args ... int)int{
	finalVal:=0
	for _,value := range args{
		finalVal-=value
	}
return finalVal
}

func next2Values(num int)(int,int){
	return num+1, num+2
}

func addThemUp (numbers[]float64) float64{
	sum:=0.0
	for _,val := range numbers{
		sum+= val

	}

return sum
}