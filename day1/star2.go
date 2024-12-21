package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sliceAtoi(sa []string) []int {
	si := make([]int, 0, len(sa))
	for _, a := range sa {
		i, err := strconv.Atoi(a)
		if err == nil {
			si = append(si, i)
		}

	}
	return si
}
func parseLists() ([]int, []int) {
	str1 := ""
	str2 := ""
	dummy := []string{"", ""}
	length := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		dummy = strings.Split(scanner.Text(), "   ")
		length += 1
		str1 = str1 + dummy[0] + ","
		str2 = str2 + dummy[1] + ","
	}

	list1 := sliceAtoi(strings.Split(str1, ","))
	list2 := sliceAtoi(strings.Split(str2, ","))
	return list1, list2
}

func count(list []int) map[int]int {
	res := map[int]int{}
	for _, v := range list {
		dummy, ok := res[v]
		if ok {
			res[v] = dummy + 1
		} else {
			res[v] = 1
		}
	}
	return res
}
func calc(list []int, count map[int]int) int {
	res := 0
	for _, v := range list {
		res += v * count[v]
	}
	return res
}

func main() {
	l1, l2 := parseLists()
	fmt.Println(calc(l1, count(l2)))
}
