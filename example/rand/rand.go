package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		x := r.Intn(100)
		if x == 100 {
			log.Println("100有问题有问题")
		} else if x == 99 {
			log.Println(x)
		}
	}
}
