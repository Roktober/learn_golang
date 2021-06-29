package fanin

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"time"
)

func FaninTest() {
	testData1 := strings.Split("Канал — модель межпроцессного взаимодействия", " ")
	testData2 := strings.Split("и синхронизации через передачу сообщений в программировании", " ")

	inChs := CreateChannels(2, 1)
	outCh := make(chan string)
	fanOut := New(inChs, outCh)
	go func() {
		fanOut.Run()
	}()
	go func() {
		for _, data := range testData1 {
			inChs[0] <- data
		}
	}()
	go func() {
		for _, data := range testData2 {
			inChs[1] <- data
		}
	}()
	go func() {
		for data := range outCh {
			Print(data)
		}
	}()
	fanOut.Stop()
	time.Sleep(time.Second * 2)
}
func CreateChannels(count int, buffer int) []chan string {
	channels := make([]chan string, count)
	for i := 0; i < count; i++ {
		channels[i] = make(chan string, buffer)
	}
	return channels
}

type FanIn struct {
	inChs []chan string
	outCh chan string
	ctx   context.Context
}

func New(inChs []chan string, outCh chan string) *FanIn {
	return &FanIn{inChs: inChs, outCh: outCh}
}

func (o *FanIn) Run() {
	o.ctx = context.TODO()
	cases := make([]reflect.SelectCase, len(o.inChs)+1)
	for i, ch := range o.inChs {
		cases[i] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(ch)}
	}
	cases[len(o.inChs)] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(o.ctx.Done())}
	for {
		chNumber, value, ok := reflect.Select(cases)
		if !ok || chNumber == len(o.inChs) {
			fmt.Println("Stop")
			return
		}
		o.outCh <- value.String()
	}
}

func (o *FanIn) Stop() {
	_, cancel := context.WithCancel(o.ctx)
	cancel()
}

func Print(msg string) string {
	fmt.Println(msg)
	return msg
}
