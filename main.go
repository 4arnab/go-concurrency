package main

import (
	"fmt"
	"os"
)

func main() {
	// storeData("another data stored ", "data.txt")
	go storeDataTwo(5000)
	go storeDataTwo(5000)

	great()
}

func great() {
	fmt.Println("hello world!")
}
func storeData(data, fileName string) {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		panic(err)
	}

	defer file.Close()
	_, err = file.WriteString(data)

	if err != nil {
		panic(err)
	}
}

func storeDataTwo(item int) {
	for i := 0; i < item; i++ {
		storeData(fmt.Sprintf("The loop is: %v\n", i), "data.txt")
	}

	fmt.Println("successfully stored")

}
