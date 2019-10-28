package main

import (
	"log"
	"strings"
)

func RemoveParameter(url string) string {
	index := strings.Index(url, "?")
	if index == -1 {
		return url
	}
	return url[:index]
}

//func RemoveUnit(dt string) string {
//	index := strings.LastIndex(dt, "/")
//	if index == -1 {
//		return dt
//	}
//	return dt[:index]
//}

func main() {
	//url := "https://segmentfault.com/q/1010000019687395?utm_source=tag-newest"
	url := "?"
	//index := strings.Index(url, "?")
	//if index == -1 {
	//	return
	//}

	log.Println(RemoveParameter(url))

	//RemoveUnit()
}
