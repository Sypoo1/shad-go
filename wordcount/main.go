//go:build !solution

package main

import (
	"fmt"
	"os"
)

func helper(err error){
	if err != nil{
		panic(err)
	}
}


func main() {
	m := make(map[string]int)
	var files []string = os.Args[1:]
	for _, file := range files {
		data, err := os.ReadFile(file)
		helper(err)
		for _, s := range data {
			ss := string(s)
			if ss != "\n"{
				m[ss] = m[ss] + 1
			}
		}
	}
	for key, value := range m {
		if value >= 2 {
			fmt.Printf("%d\t%s\n", value, key)
		}
	}
}	
