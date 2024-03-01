package main

import (
	"fmt"
	"go-assignment1/helpers"
	"os"
)

const isPrintAll bool = false

func main() {
	var inputArgs = os.Args[1:]
	fmt.Println("Assignment-1 : Biodata")
	fmt.Println("-----------------------")

	//Test add Student struct to map
	helpers.AddStudent("6", "nama test", "alamat test", "pekerjaan test", "alasan test")
	helpers.AddStudent("7", "Adyanata Komang", "Bali", "Software Engineer", "-")

	//Test delete student no 3
	helpers.DeleteStudent("3")

	if isPrintAll {
		helpers.GetDetail("", isPrintAll)
	} else if len(inputArgs) > 0 {
		helpers.GetDetail(inputArgs[0], isPrintAll)
	} else {
		fmt.Println("Please inpurt CLI Parameter!")
		fmt.Println("go run biodata.go <noAbsen>")
		fmt.Println("ex: go run biodata.go 1")
	}

}
