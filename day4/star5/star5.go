package main

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)
func parse() string {
  res :=""
  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
    res += scanner.Text()
  }
  return res
}
func mult(in string) int{
	if((len(in)<8)||(in[:4]!="mul(")||(in[len(in)-1:]!=")")){
		println("string error: ",in)
		return 0
	}else{
		strs:=strings.Split(in[4:len(in)-1],",")
		nums:= SliceAtoi(strs)
		return nums[0]*nums[1]
	}
}
func main() {
  str:=parse()
	regex:=regexp.MustCompile(`mul\([0-9]+,[0-9]+\)`)//`mul\([0-9]+,[0-9]+\)$`)
	strs:=regex.FindAllString(str,-1)
	res:=0
	for _,v :=range strs{
		res+=mult(v)
	}
  println(res)
}
