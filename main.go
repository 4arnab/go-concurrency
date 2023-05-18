package main

import (
	"fmt"
	"os"
)

func main() {
	channel := make(chan int)
	// storeData("another data stored ", "data.txt")
	go storeDataTwo(50, "data1.txt", channel)
	go storeDataTwo(50, "data2.txt", channel)

	<-channel
	<-channel

	fmt.Println(<-channel)
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

func storeDataTwo(item int, fileName string, c chan int) {
	for i := 0; i < item; i++ {
		storeData(fmt.Sprintf("%v -- Dummy data\n", i), fileName)
	}

	fmt.Println("successfully stored")

	c <- 1
}
