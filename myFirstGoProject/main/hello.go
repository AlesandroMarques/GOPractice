package main

import "fmt"

func main(){
mapsInMaps()


}

func maps(){
	// A Map is a collection of key value pairs
	// Created with varName := make(map[keyType] valueType)
	presAge := make(map[string] int)
	presAge["teddy Rose"] = 42
	fmt.Println(presAge["teddy Rose"]) //42
	fmt.Println(len(presAge))

	presAge["John Ken"]=45
	fmt.Println(len(presAge))

	delete(presAge, "John Ken")
	fmt.Println(len(presAge))

        sshInfo := make(map[string] string)
	sshInfo["server"] = "mdmubu107.torolab.ibm.com"
	sshInfo["user"] = "ws8admin"
	sshInfo["pw"]= "go2aruba"
	fmt.Println(sshInfo["server"])

}
func mapsInMaps(){
	superhero := map[string]map[string]string{
		"Superman": map[string]string{
		"realname":"Clark Kent",
		"city":"Metro" ,
		},
		"Batman":map[string]string{
		"realname":"Bruce Wayne",
		"city": "Gothem",
		},

	}

	if temp , hero := superhero["Superman"]; hero{
		fmt.Println(temp["realname"], temp["city"])

	}




}




func arrays(){
	//ARRAYS
	var favNums[5] float64
	favNums[0]=165
	favNums[1]= 78557
	favNums[2]= 691
	favNums[3]= 3.145
	favNums[4]= 1.681

	fmt.Println(favNums[3])
	favNums2 := [5]float64 {1,2,3,4,5}

	for i ,value := range favNums2{
		fmt.Println(value , i)

	}

// SLICES
	numSlice := []int {5,4,3,2,1}
	numSlice2 := numSlice[3:5] // = [2,1]
	fmt.Println(numSlice2[0])

        fmt.Println("numSlice[:2] =", numSlice[:2])

	fmt.Println("numSlice[2:] =", numSlice[2:])


	// You can also create an empty slice and define the data type,
	// length (receive value of 0), capacity (max size)
	numSlice3 := make([]int, 5, 10)

	copy(numSlice3,numSlice)
	fmt.Println(numSlice3[0])

	// Append values to the end of a slice

	numSlice3 = append(numSlice3, 0, -1)

	fmt.Println(numSlice3[6])

}


func ifElse(){
	for j := 0; j < 5; j++ {

		fmt.Println(j);

	}

	yourAge:=18

	if yourAge>=16{
		fmt.Println("you can drivw")
	}else{
		fmt.Println("you cant drive")
	}

	if yourAge >=16 {
		fmt.Println("driving")
	}else if yourAge>=18 {
		fmt.Println("voting")
	}else {  fmt.Println("have fun ")}

	switch yourAge {
	case 16: fmt.Println("go drive")
	case 18: fmt.Println("go vote")
	default: fmt.Println("go have fun")

	}




}