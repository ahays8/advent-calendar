package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func sliceAtoi(sa []string) []int {
	si := make([]int, 0, len(sa))
	slices.Sort(sa)
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
func calculation(list1 []int, list2 []int) int {
	res := 0
	for i, _ := range list1 {
		res += int(math.Abs(float64(list1[i] - list2[i])))
	}
	return res
}

func main() {
	fmt.Println(calculation(parseLists()))
}
