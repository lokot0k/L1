package main

import (
	"fmt"
	"time"
)

func Sleep(duration time.Duration) {
	<-time.After(duration)
}

func main() {
	fmt.Println("Засыпаем, время сейчас: ", time.Now())
	Sleep(5 * time.Second)
	fmt.Println("Проснулись, время сейчас:", time.Now())
}
