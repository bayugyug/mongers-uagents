package main

import (
	"fmt"
	"time"
)

func main() {
	pAppStart = time.Now()
	doIt()
	elapsed := time.Since(pAppStart)
	fmt.Println("Takes about: ", elapsed)
}
