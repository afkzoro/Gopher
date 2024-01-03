package main

import (
	"errors"
	"log"
	"time"
)

var ranOnce bool

var (
	ErrNetwork = errors.New("network error")
	ErrInput   = errors.New("input error")
)

func someFunc(data string) error {
	if !ranOnce {
		ranOnce = true
		return ErrNetwork
	}
	return ErrInput
}

// func main() {
// 	for {
// 		err := someFunc("data")
// 		if err == nil {
// 			// Success so exit the loop
// 			break
// 		}
// 		if errors.Is(err, ErrNetwork) {
// 			log.Println("recoverable network error")
// 			time.Sleep(1 * time.Second)
// 			continue
// 		}
// 		log.Println("unrecoverable error")
// 		break // exit loop, as retrying is useless
// 	}
// }
