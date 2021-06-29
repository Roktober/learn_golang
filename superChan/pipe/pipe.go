package pipe

import (
	"context"
	"fmt"
	"strings"
	"time"
)

func PipeTest() {
	testData := strings.Split("Канал — модель межпроцессного взаимодействия и "+
		"синхронизации через передачу сообщений в программировании", " ")

	ctx, cancel := context.WithCancel(context.Background())
	ch1, ch2 := make(chan string), make(chan string)

	defer cancel()
	defer close(ch1)
	defer close(ch2)

	go func() {
		Pipe(ctx, ch1, ch2, Print)
	}()
	go func() {
		for _, data := range testData {
			ch1 <- data
		}
	}()
	go func() {
		for data := range ch2 {
			_ = data
		}
	}()
	time.Sleep(time.Second)
}
func Pipe(ctx context.Context, ch1 chan string, ch2 chan string, f func(msg string) string) {
	for data := range ch1 {
		select {
		case ch2 <- f(data):
			continue
		case <-ctx.Done():
			fmt.Println("Stop Pipe!")
			return
		}
	}
}

func Print(msg string) string {
	fmt.Println(msg)
	return msg
}
