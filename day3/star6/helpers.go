package main

import (
	// "bufio"
	// "fmt"
	// "os"
	"strconv"
	// "strings"
	// "../helpers"
)

func SliceAtoi(sa []string) []int {
	si := make([]int, 0, len(sa))
	for _, a := range sa {
		i, err := strconv.Atoi(a)
		if err == nil {
			si = append(si, i)
		}

	}
	return si
}
func IntAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func Min(x int,y int) int{
	if(x<y){
		return x
	}else{
		return y
	}
}
