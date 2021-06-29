package fanout

import (
	"context"
	"fmt"
	"strings"
	"time"
)

func FanoutTest() {
	testData := strings.Split("Канал — модель межпроцессного взаимодействия и "+
		"синхронизации через передачу сообщений в программировании,"+
		" Канал — модель межпроцессного взаимодействия и "+
		"синхронизации через передачу сообщений в программировании", " ")

	inCh := make(chan string)
	outChs := CreateChannels(1, 1)
	fanOut := New(inCh, outChs, 1)
	go func() {
		fanOut.Run()
	}()

	go func() {
		for _, data := range testData {
			inCh <- data
		}
	}()

	for _, ch := range outChs {
		go func(ch chan string) {
			for data := range ch {
				Print(data)
			}
		}(ch)
	}
	time.Sleep(100)
	newCh := make(chan string)
	fanOut.AppendOutChannel(newCh)
	go func(ch chan string) {
		for data := range ch {
			Print(data + " Im new")
		}
	}(newCh)
	time.Sleep(100)
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

type FanOut struct {
	inCh        chan string
	outChs      []chan string
	bufferSize  int
	bufferChans []chan string
	ctx         context.Context
}

func New(in chan string, out []chan string, bufferSize int) *FanOut {
	return &FanOut{inCh: in, outChs: out, bufferSize: bufferSize}
}

func (o *FanOut) Run() {
	o.ctx = context.TODO()
	o.createChansWorkers()
	for data := range o.inCh {
		for _, ch := range o.bufferChans {
			ch <- data
		}
	}
}

func (o *FanOut) Stop() {
	_, cancel := context.WithCancel(o.ctx)
	cancel()
	for _, ch := range o.bufferChans {
		close(ch)
	}
}

func (o *FanOut) AppendOutChannel(ch chan string) {
	bufferCh := make(chan string, o.bufferSize)
	o.bufferChans = append(o.bufferChans, bufferCh)
	o.outChs = append(o.outChs, ch)
	createWorker(ch, bufferCh, o.ctx)
}
func (o *FanOut) createChansWorkers() {
	o.bufferChans = CreateChannels(len(o.outChs), o.bufferSize)
	for i, ch := range o.outChs {
		createWorker(ch, o.bufferChans[i], o.ctx)
	}
}

func createWorker(ch chan string, bufferChan chan string, ctx context.Context) {
	go func(ch chan string, bufferChan chan string, ctx context.Context) {
		for data := range bufferChan {
			if len(bufferChan) == cap(bufferChan) {
				fmt.Printf("Wait for ch %v\n", ch)
			}
			select {
			case ch <- data:
				continue
			case <-ctx.Done():
				return
			}

		}
	}(ch, bufferChan, ctx)
}

func Print(msg string) string {
	fmt.Println(msg)
	return msg
}
