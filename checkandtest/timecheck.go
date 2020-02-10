package main

import (
	"fmt"
	"time"
)

const (
	layoutISO = "2006-01-02"
)

func main() {
	t := time.Now()
	fmt.Println(t)
	fmt.Println(t.Format(layoutISO))
	unixt := t.Unix()
	fmt.Println(unixt)
	fmt.Println(time.Unix(unixt, 0))
}
