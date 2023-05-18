package main

import (
	"fmt"
	"os"
)

func main() {
	channel := make(chan int)

	// this means w're limiting the rate of goroutines executes at the same time
	//in this case 3 goroutines allowed ber execution whens finishes that three executions another three also began
	// channelTwo := make(chan int, 3)
	// storeData("another data stored ", "data.txt")
	go storeDataTwo(50, "data1.txt", channel)
	go storeDataTwo(50, "data2.txt", channel)
	go storeDataTwo(50, "data2.txt", channel)
	go storeDataTwo(50, "data2.txt", channel)

	// if u are using singe in a goroutine close it in its living function
	// else close it manually by using if statement in the loop

	sum := 0

	for chanData := range channel {
		sum += chanData
		if sum == 4 {
			close(channel) // this will close the channel and stops the goroutine execution
		}
		fmt.Println(chanData)
	}
	// <-channel
	// <-channel

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
	// this is the living function to close use the built in <Close> function then pass the channel you are trying to close
}
