package golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Bro"
		fmt.Println("Selesai kirim data ke Channel")
	}()

	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
	defer close(channel)
}

func GiveMeReponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Brooo"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)

	go GiveMeReponse(channel)

	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
	defer close(channel)
}
